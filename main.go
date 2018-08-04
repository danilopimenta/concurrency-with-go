package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	IsPalindrome(1, "A MALA NA CAMA E A CAMA NA CASA A BOLA DO SACO E O GALO NA PRACA O DEDO NO LACO E A GOTA NO CAPO DO CABO DO TOCO A SETA DO SACO E TOMO DA TRACA AMADA DA TABA E O MICO DO CIPO O SACO DE TROCO E TOPO DO TETO O CEBO DA RUA, A DROGA NA GRAMA O DEMO DA MOCA A PORTA DA CASA A CARA DO TIM E MITO DA RACA A SACA DA TROPA A COMA DO MEDO AMARGA NA GORDA AURA DO BECO O TETO DO POTE O CORTE DO CASO O PICO DO CIMO E A BATA DA DAMA A CARTA DO MOTE O CASO DA TESA O COTO DO BACO DO PACO NA TOGA E O CALO NO DEDO A CARPA NO LAGO E O CASO DA LOBA A SACA NA MACA E A MACA NA LAMA")
	IsPalindrome(2, "danilo")
	IsPalindrome(3, "nerdzao")
	IsPalindrome(4, "traxart")
}

func IsPalindrome(top int, text string) {
	text = sanitize(text)
	runes := []rune(text)
	var not string

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {

		if runes[j] != runes[i] {
			not = "not "
			break
		}
	}

	fmt.Println(strconv.Itoa(top) + ": is " + not + "a Palindrome")
}

func sanitize(value string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	safe := reg.ReplaceAllString(value, "")
	return strings.ToLower(strings.Trim(safe, ""))
}
