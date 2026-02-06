package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- Calculadora de IMC ---")

	fmt.Print("Digite o seu peso: \n")
	inputPeso, _ := reader.ReadString('\n')
	inputPeso = strings.TrimSpace(inputPeso)
	peso, _ := strconv.ParseFloat(inputPeso, 64)

	fmt.Print("Digite a sua altura: \n")
	inputAltura, _ := reader.ReadString('\n')
	inputAltura = strings.TrimSpace(inputAltura)
	altura, _ := strconv.ParseFloat(inputAltura, 64)

	imc := peso / (altura * altura)

	fmt.Printf("\nSeu IMC é: %2f\n", imc)

	if imc < 18.5 {
		fmt.Println("Status: Você esta abaixo do Peso")
	} else if imc >= 18.5 && imc <= 24.9 {
		fmt.Println("Status: Seu peso esta normal")
	} else {
		fmt.Println("Status: Você esta acima do peso")
	}

}
