package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Abre o Arquivo
	file, err := os.Open("document.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	// Cria um novo scanner e divide linha por linha dentro de uma lista
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	// Fecha o arquivo
	err = file.Close()
	if err != nil {
		return
	}

	// Cria um Replacer, que substitui as palavras por números
	replacer := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9", "zero", "0")

	// Declara a variavel da soma total
	totalSum := 0

	// Itera sobre cada linha do arquivo
	for _, eachline := range txtlines {
		// Declara a variavel do primeiro e do ultimo digito
		var firstDigit string = "0"
		var lastDigit string = "0"
		// Substitui as palavras por números usando o replacer
		newLine := replacer.Replace(eachline)
		// Itera sobre cada caractere da linha
		for _, char := range newLine {
			// Se o char for um número ele verifica se firstDigit é = 0, se for ele atribui o valor do char para firstDigit e lastDigit, se não ele atribui o valor do char para lastDigit
			if isNumber(string(char)) {
				if firstDigit == "0" {
					firstDigit = string(char)
					lastDigit = string(char)
				} else {
					lastDigit = string(char)
				}
			}
		}

		// Junta os dois digitos
		number := firstDigit + lastDigit

		// Transforma a string em um inteiro
		finalNumber, _ := strconv.Atoi(number)

		// Soma o número ao total
		totalSum += finalNumber
	}

	// Printa o total
	fmt.Println(totalSum)
}

// Cria a funçao para verificar se é um número
func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
