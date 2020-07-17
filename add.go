package soa

// AddFloats adds two slices [a] and [b] of float64 values and saves the result into [c].
// SSE2 required
func AddFloat64s(a, b, c []float64)
