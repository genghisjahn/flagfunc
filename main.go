package main

import (
	"fmt"
	"log"
)

func main() {

	log.Println("FlagFunc - How to use flags with functions and interfaces.")

	normalOp := NormalOperator{}
	normalOp.Num1 = 5.0
	normalOp.Num2 = 5.0

	ShowAddition(&normalOp)

}

func ShowAddition(op IOperator) {
	if result, problem, err := op.Add(); err != nil {
		log.Println("Error Addition: ", err)
	} else {
		fmt.Println(problem, result)
	}
}
