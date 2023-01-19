package main

import "fmt"

//variable declaration and variable assign yapılacak ;
//syntax: var - <name of variable> - <type of variable>

func VariableFunction() {
	var name string
	name = "Ezgi"
	//var name string = "Ezgi" : kısayoldur
	//var name = "Ezgi" : daha kısa olan kısayoldur
	//name := "Ezgi" : en kısa olan kısayoldur (short declaration (statement))
	//var firstname, lastname string = "Ezgi", "Sivas" : tek satırda 1'den fazla değişkeni bu şekilde yazabiliriz.

	var age int
	age = 22
	//var age int = 22 : kısayoldur
	//var age = 22 : daha kısa olan kısayoldur
	//age := 22 : en kısa olan kısayoldur (short declaration (statement))
	//age := 23 : dersen yeniden declaration olacağı için hata verir. ancak yanına başka bir değişken yazıldığında bu kısayolu kullanabiliyoruz; age, height := 23, 170 şeklinde.

	var isMarried bool
	isMarried = false
	//var isMarried bool = false : kısayoldur
	//var isMarried = false : daha kısa olan kısayoldur
	//isMarried := false : en kısa olan kısayoldur (short declaration (statement))

	var weight float32
	weight = 55.2
	//var weight int = 55.2 : kısayoldur
	//var weight = 55.2 : daha kısa olan kısayoldur
	//weight := 55.2 : en kısa olan kısayoldur (short declaration (statement))

	//var name, age, isMarried, weight = "Ezgi", 22, false, 55.2 : tek satırda 1'den fazla değişkeni bu şekilde hem declare eder, hem de o değişkenlere assign edebiliriz.
	//name, age, isMarried, weight := "Ezgi", 22, false, 55.2 : tek satırda 1'den fazla değişkenin en kısa olan kısayoldur (short declaration (statement))

	fmt.Println(name)
	//fmt.Println(firstname, lastname)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)

	//Printf() : değişkenlerin veri tiplerini okumak için kullanılabilir. parantez içerisindeki "%T" veri tipini yazdırmak için konur. "\n" ise aşağıdaki satıra geçiş için vardır.
	/*fmt.Printf(%T\n, name)
	fmt.Printf(%T\n, age)
	fmt.Printf(%T\n, isMarried)
	fmt.Printf(%T\n, weight)*/
}

func VariableFunction2() {
	var name string = "Ezgi"
	fmt.Println("Hello", name)
}

func MultipleVariableFunction() {
	x := 5
	y := 10
	fmt.Println("x:", x, "y:", y)
	x, y = y, x //multiple assign (çoklu değer atama) : x=y, y=x demektir.
	fmt.Println("x:", x, "y:", y)
}

func main() {
	VariableFunction()
	VariableFunction2()
	MultipleVariableFunction()
}
