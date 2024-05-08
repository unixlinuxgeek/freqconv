package freqconv

import (
	"github.com/unixlinuxgeek/floatgen"
	"testing"
)

func TestHertzToKilohertz(t *testing.T) {
	r := Hertz(floatgen.GenRan(1, 9)) // generate random float
	ok := Kilohertz(r / 1000)         // correct option
	res := HertzToKilohertz(r)        // need to check

	if ok == res {
		t.Logf("Test Passed: %v equal %v\n", ok, res)
	} else {
		t.Fatalf("Test Failed: %v not equal %v\n", ok, res)
	}
}
