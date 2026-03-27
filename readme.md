<p align="center">
  <img src="assets/logo.png" width="300" alt="GKA Logo">
</p>

# GKA (Keystore Generator Android CLI)

A Go-based terminal application to generate a Java Keystore (`.keystore`) and `key.properties` file from a configuration file.

## Prerequisites

- **Go**: Installed and reachable in your terminal.
- **Java Development Kit (JDK)**: `keytool` must be available in your system path.

## Usage Instructions

### 1. Prepare your configuration

Create a file named `file-gka.txt` in the same directory as the application.

```text
ALIAS=upload
PASSWORD=your_secure_password
CN=Nama Lengkap
OU=Unit Organisasi
O=Organisasi
L=Kota
ST=Provinsi
C=ID
```

### 2. Run the application

To generate a template configuration file:
```bash
gka init
```
This will create `example.txt`.

To generate the keystore from `file-gka.txt`:
```bash
gka
```
(Or specify a different file: `gka my_config.txt`)

### 3. Generated Files

After the command finishes, you will see a new `android` directory containing:
- `android/app/winda.keystore`: Your generated keystore file (named after your ALIAS).
- `android/key.properties`: A properties file pointing to the keystore:
  ```properties
  storePassword=...
  keyPassword=...
  keyAlias=winda
  storeFile=app/winda.keystore
  ```

## Installation

To install the application as a global command on your Mac:

1.  **Build the binary**:
    ```bash
    go build -o gka main.go
    ```

2.  **Move to your bin folder**:
    ```bash
    sudo mv gka /usr/local/bin/
    ```

Now you can run the command from anywhere:

```bash
gka [your_file.txt]
```

If you don't specify a file, it will look for `file-gka.txt` in the current directory.
