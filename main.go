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

	fmt.Println("--- Calculadora de IMC (Versão Robusta) ---")

	// Lendo o Peso
	fmt.Println("Digite o seu peso (ex: 70,5 ou 70.5): ")
	inputPeso, _ := reader.ReadString('\n')
	inputPeso = strings.TrimSpace(inputPeso)
	// Tratamento: troca vírgula por ponto para o sistema entender
	inputPeso = strings.ReplaceAll(inputPeso, ",", ".")
	peso, err := strconv.ParseFloat(inputPeso, 64)

	if err != nil {
		fmt.Println("Erro: Peso inválido. Use apenas números.")
		return
	}

	// Lendo a Altura
	fmt.Println("Digite a sua altura (ex: 1,75 ou 1.75): ")
	inputAltura, _ := reader.ReadString('\n')
	inputAltura = strings.TrimSpace(inputAltura)
	// Tratamento: troca vírgula por ponto
	inputAltura = strings.ReplaceAll(inputAltura, ",", ".")
	altura, err := strconv.ParseFloat(inputAltura, 64)

	if err != nil {
		fmt.Println("Erro: Altura inválida. Use apenas números.")
		return
	}

	// Calculo
	imc := peso / (altura * altura)

	// Exibicao com duas casa decimais usando %2f
	fmt.Printf("\nSeu IMC é: %2f\n", imc)

	// Lógica de Classificação
	if imc < 18.5 {
		fmt.Println("Status: Você está abaixo do peso")
	} else if imc >= 18.5 && imc <= 24.9 {
		fmt.Println("Status: Seu peso está normal")
	} else {
		fmt.Println("Status: Você está acima do peso")
	}
}
