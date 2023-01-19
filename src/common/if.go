package main

import "fmt"

//conditional statements (koşullu ifadeler) : en temel koşullu ifade if statement 'tir.
//if statement syntax: if <boolean expression> { code }

func IfFunction() {
	//<boolean expression> yerine direkt false yazıldığında bir çıktı vermez. çünkü o if bloğunun içine girmez. örneğin;
	if !false {
		fmt.Println("mesaj gösterilemez.")
	}
	if !true {
		fmt.Println("mesaj gösterilemez.")
	}

	if 3 > 2 {
		fmt.Println("mesaj gösterilir.")
	}

	x := 27
	if x%2 == 0 {
		fmt.Println(x, "çift sayıdır.")
	}
	if x%2 != 0 {
		fmt.Println(x, "tek sayıdır.")
	}
	fmt.Println(x) //22.satırdaki x değişkeni if koşulu dışında ve fonksiyon içinde tanımlı olduğundan bu satırdaki kod çalışır. ancak çoğu durumda bir if koşuluna ait olan değişkenin if bloğu/koşulu içinde tanımlı olması istenir. bunu sağlamak için if sonrası 2 statement yazılır ve aralarına ; konur. örneğin ;

	if x := 10; x < 0 {
		fmt.Println(x, "negatif sayıdır.")
	} else {
		if x%2 == 0 {
			fmt.Println(x, "çift sayıdır.")
		} else {
			fmt.Println(x, "tek sayıdır.")
		}
	} //bu örnek aynı zamanda iç içe if yapısı örneğidir.

	y := 10
	if y := -5; y < 0 {
		fmt.Println(y, "negatif sayıdır.")
	} else if y%2 == 0 {
		fmt.Println(y, "çift sayıdır.")
	} else {
		fmt.Println(y, "tek sayıdır.")
	}
	fmt.Println(y)
	//2 y 'nin scope 'u/kapsamı farklı olduğundan 2 y de yazılır. buna if bloğu içerisindeki "shadowing" denir.
}

//if else statement syntax: if <boolean expression> { code } else { code }

func IfElseFunction() {
	x := 27
	if x%2 == 0 {
		fmt.Println(x, "çift sayıdır.")
	} else {
		fmt.Println(x, "tek sayıdır.")
	}

	/*if b := 20; b%2 == 0 {
		fmt.Println(b, "çifttir")
	} else {
		fmt.Println(b, "tektir")
	}*/
	//daha idiomatic hale getirelim;
	b := 20
	if b%2 == 0 {
		fmt.Println(b, "çifttir")
		return // çift değilse zaten tek olacağından sadece bu koşul üzerinden else'i yazabiliriz. bu koşulun sağlanmadığı durumda return ile koşuldan çıkılır ve 542.satır çalışır
	}
	fmt.Println(b, "tektir")
}

//if - else if - else statement syntax: if <boolean expression> { code } else if <boolean expression> { code } else { code }

func IfElseIfElseFunction() {
	x := -5
	if x < 0 {
		fmt.Println(x, "negatif sayıdır.")
	} else if x%2 == 0 {
		fmt.Println(x, "çift sayıdır.")
	} else {
		fmt.Println(x, "tek sayıdır.")
	} //bu örnek şöyle de yazılabilir;

	if y := -5; y < 0 {
		fmt.Println(y, "negatif sayıdır.")
	} else if y%2 == 0 {
		fmt.Println(y, "çift sayıdır.")
	} else {
		fmt.Println(y, "tek sayıdır.")
	}
}

func main() {
	IfFunction()
	IfElseFunction()
	IfElseIfElseFunction()
}
