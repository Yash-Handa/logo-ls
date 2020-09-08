![OpenSource](https://img.shields.io/badge/Open%20Source-000000?style=for-the-badge&logo=github)
![go](https://img.shields.io/badge/-Written%20In%20Go-00add8?style=for-the-badge&logo=Go&logoColor=ffffff)
![cli](https://img.shields.io/badge/-Build%20for%20CLI-000000?style=for-the-badge&logo=Powershell&logoColor=ffffff)

# logo-ls

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Yash-Handa/logo-ls?style=flat-square)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/Yash-Handa/logo-ls?sort=semver&style=flat-square)
![PRs](https://img.shields.io/badge/PRs-welcome-56cc14?style=flat-square)
[![HitCount](http://hits.dwyl.com/{Yash-Handa}/{logo-ls}.svg)](http://hits.dwyl.com/{Yash-Handa}/{logo-ls})

modern ls command with vscode like file logos. Written in Golang

Command and Arguments supported are listed in [HELP.md](/HELP.md)

## Check the downloaded Resource

receive public key of the signing party

```cmd
$ gpg2 --keyid-format long --keyserver keyserver.ubuntu.com --recv-keys 0x28182066bcacccb2

gpg: key 28182066BCACCCB2: "Yash Handa (logo-ls) <yashhanda7@yahoo.com>" not changed
gpg: Total number processed: 1
gpg:              unchanged: 1
```

check the signature on logo-ls_SHA512sums.txt

```cmd
$ gpg2 --keyid-format long --verify logo-ls_SHA512sums.txt.sig logo-ls_SHA512sums.txt

gpg: Signature made Tue 08 Sep 2020 10:21:52 PM IST
gpg:                using RSA key D9498B225223344C0205FDF528182066BCACCCB2
gpg: Good signature from "Yash Handa (logo-ls) <yashhanda7@yahoo.com>" [ultimate]
```

A **Good** signature means that the checked file was definitely signed by the owner of the keyfile stated (if they didn’t match, the signature would be reported as **BAD**)

Now use `logo-ls_SHA512sums.txt` file to verify the authenticity of the resource downloaded

```cmd
$ sha512sum -c logo-ls_SHA512sums.txt 2>&1 | grep OK

logo-ls_Linux_x86_64.tar.gz: OK
logo-ls_Linux_arm64.tar.gz: OK
logo-ls_Darwin_x86_64.tar.gz: OK
logo-ls_Linux_i386.tar.gz: OK
logo-ls_Linux_armv6.tar.gz: OK
```

Note: The downloaded resources to verify and `logo-ls_SHA512sums.txt` should be in the same directory

## Example Commands

<div>
  <span align="center">
  <img alt="logo-ls" title="logo-ls" src="/.github/images/ls.png">
    </span>
</div><br>

<div>
  <span align="center">
  <img alt="logo-ls -aD" title="logo-ls -aD" src="/.github/images/ls-aD.png">
    </span>
</div><br>

<div>
  <span align="center">
  <img alt="logo-ls -alhD" title="logo-ls -alhD" src="/.github/images/ls-alhD.png">
    </span>
</div><br>

<div>
  <span align="center">
  <img alt="logo-ls -1Dhs" title="logo-ls -1Dhs" src="/.github/images/ls-1Dhs.png">
    </span>
</div><br>

<div>
  <span align="center">
  <img alt="logo-ls -RD" title="logo-ls -RD" src="/.github/images/ls-RD.png">
    </span>
</div><br>

<div>
  <span align="center">
  <img alt="logo-ls -a" title="logo-ls -a" src="/.github/images/ls-a.png">
    </span>
</div><br>

<div>
  <span align="center">
  <img alt="logo-ls -oahd -T Kitchen" title="logo-ls -oahd -T Kitchen" src="/.github/images/ls-oahD.png">
    </span>
</div><br>

## Icon Set

<div>
  <span align="center">
  <img alt="icons" title="Icon Set" src="/.github/images/icons.png">
    </span>
</div><br>
