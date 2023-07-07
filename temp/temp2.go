package main

// importanto bibliotecas
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	// Pegando valores do cmd
	leitor_input := bufio.NewReader(os.Stdin)
	fmt.Printf("Digite a temperatura em °F: ")
	input, _ := leitor_input.ReadString('\n')

	// Retirando os caracteres não numeros do input
	returnInputFormated := strings.TrimRightFunc(input, func(r rune) bool {
		return !unicode.IsDigit(r)
	})

	// Convertendo em float64
	tempF, err := strconv.ParseFloat(returnInputFormated, 64)

	// Caso houver algum erro em converter em float64
	if err != nil {

		fmt.Println("Erro ao converter a string para float64: ", err)
		return

	}

	// Calculo de Fahrenheit para Celcius
	tempC := (tempF - 32) * 5 / 9

	fmt.Printf("A temperatura de ebulição da agua em °F = %g (%T) e temperatura de ebulição da agua em °C = %g (%T) .", tempF, tempF, tempC, tempC)

}
