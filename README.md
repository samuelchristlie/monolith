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

**Monolith** is a stateless password manager written in Go. It allows you to generate strong and secure password on the fly, without having to store it

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
