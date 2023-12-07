<div align="center">
<div>
  
```
  __  __                   _ _ _   _     
 |  \/  | ___  _ __   ___ | (_) |_| |__  
 | |\/| |/ _ \| '_ \ / _ \| | | __| '_ \ 
 | |  | | (_) | | | | (_) | | | |_| | | |
 |_|  |_|\___/|_| |_|\___/|_|_|\__|_| |_|
                                         
```
</div>

# Monolith

![Screenshot](https://github.com/samuelchristlie/monolith/assets/127090887/09a34747-f1a1-4ee0-8e77-994fa38af524)


## Stateless password manager implementation in Go

![github stars badge](https://badgen.net/github/stars/samuelchristlie/monolith?icon=github)
![github forks badge](https://badgen.net/github/forks/samuelchristlie/monolith?icon=github)
![github issues badge](https://badgen.net/github/open-issues/samuelchristlie/monolith?icon=github)

![github last commit badge](https://badgen.net/github/last-commit/samuelchristlie/monolith?icon=github)
</div>

**Monolith** is a stateless password manager written in Go. It allows you to generate strong and secure password on the fly, without having to store it. Monolith allows you to have different passwords for each website/service by just remembering your master password. Monolith takes 3 main inputs: website, login, and master password. In addition, it also supports additional options such as custom password length and character sets. 

Monolith currently uses **PBKDF2** and the more memory-intensive **Scrypt** (pronounced "ess crypt") key derivation function to generate passwords, picked based on [NIST](https://pages.nist.gov/800-63-3/sp800-63-3.html) and [OWASP](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html) guidelines. Scrypt is used by default, but you can change this setting in the code. 

The parameter used in **Monolith** is as follows

**PBKDF2**
- SHA-256 hash algorithm
- 600,000 iterations

**Scrypt**
- 2^14 (16 MiB) CPU/memory cost parameter
- 8 (1024 bytes) block size
- 5 degree of parallelism



## Features 💪
- **No Data Stored**: Monolith doesn't store any of your data, ensuring the security of your passwords.
- **No Internet Connection Needed**: You don't need an internet connection to use Monolith, since it doesn't require synchronization.
- **Customize Your Password**: Monolith supports custom password length and allows you to choose character sets.

## Table of Contents 📝
1. 💻 [Installation](#installation)
2. ▶ [Usage](#usage)
3. 🙏 [Acknowledgements](#acknowledgements)

<a name="installation"/>

## 💻 Installation
### Building
Clone this repository
```
git clone https://github.com/samuelchristlie/monolith
cd monolith
```
Build project using Go

#### Windows
```
go build -ldflags="-s -w -H=windowsgui -extldflags=-static" monolith.go
```
#### Linux and MacOS
```
go build -ldflags="-s -w" monolith.go
```

The compiled build should be located in the current folder.

### Prebuilt Binary
Prebuilt binary release is planned.

<a name="usage"/>

## ▶ Usage
To run **Monolith**, simply open the executable. CLI version is planned.

<a name="acknowledgements"/>

## 🙏 Acknowledgements
Thanks to Patrick Gillespie for creating the ASCII text art tool used in this project
https://patorjk.com/software/taag/

Thanks to Fyne for providing the UI toolkit used in this project
https://github.com/fyne-io/fyne
