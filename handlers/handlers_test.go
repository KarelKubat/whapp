package handlers

import "testing"

func TestTypeString(t *testing.T) {
	for tp := firstUnused + 1; tp < lastUnused; tp++ {
		t.Log(int(tp))
		t.Log(tp.String())
	}
}
