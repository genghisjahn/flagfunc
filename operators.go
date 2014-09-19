package main

import (
	"errors"
	"fmt"
)

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

type OffByOneOperator struct {
	Number1 float64
	Number2 float64
}

func (op *OffByOneOperator) Add() (float64, string, error) {
	return op.Number1 + op.Number2 + 1, fmt.Sprintf("%v + %v = ", op.Number1, op.Number2), nil
}

func (op *OffByOneOperator) Subtract() (float64, string, error) {
	return op.Number1 - op.Number2 + 1, fmt.Sprintf("%v - %v = ", op.Number1, op.Number2), nil
}

func (op *OffByOneOperator) Multiply() (float64, string, error) {
	return op.Number1*op.Number2 + 1, fmt.Sprintf("%v * %v = ", op.Number1, op.Number2), nil
}

func (op *OffByOneOperator) Divide() (float64, string, error) {
	if op.Number2 == 0.0 {
		return 0.0, "", errors.New("For division (div), Num2 cannot equal 0.")
	}
	return op.Number1/op.Number2 + 1, fmt.Sprintf("%v / %v = ", op.Number1, op.Number2), nil
}
