package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		cmd := os.Args[1]
		if cmd == "-h" || cmd == "--help" {
			fmt.Println("Usage:")
			fmt.Println("  gka [config_file.yaml]  - Generate keystore from config")
			fmt.Println("  gka init                - Generate file-gka.yaml template")
			return
		}
		if cmd == "init" {
			fmt.Println("Generating file-gka.yaml...")
			err := generateExample()
			if err != nil {
				log.Fatalf("Error generating file-gka: %v", err)
			}
			fmt.Println("Done! Use file-gka.yaml as a template.")
			return
		}
	}

	inputFile := "file-gka.yaml"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		log.Fatalf("Input file '%s' not found. Please provide a configuration file (e.g., file-gka.yaml). or run 'gka init' to generate a template.", inputFile)
	}

	fmt.Println("Cleaning up existing files...")
	os.RemoveAll("android") // Start fresh

	fmt.Println("Creating directories...")
	err := os.MkdirAll("android/app", 0755)
	if err != nil {
		log.Fatalf("Error creating directories: %v", err)
	}

	config, err := parseConfig(inputFile)
	if err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	fmt.Println("Generating keystore...")
	keystorePath := fmt.Sprintf("android/app/%s.keystore", config.Alias)
	err = generateKeystore(config, keystorePath)
	if err != nil {
		log.Fatalf("Error generating keystore: %v", err)
	}

	fmt.Println("Generating android/key.properties...")
	err = generateKeyProperties(config, keystorePath)
	if err != nil {
		log.Fatalf("Error generating key.properties: %v", err)
	}

	fmt.Println("Done!")
}
