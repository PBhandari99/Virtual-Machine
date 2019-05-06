package main

import "fmt"

const (
	FALSE              = 0
	TRUE               = 1
	DEFAULT_STACK_SIZE = 100
)

type VM struct {
	Data    []int
	Code    []int
	Stack   []int
	Globals []int
	IP      int
	FP      int
	SP      int
}

func NewVM(code []int, ip int, dataSize int, numOfGlobs int) *VM {
	vm := VM{
		Data:    make([]int, dataSize),
		Code:    code,
		Stack:   make([]int, DEFAULT_STACK_SIZE),
		Globals: make([]int, numOfGlobs),
		IP:      ip,
		SP:      -1,
	}
	return &vm
}

func (vm *VM) stackPop() int {
	value := vm.Stack[vm.SP]
	vm.SP--
	return value
}

func (vm *VM) stackPush(value int) {
	vm.SP++
	vm.Stack[vm.SP] = value
}

func (vm *VM) nextInstruction() int {
	instruction := vm.Code[vm.IP]
	vm.IP++
	return instruction
}

func (vm *VM) execute() {
	var opcode int
	for vm.IP < len(vm.Code) {
		fmt.Printf("%x  %s\n", vm.IP, Instructions[vm.Code[vm.IP]].name)
		opcode = vm.Code[vm.IP]
		vm.IP++
		switch opcode {
		case ICONST:
			value := vm.Code[vm.IP]
			vm.IP++
			vm.SP++
			vm.Stack[vm.SP] = value
			break
		case IADD:
			sum := vm.stackPop() + vm.stackPop()
			vm.stackPush(sum)
			break
		case ISUB:
			a := vm.stackPop()
			b := vm.stackPop()
			vm.stackPush(b - a)
			break
		case IMUL:
			prod := vm.stackPop() * vm.stackPop()
			vm.stackPush(prod)
			break
		case BR:
			vm.IP = vm.nextInstruction()
			break
		case BRT:
			a := vm.stackPop()
			code := vm.nextInstruction()
			if a == TRUE {
				vm.IP = code
			}
			break
		case BRF:
			a := vm.stackPop()
			code := vm.nextInstruction()
			if a == FALSE {
				vm.IP = code
			}
			break
		case ILT:
			b := vm.stackPop()
			a := vm.stackPop()
			if a < b {
				vm.stackPush(TRUE)
			} else {
				vm.stackPush(FALSE)
			}
			break
		// case LOAD: // load local or arg
		// break
		case GLOAD:
			vm.stackPush(vm.Globals[vm.nextInstruction()])
			break
		case GSTORE:
			vm.Globals[vm.nextInstruction()] = vm.stackPop()
			break
		case PRINT:
			value := vm.stackPop()
			fmt.Printf("%d\n", value)
			break
		case CALL:
			funcAddress := vm.nextInstruction() // get the address of the func to jump to.
			numOfArgs := vm.nextInstruction()   // get the number of args in the function call
			vm.stackPush(numOfArgs)             // push the num of args to the call stack.
			vm.stackPush(vm.FP)                 // push frame pointer to the stack.
			vm.stackPush(vm.IP)                 // push return address to the stack.
			vm.FP = vm.SP                       // point frame pointer to the return address on the stack.
			// TODO: push all arguments to stack before jump.
			vm.IP = funcAddress // Jump to the function.
			break
		case RET:
			returnValue := vm.stackPop()
			vm.SP = vm.FP
			vm.IP = vm.stackPop()
			vm.FP = vm.stackPop()      // Reset frame pointer to the frame of caller.
			numOfArgs := vm.stackPop() // get the num of args to clean.
			vm.SP -= numOfArgs         // clear the arguments from the stack.
			vm.stackPush(returnValue)  // push the return value to the stack.
			break
		case HALT:
			return
		default:
			// TODO: implement this case for the invalid instruction.
			return
		}
	}
}
