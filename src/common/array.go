package main

import "fmt"

//array syntax : <name of array> := [ length of array ]{ values }
//array 'lerdeki tüm value 'lar aynı veri tipinde olmalıdır.

func ListFunction() {
	city1 := "istanbul"
	city2 := "new york"
	city3 := "paris"
	city4 := "milano"
	fmt.Println(city1, city2, city3, city4)
	//burada değişkenlerin sayısı 4 değil 100 olsaydı hem teker teker tanımlamak hem de yazdırmak çok zor olurdu. bu sebeple array oluştururuz.
}

func ArrayFunction() {
	cities := [4]string{"istanbul", "new york", "paris", "milano"}
	//cities diye 4 birimlik bir yer oluşturulmuş oldu. yani 4 tane kutu yanyana düşünülebilir. her bir kutuya da bir şehir ismi gelir.
	fmt.Println(cities)

	//array 'e ait elemanlar teker teker yazdırılabilir;
	fmt.Println(cities[0]) //zero based index
	fmt.Println(cities[3])

	//bir array 'i yazdırırken her zaman  array 'in kaç elemandan oluşacağını bilemeyeceğimizden şöyle tanımlama yapmak mümkündür ;
	cities2 := []string{"istanbul", "new york", "paris", "milano"}
	//cities2 := [...]string{"istanbul", "new york", "paris", "milano"} ifadesi de aynı sonucu verir.
	fmt.Println(cities2)
	fmt.Println(cities2[0])
	fmt.Println(cities2[3])
	fmt.Println(len(cities2)) //array uzunluğu böyle yazdırılır.

	//array 'lerin eleman sayısı bilinmediğinde yada çok fazla sayıda eleman olduğunda, array üzerinde bir for loop ile elemanlar teker teker yazdırılabilir ;
	cities3 := [4]string{"istanbul", "new york", "paris", "milano"}
	for i := 0; i < len(cities3); i++ {
		fmt.Println(i, cities3[i])
	}
	//array 'lerin istenilen bir elemanı aşağıdaki gibi değiştirilebilir;
	cities3[0] = "belgrad"
	fmt.Println()
	for i := 0; i < len(cities3); i++ {
		fmt.Println(i, cities3[i])
	}

	//genel olarak bir array oluştururken o array 'e gelecek elemanları bilemeyiz. örneğin bir veri tabanı vs. gelir veriler.
	//bu sebeple başlangıç değeri verilmeyen, yani başlangıç değeri kendiliğinden olan array  nasıl oluşturulur ona bakalım;
	var myArray [5]int
	fmt.Println(myArray) //burada int değerler için zero value 0 olduğundan, 0 'lardan oluşan 5 elemanlı bir array döner.
	//yani bir array oluşturulduğu zaman, o array içerdiği eleman sayısı kadar zero value verilerek oluşturulur.
	myArray[0] = 100              //ilk elemana 100 değeri verilmiş oldu.
	myArray[len(myArray)-1] = 200 //array 'in son elemanına 200 değeri verilmiş oldu.
	fmt.Println(myArray)

	//array 'lerde bulunan tüm elemanları okumak ve tamamına işlemler yaptırmak istenir. bunun için örneğin 0-9'a kadar olan rakamların karelerini alalım;
	myArray2 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//karelerini alma işlemini ayrı bir fonksiyonda yapalım;
	myArray = mySquare(myArray2) //first class functions: bir fonksiyonu bir değişkene atadık. yani artık o değişkene fonksiyon muamelesi yapabiliriz.
	fmt.Println(myArray2)
	//for i yapısı yerine for range yapısını da kullanabiiriz; syntax: for <index>, <value> := range <array ismi> (range dizimizin son indexine kadar olan elemanları yazdırmamızı sağlıyor)
	cities := [4]string{"istanbul", "new york", "paris", "milano"}
	for index, city := range cities {
		fmt.Println(index, city)
	}
	fmt.Println()
	//index'siz de yazdırabiliriz;
	for _, city := range cities {
		fmt.Println(city)
	}

	/*myArr := [3]int{1, 2, 3}
	myArr2 := [...]int{1, 2, 3, 4}
	var myArr3 [5]int //5 elemanlı ve elemanları zero value olan bir array'dir. yani eleman belirtilmediğinde elemanlar zero value alınır.
	fmt.Println(myArr)
	fmt.Println(myArr2)
	fmt.Println(myArr3)
	fmt.Println(len(myArr))
	fmt.Println(len(myArr2))
	fmt.Println(len(myArr3))*/

	/*var myArr [4]int
	fmt.Println(myArr)
	myArr[0] = 5
	fmt.Println(myArr)*/
}

func ArrayFunction2() {
	//var myArray [5]int
	//var myArray2 [4]int
	//go programlama dili için array uzunluğu da veri tipine dahildir.
	//örneğin myArray ve myArray2 'nin veri tipleri birbirleriyle aynı değil. ispatlayalım ;
	/*if myArray == myArray2 {
		fmt.Println("MESAJ")
	}*/ //bu if koşulu çalıştırılmaya çalışılırsa "mismatched error" alınır. zaten 2 veri tipi eşit olsaydı if koşulu true olacağından "MESAJ" çıktısı alınırdı.
}

func mySquare(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}

func main() {
	ListFunction()
	ArrayFunction()
	ArrayFunction2()
}
