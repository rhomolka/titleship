package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	ansi "github.com/pborman/ansi"
)

// Modules will give a string blob.  Cf
type ModuleFunc func()string
var modmap = make(map[string]ModuleFunc)
var moddefaultstring = "CHZM|HSTN|K8SC|GITS|GCWD"
var modconfig []string

// globals, yeah i know
var  currwd string
var  envHome string
var  escSeqStart string = ""
var  escSeqEnd string = ""

func initialize() {
	var err error

	// globals, ick
	envHome = os.Getenv("HOME")
	currwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	titleshipconfenv := os.Getenv("TITLESHIP_CONF")
	if (len(titleshipconfenv) != 0) {
		modconfig = strings.Split(titleshipconfenv, "|")
	} else {
		modconfig = strings.Split(moddefaultstring, "|")
	}

	// Do we do ansi sequences?
	envTerm := os.Getenv("TERM")
	if	strings.HasPrefix(envTerm, "xterm") ||
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
	finalString := ""

	for _, modname := range(modconfig) {
		modFunc := modmap[modname]
		finalString += modFunc()
	}
	return finalString
}

func main() {
	setLogger()
	initialize()
	titleString := getPrintString()
	fmt.Printf("%s%s%s", escSeqStart, titleString, escSeqEnd)
}