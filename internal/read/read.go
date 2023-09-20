package read

import "fmt"

// type Reader interface {
// 	Read()
// }

type Reader struct{}

func (_ Reader) Read() {
	fmt.Println("Esto es una prueba")
	fmt.Println("Perdoname porque tengo poca experiencia con los go modules")
	fmt.Println("Veremos la estructura y si esto funciona por ahora")
}
