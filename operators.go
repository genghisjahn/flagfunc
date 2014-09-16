package main

type NormalOperator struct {
	Num1 int
	Num2 int
}

func (op *NormalOperator) Add() (float64, string, error) {
	return op.Num1 + op.Num2, "+", nil
}
