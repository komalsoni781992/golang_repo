package main

import (
	"fmt"
	"fundamentals/practice/day4/question1/calc"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 4 {
		s := os.Args[1]
		x, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Println("invalid number", x)
			return
		}
		y, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Println("invalid number", y)
			return
		}
		switch s {
		case "+":
			fmt.Println(calc.Add(x, y))
		case "-":
			fmt.Println(calc.Sub(x, y))
		case "*":
			fmt.Println(calc.Mul(x, y))
		case "%":
			fmt.Println(calc.Mod(x, y))
		default:
			panic("Invalid operation " + s)
		}
	} else {
		panic("Invalid input provided.")
	}
}
