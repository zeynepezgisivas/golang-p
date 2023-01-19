package main

import (
	"fmt"
	"strings"
)

func ContainsFunction() {
	fmt.Println(strings.Contains("seafood", "foo")) //strings paketindeki constains methodu 2.string ifade 1.string ifadede var mı diye kontrol ediyor
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	fmt.Println(strings.Count("animatrix", "a")) //strings paketindeki count methodu 2.string ifadedeki karakterin 1.string ifadede kaç tane olduğunu veriyor
	fmt.Println(strings.ToUpper("GOPHER"))       //parametre olarak ALDIğı string ifadeyi büyük harflerle yazdırıyor
}

func main() {
	ContainsFunction()
}
