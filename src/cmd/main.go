// package clause: yazılan kodun çalıştırılabilir olması için belli başlı paketler vardır ve bunlardan biri "main" paketidir.
package main

//import statement: kullanılacak hazır paketler "import" komutuyla kendi kodumuza çağırılır.
import (
	"fmt"
	"strings"
)

func TwentynightFunction() {
	//5 string elemandan oluşan bir array ve 5 integer elemandan oluşan bir slice oluşturup index numaralarıyla birlikte gösterelim;
	/*myArr := [5]string{"new york", "edinburgh", "paris", "londra", "roma"}
	for index, value := range myArr {
		fmt.Println(index, value)
	}
	fmt.Println()
	mySlc := []int{1, 2, 3, 4, 5}
	for index, value := range mySlc {
		fmt.Println(index, value)
	}*/

	//[0,1,2,3,4,5,6,7,8] slice'ından [0,1,2,3], [4,5,6], [6,7,8] ve [2,3,6,7] slice'larını oluşturalım;
	/*mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	mySlc1 := mySlc[:4]
	fmt.Println(mySlc1)
	mySlc2 := mySlc[4:7]
	fmt.Println(mySlc2)
	mySlc3 := mySlc[6:]
	fmt.Println(mySlc3)
	mySlc4 := append(mySlc[2:4], mySlc[6:8]...)
	fmt.Println(mySlc4)*/

	//slice'lar için copy methodunu ve assign (=) ile farkını açıklayalım;
	/*mySlc := []int{1, 2, 3}
	mySlc2 := make([]int, 2)
	fmt.Println(mySlc)
	fmt.Println(mySlc2)
	copy(mySlc2, mySlc) //mySlc2'ye mySlc'yi kopyalıyoruz. yani 2.parametre kopyalanan. COPY METHODU
	fmt.Println(mySlc)  //kopyalananda bir değişim beklenmiyor zaten.
	fmt.Println(mySlc2) //mySlc'nin 1. ve 2.elemanını kopyaladı(967.kodda) ancak 3.yü kopyalayamadı çünkü mySlc2'de sadece 2 elemanlık yer var sadece.
	mySlc[0] = 100      //şeklinde 0.indexteki ilk elemanı değiştirip bakalım slice'lara;
	fmt.Println(mySlc)  //0.indexteki ilk elemanı değişti tabii ki.
	fmt.Println(mySlc2)*/ //fakat bunda 0.indexteki ilk eleman değişmedi, eskisi gibi [1, 2]. çünkü biz kopyalamakla farklı bir underline array oluşturuyoruz arka planda. yani kopyalandığı zaman aynı array'e referans vermiyorlar, farklı array'e referans veriyorlar.

	/*mySlc := []int{1, 2, 3}
	mySlc2 := make([]int, 2)
	fmt.Println(mySlc)
	fmt.Println(mySlc2)
	mySlc2 = mySlc //copylamak yerine assign yapıyoruz, yani mySlc'yi mySlc2'ye atıyoruz, assign ediyoruz.
	fmt.Println(mySlc)
	fmt.Println(mySlc2) //assign işleminden sonra mySlc2 mySlc'a tamamiyle eşit oldu. çünkü mySlc2'a mySlc'yi atamış oldum. yani bunlar aynı underline array'i tutuyorlar gibi düşünün. çünkü aşağıda ilk elemanı değiştirdiğimizde (981.satır) direkt mySlc2 de etkilenmiş oldu(983.satır).
	mySlc[0] = 100
	fmt.Println(mySlc)
	fmt.Println(mySlc2)*/
}

func ThirtiethFunction() {
	//slice ve array'ler aynı veri tipine sahip değerlerden oluşuyorlardı. aynı şekilde map'leri oluşturan key-value çiftleri de kendi içlerinde aynı veri tipine sahiler idi.
	//ancak gerçek hayatta herhangi bir şeyi tanımlarken farklı veri tiplerinden yararlanırız. go'da bu farklı veri tiplerinden yararlanma durumu yeni bir veri tipi olarak düşünülür ve bu veri tipine struct adı verilir.
	//çalışanların bilgilerinin tutulduğu bir struct yazalım;
	var employee struct { //struct veri tipinin ilgili alanlarına field denir. yani struct içindeki her bir satırı field gibi düşünebilirsin. her bir field'ın da kendine özel veri tipi var.
		name      string //bu field'ın veri tipi string
		age       int    //bu field'ın veri tipi int
		isMarried bool   //bu field'ın veri tipi bool
	} //yani struct veri tipi yukarda bulunan diğer veri tiplerinin (string, int, bool) üzerine inşaa edilir.
	fmt.Println(employee)
	fmt.Printf("%#v\n", employee)
	fmt.Println(employee.age) //bir struct'ın herhangi bir alanna ulaşmak için böyle yazarız.
	//struct'ın her bir field'ına belirli değerler atayıp yazdıralım;
	employee.name = "Ezgi"
	employee.age = 22
	employee.isMarried = false
	fmt.Printf("%#v\n", employee)
	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.isMarried)
	//peki bir şirkette binlerce çalışan olduğunda da herbiri için yukarıdaki gibi struct oluşturup değerler mi atayacağız? bu ne kadar işlevsel? yazılım biliminde böyle bir şeyi sürekli tekrarlayarak yazıyorsak muhtemelen ya yanlış ya da eksik bir bilgiyle yapıyoruzdur.
	//yukarıda da söylediğimiz gibi struct veri tipi diğer veri tiplerinin (string, int, bool) üzerine inşaa edilir. ancak bunu bu şekilde değil de tanımlı bir veri tipi şeklinde, yani kendi oluşturacağımız/tanımladığımız bir veri tipi şeklinde de tanımlayabiliriz ve bu yol struct'lardan daha farklı ve kısa bir yoldur;
	/*type employeee struct { // var/variable/değişken değil employeee adında yeni bir type/veri tipi oluşturuyoruz.yani employeee struct üzerinden oluşturulan bir veri tipidir, veri tipi struct değil yani.buradaki struct kendi veri tipimizi oluştururken faydalandığımız veri tipidir. yani struct'ın buradaki konumu underlying type'dır. (slice'lar array üzerinden oluşturulurken oradaki array'e underlying array diyoruz ya, aynı mantık.) kendi oluşturduğumuz employeee veri tipine ise defined/named type'dır.
		name      string
		age       int
		isMarried bool
	}*/
	//kendi veri tipimizin yazılım şekli;  type <my_type_name> <type_of_value>
	var e1 employeee //e1 şeklinde değişken oluşturduk ve veri tipi olarak employeee dedik.
	e1.name = "Ali"
	e1.age = 25
	e1.isMarried = false
	var e2 employeee
	e2.name = "Ayşe"
	e2.age = 33
	e2.isMarried = true
	e3 := employeee{
		name:      "Seda",
		age:       62,
		isMarried: true,
		/*kids: []string{
			"oğulcan",
			"hande",
		},*/
	}
	fmt.Printf("%#v\n", e1)
	fmt.Printf("%#v\n", e2)
	fmt.Printf("%#v\n", e3) //peki elemanlarına nasıl ulaşırız?;
	fmt.Println(e3)
	/*fmt.Println(e3.kids)
	fmt.Println(e3.kids[0])*/ //slice olduğu için istediğimiz index'ini de yazdırabiliriz.
	e4 := e3                  //oluşturduğumuz e4 struct'ına e3'ü atıyoruz.
	fmt.Println(e4)
	e4.name = "Göksun"
	fmt.Println(e4)
	fmt.Println(e3) //e4 struct'ında ismi değiştirmemize rağmen e3 struct'ı aynı kalır.
	//struct'lar birbirleri arasında değerleri paylaşırlar, referansı değil. yani pass by reference değil pass by value. yani bunlar hafızada 2 tane aynı yeri değil farklı yeri tutuyorlar. küme gibi düşün 2 ayrı küme var ve bir kümede değişiklik olduğunda öbür kümede değişiklik olmuyor.

	m1 := manager{
		employeee: employeee{
			name:      "Ayşe",
			age:       28,
			isMarried: false,
		},
		hasDegree: true,
	}
	/*m1 := manager{}
	m1.name = "Ayşe"
	m1.age = 28
	m1.isMarried = false
	m1.hasDegree = true*/ //şeklinde de oluşturabiliriz.
	fmt.Println(m1)

	//birden fazla employe yada manager olabilir. ancak bazen öyle durumlar olabilir ki tek kullanım olur (mesela patron). onun için biz anonim struct yapısını kullanabiliriz.
	theBoss := struct {
		name  string
		money bool
	}{name: "the boss", money: true}
	fmt.Println(theBoss) //sadece 1 kez kullanılacak bu yapı için gidip diğerleri gibi type theboss struct falan yapmamıza gerek yok yani. bu tarz durumlarda anonim struct kullan.
}

// kendi oluşturduğumuz veri tipine her fonksiyon içinden ulaşabilmek için tanımladığımız veri tipini main fonksiyonunun dışına almamız gerekir;
type employeee struct {
	name      string
	age       int
	isMarried bool
	//kids      []string //slice'dan oluşan bir kids veri tipi.
}

type manager struct {
	employeee
	hasDegree bool
} //manager employeee'nin özelliklerine sahiptir ama employeee değildir. yani IS A RELATION değil HAS A RELATION.

//manager ve employeee birbirlerinden bağımsız, struct'lar üzerinde kurulan veri tipleridir. yani manager struct'ında yaptığımız (employeee struct'ını kullanmamız) sadece daha kısa bir yazım şeklidir. teker teker yazmakla uğraşmadık.

func ThirtyfirstFunction() {
	var m1 mile
	m1 = 3.2
	fmt.Println(m1)
	fmt.Printf("%T, %v", m1, m1)

	f1 := float64(4.4)
	fmt.Println()
	fmt.Println(3.2 + 4.4) //hata vermez çünkü aynı veri tipindeler.
	//fmt.Println(m1 + f1)   //hata verir çünkü birbirlerinden farklı veri tiplerindeler. mile float64 üzerine kurulmasına rağmen mile ile float64 birbirlerinden ayrı veri tiplerine sahiplerdir. ancak birbirlerine dönüştürülebilirler type conversion ile tabi;
	fmt.Printf("%T, %v", m1+mile(f1), m1+mile(f1))
	fmt.Println()
	fmt.Printf("%T, %v", float64(m1)+f1, float64(m1)+f1)
	//yani underlying type(float64) veri tipiyle defined type(mile) veri tipi birbirlerinden farklıdır.ancak birbirlerinde dönüştürülebilirler.

	fmt.Println()
	k1 := kilometer(7.8)
	fmt.Printf("%T, %v", k1, k1)
	//fmt.Println(m1 + k1) //mile ve kilometer aynı underlying type üzerinden oluşturulsalar da farklı veri tiplerinde olduklarından işlem yapılamaz.

	fmt.Println()
	m2 := mile(4.6)
	fmt.Println(m1 + m2)
	fmt.Println(m1 * m2)
	fmt.Println(m1 + m2 + 2.1)
	//fmt.Println(m1 + m2 + "arin") hata verir çünkü biz bir veri tipi için kendimiz yeni bir metot oluşturmadıkça, o veri tipi yalnızca underlying veri tipine sahip olan işlemleri yapar.

	mystring := "ezgi"
	fmt.Println(strings.ToUpper(mystring)) //oluşturduğumuz veri tiplerine underlying type'larının sahip oldupu tüm paketleri uygulayabiliriz. örneğin burada to upper metodu m1'e uygulanamazdı.

	//kafalarda şöyle bir soru belirmiş olabilir: evet 3 tane yeni veri tipi oluşturdun ama işlemlerde yaptıklarının tamamını yeni veri tipleri oluşturmadan da yapabilirdin. mesela yeni veri tipi tanımlamadan yapalım;
	s1 := "sivas"
	fmt.Println(strings.ToUpper(s1)) //görüldüğü üzere bu şekilde de aynı sonucu alırsın. peki sen neden kendi veri tipini oluşturuyorsun aynı işlemleri kendi veri tipini oluşturmadan da yapabiliyorken?
	//1; zaten sadece yukarıdaki gibi işlemler yapmak için kendi veri tipimizi oluşturmayız.
	//2; go'da kendi veri tipimizi oluşturmanın en temel sebebi, o veri tipine ait olan yeni fonksiyonellikler tanımlamaktır. yani o veri tipine yeni metotlar bağlamaktır.
	//3; kodun okunurluğu açısından daha iyi.
}

type mile float64 //mile'ın veri tipi float64 değil.mile float64 veri tipi üzerine kurulmuş defined type'dır.
type kilometer float64
type mystring string

// mile veri tipine ait olan, bir uzunluğu kilometreye çevirme fonksiyonu yapalım; aldığı mile değerini kilometreye çevirecek yani;
func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6) //1 mil 1.6 kilometre
}

// kilometreyi mile çeviren bir fonksiyon yazalım;
func toMile(k kilometer) mile {
	return mile(k * 0.62)
}

// polymorphism : bir işlemin farklı biçimlerde yapılmasıdır/ çok biçimlilik. örneğin alan hesaplama mantığı her zaman aynıdır fakat daire, dikdörtgen alanları farklıdır. yani methodun adı aynı, ancak methodun uygulanması/implement edilmesi çok biçimli/polymorhf.
// en büyük avantajı methodların tekrar tekrar kullanılabilmesidir. interface'den tek farkı gelen veri tipine bu methodun farklı bir şekilde uygulanmasıdır.
// tek bir fonk ismi var ancak bu fonknun impelent edilmesi/kullanılması veri tiplerine göre farklılık gösteriyor.
type shape interface {
	area() float64
}

func printArea(shapes ...shape) { //shape veri tipini/interface'sini alacak ama kaç tane shape parametresi gelecek bilmiyoruz. o yüzden ...shape diyoruz. varyatik parametre denir.
	for _, shape := range shapes { //range fonksiyonu gelen shapes'lerin her bir elemanını shape1 olarak belirtecek.
		fmt.Println("Alan:", shape.area())
	}
}

type triangle struct {
	a float64
	h float64
}

func (t triangle) area() float64 {
	return (t.a * t.h) / 2
}

type square struct {
	a float64
}

func (s square) area() float64 {
	return s.a * s.a
}

type rectangle struct {
	a, b float64
}

func (r rectangle) area() float64 {
	return r.a * r.b
}
func ThirtyfourthFunction() {
	t := triangle{3, 4}
	s := square{4}
	r := rectangle{4, 5}
	printArea(t, s, r)
}

func selam() string {
	return "selam"
}
func FourtyFunction() {
	fmt.Println(selam())
	//fmt.Println(go selam()) //dersen yine aynı hatayı alırsın.selam fonkda return var çünkü
}

func main() {
	//bir dosyayı açmaya çalışalım;
	/*file, err := os.Open("test.txt") //methodda tanımlanmış varolan hata mesajıını kendi döndürüyor çalıştırınca. test.txt dosyası olursa hata döndürmeyecek tabiki.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dosyamız:", file)
	}*/

	TwentynightFunction()
	ThirtiethFunction()
	ThirtyfirstFunction()

	//m1=10 k1=?
	/*m1 := mile(10)
	k1 := toKilometer(m1)
	fmt.Println(k1)*/ //10 milin kilometre değerinden karşılığını buluyoruz toKilometer fonksiyonu sayesinde.

	//k2=10 m2=?
	/*k2 := kilometer(10)
	m2 := toMile(k2)
	fmt.Println(m2)*/ //10 kilometrenin mil değerinden karşılığını buluyoruz toMile fonksiyonu sayesinde.

	ThirtyfourthFunction()
	FourtyFunction()
}
