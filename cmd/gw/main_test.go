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

	defaultBuildFile       = "build.gradle"
	defaultBuildFileKotlin = "build.gradle.kts"

	groovyProjectBuildFileLocation    = groovyProjectLocation + defaultBuildFile
	groovySubProjectBuildFileLocation = groovySubProjectLocation + defaultBuildFile

	kotlinProjectBuildFileKotlinLocation    = kotlinProjectLocation + defaultBuildFileKotlin
	kotlinSubProjectBuildFileKotlinLocation = kotlinSubProjectLocation + defaultBuildFileKotlin

	gradleLocation = "test_resources/gradle/binary/"
)

func Test_Main(t *testing.T) {
	os.Chdir(groovyProjectLocation)

	// change the current working directory back so it doesn't effect the other tests
	defer os.Chdir(cwd)

	main()
}

func Test_SelectGradleBinary(t *testing.T) {
	absGradlePath, _ := filepath.Abs(gradleLocation)

	path := os.Getenv("PATH")

	// clean PATH
	os.Setenv("PATH", "")

	// test that error path before adding binary to path
	os.Chdir("/tmp")
	result := selectGradleBinary()

	if result != "" {
		t.Errorf("empty string should have been returned, got = %s", result)
	}

	// set path back
	os.Setenv("PATH", path)

	// check other paths
	os.Setenv("PATH", os.Getenv("PATH")+":"+absGradlePath)
	locations := []string{
		groovyProjectLocation,
		groovySubProjectLocation,
		kotlinProjectLocation,
		kotlinSubProjectLocation,
		"/tmp",
	}

	for _, location := range locations {
		os.Chdir(location)
		result := selectGradleBinary()

		if result == "" {
			t.Error("Correct gradle binary not selected")
		}
	}

	// change the current working directory back so it doesn't effect the other tests
	defer os.Chdir(cwd)
}

func Test_SelectGradleBuildFile(t *testing.T) {
	type scenarios struct {
		location string
		expected string
	}

	tests := []scenarios{
		// cd to groovy project and find default build file
		{
			location: getAbsPath(groovyProjectLocation),
			expected: getAbsPath(groovyProjectBuildFileLocation),
		},
		// cd to kotlin project and find default build file
		{
			location: getAbsPath(kotlinProjectLocation),
			expected: getAbsPath(kotlinProjectBuildFileKotlinLocation),
		},
		// cd to groovy sub project and find default build file
		{
			location: getAbsPath(groovySubProjectLocation),
			expected: getAbsPath(groovySubProjectBuildFileLocation),
		},
		// cd to kotlin sub project and find default build file
		{
			location: getAbsPath(kotlinSubProjectLocation),
			expected: getAbsPath(kotlinSubProjectBuildFileKotlinLocation),
		},
		// if default build files are not found check that cwd is returned
		{
			location: "/tmp",
			expected: ".",
		},
	}

	for _, test := range tests {
		os.Chdir(test.location)
		result := selectGradleBuildFile()
		if result != test.expected {
			t.Errorf("actual: %s, expected: %s", result, test.expected)
		}
	}

	defer os.Chdir(cwd)
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
