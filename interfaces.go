package main

type IOperators interface {
	Add() fnOperation
	Subtract() fnOperation
	Multiply() fnOperation
	Divide() fnOperation
}
