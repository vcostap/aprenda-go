package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

type figura interface {
	calcArea() float64
}

func main() {
	quadrado := quadrado{4}
	circulo := circulo{
		raio: 3,
	}

	fmt.Println("A área do quadrado é:", info(quadrado))
	fmt.Println("A área do círculo é:", info(circulo))
}

func (q quadrado) calcArea() float64 {
	return q.lado * q.lado
}

func (c circulo) calcArea() float64 {
	return 2 * math.Pi * c.raio
}

func info(fig figura) float64 {
	return fig.calcArea()
}