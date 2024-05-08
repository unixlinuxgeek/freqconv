package freqconv

import (
	"github.com/unixlinuxgeek/floatgen"
	"testing"
)

// Hertz To Kilohertz converter test
func TestHzToKHz(t *testing.T) {
	r := Hertz(floatgen.GenRan(1, 9)) // generate random float
	ok := Kilohertz(r / 1000)         // correct option
	chk := HzToKHz(r)                 // need to checking

	if ok == chk {
		t.Logf("Test TestHzToKHz Passed: %v equal %v\n", ok, chk)
	} else {
		t.Errorf("Test TestHzToKHz Failed: %v not equal %v\n", ok, chk)
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
		t.Errorf("Test TestHzToMHz Failed: %v not equal %v\n", ok, chk)
	}
}

// // Hertz To Gigahertz converter test
func TestHzToGHz(t *testing.T) {
	r := Hertz(floatgen.GenRan(1, 9)) // generate random float
	ok := Gigahertz(r / 1e+9)         // correct option
	chk := HzToGHz(r)                 // need to checking

	if ok == chk {
		t.Logf("Test TestHzToGHz Passed: %v equal %v\n", ok, chk)
	} else {
		t.Errorf("Test TestHzToGHz Failed: %v not equal %v\n", ok, chk)
	}
}

// Kilohertz To Hertz converter test
func TestKHzToHz(t *testing.T) {
	r := Kilohertz(floatgen.GenRan(1, 9)) // generate random float
	ok := Hertz(r * 1000)                 // correct option
	chk := KHzToHz(r)                     // need to checking

	if ok == chk {
		t.Logf("Test TestKHzToHz Passed: %v equal %v\n", ok, chk)
	} else {
		t.Errorf("Test TestKHzToHz Failed: %v not equal %v\n", ok, chk)
	}
}

// Kilohertz To Megahertz converter test
func TestKHzToMHz(t *testing.T) {
	r := Kilohertz(floatgen.GenRan(1, 9)) // generate random float
	ok := Megahertz(r / 1000)             // correct option
	chk := KHzToMhz(r)                    // need to checking

	if ok == chk {
		t.Logf("Test TestKHzToMHz Passed: %v equal %v\n", ok, chk)
	} else {
		t.Errorf("Test TestKHzToMHz Failed: %v not equal %v\n", ok, chk)
	}
}

// Kilohertz To Gigahertz converter test
func TestKhzToGhz(t *testing.T) {
	r := Kilohertz(floatgen.GenRan(1, 9)) // generate random float
	ok := Gigahertz(r / 1e+6)             // correct option
	chk := KhzToGhz(r)                    // need to checking

	if ok == chk {
		t.Logf("Test TestKhzToGhz Passed: %v equal %v\n", ok, chk)
	} else {
		t.Errorf("Test TestKhzToGhz Failed: %v not equal %v\n", ok, chk)
	}
}

func TestMHzToHz(t *testing.T) {
	r := Megahertz(floatgen.GenRan(1, 9)) // generate random float
	ok := Hertz(r * 1e+6)                 // correct option
	chk := MHzToHz(r)                     // need to checking

	if ok == chk {
		t.Logf("Test TestMHzToHz Passed: %v equal %v\n", ok, chk)
	} else {
		t.Errorf("Test TestMHzToHz Failed: %v not equal %v\n", ok, chk)
	}
}

// Megahertz to Kilohertz converter test
func TestMHzToKHz(t *testing.T) {
	r := Megahertz(floatgen.GenRan(1, 9)) // generate random float
	ok := Kilohertz(r * 1000)             // correct option
	chk := MHzToKHz(r)                    // need to checking

	if ok == chk {
		t.Logf("Test TestMHzToKHz Passed: %v equal %v\n", ok, chk)
	} else {
		t.Errorf("Test TestMHzToKHz Failed: %v not equal %v\n", ok, chk)
	}
}

func TestMHzToGHz(t *testing.T) {
	r := Megahertz(floatgen.GenRan(1, 9)) // generate random float
	ok := Gigahertz(r / 1000)             // correct option
	chk := MHzToGHz(r)                    // need to checking

	if ok == chk {
		t.Logf("Test TestMHzToGHz Passed: %v equal %v\n", ok, chk)
	} else {
		t.Errorf("Test TestMHzToGHz Failed: %v not equal %v\n", ok, chk)
	}
}
