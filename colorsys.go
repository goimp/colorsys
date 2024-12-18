package colorsys

/* -----------------------------------------------------
Conversion functions between RGB and other color systems.

This modules provides two functions for each color system ABC:

  rgb_to_abc(r, g, b) --> a, b, c
  abc_to_rgb(a, b, c) --> r, g, b

All inputs and outputs are triples of floats in the range [0.0...1.0]
(with the exception of I and Q, which covers a slightly larger range).
Inputs outside the valid range may cause exceptions or invalid outputs.

Supported color systems:
RGB: Red, Green, Blue components
YIQ: Luminance, Chrominance (used by composite video signals)
HLS: Hue, Luminance, Saturation
HSV: Hue, Saturation, Value
-------------------------------------------------------- */

// References:
// http://en.wikipedia.org/wiki/YIQ
// http://en.wikipedia.org/wiki/HLS_color_space
// http://en.wikipedia.org/wiki/HSV_color_space
// https://github.com/python/cpython/blob/main/Lib/colorsys.py

// Some floating-point constants

const (
	ONE_THIRD = 1.0 / 3.0
	ONE_SIXTH = 1.0 / 6.0
	TWO_THIRD = 2.0 / 3.0
)

// YIQ: used by composite video signals (linear combinations of RGB)
// Y: perceived grey level (0.0 == black, 1.0 == white)
// I, Q: color components
//
// There are a great many versions of the constants used in these formulae.
// The ones in this library uses constants from the FCC version of NTSC.

func RgbToYiq(r, g, b float64) (float64, float64, float64) {
	y := 0.30*r + 0.59*g + 0.11*b
	i := 0.74*(r-y) - 0.27*(b-y)
	q := 0.48*(r-y) + 0.41*(b-y)
	return y, i, q
}

func YiqToRgb(y, i, q float64) (float64, float64, float64) {
	// r = y + (0.27*q + 0.41*i) / (0.74*0.41 + 0.27*0.48)
	// b = y + (0.74*q - 0.48*i) / (0.74*0.41 + 0.27*0.48)
	// g = y - (0.30*(r-y) + 0.11*(b-y)) / 0.59

	r := y + 0.9468822170900693*i + 0.6235565819861433*q
	g := y - 0.27478764629897834*i - 0.6356910791873801*q
	b := y - 1.1085450346420322*i + 1.7090069284064666*q

	if r < 0.0 {
		r = 0.0
	}
	if g < 0.0 {
		g = 0.0
	}
	if b < 0.0 {
		b = 0.0
	}
	if r > 1.0 {
		r = 1.0
	}
	if g > 1.0 {
		g = 1.0
	}
	if b > 1.0 {
		b = 1.0
	}
	return r, g, b
}

// HLS: Hue, Luminance, Saturation
// H: position in the spectrum
// L: color lightness
// S: color saturation

func RgbToHls(r, g, b float64) (float64, float64, float64) {
	maxc := max(r, g, b)
	minc := min(r, g, b)
	sumc := (maxc + minc)
	rangec := (maxc - minc)
	l := sumc / 2.0
	if minc == maxc {
		return 0.0, l, 0.0
	}
	var s, h float64
	if l <= 0.5 {
		s = rangec / sumc
	} else {
		s = rangec / (2.0 - maxc - minc) // Not always 2.0-sumc: gh-106498.
	}
	rc := (maxc - r) / rangec
	gc := (maxc - g) / rangec
	bc := (maxc - b) / rangec
	if r == maxc {
		h = bc - gc
	} else if g == maxc {
		h = 2.0 + rc - bc
	} else {
		h = 4.0 + gc - rc
	}
	// h = (h / 6.0) - float64(int(h/6.0))
	h = _mod(h/6.0, 1.0)
	return h, l, s
}

func HlsToRgb(h, l, s float64) (float64, float64, float64) {
	var m1, m2 float64
	if s == 0.0 {
		return l, l, l
	}
	if l <= 0.5 {
		m2 = l * (1.0 + s)
	} else {
		m2 = l + s - (l * s)
	}
	m1 = 2.0*l - m2
	return _v(m1, m2, h+ONE_THIRD), _v(m1, m2, h), _v(m1, m2, h-ONE_THIRD)
}

func _v(m1, m2, h float64) float64 {
	if h < 0.0 {
		h += 1.0
	}
	if h > 1.0 {
		h -= 1.0
	}
	if h < ONE_SIXTH {
		return m1 + (m2-m1)*6.0*h
	}
	if h < 1.0/2.0 {
		return m2
	}
	if h < TWO_THIRD {
		return m1 + (m2-m1)*(TWO_THIRD-h)*6.0
	}
	return m1
}

// Custom _mod function for floating-point numbers
func _mod(number, divisor float64) float64 {
	quotient := float64(int(number / divisor))
	return number - quotient*divisor
}

// HSV: Hue, Saturation, Value
// H: position in the spectrum
// S: color saturation ("purity")
// V: color brightness

func RgbToHsv(r, g, b float64) (float64, float64, float64) {
	maxc := max(r, g, b)
	minc := min(r, g, b)
	rangec := (maxc - minc)
	v := maxc
	if minc == maxc {
		return 0.0, 0.0, v
	}
	s := rangec / maxc
	rc := (maxc - r) / rangec
	gc := (maxc - g) / rangec
	bc := (maxc - b) / rangec
	var h float64
	if r == maxc {
		h = bc - gc
	} else if g == maxc {
		h = 2.0 + rc - bc
	} else {
		h = 4.0 + gc - rc
	}
	// Normalize the hue to the [0, 1) range
	h = _mod(h/6.0, 1.0)
	if h < 0 {
		h += 1.0
	}
	return h, s, v
}

func HsvToRgb(h, s, v float64) (float64, float64, float64) {
	if s == 0.0 {
		// If saturation is 0, return grayscale
		return v, v, v
	}

	i := int(h * 6.0) // XXX assume int() truncates!
	f := (h * 6.0) - float64(i)
	p := v * (1.0 - s)
	q := v * (1.0 - s*f)
	t := v * (1.0 - s*(1.0-f))

	i = int(_mod(float64(i), 6.0)) // Wrap sector index to [0, 5]

	switch i {
	case 0:
		return v, t, p
	case 1:
		return q, v, p
	case 2:
		return p, v, t
	case 3:
		return p, q, v
	case 4:
		return t, p, v
	case 5:
		return v, p, q
	default:
		// Cannot get here
		return 0, 0, 0
	}
}
