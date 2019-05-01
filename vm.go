package main

import "fmt"

type VM struct {
	Data  []int
	Code  []Bytecode
	Stack []Bytecode
	IP    int
	FP    int
	SP    int
}

func NewVM(code []Bytecode, ip int, dataSize int) *VM {
	vm := VM{
		Data:  make([]int, dataSize),
		Code:  code,
		Stack: make([]Bytecode, 100), // TODO:Correct the size later.
		IP:    ip,
		SP:    -1,
	}
	return &vm
}

func (vm *VM) execute() {
	var opcode Bytecode
	for vm.IP < len(vm.Code) {
		opcode = vm.Code[vm.IP]
		vm.IP++
		switch opcode {
		case ICONST:
			value := vm.Code[vm.IP]
			vm.IP++
			vm.SP++
			vm.Stack[vm.SP] = value
			break
		case PRINT:
			value := vm.Stack[vm.SP]
			vm.SP--
			fmt.Printf("%d\n", value)
			break
		case HALT:
			return
		}
	}
}
