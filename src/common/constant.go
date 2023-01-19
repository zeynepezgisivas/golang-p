package main

import (
	"fmt"
	"math"
)

// constant (sabit) : programın çalışma süresi boyunca değişmeyen veri değerleridir. pi sayısı, yıllık faiz oranı vb. değerler constant 'a örnektir.
// syntax : const - <name of constant> - <type of constant>

func ConstantFunction() {
	//constant örnekleri ;
	const x int = 2
	const y float32 = 3.4
	const z string = "test"
	const t bool = false

	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
	fmt.Printf("%T %v\n", t, t)

	r := 3.0
	const pi float64 = 3.14
	areaofcircle := pi * (math.Pow(r, 2))
	fmt.Println(areaofcircle)

	const h int = 2
	//h = 5, h++, h = h + 1 yazılırsa "cannot assign error" alınır. çünkü constant'a sadece 1 kez değer ataması yapılır, sonrasında yeni değer atanamaz.
	//h = h yazılsa bile aynı error alınır çünkü h bir constant ve h'a tekrar başka bir değer atanamaz. tek bir sefer değer atanır ve o değer değişmez.
	fmt.Printf("%T %v\n", h, h)

	//var a = math.Pow(3, 4) //variable : run time
	//const a = math.Pow(3, 4) //constant : compile time   direkt hata alırsın çünkü Pow() bir fonksiyondur ve fonksiyonlar çalışma ortamında (run time) hesaplanırlar. ama const a compile time dı ya, o yüzden çalışamıyorlar. böyle yapmaya çalışırsan initialize hatası alırsın. yani ilk değerini alamıyorum diyor. çünkü compile time da istiyorum diyor ilk değeri ama fonk run time da çalışıyor. yani compile time da alamıyor.
	//fmt.Printf("%T %v\n", a, a)

	/*const d : hata alırsın çünkü bir sbt oluşturduğunda değer ataması da yapmak zorundasın. zaten değeri olmayan sbt oluşturmak mantıksız olurdu.
	fmt.Printf("%T %v\n", d, d)*/

	f := 3
	const d = 3
	//const d = f : hata alırsın çünkü değişkeni sabite atamaya çalışıyorsun.yani run time da oluşacak değeri compile time da istediimiz değere atarsak hata alırız.
	fmt.Printf("%T %v\n", f, f)
	fmt.Printf("%T %v\n", d, d)

	const v, w = 3, 5
	fmt.Printf("%T %v\n", v, v)
	fmt.Printf("%T %v\n", w, w)

	const a = 1
	var b = 3
	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", a+b, a+b) //aynı veri tipinde olduklarından run time'da toplayabilir.
}

func TypelessConstantFunction() {
	//typeless constant : belirli bir veri tipine sahip olmayan sabittir.
	//typeless constant örnekleri ;
	const x = 5
	fmt.Printf("%T %v\n", x, x) //go programlama dilinde veri tipi çok önemli olduğundan print edildiğinde, go buradaki sabite varsayılan bir veri tipi atıyor.

	const y int8 = 3
	var z int16 = 12
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
	//fmt.Printf("%T %v\n", y+z, y+z) : bu şekilde error alınır çünkü veri tipleri farklıdır.

	const p = 3
	var u int16 = 12
	fmt.Printf("%T %v\n", p, p)
	fmt.Printf("%T %v\n", u, u)
	fmt.Printf("%T %v\n", p+u, p+u) //bu şekilde hata vermez çünkü go hatasız işlem gerçekleştirmek için typeless constant'a sadece bu satıra yönelik type conversion yapıyor. yani int16(p)+u yapıyor.

	//yani typeless constant, go typeless constant'ı kullanmadığında typeless durumundadır. ama kullanıldığında varsayılan yada ihtiyaca yönelik bir veri tipi atıyor.

	const h = int16(4.8 + 5.2)
	fmt.Printf("%T %v\n", h, h)

	/*const n int32 = 3
	const m float32 = 5.6
	fmt.Printf("%T %v\n", n, n)
	fmt.Printf("%T %v\n", m, m)
	fmt.Printf("%T %v\n", n+m, n+m) : typelar farklı olduğundan bu şekilde hata alınır tabi. typeless constant haline getirelim;  */

	const n = 3
	const m = 5.6
	fmt.Printf("%T %v\n", n, n)
	fmt.Printf("%T %v\n", m, m)
	fmt.Printf("%T %v\n", n+m, n+m) //burada, bu satır bazında type conversion yapılıyor.

	const myAge = 22
	fmt.Printf("%v, %T", myAge, myAge)
}

//const x = 12

func FourteenthFunction() {
	/*const x = 24 //dışardaki değeri değil fonk içindeki değeri yazar. shadowing kavramı budur.gölgede brakır yani dışardakini. değişkenlerde de geçerlidir bu.
	fmt.Printf("%v, %T", x, x)*/

	const x = 4
	const y = 5.4
	fmt.Printf("%v, %T", x+y, x+y) //x'in verii tipi bu satır bazında float64 alınıyor. çünkü x typeless constant

	const m float64 = 6.4
	n := 4 + m //4 typeless constanttır. burada toplanabilmek için float64 tipi olarak alınır
	fmt.Printf("%v, %T", n, n)
}

func main() {
	ConstantFunction()
	TypelessConstantFunction()
}
