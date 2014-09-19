package main

type NormalOperator struct {
	Num1 float64
	Num2 float64
}

func (op *NormalOperator) Add() (float64, string, error) {
	return op.Num1 + op.Num2, "+", nil
}
