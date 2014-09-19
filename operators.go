package main

import "errors"

type NormalOperator struct {
	Num1 float64
	Num2 float64
}

func (op *NormalOperator) Add() (float64, string, error) {
	return op.Num1 + op.Num2, "+", nil
}

func (op *NormalOperator) Subtract() (float64, string, error) {
	return op.Num1 - op.Num2, "-", nil
}

func (op *NormalOperator) Multiply() (float64, string, error) {
	return op.Num1 * op.Num2, "*", nil
}

func (op *NormalOperator) Divide() (float64, string, error) {
	if op.Num2 == 0.0 {
		return 0.0, "", errors.New("For division (div), Num2 cannot equal 0.")
	}
	return op.Num1 / op.Num2, "/", nil
}
