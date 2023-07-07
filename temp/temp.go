// declarar o pacote principal
package main

//importar FMT
import "fmt"

//declaração da variável CONST da ebulição
const ebulicaoF float64 = 212.0

//função principal
func main() {

	tempF := ebulicaoF //variavel da temp em farh
	tempC := (tempF - 32) * 5 / 9

	fmt.Printf("A temperatura de ebulição da agua em ºF = %g (%T) e temperatura de ebulição da agua em ºC = %g (%T) .", tempF, tempF, tempC, tempC)

}
