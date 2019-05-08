package main

// import "fmt"

const (
	ip = 0
)

func main() {

	// Simple hello world
	numOfGlobs := 0
	hello := []int{
		ICONST, 99,
		ICONST, 100,
		IADD,
		PRINT,
		HALT,
	}
	vm := NewVM(hello, ip, 0, numOfGlobs)
	vm.execute()

	// While loop
	numOfGlobs = 2
	loop := []int{
		// 2 Globals N, I
		ICONST, 10, // N = 10
		GSTORE, 0,
		ICONST, 0, // I = 0
		GSTORE, 1,

		// WHILE I<N:
		// START 8
		GLOAD, 1,
		GLOAD, 0,
		ILT,
		BRF, 24,

		// I++
		GLOAD, 1,
		ICONST, 1,
		IADD,
		GSTORE, 1,

		BR, 8, // jump to 8

		// DONE 24
		HALT,
	}
	vm = NewVM(loop, ip, 0, numOfGlobs)
	vm.execute()
}
