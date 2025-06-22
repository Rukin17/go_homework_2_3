// Создайте функцию capitalizeWords(s string) string, которая преобразует каждое слово в строке так,
// чтобы первая буква была заглавной, а остальные — строчными. Например: "привет мир" → "Привет Мир".
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	result := capitalizeWords(input)
	fmt.Println(result)
}

func capitalize(word string) string {
	if len(word) == 0 {
		return ""
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

func capitalizeWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = capitalize(word)
	}
	return strings.Join(words, " ")
}
