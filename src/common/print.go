package main

import "fmt"

//Print(), Printf() ve Println(), () içine aldığı string ifadeyi olduğu gibi yazdırır. Println() 'in diğerlerinden tek farkı alt satıra geçmesidir.

func PrintFunction() {
	fmt.Print("Hello")
	fmt.Println("")
	fmt.Println("Hello")
	fmt.Printf("Hello")
	fmt.Println("")

	name := "Ezgi"
	fmt.Print(name)
	fmt.Println("")
	fmt.Println(name)
	fmt.Printf(name)
	fmt.Println("")

	fmt.Print("Hello ", name) //parametreye bir string ifadeyle bir değişkeni karışık olarak yazdığımızda ("", değişken), araya boşluk koymadan yazdırır.
	fmt.Println("")
	fmt.Println("Hello", name) //parametreye bir string ifadeyle bir değişkeni karışık olarak yazdığımızda ("", değişken), araya boşluk koyarak yazdırır.
	fmt.Printf("Hello", name)  //parametreye bir string ifadeyle bir değişkeni karışık olarak yazdığımızda ("", değişken), verilen değişken parametresini formatlayarak, string ifadeyi ise direkt yazdırır.
	fmt.Println("")
	fmt.Printf("Hello %v", name) //string parametresi içine yazılan özel "%v" karakteri, ardından gelen değişken parametresinin %v'sini, yani değişkenin değerini gösterir.
	fmt.Println("")

	name, age := "Ezgi", 22
	fmt.Print("Hello! My name is ", name, " and i'm ", age, " years old.")
	fmt.Println("")
	fmt.Println("Hello! My name is", name, "and i'm", age, "years old.")
	fmt.Printf("Hello! My name is %v and i'm %v years old.", name, age)
}

func main() {
	PrintFunction()
}
