package main

import "fmt"

//scope : oluşturulan bir değişkene ulaşılabilen, değişkenin varolduğu alan/ kapsam'dır.

var packVar = "Package Scope" //"package" değişkenine sahip olduğundan kullanılan paketin tamamında ulaşılabilir.
// packVar := "Package Scope" : paket değişkenlerinde bu şekilde short declaration (statement) yapamayız. yapılırsa "syntax error" alınır.

func ScopeFunction() {
	var funcVar = "Func Scope" //değişkenin alanı/ kapsamı bu blok yapısı içerisindedir.
	//funcVar := "Func Scope" : block scope'da (yani block içerisindeki scope'da) short declaration (statement) yapabiliriz.

	fmt.Println(funcVar)
	fmt.Println(packVar)
}

func ScopeFunction2() {
	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	}
	//fmt.Println(blockVar) : buraya konsaydı error alınırdı çünkü blockVar değişkeni kendi spoce'unda tanımlıdır.
	//fmt.Println(funcVar) : buraya konsaydı "undefined error" alınırdı çünkü çağırılan değişkene kendi scope'unun dışında ulaşılamaz.
}

func main() {
	ScopeFunction()
	ScopeFunction2()
}
