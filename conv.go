package freqconv

// Hertz to Kilohertz, Megahertz and Gigahertz converters

// HzToKHz Hertz to Kilohertz converter
func HzToKHz(h Hertz) Kilohertz {
	return Kilohertz(h / 1000)
}

// HzToMHz Hertz to Megahertz converter
func HzToMHz(h Hertz) Megahertz {
	return Megahertz(h / 1e+6)
}

// HzToGHz Hertz to Gigahertz converter
func HzToGHz(h Hertz) Gigahertz {
	return Gigahertz(h / 1e+9)
}

// Kilohertz to Hertz, Megahertz and Gigahertz converters

// KHzToHz Kilohertz to Hertz converter
func KHzToHz(KHz Kilohertz) Hertz {
	return Hertz(KHz * 1000)
}

// KHzToMhz Kilohertz to Megahertz converter
func KHzToMhz(KHz Kilohertz) Megahertz {
	return Megahertz(KHz / 1000)
}

// KhzToGhz Kilohertz to Gigahertz converter
func KhzToGhz(KHz Kilohertz) Gigahertz {
	return Gigahertz(KHz / 1e+6)
}

// Megahertz to Hertz, Kilohertz and Gigahertz converters

// MHzToHz Megahertz to Hertz converter
func MHzToHz(mhz Megahertz) Hertz {
	return Hertz(mhz * 1e+6)
}

// MHzToKHz Megahertz to Kilohertz converter
func MHzToKHz(MHz Megahertz) Kilohertz {
	return Kilohertz(MHz * 1000)
}

// MHzToGHz Megahertz to Gigahertz converter
func MHzToGHz(MHz Megahertz) Gigahertz {
	return Gigahertz(MHz / 1000)
}
