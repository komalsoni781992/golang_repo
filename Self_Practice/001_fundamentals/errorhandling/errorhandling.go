package errorhandling

import "errors"

// TestErrorHandling - you should set the other return values to their zero values when a non-nil error is returned except for sentinel errors.
func TestErrorHandling(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}
