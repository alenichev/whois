// Package whois provides simple whois protocol (rfc3912) implementation.
package whois

import (
	"context"
	"io/ioutil"
	"net"
	"strings"
	"time"
)

// ReadTimeout sets timeout for reading response.
// Default: 5
var ReadTimeout time.Duration = 5

// WriteTimeout sets timeout for writing query.
// Default: 5
var WriteTimeout time.Duration = 5

// TotalTimeout sets total timeout.
// Default: 60
var TotalTimeout time.Duration = 60

// IANA whois service address.
const IANASever = "whois.iana.org"

// WHOIS protocol port.
const PortNumber = "43"

// MakeWhoisQuery makes query using whois protocol
// for domain on host and return response.
func Query(host, domain string) (string, error) {
	var (
		d   net.Dialer
		out string
		err error
	)

	ctx, cancel := context.WithTimeout(context.Background(),
		TotalTimeout*time.Second)
	defer cancel()

	hostport := net.JoinHostPort(host, PortNumber)
	conn, err := d.DialContext(ctx, "tcp", hostport)
	if err != nil {
		return out, err
	}
	defer conn.Close()

	err = conn.SetWriteDeadline(time.Now().Add(WriteTimeout *
		time.Second))
	if err != nil {
		return out, err
	}

	if _, err := conn.Write([]byte(domain + "\r\n")); err != nil {
		return out, err
	}

	err = conn.SetReadDeadline(time.Now().Add(ReadTimeout *
		time.Second))
	if err != nil {
		return out, err
	}

	output, err := ioutil.ReadAll(conn)
	if err != nil {
		return out, err
	}

	out = string(output)
	return out, nil
}

// MakeWhoisQueryAll makes query using whois protocol
// to "whois.iana.org" for domain and follows refer field
// in response, if any.
func QueryAll(domain string) (string, error) {
	var (
		out string
		err error
	)

	out, err = Query(IANASever, domain)
	if err != nil {
		return out, err
	}

	if strings.Contains(out, "refer") {
		referString := strings.Split(out, "\n")[4]
		refer := strings.Split(referString, " ")
		referHost := refer[len(refer)-1]

		out, err = Query(referHost, domain)
		if err != nil {
			return out, err
		}
	}
	return out, nil
}
