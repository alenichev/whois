# whois

[![Go Reference](https://pkg.go.dev/badge/github.com/alenichev/whois.svg)](https://pkg.go.dev/github.com/alenichev/whois)
[![Test Status](https://github.com/alenichev/whois/workflows/Go/badge.svg)](https://github.com/alenichev/whois/actions)


Package whois provides simple whois protocol (rfc3912) implementation.

## Importing
```Go
import "github.com/alenichev/whois"
```

## Usage examples
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
