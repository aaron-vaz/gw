package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
		log.Fatalln("Cannot find gradle build file %s in the project", buildFile)
	}

	log.Printf("Using %s to run build file %s \n", gradleBinary, buildFile)
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

	log.Printf("No %s set up for this project \nplease refer to http://gradle.org/docs/current/userguide/gradle_wrapper.html to set it up", defaultGradlew)

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
			result = ""
			findFile(file, filepath.Dir(dir))
		}
	}
	return result
}
