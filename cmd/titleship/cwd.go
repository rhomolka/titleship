package main

import (
	"strings"
)

func init() {
	modmap["GCWD"] = getMassagedDirectory
}

func getMassagedDirectory() string {
	path := currwd
	if strings.HasPrefix(path, envHome) {
		pathfixed := strings.TrimPrefix(path, envHome)
		return "~" + pathfixed
	} else {
		return path
	}
}
