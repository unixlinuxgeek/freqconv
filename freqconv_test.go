package freqconv

import (
	"github.com/unixlinuxgeek/floatgen"
	"testing"
)

// Hertz To Kilohertz converter test
func TestHzToKHz(t *testing.T) {
	r := Hertz(floatgen.GenRan(1, 9)) // generate random float
	ok := Kilohertz(r / 1000)         // correct option
	chk := HzToKHz(r)                 // need to check

	if ok == chk {
		t.Logf("Test TestHzToKHz Passed: %v equal %v\n", ok, chk)
	} else {
		t.Fatalf("Test TestHzToKHz Failed: %v not equal %v\n", ok, chk)
	}
}

// Hertz To Megahertz converter test
func TestHzToMHz(t *testing.T) {
	r := Hertz(floatgen.GenRan(1, 9)) // generate random float
	ok := Megahertz(r / 1e+6)         // correct option
	chk := HzToMHz(r)                 // need to checking

	if ok == chk {
		t.Logf("Test TestHzToMHz Passed: %v equal %v\n", ok, chk)
	} else {
		t.Fatalf("Test TestHzToMHz Failed: %v not equal %v\n", ok, chk)
	}
}

// // Hertz To Gigahertz converter test
func TestHzToGHz(t *testing.T) {
	r := Hertz(floatgen.GenRan(1, 9)) // generate random float
	ok := Gigahertz(r / 1e+9)         // correct option
	chk := HzToGHz(r)

	if ok == chk {
		t.Logf("Test TestHzToGHz Passed: %v equal %v\n", ok, chk)
	} else {
		t.Fatalf("Test TestHzToGHz Failed: %v not equal %v\n", ok, chk)
	}
}
