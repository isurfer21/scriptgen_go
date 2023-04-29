# scriptgen_go
ScriptGen is a command-line based script generator tool that simplifies the process of running scripts in any language.

## Introduction

When a script is executed using an interpreter, the command line statement can become quite long as it includes the interpreter path, script path, and all required arguments. ScriptGen addresses this issue by consolidating the interpreter and script paths into a short command that can be easily accessed from anywhere in the system. This means that when calling the script via command, only the arguments need to be passed.

ScriptGen generates commands for all operating systems. For Windows, it generates batch and PowerShell scripts. For macOS, it generates Shell scripts and PowerShell scripts which can be used if PowerShell is installed. For Linux, it generates Shell scripts. This makes it easy to run scripts on any platform without having to manually create the appropriate command.

### Working

Here's a Go program that uses the `flag` package to parse command line options for displaying help and version information. It takes the interpreter and script-path arguments from the command line and uses them to generate `.cmd`, `.ps1` and `.sh` scripts with the specified template content.

## Build

Here is the procedure for building the `scriptgen_go` project on Windows, macOS, and Linux:

1. Open a command prompt (Windows) or terminal (macOS and Linux) and navigate to the directory where you want to clone the `scriptgen_go` repository.

2. Clone the `scriptgen_go` repository from GitHub by running the following command:
```
git clone https://github.com/isurfer21/scriptgen_go.git
```

3. Navigate to the cloned repository by running the following command:
```
cd scriptgen_go
```

### For development

4. Build the project using `go build`. On Windows, run the following command:
```
go build -o scriptgen.exe
```
On macOS and Linux, run the following command:
```
go build -o scriptgen
```

5. After running these commands, you should have a local copy of the ScriptGen project that is ready to use.

6. The executable binary can be found in the current directory. On Windows, it is named `scriptgen.exe`, while on macOS and Linux it is simply named `scriptgen`.

### For production release

7. To generate an optimized binary for this Go program, you can use the `go build` command with the `-ldflags` option to set the appropriate flags for optimization. For example, to generate an optimized binary for the `scriptgen.go` program, you can use the following command:

```
go build -ldflags="-s -w" -o scriptgen
```

This will generate a binary file named `scriptgen` (or `scriptgen.exe` on Windows) that is optimized for production release.

### Cross platform production build

To generate an optimized binary for production release of the `scriptgen.go` program for different operating systems, you can use the `GOOS` and `GOARCH` environment variables to specify the target operating system and architecture, respectively. Here are the commands to generate the production build of `scriptgen.go` for Windows, macOS, and Linux:

- For Windows (64-bit):
```
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -o scriptgen.exe
```

- For macOS (64-bit):
```
export GOOS=darwin
export GOARCH=amd64
go build -ldflags="-s -w" -o scriptgen
```

- For Linux (64-bit):
```
export GOOS=linux
export GOARCH=amd64
go build -ldflags="-s -w" -o scriptgen
```

These commands will generate an optimized binary file named `scriptgen.exe` for Windows and `scriptgen` for macOS and Linux.

## Usage

To use _ScriptGen_, run the `scriptgen` command followed by any options, the interpreter, and the path to the script you want to run. Here is the basic usage syntax:

```
scriptgen [options] [interpreter] [script-path]
```

### Options

_ScriptGen_ supports the following options:

- `-h` or `-help`: Show the help menu.
- `-v` or `-version`: Show version information.

### Examples

The following help menu shows the available options and arguments for using the tool:

```
> scriptgen -h
Usage: scriptgen [options] [interpreter] [script-path]
Options:
  -h    Show help menu (shorthand)
  -help
        Show help menu
  -v    Show version information (shorthand)
  -version
        Show version information
```

Here are some examples of how to use _ScriptGen_ with different interpreters:

- **Node.js:** To run a Node.js script, specify `node` as the interpreter and provide the path to the script. For example:
```
scriptgen node /path/to/script.js
```

- **Python:** To run a Python script, specify `python` as the interpreter and provide the path to the script. For example:
```
scriptgen python /path/to/script.py
```

- **Java:** To run a Java class file, specify `java` as the interpreter and provide the path to the class file (without the `.class` extension). For example:
```
scriptgen java /path/to/MyClass
```

## Installation 

Here are some ways to install the _ScriptGen_ tool, written in Go, on your local computer:

1. **Downloading a pre-built binary:** The developers of _ScriptGen_ provide pre-built binaries for your operating system, you can download the appropriate binary from the project's website or GitHub releases page. Once downloaded, you can place the binary in a directory that is in your system's `PATH` so that you can run the `scriptgen` command from anywhere.

2. **Building from source:** If you have a Go development environment set up on your computer, you can build _ScriptGen_ from source. To do this, clone the _ScriptGen_ repository from GitHub and navigate to the cloned repository. Then, run the `go build` command to build the project. This will create an executable binary that you can place in a directory that is in your system's `PATH`.

3. **Using `go get`:** Since the _ScriptGen_ tool is available as a Go package, you can use the `go get` command to download and install it. To do this, run the following command:
```
go get github.com/isurfer21/scriptgen_go
```
This will download the _ScriptGen_ package and its dependencies, and install the `scriptgen` binary in your `$GOPATH/bin` directory. Make sure that this directory is in your system's `PATH` so that you can run the `scriptgen` command from anywhere.
