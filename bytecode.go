package main

type Bytecode int

const (
	IADD Bytecode = iota // 1 , and so on
	ISUB
	IMUL
	ILT
	IEQ
	BR
	BRT
	BRF
	ICONST
	LOAD
	GLOAD
	STORE
	GSTORE
	PRINT
	POP
	HALT
)
