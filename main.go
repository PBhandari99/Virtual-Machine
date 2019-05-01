package main

func main() {
	code := []Bytecode{
		ICONST, 99,
		PRINT,
		HALT,
	}
	vm := NewVM(code, 0, 0)
	vm.execute()
}
