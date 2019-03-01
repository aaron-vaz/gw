package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/aaron-vaz/golang-utils/pkg/errorutil"
)

// get cwd to return back to after a test is run
var cwd = getCurrentWorkingDirectory()

const (
	groovyProjectLocation    = "test_resources/gradle/groovy_project/"
	groovySubProjectLocation = groovyProjectLocation + "com.example.app/"

	kotlinProjectLocation    = "test_resources/gradle/kotlin_project/"
	kotlinSubProjectLocation = kotlinProjectLocation + "com.example.app/"

	groovyProjectBuildFileLocation    = groovyProjectLocation + gradleGroovyBuildFile
	groovySubProjectBuildFileLocation = groovySubProjectLocation + gradleGroovyBuildFile

	kotlinProjectBuildFileKotlinLocation    = kotlinProjectLocation + gradleKotlinBuildFile
	kotlinSubProjectBuildFileKotlinLocation = kotlinSubProjectLocation + gradleKotlinBuildFile

	gradleLocation = "test_resources/gradle/binary/"
)

func TestSelectGradleBuildFile(t *testing.T) {
	tests := []struct {
		name     string
		location string
		want     string
	}{
		{
			name:     "Test that the groovy build file is returned when launched from the root of the groovy project",
			location: getAbsPath(groovyProjectLocation),
			want:     getAbsPath(groovyProjectBuildFileLocation),
		},

		{
			name:     "Test that the kotlin build file is returned when launched from the root of the kotlin project",
			location: getAbsPath(kotlinProjectLocation),
			want:     getAbsPath(kotlinProjectBuildFileKotlinLocation),
		},

		{
			name:     "Test that the groovy build file from the sub project is returned when launched from the sub project",
			location: getAbsPath(groovySubProjectLocation),
			want:     getAbsPath(groovySubProjectBuildFileLocation),
		},

		{
			name:     "Test that the kotlin build file from the sub project is returned when launched from the sub project",
			location: getAbsPath(kotlinSubProjectLocation),
			want:     getAbsPath(kotlinSubProjectBuildFileKotlinLocation),
		},

		{
			name:     "Test that the kotlin build file from the sub project is returned when launched from java dir",
			location: getAbsPath(kotlinSubProjectLocation + "src/main/java"),
			want:     getAbsPath(kotlinSubProjectBuildFileKotlinLocation),
		},

		{
			name:     "Test that the current dir (.) is returned when no build files could be found",
			location: os.TempDir(),
			want:     ".",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Chdir(tt.location)
			defer os.Chdir(cwd)

			if got := SelectGradleBuildFile(); got != tt.want {
				t.Errorf("SelectGradleBuildFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectGradleBinary(t *testing.T) {
	tests := []struct {
		name     string
		location string
		pathFunc func()
		want     string
	}{
		{
			name:     "Test gradlew script is selected when launched from root of projet",
			location: getAbsPath(kotlinProjectLocation),
			want:     getAbsPath(kotlinProjectLocation + gradlewFile),
		},

		{
			name:     "Test gradlew script is selected when launched from sub project",
			location: getAbsPath(kotlinSubProjectLocation),
			want:     getAbsPath(kotlinProjectLocation + gradlewFile),
		},

		{
			name:     "Test gradlew script is selected when launched from java dir",
			location: getAbsPath(kotlinSubProjectLocation + "src/main/java"),
			want:     getAbsPath(kotlinProjectLocation + gradlewFile),
		},

		{
			name:     "Test gradle binary is selected when no build files are found",
			location: os.TempDir(),
			pathFunc: func() {
				os.Setenv("PATH", getAbsPath(gradleLocation))
			},
			want: getAbsPath(gradleLocation + gradleBinary),
		},

		{
			name:     "Test empty string is returned when no gradlew script is found and gradle isnt installed",
			location: os.TempDir(),
			pathFunc: func() {
				os.Setenv("PATH", "")
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// cleanup after test
			path := os.Getenv("PATH")
			defer os.Setenv("PATH", path)

			if tt.pathFunc != nil {
				tt.pathFunc()
			}

			os.Chdir(tt.location)
			defer os.Chdir(cwd)

			if got := SelectGradleBinary(); got != tt.want {
				t.Errorf("SelectGradleBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStart(t *testing.T) {
	tests := []struct {
		name     string
		location string
		pathFunc func()
		want     int
	}{
		{
			name:     "Test happy path",
			location: getAbsPath(kotlinProjectLocation),
			want:     0,
		},

		{
			name:     "Test error path",
			location: os.TempDir(),
			pathFunc: func() {
				os.Setenv("PATH", "")
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// cleanup after test
			path := os.Getenv("PATH")
			defer os.Setenv("PATH", path)

			if tt.pathFunc != nil {
				tt.pathFunc()
			}

			os.Chdir(tt.location)
			defer os.Chdir(cwd)

			if got := Start(); got != tt.want {
				t.Errorf("Start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	os.Chdir(getAbsPath(kotlinProjectLocation))
	defer os.Chdir(cwd)

	main()
}

func getCurrentWorkingDirectory() string {
	cwd, err := os.Getwd()
	errorutil.ErrCheck(err, false)
	return cwd
}

func getAbsPath(file string) string {
	path, err := filepath.Abs(file)
	errorutil.ErrCheck(err, false)
	return path
}
