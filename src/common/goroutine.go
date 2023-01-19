package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// goroutine : 1 'den fazla işin aynı anda yapılmasını sağlar. yani go 'da eş zamanlı (concurrency) olarak yapılan görevlerin her birine bir goroutine denebilir.
func printX() {
	for i := 0; i < 20; i++ {
		fmt.Print("X")
	}
}
func printY() {
	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}
}
func ThirtyfifthFunction() {
	printY()
	fmt.Println()
	printX()
} //bu çalışma sırasının bir sorunu yok bildiğimiz üzere. ancak bir programın en hızlı çalışma şekli bu(sıralı çalışma) değildir. çünkü çalışan ilk fonksiyonun uzun sürmesi vs. demek diğer fonksiyonların onu beklemesi demektir ve bu iyi bir şey değildir. o yüzden biz 1'den fazla işin aynı anda yapılmasını isteriz.
// böyle sıralı çalışma olduğunda mesela x'in önce yazılmasını isesek manuel bir şekilde evet eddeğiştiririz x ve y'nin yerini. ama 200 tane fonk olduğunda bu çözüm efektif bir çözüm olmaz.
// peki birden fazla işi aynı anda/eşzamanlı yaptırmak istersek goroutin'leri kullanıcaz. mesela printx printy fonksiyonlarının önüne go eklemek goroutin oluşturmak demek. yapalım;

func ThirtysixthFunction() {
	go printY() //bu aslında 2.goroutine
	fmt.Println()
	go printX() //3.goroutine
	time.Sleep(time.Second)
} //run edersen bu fonksiyonu sadece fmt.Println() kodunun hatasının çalıştığını görürsün(1455 hariç). ee noldu go routin'ler?? aslında bizim 1 tane main goroutine'imiz var. main fonksiyonunu çağırdığımız anda/ main fonksiyonu çalışmaya başlar başlamaz fonksiyonumuzun ana goroutine'i, yani main goroutine'i oluşturmuş oluyoruz. yani aslında çalışan tüm go programlarında bizim 1 tane ana goroutine dediğimiz main goroutine'imiz çalışıyor.
// peki biz bu çıktı alamama sorununu nasıl çözeriz? main fonksiyonun çalışmayı tamamlamasını bir süre durduralım.(1455.satır)
// şöyle de bir sorun var: önce printy goroutine oluştu sonra printx goroutine. ama önce printy goroutine sonra printx goroutine çalışacak diye bir garanti olmuyor. yani goroutine'lerin birbirleri arasında ne zaman ve nasıl geçiş yapacaklarını ve bu geçişlerin sürelerini bilemiyoruz/kontrol edemiyoruz sadece goroutinelerle.
// goroutine'ler eş zamanlı(concurrency) olarak çalışır. ancak bu goroutineler aynı aynda çalışamıyor.(tek çekirdeklilerde) peki neden eş zamanlı diyoruz?ççünkü baktık 1 işlem fazla zaman alıyor hemen diğer işlemi alıyoruz. ama böyle olunca da kararsızlıklar oluyor. biz gelişigüzel çıktı almak istemeyiz sonuçta
// farklı işlemlerin/proceslerin aynı anda yapılması paralellism'dir.

// mesela ben her zaman en önce x'lerin yazılmasını istiyorum.
var wg sync.WaitGroup //waitgrouplarda 3 tane öemli method var; add, wait, done
func ThirtyseventhFunction() {
	wg.Add(1) //maingoroutine'e senden farklı 1 tane bekleyen gorouitne var diyoruz.
	go printXX()
	wg.Wait() //kendinden önceki goroutine tamamlanana kadar bekle diyoruz. sonra printyy den devam edicek. böylelikle önce x sonra y yazdırmış olduk. önceki gibi karışık olmayacak.
	printYY()
}
func printXX() {
	for i := 0; i < 20; i++ {
		fmt.Print("X")
	}
	wg.Done()
}
func printYY() {
	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}
}

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("a circle")
	wg.Done()
}
func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}
func ThirtyeightFunction() {
	wg.Add(1)
	c1 := circle{5} //daire bir değişkene atıldı
	//c1.display()             //display methodu çağırıldı
	area1 := c1.area()          //alan hesaplanıp bir değişkene atandıyor
	fmt.Printf("%.2f\n", area1) //değişkeni yazdırıyoruz
	go c1.display()
	wg.Wait() //tüm goroutine'ler çalışana kadar bekler
}
func ThirtynineFunction() {
	c1 := circle{5}
	//area1 := go c1.area() //go'da goroutine oluşturduğumuz zaman, goroutine'nin iliştirildiği fonksiyonda herhangi bir return değeri alınamaz. o yüzden bu şekilde hata alınır. ama nede? çünkü ana goroutine bir değere bağımlı olmuş oluyor, bekliyor yani. bu da akışı bozuyor.
	area1 := c1.area()
	fmt.Printf("%.2f\n", area1)
	go c1.display() //mesela burda bir return değeri olmadığı için dislay fonknun bir sorun yaratmıyor.

