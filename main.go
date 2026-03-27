package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	Alias    string
	Password string
	CN       string
	OU       string
	O        string
	L        string
	ST       string
	C        string
}

func main() {
	if len(os.Args) > 1 {
		cmd := os.Args[1]
		if cmd == "-h" || cmd == "--help" {
			fmt.Println("Usage:")
			fmt.Println("  gka [config_file.txt]  - Generate keystore from config")
			fmt.Println("  gka init               - Generate file-gka.txt template")
			return
		}
		if cmd == "init" {
			fmt.Println("Generating file-gka.txt...")
			err := generateExample()
			if err != nil {
				log.Fatalf("Error generating file-gka: %v", err)
			}
			fmt.Println("Done! Use file-gka.txt as a template.")
			return
		}
	}

	inputFile := "file-gka.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		log.Fatalf("Input file '%s' not found. Please provide a configuration file (e.g., file-gka.txt).", inputFile)
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

func parseConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		switch key {
		case "ALIAS":
			config.Alias = val
		case "PASSWORD":
			config.Password = val
		case "CN":
			config.CN = val
		case "OU":
			config.OU = val
		case "O":
			config.O = val
		case "L":
			config.L = val
		case "ST":
			config.ST = val
		case "C":
			config.C = val
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func generateKeystore(config *Config, outputPath string) error {
	dname := fmt.Sprintf("CN=%s, OU=%s, O=%s, L=%s, ST=%s, C=%s",
		config.CN, config.OU, config.O, config.L, config.ST, config.C)

	cmd := exec.Command("keytool",
		"-genkey", "-v",
		"-keystore", outputPath,
		"-alias", config.Alias,
		"-keyalg", "RSA",
		"-keysize", "2048",
		"-validity", "10000",
		"-dname", dname,
		"-storepass", config.Password,
		"-keypass", config.Password,
		"-deststoretype", "pkcs12",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func generateKeyProperties(config *Config, keystorePath string) error {
	storeFile := strings.TrimPrefix(keystorePath, "android/")
	content := fmt.Sprintf("storePassword=%s\nkeyPassword=%s\nkeyAlias=%s\nstoreFile=%s\n",
		config.Password, config.Password, config.Alias, storeFile)
	return os.WriteFile("android/key.properties", []byte(content), 0644)
}

func generateExample() error {
	content := `# GKA Configuration Template
# Each key-value pair should follow the format: KEY=VALUE

ALIAS=upload        # Key alias (e.g., upload)
PASSWORD=supersecret # Password for both keystore and key
CN=John Doe          # Common Name (e.g., Your Name)
OU=Android           # Organizational Unit (e.g., Android)
O=MyCompany           # Organization (e.g., My Company)
L=Jakarta            # Locality (City)
ST=DKI Jakarta       # State or Province
C=ID                 # Country Code (e.g., ID, US, UK)
`
	return os.WriteFile("file-gka.txt", []byte(content), 0644)
}
