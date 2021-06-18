package whois

import (
	"strings"
	"testing"
)

const (
	ianaWhois     = "whois.iana.org"
	exampleDomain = "example.com"
)

var exampleComOut = `% IANA WHOIS server
% for more information on IANA, visit http://www.iana.org
% This query returned 1 object

domain:       EXAMPLE.COM

organisation: Internet Assigned Numbers Authority

created:      1992-01-01
source:       IANA

`

func TestExampleCom(t *testing.T) {
	out, err := MakeWhoisQuery(ianaWhois, exampleDomain)
	if err != nil {
		t.Errorf("MakeWhoisQuery() error: %s", err)
	}

	if strings.Compare(exampleComOut, out) != 0 {
		t.Errorf("output differs: %v %v", exampleComOut, out)
	}
}
