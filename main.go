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

	// O loop for sem condições funciona como um "enquanto for verdade" (loop infinito)
	for {
		fmt.Println("\n--- Calculadora de IMC (Digite 'sair' para encerrar) ---")

		fmt.Print("Digite o seu peso: ")
		inputPeso, _ := reader.ReadString('\n')
		inputPeso = strings.TrimSpace(inputPeso)

		// Opção para o usuário encerrar o programa com elegância
		if strings.ToLower(inputPeso) == "sair" {
			fmt.Println("Encerrando... Até logo!")
			break // Sai do loop for
		}

		inputPeso = strings.ReplaceAll(inputPeso, ",", ".")
		peso, err := strconv.ParseFloat(inputPeso, 64)

		if err != nil || peso <= 0 {
			fmt.Println("Erro: Digite um peso válido e maior que zero.")
			continue // Pula o resto e volta para o início do for
		}

		fmt.Print("Digite a sua altura: ")
		inputAltura, _ := reader.ReadString('\n')
		inputAltura = strings.TrimSpace(inputAltura)
		inputAltura = strings.ReplaceAll(inputAltura, ",", ".")
		altura, err := strconv.ParseFloat(inputAltura, 64)

		if err != nil || altura <= 0 {
			fmt.Println("Erro: Digite uma altura válida e maior que zero.")
			continue
		}

		imc := peso / (altura * altura)
		fmt.Printf("Seu IMC é: %.2f\n", imc)

		// Lógica simplificada de status
		switch {
		case imc < 18.5:
			fmt.Println("Status: Abaixo do peso")
		case imc <= 24.9:
			fmt.Println("Status: Peso normal")
		case imc <= 29.9:
			fmt.Println("Status: Sobrepeso")
		default:
			fmt.Println("Status: Obesidade")
		}
	}
}
