package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Info struct {
	Version    string `json:"version"`
	Environment string `json:"environment"`
}

func main() {
	// Check if the argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide the argument.")
		return
	}

	// Simulate the provided argument
	argument := os.Args[1]

	// The input string
	//input := `### Version
	
	//v4.15.6
	
	//### Environment
	
	//ite`

	// Extract version and environment
	version, environment := extractVersionAndEnvironment(argument)

	// Print the results
	fmt.Println("Version:", version)
	fmt.Println("Environment:", environment)
}

func extractVersionAndEnvironment(input string) (string, string) {
	// Define regex patterns for version and environment
	versionPattern := regexp.MustCompile(`(?m)### Version\s*\s*(v[\d.]+)`)
	environmentPattern := regexp.MustCompile(`(?m)### Environment\s*\s*(\w+)`) //(`(?m)### Environment\s*\n\s*(\w+)`)

	// Extract the version using the regex pattern
	versionMatch := versionPattern.FindStringSubmatch(input)
	version := ""
	if len(versionMatch) > 1 {
		version = strings.TrimSpace(versionMatch[1])
	}

	// Extract the environment using the regex pattern
	environmentMatch := environmentPattern.FindStringSubmatch(input)
	environment := ""
	if len(environmentMatch) > 1 {
		environment = strings.TrimSpace(environmentMatch[1])
	}

	return version, environment
}