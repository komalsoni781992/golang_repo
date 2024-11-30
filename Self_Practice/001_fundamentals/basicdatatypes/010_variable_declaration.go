package basicdatatypes

import "fmt"

// DataTypeDeclaration - ways to declare variables
func DataTypeDeclaration() {
	//If you are declaring a variable at package level, you must use var; := is not legal outside of functions
	var a int = 10 //var-type-val
	var b = 10     //var-val
	var c int      //var-type // should be used this way to imply zero value initialzation

	//x := byte(20), it is idiomatic to write var x byte = 20

	var d, e int = 10, 20  //multiple var-type-val
	var f, g int           //multiple var-type
	var h, i = 10, "hello" //multiple var-val

	l, m := 10, "hello" // multiple shorthand var-val
	n := 10             // shorthand var-val
	var (
		q    int                  // var-type
		r           = 20          // var-val
		s    int    = 30          // var-type-val
		t, u        = 40, "hello" // multiple var-val
		v, w string               // multiple var-type
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, l, m, n, q, r, s, t, u, v, w)
}
