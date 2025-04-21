package utils

import (
	"testing"
)

func TestGenFileName(t *testing.T) {
	filename := GenFileName("/a/b/c/d/f/e/g.xxx")
	t.Log(filename)
	if filename != "eg.xxx" {
		t.Error("GenFileName error")
	}
}
