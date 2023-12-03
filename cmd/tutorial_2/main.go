package main

import "fmt"

func main() {
	var intNum int16 = 1
	intNum = intNum + 1
	fmt.Println("Int16: ",intNum)

	var floatNum float32 = 1.0
	floatNum = floatNum + 1.0
	fmt.Println("Float32: ",floatNum)

	var stringText string = "Hello, World!"
	fmt.Println("String: ",stringText)

	var boolValue bool = true
	fmt.Println("Boolean: ",boolValue)

	var byteValue byte = 255
	fmt.Println("Byte: ",byteValue)

	var floatNum32 float32 = 10.0
	var intNum32 int32 = 10
	var result float32 = float32(intNum32) + floatNum32
	fmt.Println("Result: ",result) 

}
