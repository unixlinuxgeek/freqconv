package freqconv

import "fmt"

type Hertz float64
type Kilohertz float64
type Megahertz float64
type Gigahertz float64

func (h Hertz) String() string {
	return fmt.Sprintf("%g Hz\n", h)
}

func (k Kilohertz) String() string {
	return fmt.Sprintf("%g kHz\n", k)
}

func (MHz Megahertz) String() string {
	return fmt.Sprintf("%g MHz\n", MHz)
}

func (GHz Gigahertz) String() string {
	return fmt.Sprintf("%g GHz\n", GHz)
}
