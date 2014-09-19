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

	offbyone := OffByOneOperator{}
	offbyone.Number1 = 10.0
	offbyone.Number2 = 15.0

	ShowAddition(&normalOp)
	ShowAddition(&offbyone)

}

func ShowAddition(op IOperator) {
	if result, problem, err := op.Add(); err != nil {
		log.Println("Error Addition: ", err)
	} else {
		fmt.Printf("%T %v %v\n", op, problem, result)
	}
}
