package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func init() {
	modmap["K8SC"] = getKubeContextString
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
