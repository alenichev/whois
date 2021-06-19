# whois

[![whois release (latest SemVer)](https://img.shields.io/github/v/release/alenichev/whois?sort=semver)](https://github.com/alenichev/whois/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/alenichev/whois.svg)](https://pkg.go.dev/github.com/alenichev/whois)
[![tests](https://github.com/alenichev/whois/actions/workflows/tests.yml/badge.svg)](https://github.com/alenichev/whois/actions/workflows/tests.yml)

Package whois provides simple whois protocol ([RFC 3912](https://tools.ietf.org/html/rfc3912)) implementation.

## Installation
```
go get github.com/alenichev/whois
```
or import as package:
```Go
import "github.com/alenichev/whois"
```
an run `go get` without arguments.

## Usage examples
```Go
import "github.com/alenichev/whois"
```
ask "whois.iana.org" about "example.com":
```Go
output, err := whois.MakeWhoisQuery("whois.iana.org", "example.com")
if err != nil {
    // error handling
}
```
or (final result, using refer field from IANA's response, if needed):
```Go
output, err := whois.MakeWhoisQueryAll("example.com")
if err != nil {
    // error handling
}
```
