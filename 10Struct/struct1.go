package main

//Yapılar,nesne tabanlı programlamada kullanılan NESNE kavramına karsılık gelir.
//Yapıların kendıne aıt metotları da tanımlanabılmektdır.
import "fmt"

type insan struct {
	boy, kilo float32
	cinsiyet  string
}

func main() {
	Kasim := insan{boy: 173, cinsiyet: "erkek"}
	Melike := insan{boy: 170, cinsiyet: "Kadın"}
	ideal_kilo(&Kasim)
	ideal_kilo(&Melike)
	fmt.Println(Kasim.kilo)
	fmt.Println(Melike.kilo)

}
func ideal_kilo(birisi *insan) {
	switch birisi.cinsiyet {
	case "erkek":
		birisi.kilo = 0.9*birisi.boy - 85
	case "Kadın":
		birisi.kilo = 0.9*birisi.boy - 90
	default:
		fmt.Println("cinsiyet alaninda sorun var")

	}
}
