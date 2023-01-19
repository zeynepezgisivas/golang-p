package main

import (
	"fmt"
	"math"
)

// goroutine 'lerin birbirleriyle iletişime geçmesi için channels 'ler (kanallar) kullanılır.
// channels 'ler aynı zamanda bir goroutine tarafından gönderilen değerin, diğer goroutine tarafından kullanılmaya başlamasından önce gönderildiğinden de emin olur, yani garanti eder.

var myChannel chan string //veri tipi string olan bir channel oluşturuldu.

func mrb(channel chan string) { //parametresi "channel" olan bir fonksiyon oluşturuldu.
	channel <- "mrb" //artık string 'i bir return ile değil bir kanala gönderiyoruz. atama değil bak, gönderme
	//parametre olarak string kabul eden bir kanal alıyor bu fonk yani
}
func ChannelFunction() {
	myChannel = make(chan string)
	// myChannel := make(chan string) ---------------kısa yöntem
	go mrb(myChannel)
	fmt.Println(<-myChannel) //kanaldaki değeri almamızı sağlıyor.

	//chan2 := make(chan string)
	//chan2 <- "ezgi" //bir channel bulunduğu goroutine içerisinde başka bir goroutine'den değer dönmeden(aynı channel vasıtasıyla) kullanıldğı goroutine'ı blokluyor. o yüzden ezgi'yi yazdıramıyoruz, hata alıyoruz.
	//fmt.Println(<-chan 2)
}
func areaa(c circle, myChannel chan float64) {
	result := math.Pi * c.r * c.r
	myChannel <- result
}
func FourtytwoFunc() {
	c1 := circle{5}
	chan1 := make(chan float64)
	go areaa(c1, chan1)
	fmt.Printf("%.2f\n", <-chan1)
	c1.display()
}

func main() {

}
