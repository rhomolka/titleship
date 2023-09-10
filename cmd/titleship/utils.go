package main

import (
	"log"
	"os"
)

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
