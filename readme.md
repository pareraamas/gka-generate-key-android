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

Create a file named `file-gka.yaml` in the same directory as the application.

```yaml
alias: upload
password: your_secure_password
cn: Nama Lengkap
ou: Unit Organisasi
o: Organisasi
l: Kota
st: Provinsi
c: ID
```

### 2. Run the application

To generate a template configuration file:
```bash
gka init
```
This will create `file-gka.yaml`.

To generate the keystore from `file-gka.yaml`:
```bash
gka
```
(Or specify a different file: `gka my_config.yaml`)

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
    go build -o gka .
    ```

2.  **Move to your bin folder**:
    ```bash
    sudo mv gka /usr/local/bin/
    ```

Now you can run the command from anywhere:

```bash
gka [your_file.txt]
```

If you don't specify a file, it will look for `file-gka.yaml` in the current directory.

---

## Installation from Binary

### 🍎 **macOS**
1.  Download the binary for your architecture (`amd64` for Intel, `arm64` for Apple Silicon).
2.  Move it to your local bin:
    ```bash
    sudo mv gka_darwin_arm64 /usr/local/bin/gka
    sudo chmod +x /usr/local/bin/gka
    ```

### 🐧 **Linux**
1.  Download the `gka_linux_amd64` binary.
2.  Move it to your local bin:
    ```bash
    sudo mv gka_linux_amd64 /usr/local/bin/gka
    sudo chmod +x /usr/local/bin/gka
    ```

### 🪟 **Windows**
1.  Download the `gka_windows_amd64.exe` binary.
2.  Rename it to `gka.exe`.
3.  Add the folder containing `gka.exe` to your System **PATH** environment variable.
4.  Open a new terminal and type `gka`.



## Download Binaries

You can download the compiled binaries for your platform directly from the `build/` folder:

- 🍎 **macOS (Intel)**: [gka_darwin_amd64](build/gka_darwin_amd64)
- 🍎 **macOS (Apple Silicon)**: [gka_darwin_arm64](build/gka_darwin_arm64)
- 🐧 **Linux (amd64)**: [gka_linux_amd64](build/gka_linux_amd64)
- 🪟 **Windows (amd64)**: [gka_windows_amd64.exe](build/gka_windows_amd64.exe)

---

## Build from Source

If you want to build the binaries yourself:

1.  **macOS / Linux**:
    ```bash
    chmod +x build.sh
    ./build.sh
    ```

2.  **Windows (Git Bash or WSL)**:
    ```bash
    ./build.sh
    ```

3.  **Windows (PowerShell/CMD)**:
    ```powershell
    # Build for Windows
    go build -o build/gka_windows_amd64.exe .
    ```
