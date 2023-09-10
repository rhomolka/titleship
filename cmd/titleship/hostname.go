package main

import (
	"log"
	"os"
	"strings"
)

func init() {
	modmap["HSTN"] = getUnqualifiedHostname
}

func getUnqualifiedHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(hostname, ".")
	return parts[0] + ": "
}
