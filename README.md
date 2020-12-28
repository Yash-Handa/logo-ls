![OpenSource](https://img.shields.io/badge/Open%20Source-000000?style=for-the-badge&logo=github)
![go](https://img.shields.io/badge/-Written%20In%20Go-00add8?style=for-the-badge&logo=Go&logoColor=ffffff)
![cli](https://img.shields.io/badge/-Build%20for%20CLI-000000?style=for-the-badge&logo=Powershell&logoColor=ffffff)

# logo-ls

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Yash-Handa/logo-ls?style=flat-square)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/Yash-Handa/logo-ls?sort=semver&style=flat-square)
![PRs](https://img.shields.io/badge/PRs-welcome-56cc14?style=flat-square)
![AUR](https://img.shields.io/aur/version/logo-ls?style=flat-square&logo=Arch-Linux&labelColor=abcdef&label=AUR)
![deb](https://img.shields.io/badge/-Deb%20Package-A81D33?style=flat-square&logo=Debian&link=https://github.com/Yash-Handa/logo-ls/releases/)
![rpm](https://img.shields.io/badge/-RPM%20Package-EE0000?style=flat-square&logo=Red-Hat&link=https://github.com/Yash-Handa/logo-ls/releases/)
![linux](https://img.shields.io/badge/-Linux%20Binary-FCC624?style=flat-square&logo=Linux&logoColor=000000&link=https://github.com/Yash-Handa/logo-ls/releases/)
![apple](https://img.shields.io/badge/-Darwin%20Binary-999999?style=flat-square&logo=Apple&logoColor=ffffff&link=https://github.com/Yash-Handa/logo-ls/releases/)
![Go](https://github.com/Yash-Handa/logo-ls/workflows/Go/badge.svg)

modern ls command with beautiful Icons and Git Integrations . Written in Golang

<div>
  <span align="center">
  <img alt="logo-ls" title="logo-ls" src="/.github/images/ls.png">
    </span>
</div><br>

Command and Arguments supported are listed in [HELP.md](/HELP.md)

## Table of contents

- [Features](#features)
- [Usage](#usage)
  - [Flags](#flags)
    - `-1`
    - `-a` (or) `--all`
    - `-A` (or) `--almost-all`
    - `-D` (or) `--git-status`
    - `-l`, `-o`, `-g`, `-G`
    - `-R` (or) `--recursive`
    - `-?` (or) `--help`
    - Sorting:
      - `-t`, `-S`, `-X`, `-U`, `-v` and `-r` (`--reverse`)
    - `-T` (or) `--time-style=value`
    - `-i` (or) `--disable-icon`
    - `-c` (or) `--disable-color`
  - [Combination of flags](#combination-of-flags)
  - [Multiple Files and Directories](#multiple-files-and-directories)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [Debian (.deb package)](#debian-deb-package)
  - [Arch Linux (AUR)](#arch-linux)
  - [Red Hat (.rpm package)](#red-hat-rpm-package)
  - [MacOS (Darwin)](#macos-darwin)
  - [Linux](#linux)
  - [go get](#go-get)
  - [Build from Source](#build-from-source)
  - [Check the downloaded Resource](#check-the-downloaded-resource) 
- [Recommended configurations](#recommended-configurations)
- [Updating](#updating)
- [Icon Set](#icon-set)
- [Contributing](#contributing)
- [License](#license)

## Features

[:arrow_up: TOC](#table-of-contents)

This project is build to add esthetics to ls(coreutiles) command

- Over 250+ icons :sunglasses:
- Supporting 600+ files, extensions and directories
- 16 Million, true colors supported
- Git Status Integration
- :racehorse: Near native speed. (thanks to Golang)
- Language agnostic binaries

This project is highly inspired by [ls(coreutiles)](https://www.gnu.org/software/coreutils/manual/html_node/ls-invocation.html#ls-invocation) and [color ls](https://github.com/athityakumar/colorls). The project tries to find a happy path between speed and aesthetics.

## Usage

[:arrow_up: TOC](#table-of-contents)

All supported flags can be found by using help flag `$ logo-ls -?`. The same has been provided as [HELP.md](/HELP.md).

The project also has its manpage which can be accessed by `man logo-ls`

### Flags

Almost all flags are same as that of the classic ls command and behave similarly. The project can be used as a drop-in replacement for the ls(coreutiles)

- With `-1`: List one entry per line

<div>
  <span align="center">
  <img alt="logo-ls -1" title="logo-ls -1" src="/.github/images/ls-1.png">
    </span>
</div><br>

- With `-a` (or) `--all` : Does not ignore entries starting with '.'

<div>
  <span align="center">
  <img alt="logo-ls -a" title="logo-ls -a" src="/.github/images/ls-a.png">
    </span>
</div><br>

- With `-A` (or) `--almost-all` : Does not ignore entries starting with '.', except `./` and `../`

<div>
  <span align="center">
  <img alt="logo-ls -A" title="logo-ls -A" src="/.github/images/ls-A.png">
    </span>
</div><br>

- With `-D` (or) `--git-status`: Add Git Status to the listed Files and Directory
  - *Note*: As much I would love to make this the default behavior of the command but showing git status is an intensive task and may slow (a tiny bit) the command itself. If you want you can make alias to the command with `-D` applied to it.

<div>
  <span align="center">
  <img alt="logo-ls -D" title="logo-ls -D" src="/.github/images/ls-D.png">
    </span>
</div><br>

- With `-l`: Shows in long listing format
  other similar commands are:
  - `-o`: like `-l`, but do not list group information
  - `-g`: like `-l`, but do not list owner
  - `-G` (or) `--no-group`: in a long listing, don't print group names

<div>
  <span align="center">
  <img alt="logo-ls -l" title="logo-ls -l" src="/.github/images/ls-l.png">
    </span>
</div><br>

- With `-R` (or) `--recursive`: list subdirectories recursively

<div>
  <span align="center">
  <img alt="logo-ls -R" title="logo-ls -R" src="/.github/images/ls-R.png">
    </span>
</div><br>

- With `-?` (or) `--help`: print the help message
  - similar message can be found at [HELP.md](/HELP.md).

<div>
  <span align="center">
  <img alt="logo-ls -?" title="logo-ls -?" src="/.github/images/ls-h.png">
    </span>
</div><br>

- **Sorting**: There are many sorting flags available [default is *alphabetic order*]
  - With `-t`: sort by modification time, newest first
  - With `-S`: sort by file size, largest first
  - With `-X`: sort alphabetically by entry extension
  - With `-U`: do not sort; list entries in directory order
  - With `-v`: natural sort of (version) numbers within text
  - With `-r` (or) `-reverse`: reverse order while sorting

<div>
  <span align="center">
  <img alt="logo-ls -X" title="logo-ls -X" src="/.github/images/ls-X.png">
    </span>
</div><br>

- With `-T` (or) `--time-style=value`: set time/date format in long formats (`-l`, `-o`, `-g`). There are many options to chose from all are listed in [HELP.md](/HELP.md).

<div>
  <span align="center">
  <img alt="logo-ls -T" title="logo-ls -T" src="/.github/images/ls-T.png">
    </span>
</div><br>

- With `-i` (or) `--disable-icon`: don't print icons of the files

<div>
  <span align="center">
  <img alt="logo-ls -i" title="logo-ls -i" src="/.github/images/ls-i.png">
    </span>
</div><br>

- With `-c` (or) `--disable-color`: don't color icons, filenames and git status
  - *Note*: use a combination of `-ic` to print output to a file `$ logo-ls -ic > t.txt`

<div>
  <span align="center">
  <img alt="logo-ls -c" title="logo-ls -c" src="/.github/images/ls-c.png">
    </span>
</div><br>
  

For all available commands see manpage or [HELP.md](/HELP.md)

### Combination of flags

This project uses [getopt](https://github.com/pborman/getopt) which is a golang variant of the classic getopt utility used in ls(coreutiles). Thus any combination of flags are possible and can be used.

<div>
  <span align="center">
  <img alt="logo-ls with flag combination" title="logo-ls with flag combination" src="/.github/images/ls-AshDt.png">
    </span>
</div><br>

### Multiple Files and Directories

You can provide multiple files and directories as command argument [default to PWD] and all will be displayed accordingly.

<div>
  <span align="center">
  <img alt="logo-ls  with multiple files" title="logo-ls with multiple files" src="/.github/images/ls multi.png">
    </span>
</div><br>

## Prerequisites

[:arrow_up: TOC](#table-of-contents)

For proper working of `logo-ls` the following should be set

- **UTF-8**: The terminal should be UTF-8 encoded (Can display Unicode-Code Point)
- **True Color** Support: The terminal can display color (16 Million Colors). for more information of True Color and supported Terminals [see here](https://gist.github.com/XVilka/8346728)
- **Nerd Fonts**: Nerd fonts are required to display Icons on screen. Basically Nerd Fonts patches your current font i.e., the last few unicode points (approx 2,824 out of 143,859) in the font are replaced with nerdy icons :nerd_face:. The complete patching process is given [here](https://github.com/ryanoasis/nerd-fonts)

## Installation

[:arrow_up: TOC](#table-of-contents)

Installation is very easy and straight forward with many options to choose from.

As of now almost all installation methods require downloading resources from the Github Release page of the project, so to be a bit more secure consider **Checking the Signature** of `logo-ls_SHA512sums.txt` text file and then use this file to check weather the resource have been tampered. This complete process is explained in more detail in [Check the downloaded Resource](#check-the-downloaded-resource) section below. (This is a recommended step and not at all required)

### Debian (.deb package)

If you are on Debian or any other Debian based distribution then installation is simple.

#### Step 1

Download the `.deb` package from [Github Release Page](https://github.com/Yash-Handa/logo-ls/releases). Available OS_Architectures include: `i386`, `amd64`, `arm64` and `armV6`. Check your downloaded resource(s) if you like.

#### Step 2

Use the `dpkg -i path/to/downloaded/resource/` to install the binary and the manpage

### Red Hat (.rpm package)

If you are on Red Hat or any other Red Hat based distribution (like fedora) then installation is simple.

#### Step 1

Download the `.rpm` package from [Github Release Page](https://github.com/Yash-Handa/logo-ls/releases). Available OS_Architectures include: `i386`, `amd64`, `arm64` and `armV6`. Check your downloaded resource(s) if you like.

#### Step 2

Use the `rpm -i path/to/downloaded/resource/` to install the binary and the manpage

### Arch linux

Simply run `yay -S logo-ls`. (Or use the AUR helper of your choice)<br><br>
Alternatively you can clone the PKGBUILD and build the package manually:
```bash
git clone https://aur.archlinux.org/logo-ls.git
cd logo-ls
makepkg -si
```

### MacOS (Darwin)

To install `logo-ls` on darwin you have to download the binary. Support for Homebrew will come soon

#### Step 1

Download `logo-ls_Darwin_x86_64.tar.gz` from [Github Release Page](https://github.com/Yash-Handa/logo-ls/releases).

#### Step 2

Extract a gzipped archive in the current directory.

```cmd
$ tar -xzf logo-ls_Darwin_x86_64.tar.gz
```

This will produce `logo-ls_Darwin_x86_64` directory in the current directory with the following files: `HELP.md`, `LICENSE`, `logo-ls` and `logo-ls.1.gz`

#### Step 3

Install the binary `logo-ls` by placing it in `/usr/local/bin`

```cmd
$ cd logo-ls_Darwin_x86_64
$ sudo cp logo-ls /usr/local/bin
```

#### Step 4

If you want the man page of `logo-ls` place `logo-ls.1.gz` in `/usr/local/share/man/man1/`

```cmd
$ sudo cp logo-ls.1.gz /usr/local/share/man/man1/
```

### Linux

To install `logo-ls` on any other Linux Distribution you have to download the binary.

#### Step 1

Download the `.tar.gz` archive from [Github Release Page](https://github.com/Yash-Handa/logo-ls/releases). Available OS_Arch include: `i386`, `x86_64` (`amd64`), `arm64` and `armV6`. Check your downloaded resource(s) if you like.

#### Step 2

Extract a gzipped archive in the current directory.

```cmd
$ tar -xzf logo-ls_Linux_[ARCH].tar.gz
```

This will produce `logo-ls_Linux_[ARCH]` directory in the current directory with the following files: `HELP.md`, `LICENSE`, `logo-ls` and `logo-ls.1.gz`

#### Step 3

Install the binary `logo-ls` by placing it in `/usr/local/bin`

```cmd
$ cd logo-ls_Linux_[ARCH]
$ sudo cp logo-ls /usr/local/bin
```

#### Step 4

If you want the man page of `logo-ls` place `logo-ls.1.gz` in `/usr/share/man/man1/`

```cmd
$ sudo cp logo-ls.1.gz /usr/share/man/man1/
```

### go get

If you have [Golang](https://golang.org/) installed on your system then the `logo-ls` can be downloaded and installed using the the `go` cli tool provided by the language

#### Step 1

The `go` cli installs the binary in the directory specified by `$GOBIN` env variable [default: `~/go/bin/`]. In order to call `logo-ls` from anywhere in user-space, this directory which holds the binary must be present in `$PATH` env variable.

```cmd
$ echo $PATH
```

#### Step 2

Now simply install the binary using `go get` in home directory

```cmd
$ cd ~
$ go get github.com/Yash-Handa/logo-ls
```

### Build from Source

It is quite simple to build the command from source. You should have [Golang](https://golang.org/) installed on your system.

#### Step 1

Clone this repository.

```cmd
$ git clone https://github.com/Yash-Handa/logo-ls.git
```

#### Step 2

`cd` into the directory `logo-ls`. Then run `go mod tidy` to download the dependencies and after that `go build` to build your system specific binary

```cmd
$ cd logo-ls
$ go mod tidy
$ go build
```

the `go build` command will produce `logo-ls` executable binary in the current directory.

#### Step 3

Place the `logo-ls` executable in a directory reachable from `$PATH` env variable.

```cmd
$ sudo cp logo-ls /usr/local/bin
```

#### Step 4

If you want the man page of `logo-ls` place `logo-ls.1.gz` in `/usr/share/man/man1/`

```cmd
$ sudo cp logo-ls.1.gz /usr/share/man/man1/
```

### Check the downloaded Resource

After downloading the Resource(s) from the [Github Release Page](https://github.com/Yash-Handa/logo-ls/releases) follow the below steps to check its authenticity.

#### Step 1

Download `logo-ls_SHA512sums.txt.sig` and `logo-ls_SHA512sums.txt` from [Github Release Page](https://github.com/Yash-Handa/logo-ls/releases) and place them in the same directory as of the downloaded resource(s)

#### Step 2

Receive the public key of the signing party from `keyserver.ubuntu.com` server. This will add the developer's public key to your `keyring`.

```cmd
$ gpg2 --keyid-format long --keyserver keyserver.ubuntu.com --recv-keys 0x28182066bcacccb2

gpg: key 28182066BCACCCB2: "Yash Handa (logo-ls) <yashhanda7@yahoo.com>" not changed
gpg: Total number processed: 1
gpg:              unchanged: 1
```

#### Step 3

check the signature on logo-ls_SHA512sums.txt

```cmd
$ gpg2 --keyid-format long --verify logo-ls_SHA512sums.txt.sig logo-ls_SHA512sums.txt

gpg: Signature made Tue 08 Sep 2020 10:21:52 PM IST
gpg:                using RSA key D9498B225223344C0205FDF528182066BCACCCB2
gpg: Good signature from "Yash Handa (logo-ls) <yashhanda7@yahoo.com>" [ultimate]
```

A **Good** signature means that the checked file was definitely signed by the owner of the keyfile stated (if they didn’t match, the signature would be reported as **BAD**)

#### Step 4

Use `logo-ls_SHA512sums.txt` file to verify the authenticity of the downloaded resource(s)

```cmd
$ sha512sum -c logo-ls_SHA512sums.txt 2>&1 | grep OK

logo-ls_Linux_x86_64.tar.gz: OK
logo-ls_Linux_arm64.tar.gz: OK
logo-ls_Darwin_x86_64.tar.gz: OK
logo-ls_Linux_i386.tar.gz: OK
logo-ls_Linux_armv6.tar.gz: OK
```

## Recommended configurations

[:arrow_up: TOC](#table-of-contents)

To add some short command (say, `ils` or `ls`) with some flag options by default, add this to your shell configuration file (`~/.bashrc`, `~/.zshrc`, etc.) :

```bash
alias ils='logo-ls'
alias ila='logo-ls -A'
alias ill='logo-ls -al'
# equivalents with Git Status on by Default
alias ilsg='logo-ls -D'
alias ilag='logo-ls -AD'
alias illg='logo-ls -alD'
```

## Updating

[:arrow_up: TOC](#table-of-contents)

As of now Updating is exactly similar as installation for [Debian (.deb package)](#debian-deb-package), [Red Hat (.rpm package)](#red-hat-rpm-package), [MacOS (Darwin)](#macos-darwin) and [Linux](#linux)

## Icon Set

[:arrow_up: TOC](#table-of-contents)

<div>
  <span align="center">
  <img alt="icons" title="Icon Set" src="/.github/images/icons.png">
    </span>
</div><br>

## Contributing

[:arrow_up: TOC](#table-of-contents)

The project is always open to contributions. Especially for:

- More Icons
- Better ways to distribute software
- Translations to other Languages (linguistic)

## License

[:arrow_up: TOC](#table-of-contents)

The project is licensed under **MIT**. The Licence is available [here](/LICENSE).
