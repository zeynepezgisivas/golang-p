package main

import (
	"fmt"
	"strconv"
)

func typeConversionFunction() {
	x := 20
	y := 22.2
	fmt.Printf("%v %T\n", x, x) //Printf() : " " içine yazılan string'in formatlanmış halini yazdırır. "%v" yanındaki x değişkeninin veri tipini, "%T" ise aldığı değeri belirtir.
	fmt.Printf("%v %T\n", y, y)

	z := strconv.Itoa(x) //x değişkenini string olarak ("20") yazdırır.
	fmt.Printf("%v %T\n", z, z)

	//fmt.Println(x + y) : yazılsaydı type'lar farklı olduğundan "type mismatched error" alınırdı. farklı veri tiplerine sahip olan değişkenleri birbirleriyle hiçbir işleme sokamayız. bunu yapmanın tek yolu type conversion'dır (veri tipi dönüştürme).
	//type conversion syntax : type(value) : "type" yerine olmasını istenilen veri tipini, "(value)" yerine ise o veri tipine uygulamak istenilen değer yazılmalıdır.
	fmt.Println(x + int(y)) //dönüştürme işlemi sadece buradaki işlem bazında yapılır.
	//y değişkeninin veri tipi hala float64. yani type conversion değişkene yeni bir değer oluşturur. değişkenin asıl veri tipini ve değerini değiştirmez.

	num1 := 106
	str1 := string(num1) //burada integer veri tipini stringer'a çevirdik. normalde type conversion ile integer'ı string'e çeviremiyoruz. fakat buradaki istisna decimal gösterimde 106 j'ye denk geliyor, yani tam sayı.
	fmt.Printf("%v %T\n", num1, num1)
	fmt.Printf("%v %T\n", str1, str1)
}

func main() {
	typeConversionFunction()
}
