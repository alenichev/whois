package whois

import (
	"testing"
)

func TestExampleComParser(t *testing.T) {
	r, err := GetReply(exampleCom)
	if err != nil {
		t.Errorf("GetReply() error: %s", err)
	}
	if r.Domain != "EXAMPLE.COM" {
		t.Errorf("wrong domain field: except \"EXAMPLE.COM\", got %s",
			r.Domain)
	}
	if r.Source != "IANA" {
		t.Errorf("wrong source field: except \"TCI\", got %s",
			r.Source)
	}
}

func TestRockersSuParser(t *testing.T) {
	r, err := GetReply(rockersSu)
	if err != nil {
		t.Errorf("GetReply() error: %s", err)
	}
	if r.Domain != "ROCKERS.SU" {
		t.Errorf("wrong domain field: except \"ROCKERS.SU\", got %s",
			r.Domain)
	}
	if r.Source != "TCI" {
		t.Errorf("wrong source field: except \"TCI\", got %s",
			r.Source)
	}
}
