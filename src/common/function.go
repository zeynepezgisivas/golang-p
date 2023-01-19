package main

import "fmt"

func Function() {
	var x, y, sum int
	x = 5
	y = 10
	sum = x + y
	fmt.Printf("%d ve %d toplamı %d'dir.\n", x, y, sum)

	x = 8
	y = 10
	sum = x + y //bu tekrar yazılmazsa toplam olarak yine 15 'i gösterir. yani verilen değişkenlere göre işlem tekrar tekrar yapılmalı. fakat toplama işlemleri tekrar tekrar yapılacağına, toplama işlemine ait bir fonksiyon oluşturulursa ve her ihtiyaç duyulduğunda çağırılırsa program daha modüler hale gelmiş olur.
	fmt.Printf("%d ve %d toplamı %d'dir.\n", x, y, sum)
}

//function syntax : func funcName(params) return type { code }
//her zaman değişken döndürmek ve parametre kullanmak zorunluluğu yoktur.
//return type olduğunda "return" anahtar kelimesi kullanılmak zorunda. çünkü bu kelime bir değişken döndüğünü belirtir.

func sum(x, y int) int {
	return x + y
}
func sum2(x, y int) {
	fmt.Println(x + y)
}

//bir fonksiyonda kullanılan değişkenlerin scope 'u o fonksiyon içerisindedir. yani { } arasındadır.
//sum ve sum2 fonksiyonlarında bulunan x ve y 'ler de sadece kendi fonksiyonlarının içerisinde geçerlidirler. func main içinde tek başına "fmt.Println(x)" yazılırsa "undefined error" alınır.

// func sum(x int, y int) = func sum(x, y int)

func hello() {
	fmt.Println("hello")
}
func hello2(name string, age int) {
	fmt.Printf("adım %s, yaşım %d\n", name, age)
}

func result(grade int) string {
	if grade >= 50 {
		return "geçtiniz"
		fmt.Println("if koşulu içindeeyim") //unreachable code : çünkü return görünce direkt fonksiyondan çıkıyor.
	}
	return "kaldınız"
	//fmt.Println("fonksiyon içindeeyim") dersen zaten direkt hata alırsın çünkü sen bir fonksiyondasın ve fonksiyonun en son statement'i return olmak zorundadır eğer bir return type var ise. yani kapsamı fonk olan return den sonra farklı bir statement olamaz
}

func main() {
	Function()
	fmt.Println(sum(5, 2)) //sadece "sum(5, 2)" yazılması bir işe yaramaz. çünkü sum fonksiyonu sadece toplamı bize geri dönüyor. bu yüzden toplamın ne yapılacağı belirtilmelidir.
	fmt.Println(sum(3, 9))
	fmt.Println(sum(1, 6)) //yazılan fonksiyonlar tekrar tekar çalıştırılabilir.

	//return vs print arasındaki fark;
	sum(2, 5) //hiçbir çıktı vermez
	z := sum(2, 5)
	fmt.Println(z)
	sum2(6, 11)

	hello()
	hello2("ezgi", 23)
	hello2("Ezgi", 22) //yazılan fonksiyonlar tekrar tekar çalıştırılabilir.

	name := "Yağmur"
	age := 23
	fmt.Printf("adım %s, yaşım %d", name, age) //merhaba2 fonksiyonu haricinde main fonksiyonunda tanımlandığı için bu değişkenler bunları da yazdırdı.

	fmt.Println(result(45))
}
