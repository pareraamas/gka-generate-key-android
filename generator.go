package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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
	// Calculate relative path for storeFile in key.properties
	// Since key.properties is in android/ and keystore is in android/app/
	// The path should be 'app/alias.keystore'
	storeFile := strings.TrimPrefix(keystorePath, "android/")

	content := fmt.Sprintf("storePassword=%s\nkeyPassword=%s\nkeyAlias=%s\nstoreFile=%s\n",
		config.Password, config.Password, config.Alias, storeFile)

	return os.WriteFile("android/key.properties", []byte(content), 0644)
}

func generateExample() error {
	content := `# GKA Configuration Template (YAML)
# Complete the details below to generate your Android Keystore

# Key alias (e.g., upload)
alias: upload

# Password for both keystore and key
password: supersecret

# Distinguished Name details
cn: John Doe          # Common Name (Full Name)
ou: Android           # Organizational Unit
o: MyCompany          # Organization
l: Jakarta            # Locality (City)
st: DKI Jakarta       # State or Province
c: ID                 # Country Code (e.g., ID, US)
`
	return os.WriteFile("file-gka.yaml", []byte(content), 0644)
}
