package main

import "fmt"

// 1) underlying type "int" olacak şekilde kendi veri tipinizi oluşturun. veri tipi ve değerini "10" olarak yazdırın.
type myType int //istenildiği gibi "myType" şeklinde kendi veri tipimizi oluşturduk ve underlying type olarak da int dedik.

func underlyingType() {
	var x myType //veri tipi bu şekilde belirtilmezse, varsayılan olarak alınır. mesela bu örnekte int alınırdı.
	x = 10
	fmt.Printf("%T, %v", x, x)
}

// 2) underlying type "struct" olan bir "rectangle type" oluşturun. display, circumference ve area metodlarını yazın.
type rectangle struct { //istenildiği gibi "rectangle" şeklinde kendi veri tipimizi oluşturduk ve underlying type olarak da struct dedik.
	a, b int
}

func display(r rectangle) {
	fmt.Println("Kenar uzunlukları:", r.a, "ve", r.b)
}

func circumference(r rectangle) {
	fmt.Println("Çevre:", 2*(r.a+r.b))
}

func area(r rectangle) {
	fmt.Println("Alan:", r.a*r.b)
}

func underlyingType2() {
	r1 := rectangle{3, 8}
	display(r1)
	area(r1)
	circumference(r1)
}

// 3) underlying type "struct" olan bir "family type" oluşturun (field 'lar age, name ve isMarried olsun) ve aile bireylerinin hepsini farklı şekilde tanımlayın. sonrasında for döngüsünde yazdırın.
type family struct { //istenildiği gibi "family" şeklinde kendi veri tipimizi oluşturduk ve underlying type olarak da struct dedik.
	name      string
	age       int
	isMarried bool
}

func showFamiy() []family { //family veri tipinden oluşan bir slice dönüyor.
	fam1 := family{
		name:      "ayşe",
		age:       12,
		isMarried: false,
	}

	fam2 := family{
		name: "ali",
		age:  3,
		//yazılmayan özellik için o özelliklerin zero value 'si alınır.
	}

	fam3 := family{
		"aslı",
		30,
		true,
	} //field özellikleri yazılırken illa ki teker teker belirtilmesine gerek yok. bu şekilde aynı sıralamada, key 'lere gerek kalmadan value 'lar yazdırılabilir.

	//bir de field özellikleri farklı yöntemle dışarıdan ekleyelim;
	var fam4 family
	fam4.name = "ozan"
	fam4.age = 28
	fam4.isMarried = false

	return []family{fam1, fam2, fam3, fam4} //family 'den oluşan bir slice dönülüyor.
}

func underlyingType3() {
	families := showFamiy()
	for i := 0; i < len(families); i++ {
		fmt.Printf("%v, %T\n", families[i], families[i])
	}
}

func main() {
	underlyingType()
	underlyingType2()
	underlyingType3()
}
