package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
//	yaml "gopkg.in/yaml.v3"
)

func init() {
	modmap["K8SC"] = getKubeContextString
	modmap["K8CN"] = getKubeClusterNamespaceString
}

func getKubeClusterNamespaceString() string {
	kubeConfigEnv := os.Getenv("KUBECONFIG")
	if (len(kubeConfigEnv) == 0) {
		kubeConfigEnv = envHome + "/.kube/config"
	}
	
	configfiles := strings.Split(kubeConfigEnv, ":")
	for _, configfilename := range(configfiles) {
		fmt.Println(configfilename)
	}

	return ""
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
