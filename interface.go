package figuras

import "fmt"

type geometria interface {
	area() float64
	perimetro() float64
}

func Medidas(g geometria) {
	//fmt.Println(g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perímetro: ", g.perimetro())
}
