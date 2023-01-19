package main

import (
	"fmt"
	"math"
)

// interface kullanılacak metodları ve yapılması gereken işlemleri gösterir. metodların çalıştırılması interface fonksiyonuna parametre olarak gönderdiğimiz veri tiplerine aittir.

type shape interface { //interface metodların imzasını taşır.
	// yani methodlarının/fonksiyonlarının nasıl çalışacağını/implementasyonunu bilmez interface.işlemi/implementasyonu/o fonksiyonun çalıştırılması yine parametre olarak alınan veri tipi yapar. burada rectangle mesela yani. interface olarak görevi yalnızca methodlarla ilgilenmek/ne yapacağını söylemek.  //circle veri tipini de bu shape interface'sine parametre olarak göderebiliriz.
	area() float64          //burdaki method isimleri-ver tiplerinin fonksiyonlarla uyuşması lazım. yoksa çalışmaz interface.
	circumference() float64 //burdaki method isimleri-ver tiplerinin fonksiyonlarla uyuşması lazım. yoksa çalışmaz interface.
} //diameter() fonksiyonu yok ama yine de çalışır. yani interface'in karşılaması için fazladan bir method olabilir ama eksik olamaz. eksik olsaydı çalışmazdı.yada area1 falan yani farklı olsaydı da çalışmazdı.
// interface 'e bağlı olan fonksiyon oluşturalım;

func InterfaceFunction(i shape) {
	fmt.Println(i)                 //aldığımız parametrenin bize ne olduğunu göstersin. //shape interface veri tipimizin adı. bu interface'e (1362ye yani) farli bir veri tipi geçebilmemizin sebebi, interface methodlarının/fonksiyonlarının o veri tipinden olmasıdır.
	fmt.Println(i.area())          //i'ye ait olan area methodunu çalıştırsın. //i bir shape ve shape bir interface. interfacede de area ve circumference olan 2 tane methodu var
	fmt.Println(i.circumference()) //i'ye ait olan circumference() methodunu çalıştırsın.
	fmt.Printf("%T", i)
}

/*type rectangle struct {
	a, b float64
}*/

func (r rectangle) circumference() float64 { //burdaki method isimleri-ver tiplerinin fonksiyonlarla uyuşması lazım. yoksa çalışmaz interface.
	return 2 * (r.a + r.b)
}

func (r rectangle) area() float64 { //burdaki method isimleri-ver tiplerinin fonksiyonlarla uyuşması lazım. yoksa çalışmaz interface.
	return r.a * r.b
}

type circle struct {
	r float64
}

func (c circle) circumference() float64 { //burdaki method isimleri-ver tiplerinin fonksiyonlarla uyuşması lazım. yoksa çalışmaz interface.
	return 2 * math.Pi * c.r
}

func (c circle) area() float64 { //burdaki method isimleri-ver tiplerinin fonksiyonlarla uyuşması lazım. yoksa çalışmaz interface.
	return math.Pi * c.r * c.r
}

func (c circle) diameter() float64 {
	return 2 * c.r
}

func InterfaceFunction2() {
	r1 := rectangle{3, 8}
	fmt.Println("Area:", r1.area())
	fmt.Println("Circumference:", r1.circumference())
	//bazı durumlarda interface tanımını yapmak istersek, bizim açımızdan önemli olan sadece methodlardır. methodların hangi veri tipine ait olduklarının bizim oluşturduğumuz interface açısından bir önemi yoktur. interface için önemli olan methodun kendisi.
	//1358.satırda kendi interface'imizi oluşturduk.
	InterfaceFunction(r1) //r1'i parametre olarak nasıl gönderebiliyoruz? interface fonksiyonlara bakıyordu. biz interface'e parametre olarak r1 gödermişiz.r1'i işleme alabilmemiz için/yani interface'e argüman olarak gönderebilmemiz için r1'in area ve circumference diye fonksiyonları olması lazım.evet 1341, 1344te de rectange'in area ve  circumference fonksiyonları var. o zaman biz rectengle veri tipini interface'e gönderebiliyoruz.
	fmt.Println()
	c1 := circle{5}
	InterfaceFunction(c1)
}

func main() {
	InterfaceFunction()
	InterfaceFunction2()
}
