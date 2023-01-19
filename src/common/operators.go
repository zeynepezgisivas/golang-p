package main

import "fmt"

func OperatorFunction() {
	x, y := 10, 20
	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)

	fmt.Printf("%T %v\n", x+y, x+y) //10 ve 20 operand, + operator
	fmt.Printf("%T %v\n", x-y, x-y)
	fmt.Printf("%T %v\n", x*y, x*y)
	fmt.Printf("%T %v\n", x/y, x/y) //x ve y integer olduğundan bölümleri de integer'dır. integer olduğu için de sonucun tam, yani integer kısmı alınır. "x, y := 10.0, 20" dersen, yani değişken olarak yazarsan "type mismatched" hatası verir. onun yerine şöyle yapalım;

	fmt.Printf("%T %v", 9.0/3, 9.0/3) //değişken olarak atamadığımız için float ile integer 'ı işleme sokabiliriz.

	z := 5.0 / 2
	fmt.Printf("%T %v\n", z, z)   //float64(2.5) çıktısını alırız.
	fmt.Printf("%T %v", x%y, x%y) //% kalanı (remainder) gösterir.
}

func OperatorFunction2() {
	x := 10
	x++        //"x = x + 1" ifadesine eşittir. (increment)
	x--        //"x = x - 1" ifadesine eşittir. (decrement)
	println(x) //"fmt.Println(x++)" veya "fmt.Println(x--)" yapılamaz.

	//postfix : değişkenden sonra yazılan ++ veya -- gibi işlemlerdir. peki neden önce değil de sonra? çünkü increment ve decrement birer statement'dir. programlamada 1 satırda yalnızca 1 statement bulunacağı için "fmt.Println(x++)" yapılamaz. ("Println()" de bir statement'tir.)
	//statement ile expression arasındaki fark ; statement programda bir şey yapılmasını belirten en küçük ifadeyken, yani işlemken expression ise yalnızca açıklamadır.

	//print("Hello"), x = 1 : (statement)
	//5 * 5 : (expression)
	//print(5 * 5) : (expression statement)

	y := 50
	y = y - 10 //assignment statement
	//x -= 10 kısayoludur. assignment operation. atamayla çıkarma işlemi aynı anda yapıldığı için isim farklı.
	fmt.Printf("%v, %T", y, y)

	F := -40
	K := float64(F-32)/1.8 + 273 // type conversion ile float64 tipine çevirdik ki 1.8'e bölünebilsin. float64 tipindeki sayıyı 273 ile toplayabilirim çünkü go compiler 273ü de float64müş gibi düşünüp işleme sokabiliyor çünkü 273 typeless constant. go da işlemi yapmak için float64 tipinde alıyor 273ü
	fmt.Printf("%v, %T", K, K)
}

func OperatorFunction3() {
	//comparison operators : == , != , < , > , <= , >= : karşılaştırma operatörleri sonucunda bool bir ifade elde edilir.
	x, y := 8, 9
	//x, y := "a", "b" böyle de karşılatırma yapılabiliyor çünkü buradaki string ifadeler aslında nümerik ifadelerdir. (ascii code table: a = 97, b = 98)
	//x, y := 8, 9.0 yazarsak "mismatched types" hatası alırız.
	//x, y := 8, true yazarsak "mismatched types" hatası alırız.
	//x, y := 8, "b" yazarsak "mismatched types" hatası alırız.
	//yani birbirlerine atanamayan değerler comparison operators ile karşılaştıramayız. farklı veri tipleri de birbirine atanamayacağından karşılaştırılamaz.
	fmt.Printf("%T, %v\n", x == y, x == y)
	fmt.Printf("%T, %v\n", x != y, x != y)
	fmt.Printf("%T, %v\n", x < y, x < y)
	fmt.Printf("%T, %v\n", x > y, x > y)
	fmt.Printf("%T, %v\n", x <= y, x <= y)
	fmt.Printf("%T, %v\n", x >= y, x >= y)

	//logical operators : && (conditional and/ şartlı ve) , || (or) , ! (not) : mantıksal operatörlerde 2 veya daha fazla T/F, yani bool ifadelerin durmunu karşılaştırırız.
	a, b := 1, 2
	set1 := (a == b)
	set2 := (a < b)
	//set3 := 40 yazıp karşılaştırmaya çalışsan 40 bool değer olmadığı için hata alırsın.
	set3 := true
	fmt.Printf("%T, %v\n", set1, set1)
	fmt.Printf("%T, %v\n", set2, set2)

	fmt.Printf("%T, %v\n", (set1 && set2), (set1 && set2)) //false && true --> false
	fmt.Printf("%T, %v\n", (set1 || set2), (set1 || set2)) //2 durum da false ise --> false
	fmt.Printf("%T, %v\n", !set3, !set3)

	//yani comparison operators ve logical operators sonucunda  T/F diye belirli bir koşulun doğruluğunu yada yanlışlığını ifade ediyoruz.
}

func main() {
	OperatorFunction()
	OperatorFunction2()
	OperatorFunction3()

}
