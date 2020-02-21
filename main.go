package main

import "golua/lua"
import "fmt"

func test(L *lua.State) int {
	fmt.Println("hello world! from go!")
	return 0
}

func test2(L *lua.State) int {
	arg := L.CheckInteger(-1)
	argfrombottom := L.CheckInteger(1)
	fmt.Print("test2 arg: ")
	fmt.Println(arg)
	fmt.Print("from bottom: ")
	fmt.Println(argfrombottom)
	return 0
}

func adder(L *lua.State) int {
	a := L.ToInteger(1)
	b := L.ToInteger(2)
	L.PushInteger(int64(a + b))
	L.PushString("adder is ok")
	return 2 // number of return values
}

func main()  {
	L := lua.NewState()
	defer L.Close()
	L.OpenLibs()

	L.GetGlobal("print")
	L.PushString("Hello World!")
	L.Call(1,0)

	L.PushGoFunction(test)
	L.PushGoFunction(test)
	L.PushGoFunction(test)
	L.PushGoFunction(test)

	L.PushGoFunction(test2)
	L.PushInteger(42)
	L.Call(1,0)


	L.Call(0,0)
	L.Call(0,0)
	L.Call(0,0)

	// this will fail as we didn't register test2 function
	err := L.DoString("test2(42)")

	fmt.Printf("err %v\n", err)

	L.Register("adder", adder)
	L.DoString("print(adder(112, 2))")
}