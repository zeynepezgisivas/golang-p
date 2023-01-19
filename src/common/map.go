package main

import "fmt"

//go dilinde map veri tipini her bir değerin bizim açımızdan anlamı farklıysa oluşturulur. map 'le ilgili bilinmesi gerken temel şey, map 'te  bulunan her bir değerin özel bir anlama gelmesidir. her birinin klasik slice ve array 'deki gibi 0, 1 gibi indexlemek yerine farklı key 'lerle mapliyoruz yani haritalandırıyoruz/birbirinne eşliyoruz.
//key-value çiftleri map veri tipini oluşturur. yani verilerin tamamı ayrı ayrı özel bir anlam ifade eder. mesela aşağıdaki örnekde 40 verisi ahmet 'in yaşına gelen bir değerdir. yani 0.indexteki bir değer değil. veya ayşe 1.indexteki bir değer değil. ayşe 17 değerine denk gelen bir key'imiz.
//array ve slices 'larda indexler yalnızca integer ile belirtiliyordu. ancak map 'lerde böyle değildir. biz index 'i belirten key değeri için veri tipi olarak string/int/array kullanabiliriz. slice kullanamayız çünkü slice'ın belirli bir eleman sayısı yok.
//map syntax : <map name> := map[ type of key ]< değerlerin alacağı veri tipi > { map 'e ait olan key-value çiftleri }
//key - value aynı veri tipinde olmasına gerek yok. ancak kendi içlerinde key'ler ve value'lar belirtilen/aynı veri tipinde olmak zorunda tabi.

func MapFunction() {
	//örneğin bazı kişilerin yaşlarını veri olarak tutmak isteyelim ;
	myMap := map[string]int{ //LITERAL METHOD
		"Ahmet":   40,
		"Hilal":   17,
		"Selim":   14,
		"Mustafa": 70,
	}
	fmt.Println(myMap)
	fmt.Println(myMap["Ahmet"])
	fmt.Println(myMap["Ahmet"], myMap["Selim"]) //bu şekilde çoklu olarak da yazdırabilir.

	fmt.Println(myMap["Ahmet Mehmet"])
	fmt.Println(myMap["Sıla"]) //23. satır ve bu şekilde olmayan bir değerleri yazıldığında error alınmaz ancak o veri tipinin zero value 'si döner.(burada int olduğundan 0)

	//fmt.Println(myMap[0]) denirse direkt error alınır çünkü [] içine yazılan bir key değeridir ve key değerlerinin string olması isteniyor "myMap" tanımlanmasında. bu sebeple 0 yazılamaz.

	//yukarıdaki örnekteki kişilerin evlilik durumlarını veri olarak tutmak isteyelim ;
	isMarried := map[string]bool{
		"Ahmet":   true,
		"Hilal":   false,
		"Selim":   false,
		"Mustafa": true,
	}
	fmt.Println(isMarried)

	myMap2 := make(map[string]int) //MAKE METHOD
	fmt.Println(myMap2)
	fmt.Println(myMap2["Test"])

	//yukarıdaki örnekteki kişileri ve aldıkları notları eşleştiren bir map yapalım ;
	studentGrades := map[string]int{
		"Ahmet":   80,
		"Hilal":   29,
		"Selim":   72,
		"Mustafa": 0,
	}
	fmt.Println(studentGrades)
	fmt.Println(studentGrades["Ahmet"])
	fmt.Println(studentGrades["Mustafa"])

	fmt.Println(studentGrades["Mustfa"]) //bu şekilde olmayan yada yanlış yazılan key değerleri için value olarak zero value döner. yani aslında dönen değer mustafa 'nın notu olan 0 değildir.
	//peki key değerine bakılan elemanın gerçekten map 'in içinde olup olmadığını nereden bileceğiz? "Mustfa" da yazılsa, "ezgi" de yazılsa 0 dönecek ama map 'in içinde değiller? şöyle anlaşılır ;

	value, ok := studentGrades["Mustfa"] //map 'in içerisindeki bir eleman bu şekilde alınmaya çalışılırsa istenilen öğrenilebilir.
	fmt.Println(value, ok)               //"ok" ifadesi "Mustfa" key 'inin map 'de olup olmadığını gösterir. "ok" ifadesi "false" ise aranan key yoktur.

	_, ok = studentGrades["Muustafa"] //genelde kontrol etme kısmında value ifadesi yerine _ kullanılır. çünkü buradaki asıl amaç değeri değil, key 'e sahip olan elemanın map 'in içinde olup olmadığını bilmektir.
	fmt.Println(ok)

	//studentGrades map 'ine eleman ekleyelim ;
	studentGrades["Can"] = 55
	fmt.Println(studentGrades)

	//studentGrades map 'inden eleman çıkartalım gömülü delete metodu ile ;
	delete(studentGrades, "Selim") //delete(hangi mapten eleman silmek istediğim, hangi key değerini silmek istediğim)
	fmt.Println(studentGrades)

	//studentGrades map 'inin uzunluğunu, yani kaç key-value çifti olduğunu öğrenelim ;
	fmt.Println(len(studentGrades))

	//map 'ler de slices 'lar gibi array 'lerden farklı olarak değeri değil referansı paylaşır.
	sg := studentGrades //sg yeni bir map olsun ama var olan diğer map'imizden kopyalamış olalım. ikisi de aynı map'i çıkartır tabi.
	fmt.Println(studentGrades)
	fmt.Println(sg)
	delete(sg, "Mustafa") //sg map 'inden "Mustafa" silindi.
	fmt.Println(sg)
	fmt.Println(studentGrades) //ama studentGrades map 'inden de silinmiş oldu "Mustafa".
	//yani bir map 'te değişiklik yapıldığında, o map 'in tüm kopyalarında da o değişiklik yapılmış olur. yani map 'ler de slices 'lar gibi pass by reference 'dırlar.

	//for-range yapısı ile studentGrades map 'indeki key-value çiftlerini yazdıralım ;
	for k, v := range studentGrades {
		fmt.Println(k, v)
	}
}

func MapFunction2() {
	//örneğin for-range yapısı ile yazar ve yazarlara ait kitapların isimlerini yazdıralım ;
	writerBook := map[string][]string{ //map 'in key değeri string, value değeri slice ([]string) oldu.
		"Yaşar Kemal":     []string{"İnce Memed", "Yer Demir Gök Bakır"},
		"Sabahattin Ali":  []string{"Kuyucaklı Yusuf", "Kürk Mantolu Madonna", "Değirmen"},
		"Haruki Murakami": []string{"1Q84", "Kumandanı Öldürmek", "Dans Dans Dans"},
	}
	fmt.Println(writerBook)
	fmt.Println(writerBook["Sabahattin Ali"])
	fmt.Println(writerBook["Haruki Murakami"][0])

	//önce key-value 'lar, sonra slice 'lar için iç içe 2 for döngüsü yazılmalı ;
	for key, value := range writerBook {
		fmt.Println("Yazarımız:", key)
		fmt.Println("bazı kitapları aşağıdadır:")
		for i, v := range value { //value 'dan alındı çünkü key-value 'daki value 'lardan yapılacak 2.for loop 'u.
			fmt.Println("\t", i+1, v) //indexi 0'dan başlatmamak için i+1 dendi.
		}
	}
}

func main() {
	MapFunction()
	MapFunction2()
}
