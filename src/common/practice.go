package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// 1) basit bir hesap makinesi yapın.
func calculation(num1, num2 int) (int, int, int) { //func calculation(num1, num2 int) (int, int, int) = func calculation(num1, num2 int) (sum int, dif int, prod int)
	sum := num1 + num2
	dif := num1 - num2
	prod := num1 * num2

	fmt.Println("Toplam:", sum)
	fmt.Println("Fark:", dif)
	fmt.Println("Çarpım:", prod)

	return sum, dif, prod
}

// 2) 1 ile 50 arasındaki asal sayıları gösteren program yazın.
func primeNumberFunction() {
	var a, b int
	for a = 2; a < 50; a++ {
		for b = 2; b < (a / b); b++ {
			if a%b == 0 {
				break
			}
		}
		if b > (a / b) {
			fmt.Printf("%d bir asal sayıdır\n", a)
		}
	} //önce içteki döngü çalışır
}

// 3) başlangıç değeri 10 olan bir x değişkeni oluşturun. sonrasında bulunduğu adres üzerinden y değişkenini tanımlayıp x değerini 20 yapın.
func practiceFunction() {
	x := 10
	fmt.Println(x)
	y := &x
	fmt.Println(y)
	*y = 20
	fmt.Println(*y)
	fmt.Println(x)
}

// 4) şuanki zamanı gösteren bir fonksiyon yazın.
func momentTime() {
	var moment time.Time = time.Now() //veri tipi "time.Time" olan bir variable oluşturuldu. "Now()" bir metodu "time" paketine ait olan bir fonksiyondur.
	fmt.Println(moment)
}

// 5) 1 ile 100 arasındaki bir sayıyı tahmin etme uygulaması yazın. toplamda 10 tahmin hakkı olsun.
func numRand(min, max int) (num int) {
	rand.Seed(time.Now().Unix())    //fonksiyon her çalıştırıldığında yeni bir sayı gelmesi istendiğinden "rand" paketinden "Seed()" metodu kullanılır.
	return rand.Intn(max-min) + min //herhangi bir sayı return edilmesi gerekildiğinden "rand.Intn()" yazılır ve () içine dönmesi istenen sayı yazılır.
}

func prediction() {
	target := numRand(1, 100) //gelen sayıyı target değişkeninde saklayalım
	fmt.Println("1 ile 100 arasındaki sayıyı bulmaya çalışın")
	reader := bufio.NewReader(os.Stdin)
	for attempts := 0; attempts < 10; attempts++ {
		fmt.Println(10-attempts, "hakkınız kaldı.")
		fmt.Print("lütfen tahmininizi yapınız: ")

		input, err := reader.ReadString('\n') //gelen değeri bu iki değere atıyoruz yani.
		if err != nil {
			fmt.Println(err)
		} //girilen değeri aldık

		input = strings.TrimSpace(input) //alınan verinin başında ve sonundaki boşlukları siler
		num, err := strconv.Atoi(input)  //gelen değer string olarak okunacak.bunun int değere dönüşmesini istediğimiz için strconv.Atoi(input) bunu yapıyoruz.
		if err != nil {
			fmt.Println(err)
		} //girilen değeri aldıktan sonra gelen stringi integer'a çevirdik

		if num > target {
			fmt.Println("tahmininiz daha büyük, daha küçük bir sayı giriniz.")
		} else if num < target {
			fmt.Println("tahmininiz daha küçük, daha büyük bir sayı giriniz.")
		} else {
			fmt.Println("doğru tahmin ettiniz. Hedef sayı ", target, "idi.", attempts, " seferde buldunuz.")
			break //doğru bulunduğunda bitsin, devam etmesin diye yazdık
		}
	}
}

// 6) girilen sınav sonucunu ekrandan alıp, öğrencinin geçip geçmediğini yazan bir program yazın.
func examNotes() (int, error) {
	fmt.Print("Sınav sonucunuzu giriniz:") //sınav sonucunu girmesini istediğim satırla girdiği satırın aynı olması için Print(aynı satır için print yani) kullandık.Println yeni satıra geçerdi.
	reader := bufio.NewReader(os.Stdin)    //kullanıcının ekrana girdiği ifadeyi okuyacak reader değişkeni. os.Stdin de klavyeden girilen değeri reader değişkenine atar.
	//value, _ := reader.ReadString('\n')    //reader değerini, yani ekrana girilen değeri bastırmak için value değişkeni yaptık.
	input, err := reader.ReadString('\n') //gelen değeri bu iki değere atıyoruz yani.
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSpace(input) //alınan verinin başında ve sonundaki boşlukları siler
	num, err := strconv.Atoi(input)  //gelen değer string olarak okunacak.bunun int değere dönüşmesini istediğimiz için strconv.Atoi(input) bunu yapıyoruz.
	if err != nil {
		fmt.Println(err)
	}
	var result string
	if num >= 50 {
		result = "geçtiniz"
		fmt.Println(result)
	} else {
		result = "kaldınız"
		fmt.Println(result)
	}
	return num, nil
	//grade, _ := getGrade() //zaten getGrade() foksiyonunda hata kontrolu yapıldığı ve nil geri dönüşü yapıldığı için tekrar grade, err := getGrade() diyip if err != nil { fmt.Println(err) } yapmak gereksiz. o nedenle burada tekrar hata yakalamaktansa blank identify yani _ kullanmak daha mantıklı
	//fmt.Println(result) //result'ın scope'u sadece if methodu içinde olduğundan result ifadesinde unused hatası veriyordu. o sebeple dışarda da yazdırmak lazımdı. bu yüzden yaptık bunu
}

// 7) bir bölme işlemi sonucunda bölüm ve kalanı yazdıran bir program yazın.
func bolme(bolunen, bolen int) (bolum, kalan int) { //multiple return
	bolum = bolunen / bolen
	kalan = bolunen % bolen
	fmt.Println(bolum, kalan)
	return bolum, kalan
}

// 8) bir sayının karekökünün alınacağı bir fonksiyon yazın.
/*func squareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("negatif sayıların karekökü alınamaz")
	}
	fmt.Println("Karekök sonucu:", math.Sqrt(num))
	return math.Sqrt(num), nil
} //peki bu fonksiyonda error kontrol edilmeseydi ne olurdu bakalım ;
*/
/*func evenNum(num int) (float64, error) {
	if num%2 != 0 {
		return 0, errors.New("HATA: çift sayı girmediniz")
	}
	return 0, nil
}
*/
/*func squareRoot2(num float64) float64 {
	result, err := evenNum(12) //result ve err diye 2 değer tanımlanır çünkü "evenNum" fonksiyonu 2 değer dönüyor.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Girdiğiniz sayı: ", result)
	}

	result, err = squareRoot(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = squareRoot(-5) //Nan beklenmedik sonucunu alırız erroru kontrol etmezsek. bunun için en başında bizim bir hata değeri oluşturmamız gerekiyor.
	{
		fmt.Println(result)
	}
	return num
	return math.Sqrt(num)
}
*/

func main() {
	/*calculation(10, 4)
	primeNumberFunction()
	practiceFunction()
	momentTime()
	numRand(1, 100)
	prediction()
	examNotes()
	bolme(100, 5)*/
	//squareRoot(-4)
	//squareRoot2(4)
}
