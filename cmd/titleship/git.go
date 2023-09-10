package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func init() {
	modmap["GITS"] = getGitInfoString
}

// git status -b -s --porcelain


func getGitRefString(pathToHEAD string) string {
	fileH, err := os.Open(pathToHEAD) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer fileH.Close()

	retstring := ""
	scanner := bufio.NewScanner(fileH)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ref: ") {
			after, found := strings.CutPrefix(line, "ref: ")
			if found == true {
				parts := strings.Split(after, "/")
				retstring = parts[len(parts) - 1]
			}
			break
		}
	}

	return retstring
}

func getGitRefFile(startDir string) string {
	dirIter := startDir
	
	subpaths := strings.Split(startDir, "/")

	for true {
		var found bool
		gitDir := dirIter + "/.git"
		gitRef := gitDir + "/HEAD"
		if fileExists(gitRef) {
			return gitRef
		}

		if (len(subpaths) == 2) {
			break
		}

		dirIter, found = strings.CutSuffix(dirIter, "/" + subpaths[len(subpaths)-1])
		if !found {
			log.Fatal("Could not find directory slice")
		}
		subpaths = subpaths[:len(subpaths)-1]
	}
	// nothing found
	return ""
}

func getGitInfoString() string {
	gitRefFile := getGitRefFile(currwd)
	retString := ""
	if gitRefFile != "" {
		retString = "📦 " + getGitRefString(gitRefFile) + " "
	}

	return retString
	// TODO traverse up directory tree
}
