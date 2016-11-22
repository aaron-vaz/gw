package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// default gradle values
const defaultGradle = "gradle"
const defaultGradlew = "gredlew"
const defaultGradleBuildFile = "build.gradle"

func main() {
	buildFile := findFile(defaultGradleBuildFile, "")
	gradleBinary := selectGradleBinary()
	buildArgs := os.Args

	if buildFile != "" {
		os.Chdir(filepath.Dir(buildFile))
	} else {
		fmt.Printf("Cannot find gradle build file %s in the project", buildFile)
	}

	fmt.Printf("Using %s to run build file %s", gradleBinary, buildFile)
	out, err := exec.Command(gradleBinary, buildArgs...).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// prints the output from the gradle command to standard out
	fmt.Fprintln(os.Stdout, out)
}

// selectGradleBinary find which gradle binary to use for the project
func selectGradleBinary() string {
	// look for project gradlew file
	foundGradlew := findFile(defaultGradlew, "")
	if foundGradlew != "" {
		return foundGradlew
	}

	log.Printf("\nNo %s set up for this project \n", defaultGradlew)
	log.Println("please refer to http://gradle.org/docs/current/userguide/gradle_wrapper.html to set it up")

	// if gradlew is not found revert to using the gradle binary
	foundGradle, err := exec.LookPath(defaultGradle)
	if err == nil {
		return foundGradle
	}

	log.Printf("\n%s not found in your PATH: ", defaultGradle)
	log.Println(os.Getenv("PATH"))

	return ""
}

// findFile recurcively searches upwards for a file staring from a directory
func findFile(file string, dir string) string {
	var result string

	// if no dir value is supplied default to the current working directory
	if dir == "" {
		var err error
		dir, err = os.Getwd()

		if err != nil {
			log.Fatal(err)
		}
	}

	result = filepath.Join(dir, file)
	if dir != "/" {
		if _, err := os.Stat(result); os.IsNotExist(err) {
			findFile(file, filepath.Dir(dir))
			result = ""
		}
	}
	return result
}
