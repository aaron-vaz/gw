package fsutil

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/aaron-vaz/golang-utils/pkg/errorutil"
)

// FindFile recursively searches upwards for a file staring from a directory
func FindFile(file string, dir string) string {
	var result string

	cwd, err := os.Getwd()
	errorutil.ErrCheck(err, true)

	// find filesystem root
	root := FindRootVolume(cwd)

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

// FindRootVolume find the root volume of the path supplied using filepath.VolumeName
// if filepath.VolumeName returns an empty string (on most systems) assume it is linux or darwin based and return /
// if it is in the windows environment filepath.VolumeName will return the drive letter without slashes so the slashes are added before returning the value
func FindRootVolume(path string) string {
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
	log.Fatalln("No root volume found, exiting...")
	return rootVolume
}
