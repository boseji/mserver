package mserver

import (
	"bytes"
	"testing"
)

func TestBufToHexString(t *testing.T) {
	in := "\x02\x05\x08\x00"
	b := bytes.NewBufferString(in)
	s := BufToHexString(b)
	result := "02050800"
	if s != result {
		t.Fatalf("BufToHexString(%s) = %s want %s", in, s, result)
	}
}
