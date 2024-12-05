package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
)

var version = "(devel)"

// printVersion prints the application version
func printVersion() {
	if version != "(devel)" {
		// Use version injected by -ldflags
		fmt.Printf("Version: %s\n", version)
		return
	}

	// Fallback to build info when version is not set
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Version: unknown (unable to read build info)")
		return
	}

	if buildInfo.Main.Version != "" {
		fmt.Printf("Version: %s\n", buildInfo.Main.Version)
	} else {
		fmt.Println("Version: (devel)")
	}
}

func main() {
	// -v, --version flag retrives the currently installed version
	currentVersionFlag := mainFlagSet.Bool("version", false, "Prints the current version")
	shortCurrentVersionFlag := mainFlagSet.Bool("v", false, "Prints the current version")
	flag.Parse()

	if *currentVersionFlag || *shortCurrentVersionFlag {
		printVersion()
		os.Exit(0)
	}
}
