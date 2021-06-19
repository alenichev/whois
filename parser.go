// Package whois provides simple whois protocol (rfc3912) implementation.
package whois

import (
	"strings"
	"time"
)

// Reply contains fields from WHOIS server response.
type Reply struct {
	Domain       string
	Organisation string
	Person       string
	Created      string
	Paid         time.Time
	Source       string
}

// GetReply returns WhoisReply for domain.
func GetReply(domain string) (Reply, error) {
	var (
		out string
		err error
		r   Reply
	)

	out, err = QueryAll(domain)
	if err != nil {
		return r, err
	}

	buffer := strings.Split(out, "\n")
	for _, line := range buffer {
		field := strings.SplitN(line, ":", 2)
		if len(field) == 1 {
			continue
		}

		value := strings.TrimSpace(field[1])
		switch field[0] {
		case "domain":
			r.Domain = value
		case "organisation":
			r.Organisation = value
		case "person":
			r.Person = value
		case "created":
			r.Created = value
		case "paid":
			r.Paid, _ = time.Parse(time.RFC3339, value)
		case "source":
			r.Source = value
		default:
			continue
		}
	}

	return r, nil
}
