package basicdatatypes

import (
	"fmt"
	"math/cmplx"
)

// TestComplexNumbers tests complex data type complex64 and complex128
func TestComplexNumbers() {
	// Internally usses float, so it has same comparison limitations
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x)) // modulus of complex number
}
