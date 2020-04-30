# HIBP_PasswordList_Slimmer
![Banner](https://zupimages.net/up/20/18/xfr4.png)![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg) ![made-with-go](https://img.shields.io/badge/made%20with-go-blue)  ![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)

Tool in Go designed to reduce the size of Have I Been Pwned password lists and speed up searching.

Strongly inspired by [pwnedpw_passfilt](https://github.com/darrellenns/pwnedpw_passfilt) but rewritten in Go.

On my side (Intel I7 8565U) the conversion takes 54 minutes for the full NTLM hash list and the size increases from about 19GB to about 10GB.

## Usage
Clone the repo or download one of the binaries in the releases.

```bash
go run .\PwnDBConverter.go {InputHashFile} {HashType} {OutputFile}
.\PwnDBConverter.exe {InputHashFile} {HashType} {OutputFile}
```
![Usage](https://zupimages.net/up/20/18/op6u.png)