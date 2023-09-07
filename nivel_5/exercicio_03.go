import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type camionete struct {
	veiculo           veiculo
	tracaoQuatroRodas bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	camionete := camionete{
		veiculo: veiculo{
			portas: 2,
			cor:    "preto",
		},
		tracaoQuatroRodas: true,
	}

	sedan := sedan{veiculo{4, "branco"}, false}

	fmt.Println(camionete)
	fmt.Println(sedan)
	fmt.Println(camionete.veiculo.cor)
	fmt.Println(sedan.modeloLuxo)
}