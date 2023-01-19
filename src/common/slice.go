package main

import "fmt"

//slice : array 'lerin biraz daha genişletilmiş halidir.
//slice syntax (LITERAL YÖNTEMİ) : tıpkı array 'ler gibidir. ama slice 'lar oluşturulurken eleman sayıları bilinmediği için " [] " boş bırakılır.
//slice syntax2 (MAKE METHOD/YÖNTEMİ) : <name of slice> := make([]<type of slice>, <number of slice>)

func SliceFunction() {
	mySlc := []int{1, 2, 3}
	fmt.Println(mySlc)
	fmt.Println(len(mySlc))

	var mySlc2 []int    //buradaki declare ile aslında slice değil, sadece boş bir slice parçası oluşturuluyor.
	fmt.Println(mySlc2) //yazdırıldığında eleman sayısı bilinmediğinden içersinde hiçbir eleman bulunmayan, boş bir array parçası/ slice döner bize. yani array'ler gibi zero value değerli bir array gelmez.
	//mySlc2[0] = 10  --> normalde bir array 14. satırdaki kod gibi declare edilirse istenilen atama hiçbir sorun olmadan bu satırdaki kod ile yapılır. ama slice 'larda bu şekilde eleman ataması yapılırsa "panic error" alınır. alınan bu error slice 'ın uzunluğunun geçildiğini söylemektedir. çünkü oluşturulan slice 'da herhangi bir eleman yoktu. yani slice bize dönüyor/ belirtiliyor ama oluşturulmamış bir halde üstteki kodda (fmt.Println(mySlc)). make metoduyla slice'ı oluşturalım şimdi de;
	mySlc2 = make([]int, 4) //eğer biz make metodu ile slice oluşturursak, slice elemanlarının zero valueları olur. slice elemanlarına aittir yani zero valuelar.
	fmt.Println(mySlc2)
	mySlc2[0] = 10 //make metodu ile oluşturulan slice 'a bu satırdaki kod ile ilk elemanı atandı.
	fmt.Println(mySlc2)

	var mySlc3 []bool
	mySlc3 = make([]bool, 2) //slice elemanlarının zero value 'ları "false" olarak döner.
	fmt.Println(mySlc3)

	var mySlc4 []int
	fmt.Printf("%#v", mySlc4) //bu kod ile mySlc4' ün veri tipinin nil durumunda oluduğu görülür. yani slice 'a herhangi bir değer atanmamış. boş değil, direkt yok. declare edildi ama oluşturulmadı, yok.
	fmt.Println()
	mySlc4 = make([]int, 0)
	fmt.Printf("%#v", mySlc4) //mySlc4 oluşturulmuş bir slice. slice var ama elemanı yok, boş.
	fmt.Println()
	var mySlc5 []int
	fmt.Println(mySlc5)
}

func SliceFunction2() {
	myArr := [3]int{1, 2, 3}
	fmt.Println(myArr)
	myArr2 := myArr
	fmt.Println(myArr2)
	myArr2[0] = 100
	fmt.Println(myArr2)
	fmt.Println(myArr) //array'imizin diğer birçok programlama dilinde olduğu gibi, referans olarak değil, değerlerini/ kendi elemanlarını kopyalanıyor (pass by value=değerler paylaşılıyor. yani birinin değerinin değişmesi öbürünü etkilemiyor). slice'ların ise referansları kopyalanır;

	mySlc1 := []int{1, 2, 3}
	fmt.Println(mySlc1)
	mySlc2 := mySlc1
	fmt.Println(mySlc2)
	mySlc2[0] = 33
	fmt.Println(mySlc2)
	fmt.Println(mySlc1) //mySlc1 de ilk elemanı 33 olacak şekilde değişti. array'lerde böyle olmamıştı üstte gördüğümüz gibi. slice'lar pass by reference=değer kopyalanmıyor. yani birinin değeri değişirse öbürünün de değişiyor.
}

func SliceFunction3() {
	underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(underArray)  //oluşturulan bu array 'den bir slice elde edelim ;
	mySlc := underArray[2:5] //2 ve 5 underlying array 'deki index numaralarıdır. yani 2 numaralı index 'ten başla ve 5, yani son index 'i dahil etme.
	fmt.Println(mySlc)
	mySlc2 := underArray[:6] //başlangıç indexi olmaması en baştan başla demektir.
	fmt.Println(mySlc2)
	//buradaki slicing/ dilimleme işlemi pyhton'da da aynı şekilde var. bu işlem sayesinde array'den slices oluşturuyoruz.
	mySlc3 := underArray[3:] //bitiş indexinin olmaması en sona kadar git demektir.
	fmt.Println(mySlc3)
	mySlc4 := underArray[:]
	fmt.Println(mySlc4) //hiçbir index belirtilmemesi ise baştan sona hepsini al demektir.
	mySlc[0] = 100
	fmt.Println(mySlc) //mySlc slices'ının ilk elemanı 100 oldu. bu değişiklikten sonra mySlc2 ve mySlc4'ün de 2. elemanları 100 olarak değişir. Yani biz eğer bir slices'da değişiklik yaparsak aynı underArray'den türetilmiş olan diğer tüm slices'larda da değişiklik olur. Bakalım;
	fmt.Println(mySlc2)
	fmt.Println(mySlc4)
	//Sonuç olarak slice elde ederken, array'in split(bölünme) edilmesi gerekir.

	/*mySlc := []int{1, 2, 3}*/ //biz görmesek de arka planda bir mySliceUnderArray oluşturuldu.
	/*fmt.Println(mySlc)*/
	//slices'a yeni bir eleman nasıl ekleriz? append methodu ile. bakalım;
	/*mySlc = append(mySlc, 4, 5) //1.parametre ekleme yapacağımız slice'dır.2. olarak eklemek istenilen elemanlar yazılır sırasıyla.
	fmt.Println(mySlc)*/
	/*mySlc2 := append(mySlc, 4, 5)*/ //yeni bir slice yarattık bu sefer de. //burada oluşturulan mySliceUnderArray  811.satırdakiyle aynı değil. çünkü bu satırda append ile slice'ı genişletiyoruz. bir array'in eleman sayısı değiştirilemediğinden burada mySlc2 oluşturulurken aslında mySlice2UnderArray oluşturuluyor arka planda yani. bu sebeple mySlc'nin elemanı değiştiğinde mySlc2'nin elemanı etkilenmiyor. çünkü bunlar artık farklı underline array'lere bakıyor.
	/*fmt.Println(mySlc2)*/

	/*mySlc[0] = 100
	fmt.Println(mySlc)  //ilk eleman 100 oldu
	fmt.Println(mySlc2) //ilk eleman 1, değişmedi yani.*/

	//şimdi de 2 slice'ın elemanlarını birbirlerine eklemeye çalışalım; mySlc2'yi mySlc'ye ekleyelim;
	/*mySlc := []int{1, 2, 3}
	mySlc2 := []int{4, 5}
	//mySlc = append(mySlc, mySlc2) şeklinde düşünürsen hata alırsın 2. parametreden dolayı. yukarda append kullandığımızda aslında elemanları int veri tipine sahip olan slice'a int ekliyorduk. fakat burada elemanları int olan slice'a, slice int veri tipine sahip olan bir slice eklemeye çalışıyoruz. o yüzden aslında farklı veri tipleriyle işlem yapmaktan hata alırız. yani sen mySlc2 yerine "ezgi" yazsan da hata alırsın. yazacağın şey mySlc'nin veri tipiyle aynı olmalı. şöyle yaparsak sorun düzelir ama :) ;
	mySlc = append(mySlc, mySlc2...) //peki burada aslında ... ile ne yapıyoruz? burada teorik olarak mySlc2'yi elemanlarına ayırdık.
	fmt.Println(mySlc)*/

	//şimdi de bir slice'dan eleman silmeye çalışalım; slicing/dilimleme işlemini kullanacağız;
	/*mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //ilk 2 elemanı silelim mesela;
	fmt.Println(mySlc)*/
	/*mySlc = mySlc[2:] //burada aslında var olan slices'dan yeni bir slices oluşturuyoruz. ilk n elemanı silmek için [n:]
	fmt.Println(mySlc)*/
	//son 2 elemanı silelim mesela;
	//mySlc = mySlc[:len(mySlc)-2] //son n elemanı silmek için [:len(mySlc)-n]
	//fmt.Println(mySlc)
	//sırasıyla önce baştaki 2, sonra sondaki 3 elemandan kurtlulalım;
	mySlc = mySlc[2:]
	mySlc = mySlc[:len(mySlc)-3]
	fmt.Println(mySlc)
}

func main() {
	SliceFunction()
	SliceFunction2()
	SliceFunction3()
}
