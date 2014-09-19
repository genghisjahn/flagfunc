package main

import (
	"errors"
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

type NormalOperator struct {
	Num1 float64
	Num2 float64
}

func (op *NormalOperator) Add() (float64, string, error) {
	return op.Num1 + op.Num2, fmt.Sprintf("%v + %v = ", op.Num1, op.Num2), nil
}

func (op *NormalOperator) Subtract() (float64, string, error) {
	return op.Num1 - op.Num2, fmt.Sprintf("%v - %v = ", op.Num1, op.Num2), nil
}

func (op *NormalOperator) Multiply() (float64, string, error) {
	return op.Num1 * op.Num2, fmt.Sprintf("%v * %v = ", op.Num1, op.Num2), nil
}

func (op *NormalOperator) Divide() (float64, string, error) {
	if op.Num2 == 0.0 {
		return 0.0, "", errors.New("For division (div), Num2 cannot equal 0.")
	}
	return op.Num1 / op.Num2, fmt.Sprintf("%v / %v = ", op.Num1, op.Num2), nil
}

type IOperator interface {
	Add() (float64, string, error)
	Subtract() (float64, string, error)
	Multiply() (float64, string, error)
	Divide() (float64, string, error)
}
