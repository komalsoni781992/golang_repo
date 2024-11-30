package basicdatatypes

import (
	"fmt"
)

// TestVariableConstant - introduction to constant and variable block - https://github.com/Diwakarsingh052/genspark/blob/main/1-variables/2-var.go
func TestVariableConstant() {
	// if we don't provide any values to the variable, they get initialized with there default values
	// var block/declaration list
	var (
		name  = "ajay"
		age   = 30
		marks = 80.5
	)
	const key = "some key"
	fmt.Println(name, age, marks, key)

	// peek into it for design pattern
	//os.OpenFile()
	//os.O_RDONLY
	//time.Second
	//http.StatusNotFound

	//l := log.New(os.Stdout, "INFO: ",log.LstdFlags)
	//l.Println("hello")
	//math.MaxUint8
}
