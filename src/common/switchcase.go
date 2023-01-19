package main

import "fmt"

func SwitchCaseFunction() {
	grade := -1
	switch grade {
	case 5: //if grade == 5 { fmt.Println("pek iyi") } yani koşulla case birbirine eşit mi değil mi ona bakılıyor //case ifadesinde kullandığın, yani karşılaştırdığın veri tipiyle değişkenin veri tipi aynı olmak zorunda tabi. yani case "beş": dersen hata verir
		fmt.Println("pek iyi")
		//break : aslında her bu ifadenin altında, kod burada çalışırsa yapının dışına çıkartmasını sağlayan "break" var.
	case 4:
		fmt.Println("iyi") //buradaki statementler sadece kendi caseleri tarafından görülür.yani bu satementi sadece case 4 görür. yani her case in ayrı kapsamları var diyebiliriz.
	case 3:
		fmt.Println("orta")
		//y := 100
		//fmt.Println(y) sadece case 3 çalıştığında çalışır. y farklı casede tanımlanırsa ama statement bu case de olsa dahi çalışmaz
	case 2:
		fmt.Println("kötü")
	case 1:
		fmt.Println("çok kötü")
	default: //if else yapısındaki else ile denk olan ve beklenenden daha farklı ifadeyi(test ettiğimiz koşullar harici bir sonuç için) default ile yazarız switch case'de. default switch case içinde her yere yazılabilir illa sonda olacak diye de bir şey yok yani
		//default olmadığında ve koşullar(caseler)da karşılaştıracağın değişkenin değerine uymadığında caselere uyulmadığından statementler çalıştırılmayacak ve bir çıktı vemeyecek.
		fmt.Println("geçersiz not")
	}
	fmt.Println(grade) //değişkeni switch yapısı dışında tanımladığında aynı else if yapısında olduğu gibi switch yapısı bittiğinde de o değişken ulaşılabilir olur. if yapısında olduğu gibi değişkeni switch case yapısında tanımlayabiliriz;*/

	switch grade := -3; grade {
	case 5:
		fmt.Println("pek iyi")
	case 4:
		fmt.Println("iyi")
	case 3:
		fmt.Println("orta")
	case 2:
		fmt.Println("kötü")
	case 1:
		fmt.Println("çok kötü")
	default:
		fmt.Println("geçersiz not")
	}
	//switch case yapısının aynısını if-else if yapısıyla yazalım;
	/*if grade == 5 {
		fmt.Println("pek iyi")
	} else if grade == 4 {
		fmt.Println("iyi")
	} else if grade == 3 {
		fmt.Println("orta")
	} else if grade == 2 {
		fmt.Println("kötü")
	} else if grade == 1 {
		fmt.Println("çok kötü")
	} else {
		fmt.Println("geçersiz not")
	}*/
	//case yapısı ile else if yapısı ile denk yani. (baştaki if hariç diğer case'ler yani)

	//switch koşulunu yazarken herhangi bir koşul yazmadan da oluşturabilir. örneğin ;
	switch {
	case false:
		fmt.Println("bu yazdığımız konsolda görünmeyecek.")
	case true:
		fmt.Println("bu yazdığımız konsolda görünecek.")
	}
	// bu şekilde de boolean ifade olarak true ya denk gelen case durumlarını yazdırabiliriz

	//switch fallthrough :
	switch a := 25; {
	case a < 20:
		fmt.Printf("%d 20'den küçüktür.\n", a)
		fallthrough //normalde switch case yapısında true case ile karşılaşıldığında o case'deki işlemin yapılıp, sonrasında o switch yapısından çıkılması gerekirdi. fakat diğer durumların da test edilmesini istedğimizde, yani koşulun devam etmesi durumlarında fallthrough kullanılır
	case a < 50:
		fmt.Printf("%d 50'den küçüktür.\n", a)
		fallthrough
	case a < 100:
		fmt.Printf("%d 100'den küçüktür.\n", a)
		fallthrough
	case a < 200:
		fmt.Printf("%d 200'den küçüktür.\n", a)
	}
}

func main() {
	SwitchCaseFunction()
}
