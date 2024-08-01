# Enigma

Enigma is an offline password cracker written in Go. It supports multiple hash types and uses dictionary attacks to crack password hashes.

## Features

- **Multi-Hash Support:** Supports MD5, SHA-1, and SHA-256 hash types.
- **Auto-Detection:** Automatically detects the hash type if not specified.
- **Dictionary Attack:** Uses a dictionary attack strategy to crack password hashes.
- **Scalable:** Can handle large wordlists efficiently.

## Installation

### Prerequisites

- Go 1.20 or later
- Git

### Steps

### Download the Executable
 Download the latest version of `enigma` from the [releases page](https://github.com/Spandan7724/Enigma/releases).

 ### Or to build it yourself

1. Clone the repository:

    ```sh
    git clone https://github.com/Spandan7724/Enigma.git
    cd enigma
    ```

2. Build the project:

    ```sh
    go build -o enigma ./cmd
    ```

## Usage

### Basic Usage

To run Enigma and crack a password hash, use the following command:

```sh
./enigma crack --target <target_hash> --wordlist <path_to_wordlist>
```

### Examples
#### Crack an MD5 Hash
```bash
./enigma crack --target 482c811da5d5b4bc6d497ffa98491e38 --wordlist rockyou.txt
```
#### Crack a SHA-1 Hash
```bash
./enigma crack --target cbfdac6008f9cab4083784cbd1874f76618d2a97 --wordlist rockyou.txt
```
#### Crack a SHA-256 Hash
```bash
./enigma crack --target ef92b778bafee0b79c30b84163206e0eb5c3d955de5b7c00362a9c8c775c6e9a --wordlist rockyou.txt
```
### Optional Parameters

#### Specify Hash Type (Optional)
If you want to specify the hash type explicitly, use the `--hash` flag:
```bash
./enigma crack --hash sha256 --target <target_hash> --wordlist <path_to_wordlist>
```
### Wordlist
Enigma uses a dictionary attack to crack passwords. You can use any wordlist file, such as the popular `rockyou.txt`. Ensure your wordlist file is formatted with one password per line.

The `rockyou.txt` file is available in the [releases page](https://github.com/Spandan7724/Enigma/releases).

## Acknowledgements

-   [muesli/termenv](https://github.com/muesli/termenv) for the color in the CLI
-   [spf13/cobra](https://github.com/spf13/cobra) for the CLI framework