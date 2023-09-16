import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v tem %v anos e ele é um %v warrior.", y, x, z)
	fmt.Printf("%v -> %T", s, s)
}