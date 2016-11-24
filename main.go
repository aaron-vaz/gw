package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// default gradle values
const defaultGradle = "gradle"
const defaultGradlew = "gradlew"
const defaultGradleBuildFile = "build.gradle"

func main() {
	buildFile := findFile(defaultGradleBuildFile, "")
	gradleBinary := selectGradleBinary()

	if buildFile != "" {
		os.Chdir(filepath.Dir(buildFile))
	} else {
		log.Fatalf("Cannot find gradle build file %s in the project", defaultGradleBuildFile)
	}

	log.Printf("Using %s to run build file %s \n", gradleBinary, buildFile)
	fmt.Println("")
	cmd := exec.Command(gradleBinary, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// run the command
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

// selectGradleBinary find which gradle binary to use for the project
func selectGradleBinary() string {
	// look for project gradlew file
	foundGradlew := findFile(defaultGradlew, "")
	if foundGradlew != "" {
		return foundGradlew
	}

	log.Printf("No %s set up for this project \nPlease refer to http://gradle.org/docs/current/userguide/gradle_wrapper.html to set it up", defaultGradlew)
	fmt.Println("")

	// if gradlew is not found revert to using the gradle binary
	foundGradle, err := exec.LookPath(defaultGradle)
	if err == nil {
		return foundGradle
	}

	log.Printf("%s binary not found in your PATH: \n%s", defaultGradle, os.Getenv("PATH"))
	fmt.Println("")

	return ""
}

// findFile recurcively searches upwards for a file staring from a directory
func findFile(file string, dir string) string {
	var result string

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	root := findRootVolume(cwd)

	// if no dir value is supplied default to the current working directory
	if dir == "" {
		dir = cwd
	}

	// traverse up the directory structure looking for the file
	// stops when the file is found or when the root directory has been reached
	for dir != root {
		result = filepath.Join(dir, file)
		if _, err := os.Stat(result); err == nil {
			return result
		}
		dir = filepath.Dir(dir)
	}
	return ""
}

// findRootVolume find the root volume of the path supplied using filepath.VolumeName
// if filepath.VolumeName returns an empty string (on most systems) assume it is unix based and return /
func findRootVolume(path string) string {
	rootVolume := filepath.VolumeName(path)
	if rootVolume == "" {
		if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
			return "/"
		}
	} else {
		if runtime.GOOS == "windows" {
			return rootVolume + "\\"
		}
	}
	log.Fatalln("No root volume found, exiting")
	return rootVolume
}
