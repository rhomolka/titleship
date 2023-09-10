package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	ansi "github.com/pborman/ansi"
)

// globals, yeah i know
var  currwd string
var  envHome string
var  escSeqStart string = ""
var  escSeqEnd string = ""

func setGlobals() {
	var err error

	envHome = os.Getenv("HOME")

	currwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	envTerm := os.Getenv("TERM")
	if	strings.HasPrefix(envTerm, "xterm") ||
		strings.HasPrefix(envTerm, "rxvt") || 
		strings.HasPrefix(envTerm, "rxvt") {
		escSeqStart = string(ansi.OSC) + "0;"
		escSeqEnd = string(ansi.BEL)
	}
}

func setLogger() {
	f, err := os.OpenFile("/tmp/titleship.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
}

// this will get the string, absent of any term type decorators.
func getPrintString() string {
	chezMoiString := getChezMoiString()
	hostname := getUnqualifiedHostname()
	gitInfoString := getGitInfoString()
	kubeInfoString := getKubeContextString()
	pathstring := getMassagedDirectory(currwd)

	finalString := chezMoiString + hostname + kubeInfoString + gitInfoString + pathstring
	return finalString
}

func getChezMoiString() string {
	if (len(os.Getenv("CHEZMOI")) > 0) {
		return "🇫🇷🏠:"
	} else {
		return ""
	}
}

func getMassagedDirectory(path string) string {
	if strings.HasPrefix(path, envHome) {
		pathfixed := strings.TrimPrefix(path, envHome)
		return "~" + pathfixed
	} else {
		return path
	}
}

func getUnqualifiedHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(hostname, ".")
	return parts[0] + ": "
}

func fileExists(pathToFile string) bool {
	_, err := os.Stat(pathToFile);

	if err == nil {
		return true
	} else {
		if os.IsNotExist(err) {
			return false
		} else {
			log.Fatal(err)
			return false // TODO decorator to make me not need a return?
		}
	}
}

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

func getKubeContextString() string {
	kubeConfigFile := envHome + "/.kube/config"
	if !fileExists(kubeConfigFile) {
		return ""
	}

	fileH, err := os.Open(kubeConfigFile) // For read access.
	if err != nil {
			log.Fatal(err)
	}
	defer fileH.Close()

	scanner := bufio.NewScanner(fileH)
	for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "current-context: ") {
					after, found := strings.CutPrefix(line, "current-context: ")
					if found == true {
							parts := strings.Split(after, "/")
							return "☸️ " + parts[len(parts) - 1] + " "
					}
					break
			}
	}

	return ""
}

func main() {
	setLogger()
	setGlobals()
	titleString := getPrintString()
	fmt.Printf("%s%s%s", escSeqStart, titleString, escSeqEnd)
}