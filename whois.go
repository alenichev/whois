// Package whois provides simple whois protocol (rfc3912) implementation.
package whois

import (
	"context"
	"io/ioutil"
	"net"
	"strings"
	"time"
)

var (
	WhoisReadTimeout  time.Duration = 5
	WhoisWriteTimeout time.Duration = 5
	WhoisTotalTimeout time.Duration = 60
)

// MakeWhoisQuery makes query using whois protocol
// for domain on host and return response.
func MakeWhoisQuery(host, domain string) (string, error) {
	var (
		d   net.Dialer
		out string
		err error
	)

	ctx, cancel := context.WithTimeout(context.Background(),
		WhoisTotalTimeout*time.Second)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", host+":43")
	if err != nil {
		return out, err
	}
	defer conn.Close()

	err = conn.SetWriteDeadline(time.Now().Add(WhoisWriteTimeout *
		time.Second))
	if err != nil {
		return out, err
	}

	if _, err := conn.Write([]byte(domain + "\r\n")); err != nil {
		return out, err
	}

	err = conn.SetReadDeadline(time.Now().Add(WhoisReadTimeout *
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
func MakeWhoisQueryAll(domain string) (string, error) {
	var (
		out string
		err error
	)

	out, err = MakeWhoisQuery("whois.iana.org", domain)
	if err != nil {
		return out, err
	}

	if strings.Contains(out, "refer") {
		referString := strings.Split(out, "\n")[4]
		refer := strings.Split(referString, " ")
		referHost := refer[len(refer)-1]

		out, err = MakeWhoisQuery(referHost, domain)
		if err != nil {
			return out, err
		}
	}
	return out, err
}
