package colorsys

import (
	"math"
	"testing"
)

// Helper function to compare floats with tolerance
func epsilonEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

// Helper function to compare tuples of floats
func tuplesEqual(t1, t2 []float64, epsilon float64) bool {
	if len(t1) != len(t2) {
		return false
	}
	for i := 0; i < len(t1); i++ {
		if !epsilonEqual(t1[i], t2[i], epsilon) {
			return false
		}
	}
	return true
}

// func TestRgbToYiq(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		r, g, b  float64
// 		expected []float64
// 	}{
// 		{"Black", 0.0, 0.0, 0.0, []float64{0.0, 0.0, 0.0}},
// 		{"White", 1.0, 1.0, 1.0, []float64{1.0, 0.0, 0.0}},
// 		{"Red", 1.0, 0.0, 0.0, []float64{0.30, 0.50, 0.213}},
// 		{"Green", 0.0, 1.0, 0.0, []float64{0.59, -0.44, -0.5251}},
// 		{"Blue", 0.0, 0.0, 1.0, []float64{0.11, -0.65, 0.3121}},
// 		{"Yellow", 1.0, 1.0, 0.0, []float64{0.89, 0.50, -0.3121}},
// 		{"Purple", 0.5, 0.0, 0.5, []float64{0.23, 0.52, 0.2403}},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			y, i, q := RgbToYiq(c.r, c.g, c.b)
// 			if !approximatelyEqual(y, c.expected[0]) {
// 				t.Errorf("RgbToYiq failed for %s: got y = %v, want %v", c.name, y, c.expected[0])
// 			}
// 			if !approximatelyEqual(i, c.expected[1]) {
// 				t.Errorf("RgbToYiq failed for %s: got i = %v, want %v", c.name, i, c.expected[1])
// 			}
// 			if !approximatelyEqual(q, c.expected[2]) {
// 				t.Errorf("RgbToYiq failed for %s: got q = %v, want %v", c.name, q, c.expected[2])
// 			}
// 		})
// 	}
// }

// func TestYiqToRgb(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		y, i, q  float64
// 		expected []float64
// 	}{
// 		{"Black", 0.0, 0.0, 0.0, []float64{0.0, 0.0, 0.0}},
// 		{"White", 1.0, 0.0, 0.0, []float64{1.0, 1.0, 1.0}},
// 		{"Red", 0.30, 0.50, 0.29, []float64{1.0, 0.0, 0.0}},
// 		{"Green", 0.59, -0.44, 0.05, []float64{0.0, 1.0, 0.0}},
// 		{"Blue", 0.11, -0.65, -0.49, []float64{0.0, 0.0, 1.0}},
// 		{"Yellow", 0.89, 0.50, -0.29, []float64{1.0, 1.0, 0.0}},
// 		{"Purple", 0.23, 0.52, -0.05, []float64{0.5, 0.0, 0.5}},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			r, g, b := YiqToRgb(c.y, c.i, c.q)
// 			if !approximatelyEqual(r, c.expected[0]) {
// 				t.Errorf("YiqToRgb failed for %s: got r = %v, want %v", c.name, r, c.expected[0])
// 			}
// 			if !approximatelyEqual(g, c.expected[1]) {
// 				t.Errorf("YiqToRgb failed for %s: got g = %v, want %v", c.name, g, c.expected[1])
// 			}
// 			if !approximatelyEqual(b, c.expected[2]) {
// 				t.Errorf("YiqToRgb failed for %s: got b = %v, want %v", c.name, b, c.expected[2])
// 			}
// 		})
// 	}
// }

// Test RGB <-> HLS
func TestRgbToHls(t *testing.T) {
	cases := []struct {
		name     string
		rgb      []float64
		expected []float64
	}{
		{"Black", []float64{0.0, 0.0, 0.0}, []float64{0.0, 0.0, 0.0}},
		{"White", []float64{1.0, 1.0, 1.0}, []float64{0.0, 1.0, 0.0}},
		{"Red", []float64{1.0, 0.0, 0.0}, []float64{0.0, 0.5, 1.0}},
		{"Green", []float64{0.0, 1.0, 0.0}, []float64{1.0 / 3.0, 0.5, 1.0}},
		{"Blue", []float64{0.0, 0.0, 1.0}, []float64{2.0 / 3.0, 0.5, 1.0}},
	}

	epsilon := 1e-6
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			h, l, s := RgbToHls(c.rgb[0], c.rgb[1], c.rgb[2])
			actual := []float64{h, l, s}
			if !tuplesEqual(actual, c.expected, epsilon) {
				t.Errorf("RgbToHls failed for %s: got %v, want %v", c.name, actual, c.expected)
			}
		})
	}
}
func TestHlsToRgb(t *testing.T) {
	cases := []struct {
		name     string
		hls      []float64
		expected []float64
	}{
		{"Black", []float64{0.0, 0.0, 0.0}, []float64{0.0, 0.0, 0.0}},
		{"White", []float64{0.0, 1.0, 0.0}, []float64{1.0, 1.0, 1.0}},
		{"Red", []float64{0.0, 0.5, 1.0}, []float64{1.0, 0.0, 0.0}},
		{"Green", []float64{1.0 / 3.0, 0.5, 1.0}, []float64{0.0, 1.0, 0.0}},
		{"Blue", []float64{2.0 / 3.0, 0.5, 1.0}, []float64{0.0, 0.0, 1.0}},
	}

	epsilon := 1e-6 // Tolerance for floating-point comparison
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r, g, b := HlsToRgb(c.hls[0], c.hls[1], c.hls[2])
			actual := []float64{r, g, b}

			// Compare each component with tolerance
			for i := 0; i < 3; i++ {
				if math.Abs(actual[i]-c.expected[i]) > epsilon {
					t.Errorf("HlsToRgb failed for %s: got %v, want %v", c.name, actual, c.expected)
				}
			}
		})
	}
}

const epsilon = 1e-6 // Tolerance for floating-point comparison

// Helper function to compare floats with tolerance
func approximatelyEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestRgbToHsv(t *testing.T) {
	cases := []struct {
		name     string
		r, g, b  float64
		expected []float64
	}{
		{"Black", 0.0, 0.0, 0.0, []float64{0.0, 0.0, 0.0}},
		{"White", 1.0, 1.0, 1.0, []float64{0.0, 0.0, 1.0}},
		{"Red", 1.0, 0.0, 0.0, []float64{0.0, 1.0, 1.0}},
		{"Green", 0.0, 1.0, 0.0, []float64{1.0 / 3.0, 1.0, 1.0}},
		{"Blue", 0.0, 0.0, 1.0, []float64{2.0 / 3.0, 1.0, 1.0}},
		{"Yellow", 1.0, 1.0, 0.0, []float64{1.0 / 6.0, 1.0, 1.0}},
		{"Purple", 0.5, 0.0, 0.5, []float64{5.0 / 6.0, 1.0, 0.5}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			h, s, v := RgbToHsv(c.r, c.g, c.b)
			if !approximatelyEqual(h, c.expected[0]) {
				t.Errorf("RgbToHsv failed for %s: got h = %v, want %v", c.name, h, c.expected[0])
			}
			if !approximatelyEqual(s, c.expected[1]) {
				t.Errorf("RgbToHsv failed for %s: got s = %v, want %v", c.name, s, c.expected[1])
			}
			if !approximatelyEqual(v, c.expected[2]) {
				t.Errorf("RgbToHsv failed for %s: got v = %v, want %v", c.name, v, c.expected[2])
			}
		})
	}
}

func TestHsvToRgb(t *testing.T) {
	cases := []struct {
		name     string
		h, s, v  float64
		expected []float64
	}{
		{"Black", 0.0, 0.0, 0.0, []float64{0.0, 0.0, 0.0}},
		{"White", 0.0, 0.0, 1.0, []float64{1.0, 1.0, 1.0}},
		{"Red", 0.0, 1.0, 1.0, []float64{1.0, 0.0, 0.0}},
		{"Green", 1.0 / 3.0, 1.0, 1.0, []float64{0.0, 1.0, 0.0}},
		{"Blue", 2.0 / 3.0, 1.0, 1.0, []float64{0.0, 0.0, 1.0}},
		{"Yellow", 1.0 / 6.0, 1.0, 1.0, []float64{1.0, 1.0, 0.0}},
		{"Purple", 5.0 / 6.0, 1.0, 0.5, []float64{0.5, 0.0, 0.5}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r, g, b := HsvToRgb(c.h, c.s, c.v)
			if !approximatelyEqual(r, c.expected[0]) {
				t.Errorf("HsvToRgb failed for %s: got r = %v, want %v", c.name, r, c.expected[0])
			}
			if !approximatelyEqual(g, c.expected[1]) {
				t.Errorf("HsvToRgb failed for %s: got g = %v, want %v", c.name, g, c.expected[1])
			}
			if !approximatelyEqual(b, c.expected[2]) {
				t.Errorf("HsvToRgb failed for %s: got b = %v, want %v", c.name, b, c.expected[2])
			}
		})
	}
}
