package whois

import (
	"strings"
	"testing"
)

const (
	exampleCom = "example.com"
	rockersSu  = "rockers.su"
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
	out, err := Query(IANASever, exampleCom)
	if err != nil {
		t.Errorf("Query() error: %s", err)
	}

	if strings.Compare(exampleComOut, out) != 0 {
		t.Errorf("output differs: %v %v", exampleComOut, out)
	}
}

func TestRockersSu(t *testing.T) {
	out, err := QueryAll(rockersSu)
	if err != nil {
		t.Errorf("QueryAll() error: %s", err)
	}

	if strings.Contains(out, "TCI") != true {
		t.Errorf("output does not contains TCI: %v", out)
	}
}
