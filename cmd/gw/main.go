package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/aaron-vaz/golang-utils/pkg/fsutil"

	"github.com/aaron-vaz/golang-utils/pkg/errorutil"
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

	if gradleBinary == "" {
		os.Exit(1)
	}

	buildFile := selectGradleBuildFile()

	log.Printf("Using '%s' to run build file '%s' \n", gradleBinary, buildFile)
	fmt.Println("")

	err := os.Chdir(filepath.Dir(buildFile))
	errorutil.ErrCheck(err, true)

	cmd := exec.Command(gradleBinary, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// run the command
	err = cmd.Run()
	errorutil.ErrCheck(err, true)
}

// selectGradleBinary find which gradle binary to use for the project
func selectGradleBinary() string {
	// look for project gradlew file
	foundGradlew := fsutil.FindFile(gradlewFile, "")
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

	log.Printf("%s binary not found in your PATH: \n%s", gradleBinary, os.Getenv("PATH"))

	return ""
}

// getGradleBuildFileLocation first checks for a groovy build file in the working directory
// if the groovy build file is not found it next checks for a kotlin build file
// if both checks find nothing we return the current working directory and let gradle figure out whether the cwd is a gradle project
func selectGradleBuildFile() string {
	groovyBuildFile := fsutil.FindFile(gradleGroovyBuildFile, "")
	kotlinBuildFile := fsutil.FindFile(gradleKotlinBuildFile, "")

	// check if the groovy build file was found
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
