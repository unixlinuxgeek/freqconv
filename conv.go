package freqconv

// HzToKh Hertz to Kilohertz converter
func HzToKh(h Hertz) Kilohertz {
	return Kilohertz(h / 1000)
}

func HzTo(h Hertz) Kilohertz {
	return Kilohertz(h / 1000)
}

func HzToMhz(h Megahertz) Megahertz {
	return Megahertz(1)
}

// KhToHz Kilohertz to Hertz converter
func KhToHz(kh Kilohertz) Hertz {
	return Hertz(kh * 1000)
}
