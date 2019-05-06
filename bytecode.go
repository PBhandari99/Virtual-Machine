package main

type Instruction struct {
	name     string
	noOfArgs int
}

const (
	IADD int = iota + 1 // 1 , and so on
	ISUB
	IMUL
	ILT
	IEQ
	BR
	BRT
	BRF
	ICONST
	LOAD
	CALL
	RET
	GLOAD
	STORE
	GSTORE
	PRINT
	POP
	HALT
)

var Instructions = []Instruction{
	Instruction{name: "Invalid_Instruction", noOfArgs: 0},
	Instruction{name: "iadd", noOfArgs: 0},
	Instruction{name: "isub", noOfArgs: 0},
	Instruction{name: "imul", noOfArgs: 0},
	Instruction{name: "itl", noOfArgs: 0},
	Instruction{name: "ieq", noOfArgs: 0},
	Instruction{name: "br", noOfArgs: 1},
	Instruction{name: "brt", noOfArgs: 1},
	Instruction{name: "brf", noOfArgs: 1},
	Instruction{name: "iconst", noOfArgs: 1},
	Instruction{name: "load", noOfArgs: 1},
	Instruction{name: "call", noOfArgs: 2},
	Instruction{name: "ret", noOfArgs: 0},
	Instruction{name: "gload", noOfArgs: 1},
	Instruction{name: "store", noOfArgs: 1},
	Instruction{name: "gstore", noOfArgs: 1},
	Instruction{name: "print", noOfArgs: 0},
	Instruction{name: "pop", noOfArgs: 0},
	Instruction{name: "halt", noOfArgs: 0},
}
