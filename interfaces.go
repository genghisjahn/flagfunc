package main

type IOperator interface {
	Add() (float64, string, error)
	Subtract() (float64, string, error)
	Multiply() (float64, string, error)
	Divide() (float64, string, error)
}
