package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
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
	defaultGradlewFile     = "gradlew"

	groovyProjectBuildFileLocation    = groovyProjectLocation + defaultBuildFile
	groovySubProjectBuildFileLocation = groovySubProjectLocation + defaultBuildFile

	kotlinProjectBuildFileKotlinLocation    = kotlinProjectLocation + defaultBuildFileKotlin
	kotlinSubProjectBuildFileKotlinLocation = kotlinSubProjectLocation + defaultBuildFileKotlin

	gradleLocation  = "test_resources/gradle/binary/"
	gradlewLocation = groovyProjectLocation + defaultGradlewFile

	javaSrcDir = groovySubProjectLocation + "src/main/java/"
)

type scenarios struct {
	file     string
	location string
	expected string
}

func TestMain(t *testing.T) {
	os.Chdir(groovyProjectLocation)
	main()

	// change the current working directory back so it doesnt effect the other tests
	defer os.Chdir(cwd)
}

func TestSelectGradleBinary(t *testing.T) {
	absGradlePath, _ := filepath.Abs(gradleLocation)

	// test that error path before adding binary to path
	os.Chdir("/tmp")
	result := selectGradleBinary()

	if result != "" {
		t.Error("empty string should have been returned")
	}

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

	// change the current working directory back so it doesnt effect the other tests
	defer os.Chdir(cwd)
}

func TestSelectGradleBuildFile(t *testing.T) {
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

func TestFindFile(t *testing.T) {
	tests := []scenarios{
		// find project default build file
		{
			file:     defaultBuildFile,
			location: groovyProjectLocation,
			expected: groovyProjectBuildFileLocation,
		},
		// find project default kotlin build file
		{
			file:     defaultBuildFileKotlin,
			location: kotlinProjectLocation,
			expected: kotlinProjectBuildFileKotlinLocation,
		},
		// find groovy sub project default build file
		{
			file:     defaultBuildFile,
			location: groovySubProjectLocation,
			expected: groovySubProjectBuildFileLocation,
		},
		// find kotlin sub project default build file
		{
			file:     defaultBuildFileKotlin,
			location: kotlinSubProjectLocation,
			expected: kotlinSubProjectBuildFileKotlinLocation,
		},
		// find project gradlew
		{
			file:     defaultGradlewFile,
			location: groovyProjectLocation,
			expected: gradlewLocation,
		},
		// find project gradlew from sub directory
		{
			file:     defaultGradlewFile,
			location: javaSrcDir,
			expected: gradlewLocation,
		},
	}

	for _, test := range tests {
		result := findFile(test.file, test.location)
		if result != test.expected {
			t.Errorf("actual: %s, expected: %s", result, test.expected)
		}
	}
}

func TestFindRootVolume(t *testing.T) {
	result := findRootVolume("/tmp/something")
	if result != "/" {
		t.Error("Didnt find the correct root volume")
	}
}

func getCurrentWorkingDirectory() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return cwd
}

func getAbsPath(file string) string {
	path, err := filepath.Abs(file)
	if err != nil {
		fmt.Println(err)
	}
	return path
}
