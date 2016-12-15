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
const (
	gradleBinary          = "gradle"
	gradlewFile           = "gradlew"
	gradleGroovyBuildFile = "build.gradle"
	gradleKotlinBuildFile = "build.gradle.kts"
)

func main() {
	gradleBinary := selectGradleBinary()
	buildFile := selectGradleBuildFile()

	log.Printf("Using '%s' to run build file '%s' \n", gradleBinary, buildFile)
	fmt.Println("")

	os.Chdir(filepath.Dir(buildFile))

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
	foundGradlew := findFile(gradlewFile, "")
	if foundGradlew != "" {
		return foundGradlew
	}

	log.Printf("No %s set up for this project \nPlease refer to http://gradle.org/docs/current/userguide/gradle_wrapper.html to set it up", gradlewFile)
	fmt.Println("")

	// if gradlew is not found revert to using the gradle binary
	foundGradle, err := exec.LookPath(gradleBinary)
	if err == nil {
		return foundGradle
	}

	log.Fatalln("%s binary not found in your PATH: \n%s", gradleBinary, os.Getenv("PATH"))

	return ""
}

// getGradleBuildFileLocation first checks for a groovy build file in the working directory
// if the groovy build file is not found it next checks for a kotlin build file
// if both checks find nothing we return the current working directory and let gradle figure out whether the cwd is a gradle project
func selectGradleBuildFile() string {
	groovyBuildFile := findFile(gradleGroovyBuildFile, "")
	kotlinBuildFile := findFile(gradleKotlinBuildFile, "")

	// check if the goovy build file was found
	if groovyBuildFile != "" {
		return groovyBuildFile

		// if the groovy build file was not found look for a kotlin one
	} else if kotlinBuildFile != "" {
		return kotlinBuildFile

		// if both dont exist we assume that the project is not a gradle project
	} else {
		log.Printf("Cannot find gradle build file %s or %s in the project", gradleGroovyBuildFile, gradleKotlinBuildFile)
		return "."
	}
}

// findFile recurcively searches upwards for a file staring from a directory
func findFile(file string, dir string) string {
	var result string

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// find filesystem root
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
// if filepath.VolumeName returns an empty string (on most systems) assume it is linux or darwin based and return /
// if it is windows environement filepath.VolumeName will return the drive letter without slashes so the slashes are added before returning the value
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
