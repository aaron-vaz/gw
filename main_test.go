package main

import (
	"fmt"
	"os"
	"testing"
)

// get cwd to return back to after a test is run
var cwd = getCurrentWorkingDirectory()

const projectLocation = "test_resources/gradle/project/"
const subProjectLocation = projectLocation + "com.example.app/"
const projectBuildFileLocation = projectLocation + defaultGradleBuildFile
const subProjectBuildFileLocation = subProjectLocation + defaultGradleBuildFile
const gradlewLocation = projectLocation + defaultGradlew
const javaSrcDir = subProjectLocation + "src/main/java/"

type senarios struct {
	file     string
	location string
	expected string
}

func TestMain(t *testing.T) {
	os.Chdir(projectLocation)
	main()

	// change the current working directory back so it doesnt effect the other tests
	os.Chdir(cwd)
}

func TestSelectGradleBinary(t *testing.T) {
	locations := []string{".", projectLocation}

	for _, location := range locations {
		os.Chdir(location)
		result := selectGradleBinary()

		if result == "" {
			t.Error("Correct gradle binary not selected")
		}
	}

	// change the current working directory back so it doesnt effect the other tests
	os.Chdir(cwd)
}

func TestFindFile(t *testing.T) {
	tests := []senarios{
		// find project default build file
		senarios{
			file:     defaultGradleBuildFile,
			location: projectLocation,
			expected: projectBuildFileLocation,
		},
		// find sub project default build file
		senarios{
			file:     defaultGradleBuildFile,
			location: subProjectLocation,
			expected: subProjectBuildFileLocation,
		},
		// find project gradlew
		senarios{
			file:     defaultGradlew,
			location: projectLocation,
			expected: gradlewLocation,
		},
		// find project gradlew from sub directory
		senarios{
			file:     defaultGradlew,
			location: javaSrcDir,
			expected: gradlewLocation,
		},
	}

	for _, test := range tests {
		result := findFile(test.file, test.location)
		if result != test.expected {
			t.Error("build.gradle was not found")
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
