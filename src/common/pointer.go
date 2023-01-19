package main

import "fmt"

//pointer : başka bir değişkenin adresini tutan, hafızadaki adresini gösteren değişkendir. bir belirteç, işaretleyici olarak görev yapar ve çoğunlukla referans tutmak için kullanılan bir yapıdır. peki bu ne demek? pointer ile bir değişken tanımladığını varsayalım. bu değişkeni farklı bir değişkene atadığımız zaman, bir değişken diğer bir değişkenin adresini tutar. yani Bir değişkende değişiklik yapıldığı zaman, diğer değişken bundan etkilenir (pass by reference)
//& ---> address operator. bu operatörü bir değişkenin önüne konarak o değişkenin hafızadaki yeri/adresi gösterilebilir.

func PointerFunction() {
	name := "Ezgi"
	fmt.Println(name)
	fmt.Println(&name) //--> çıktı anlamsız bir şeydir ve bu şey "name" değişkeninin adresidir. yani bu çıktı, "Ezgi" değerine sahip "name" değişkeninin bulunduğu yerin hafızadaki adresidir.

	x := 22
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", &x, &x) //pointer 'ın veri tipi ve değerini verir.

	//var y int = &x denirse veri tipi uyuşmazlığından dolayı error alınır. çünkü y 'nin veri tipi int deniyor ama atanan &x 'in veri tipi *int 'dir ve ikisi aynı şey değildir.
	y := &x //y artık veri tipi *int olan bir pointer 'dır.
	// note : * =asteriks operatörü pointer 'ın veri tipini belirtir. yani pointer 'ların da veri tipleri değişebiliyor.
	fmt.Printf("%T, %v\n", y, y)

	z := &name //z bir pointer 'dır.
	fmt.Printf("%T, %v\n", z, z)

	d := 22
	fmt.Println(d)           //22
	fmt.Println(&d)          //22'nin adresi
	fmt.Println(*(&d))       //yine 22 sonucunu döner. burada asteriks operatörünün yaptığı iş dereferencing. yani &x bir pointer ve adres gösteriyor ya bize, * operatörü ise bu adresteki ifadenin değerini bize tekrar dönüyor.
	fmt.Println((&*(&d)))    //dereferencing
	fmt.Println((*(&*(&d)))) // 22'nin adresi.

	fmt.Println(3 * 5)

	m := 10
	n := m
	fmt.Println(m, n)
	m = 5
	fmt.Println(m, n) //dönen sonuçta m değişir fakat n değişmez. sebebi normalde go'da değişkenlerin pass by value/ değer paylaşmalarıdır. o nedenle m'nin değeri değiştiğinde n aynı kalır. ancak pointer, referansları paylaşmanın kısayoludur diyebiliriz;
	//burada değer kopyalaması yapılıyor: bir kutu düşün içinde 10 var. sonra başka bir kutu oluşturuyorsun hafızada ve ona da 10 koyuyorsun.

	a := 10
	b := &a //a 'nın adresini b 'ye atıyoruz. yani b bir pointer oldu.
	fmt.Println(a, b)
	fmt.Println(a, *b) //*b de 10'dur çünkü b, a 'nın bulunduğu yeri gösteriyor ve * sayesinde de a 'nın bulunduğu yerdeki değeri alıyor.
	*b = 3             //dereferencing //a 'nın olduğu yerdeki değeri (yani b 'yi) 3 yapıyoruz. //b bir pointer ve bu pointer a 'nın adresini tutuyor. *b olduğunda a 'nın adresinin tutulduğu yerin değerine dönüşmüş oldu ve 3'e atadığımız için 3'e dönüştü.
	fmt.Println(a, *b) //3 olur ikisi de.
	//burada a 'nın içinde 10 var. sonra b diye bir değer var ve o dışardan a kutusunu tutuyor.
	c := &a
	*c = 5 //dereferencing
	fmt.Println(a, *b, *c)

	h := [4]int{1, 10, 100, 1000}
	j := h
	fmt.Println(h, j)
	j[0] = 3
	fmt.Println(j) //j 'nin 0.indeksi değişir haliyle.
	fmt.Println(h) //h değişmedi. çünkü buradaki array'ler pass by value/değer paylaşır.

	r := []int{1, 10, 100, 1000}
	s := r
	fmt.Println(r, s)
	s[0] = 3
	fmt.Println(s)
	fmt.Println(r) //r değişir. çünkü slice pass by reference
	//peki bunların pointer ile ilgisi ne? pointer'lar değişkeni değil değişkenin adresini tutardı. slice'lar underlying array'lerden oluşuyor idi. biz slice'ları birbirlerine kopyalayarak aslında oradaki underlying array'leri kopyalıyoruz. yani buradaki 2 slice da gidip aynı array'e referans veriyor, point ediyor. yani aslında biz görmesek de hook altında go derleyicisi yeni bir slice oluştururken yaptığı işlem aynı underlying array'e point atmak, yani tutmak.o nedenle biz bir slice'ı değiştirdiğimizde diğeri de değişir.

	w := 5
	fmt.Println(w) //bu sayının 2 katını yazdıran basit bir double() func yazalım dışarda;
	double(w)      //double fonksiyonunun argümanı w 'dir ve fonksiyonda num parametresinin yerine geçer. yani değişken fonksiyon çağrıldığı zaman argüman, oluşturulduğu zaman parametredir. mesela x=5 olsun. fonksiyondaki num da 5 olur tabi. ancak num=5 olurken buradaki x num'a gitmiyor. x'in kendisi değil bir kopyası num değerine gidiyor yani. o nedenle go pass by value olan bir dildir. yani bir fonksiyonu çağırırken ona gelen parametre argümanın bir kopyasıdır.ancak slice'lar için pass by reference demiştik. bu örneğin aynısını slice'lar ile yapalım 2 satır sonra;
	fmt.Println(w) //w 'yi yine 5 diye yazdı. sebebi: go pass by value şeklinde. yani değer geçen bir programlama dilidir. yani bir fonk çalıştırıldığı zaman aldığı parametre, argümanın bir kopyasıdır.

	mySlc := []int{1, 10, 100}
	fmt.Println(mySlc) //func doubletwo olsun fonksiyonu ve parametre olarak slice alacak tabii
	doubletwo(mySlc)
	fmt.Println(mySlc) //üstte olduğu gibi yine 2yle çarpılmadan önceki halini yazdıracağını düşünürüz ama 2yle çarpılmış halini yazdırır. ama hani go pass by value idi? yani değeri kopyalıyordu? ama burada orijinal slice'ın da değeri değişti. ama buradaki olay slice'ların kendi yapısıyla ilgili. doubletwo fonksiyonunda fonksiyon slice'ın kopyasını değiştirdi. ama orjinali de değişti bu satırda çünkü slice'ların arka planda bir pointerları vardı.yani aynı underlying array'leri tutuyorlardı ikiside. bu sebeple slice'ın biri değiştiğinde o underlying array de değişiyordu. underlying array'in değişmesi slice'ın da değişmesine sebep oluyor.yani aslında referansın değiştiğini düşünebiliriz ama hayır burada da değer değişiyor. eğer çağırılan bir fonksiyona parametre olarak bir argüman gönderiyorsak o argümanın kendisi değil kopyası gider fonksiyonun kendisine.
	//biz ama slicelar için pass by reference demiştik. burada söyleme çalışılan şey şu: bir slice'ı diğerine eşitlerken/atarken bu ifadeyi kullanıyoruz (1002.satır).ancak go'nun geneline bakıldığında fonksiyolarda argümanın kopyasını alıyor. bu 2si birbirine karıştırılmasın.

	//array için yapalım bir de aynı şeyleri; doublethree olsun fonksiyonu;
	myArray := [3]int{1, 10, 100}
	fmt.Println(myArray)
	doublethree(myArray)
	fmt.Println(myArray) //orijinal array aynı şekilde kalır. çünkü pass by value.peki neden slice gibi olmadı? çünkü kopyasını gönderdiğimiz arrayler birbirlerinden bağımsız. slicelar gibi aynu underlying arraye tutunmuyorlar, farklı underlying arrayleri var. o yüzden fonksiyonlar birini değiştirdiğimizde orijinal değer sabit şekilde kalıyor.
	//yani go pass by value şeklinde çalışıyor slice'ı aldığında dahi. ama sliceın kendi yapısından dolayı son değer değişiyor. peki bunların pointer ile alakası ne?
	//biz pointerları neden kullanıyorduk? bir değeri olduğu yerde değiştirebilmek için kullanıyoruz. peki bunun bize avantajı nedir? örneğin verimiz çok büyük bir array yada slicedan oluşsun. bunun işlem yaparken kopyasını almak efektif bir çözüm değil haliyle. kopyasını almak yerine olduğu yerde değiştirebilmek, yani pointerını alabilmek bizim açımızdan daha avantajlı olur. bu nedenle pointer kullanıyoruz. bir veriyi kopyalamaya gerek kalmadan, bulunduğu yerde değiştirebilmek için.

	u := 5
	fmt.Println(u) //bu kez u 'yu değil de u 'nun adresini alalım. adresini pointer ile belli ediyorduk.pointeri veri tipine dönüştürmek için * koyacağız.(*int bir pointerdır.) o yüzden fonksiyonun parametresini num *int yapıyoruz (func doublefour(num *int) );
	doublefour(&u) //pointer bir argüman göndermek için doublefour fonksiyonuna &y yazıyoruz burada.(& bir değeri pointera çeviriyordu.).yani artık bir pointerı paramtere olarak gönderiyoruz.yani u 'nun kendisi değil adresi gönderiliyor burada.
	fmt.Println(u) //u değeri de değişmiş oldu. yukarda adreste değişiklik yaptırdık çünkü.yani go'da herhangi bir şekilde değerin kendisini olduğu yerde değiştirebilmek için pointer kullanıyoruz.

	//Go, pass by value, yani değer tutucu olarak çalışan bir programlama dilidir. Referans tutmak için pointer gibi yapılardan faydalanırız. Yaptığımız örnekte, fonksiyona gönderilen değişkenin bir value kopyası gönderiliyor. Daha sonra fonksiyon içerisinde çarpıldıktan sonra o fonksiyon içerisinde x değişkenin değeri bastırılıyor. Bu işlemler bittikten sonra ise x değişken kopyası yok oluyor ve x değeri tekrar eski haline dönüyor. Go kendi bir değer tutucu olarak çalışmasına karşın, slice, pointer gibi yapıları kullanarak OOP mantığında çalışma biçimi ortaya çıkarmıştır.
}

func double(num int) {
	num *= 2
	fmt.Println(num)
}
func doubletwo(slc []int) {
	for i := 0; i < len(slc); i++ {
		slc[i] *= 2
	}
	fmt.Println(slc)
}
func doublethree(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
	fmt.Println(arr)
}
func doublefour(num *int) { //pointer method(parametre olarak pointer veri tipi alıyor)
	fmt.Println(num)  //y'nin adresini yazdırıyor.çünkü pointer/adresini gönderiyor doublefour(&y) diyerek.
	*num *= 2         //adresteki değeri böyle yakalıyorduk/derefer ediyorduz yani
	fmt.Println(*num) //adresteki değeri yazdırır
}

func main() {
	PointerFunction()
}
