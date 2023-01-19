package main

import "fmt"

//for loop syntax : for <init statement (başlangıç değeri)> ; <condition (şart)> ; <post statement (döngünün nasıl devam edeceği)> { code }
//for loop , sadece condition == true olduğunda çalışır. false ise çalışmaz.

func ForLoopFunction() {
	for a := 1; a <= 10; a++ {
		fmt.Println(a)
	}
	//fmt.Println(i) : yazılırsa "undefined error" alınır. çünkü i sadece for loop içerisinde tanımlıdır. yani i 'nin scope 'u sadece for loop içindedir.

	b := 0
	for ; b <= 10; b += 5 { //burada baştaki " ; " tanımlama yapıldı anlamındadır. silinirse error alınır. ayrıca condition statement 'ten önce başka bir statement olduğu da belli olur böylelikle.
		fmt.Println(b)
	}
	fmt.Println(b) //burada i 'nin scope 'u for loop içinde değil, fonksiyon içinde (15. satır) olduğundan bu kod çalışır.

	//for loop yanında sadece şartı/koşulu da belirtebiliriz;
	m := 10
	for m >= 0 { //başa " ; " koyulmasına gerek yok çünkü post statement aşağıda, yani tek bir ifade var.
		fmt.Println(m)
		m--
	}

	for n := 0; n <= 10; n++ {
		if n%3 == 0 {
			continue //bu koşulu sağladığında pas geç ve bulunduğun yerden tekrar for döngüsünün başına geç demek. o yüzden yazdıramıyor(çünkü başa dönüyor) yani bunu sağlayan ifadeyi. continue -->döngünün başına git
		}
		fmt.Println(n)
	}

	for d := 0; d <= 10; d++ {
		if d == 3 {
			break //eşitliği sağladığında for döngüsünün sonuna gidilir yani for döngüsünden çıkılırve kodun devamını yazdırır. break --> döngüden çıkılır
		}
		fmt.Println(d)
	}

	//for loop ile diğer programlama dillerinde olan while döngüsü yazılabilir.
	//while döngüsünde yalnızca koşul sağlanana kadar çalışmasını düşünürüz.
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "çift sayıdır.")
		} else {
			fmt.Println(i, "tek sayıdır.")
		}
	}

	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	} //aslında bu kod while döngüsüne hemen heme eşit çünkü x<10 gibi bir koşul var.x < 10 old sürece sen içerdeki işlemleri yapacaksın (xi arttırarak yazdır) diyoruz
}

//infinite loop : hiçbir sınır, şart ve başlangıç değeri, post statement verilmediğinden sonsuz döngüdür. go derleyici fordan sonra otomatik true varmış gibi algılar

func InfiniteForLoopFunction() {
	//for loop kendi başına da yazılabilir. örneğin ;
	for {
		fmt.Println("Hello World!")
	}

	for i := 1; true; i += 5 {
		fmt.Println(i)
	} //sonsuza kadar çalışır çünkü condition == true 'dur.
}

func main() {
	ForLoopFunction()
	InfiniteForLoopFunction()
}
