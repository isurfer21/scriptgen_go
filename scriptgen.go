// Here's a Go program that uses the `flag` package to parse command line options for displaying help and version information. It takes the interpreter and script-path arguments from the command line and uses them to generate .cmd, .ps1 and .sh scripts with the specified template content.
// You can compile this program with `go build -o scriptgen` and run it with `./scriptgen [options] [interpreter] [script-path]`.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

var (
	help    bool
	version bool
)

const (
	Version = "1.0.0"
)

func init() {
	flag.BoolVar(&help, "help", false, "Show help menu")
	flag.BoolVar(&help, "h", false, "Show help menu (shorthand)")
	flag.BoolVar(&version, "version", false, "Show version information")
	flag.BoolVar(&version, "v", false, "Show version information (shorthand)")
}

func main() {
	flag.Parse()

	if help {
		fmt.Println("Usage: scriptgen [options] [interpreter] [script-path]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		return
	}

	if version {
		fmt.Printf("scriptgen version %s\n", Version)
		return
	}

	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("Error: Missing required arguments: interpreter and script-path")
	}

	interpreter := args[0]
	scriptPath := args[1]

	scriptName := strings.TrimSuffix(filepath.Base(scriptPath), filepath.Ext(scriptPath))

	template := fmt.Sprintf(`@ECHO off
SETLOCAL
SET dp0=%~dp0

SET "_prog=%s"
SET PATHEXT=%%PATHEXT:;.%s;=;%

ENDLOCAL & GOTO #_undefined_# 2>NUL || title %%COMSPEC% & "%_prog%"  "%dp0%\%s" %*
`, interpreter, interpreter, scriptPath)

	err := ioutil.WriteFile(scriptName+".cmd", []byte(template), 0644)
	if err != nil {
		log.Fatal(err)
	}

	template = fmt.Sprintf(`#!/usr/bin/env pwsh
$basedir=Split-Path $MyInvocation.MyCommand.Definition -Parent

$exe=""
if ($PSVersionTable.PSVersion -lt "6.0" -or $IsWindows) {
  # Fix case when both the Windows and Linux builds of Node
  # are installed in the same directory
  $exe=".exe"
}
$ret=0

# Support pipeline input
if ($MyInvocation.ExpectingInput) {
  $input | & "%s$exe"  "$basedir/%s" $args
} else {
  & "%s$exe"  "$basedir/%s" $args
}
$ret=$LASTEXITCODE

exit $ret
%s %s`, interpreter, scriptPath, interpreter, scriptPath)

	err = ioutil.WriteFile(scriptName+".ps1", []byte(template), 0644)
	if err != nil {
		log.Fatal(err)
	}

	template = fmt.Sprintf(`#!/bin/sh
basedir=$(dirname "$(echo "$0" | sed -e 's,\\,/,g')")

case 'uname' in
    *CYGWIN*|*MINGW*|*MSYS*) basedir='cygpath -w "$basedir"';;
esac

exec node  "$basedir/%s" "$@"
`, interpreter, scriptPath)

	err = ioutil.WriteFile(scriptName+".sh", []byte(template), 0755)
	if err != nil {
		log.Fatal(err)
	}
}