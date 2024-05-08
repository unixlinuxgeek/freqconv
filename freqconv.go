package freqconv

import "fmt"

type Hertz float64
type Kilohertz float64
type Megahertz float64
type Gigahertz float64

func (h Hertz) String() string {
	return fmt.Sprintf("%g Hertz", h)
}

func (k Kilohertz) String() string {
	return fmt.Sprintf("%g kHz", k)
}

func (MHz Megahertz) String() string {
	return fmt.Sprintf("%g MHz", MHz)
}

func (GHz Gigahertz) String() string {
	return fmt.Sprintf("%g GHz", GHz)
}
