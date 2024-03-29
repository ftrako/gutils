package gutils

import (
	"testing"
)

func TestIPStr2Int(t *testing.T) {
	ip := "192.168.0.2"
	n := IPStr2Int(ip)
	if n == 0 {
		t.FailNow()
	}
}

func TestIPInt2Str(t *testing.T) {
	n := 3233235523
	ip := IPInt2Str(n)
	if ip == "" {
		t.FailNow()
	}
}

func TestIsLanIP(t *testing.T) {
	ip := "172.31.0.1"
	if !IsLanIP(ip) {
		t.FailNow()
	}
}
