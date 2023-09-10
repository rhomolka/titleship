package main

import (
	"os"
)

func init() {
	modmap["CHZM"] = getChezMoiString
}

func getChezMoiString() string {
	if (len(os.Getenv("CHEZMOI")) > 0) {
		return "🇫🇷:"
	} else {
		return ""
	}
}
