package basicdatatypes

import (
	"fmt"

	"math/rand"
)

// TestRandomNumber - generates random number betwwen a and b
func TestRandomNumber() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(401000000-56000000+1) + 56000000
	fmt.Println(num)
}
