package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func init() {
}

func Q01() {
	fmt.Print("Notas: ")
	// Jeito massa fácil gostoso (só aguenta três notas)
	// var n1, n2, n3 float32
	// fmt.Scanln(&n1, &n2, &n3)
	// notes := []float32{n1,n2,n3}

	// Jeito masoquista louco pirado (aguenta infinitas)
	notes := inputToList()

	var s float64
	for i := range notes {
		s += notes[i]
	}
	average := toFixed(s/float64(len(notes)), 2)
	if average >= 6 {
		fmt.Printf("Aprovado, sua nota foi de %v.", average)
	} else {
		fmt.Printf("Reprovado, sua nota foi de %v.", average)
	}
	return
}
func Q02() {

	fmt.Println("Número de jogos: ")
	var n int
	fmt.Scan(&n)
	for i := 1; i < (n + 1); i++ {
		fmt.Printf("Jogo %v: ", i)
		data := inputToList()
		fmt.Printf("A renda do jogo foi de R$%v\n", toFixed(data[0]/100*(data[1]*1+data[2]*5+data[3]*10+data[4]*20), 2))
	}
}
func Q03() {

	fmt.Println("Digite três algarismos: ")
	var n1, n2, n3 string
	fmt.Scanln(&n1, &n2, &n3)
	if len(n1) != 1 || len(n2) != 1 || len(n3) != 1 {
		fmt.Println("Dígito inválido.")
		return
	}
	var n, _ = strconv.ParseFloat(n1+n2+n3, 32)
	fmt.Printf("O número é %v, seu quadrado é %v", n, math.Pow(n, 2))
}
func Q04() {
	var s, q float64
	fmt.Println("Salário Mínimo: ")
	fmt.Scan(&s)
	fmt.Println("Consumo em kW: ")
	fmt.Scan(&q)
	fmt.Printf("Custo por kW: R$ %v\n", toFixed(s*0.007, 2))
	fmt.Printf("Custo do consumo: R$ %v\n", toFixed(s*0.007*q, 2))
	fmt.Printf("Custo com desconto: R$ %v", toFixed(s*0.007*q*0.9, 2))
}
func Q05() {
	var acc, waste float64
	var kind string
	var value float64
	fmt.Scanln(&acc, &waste, &kind)
	if kind == "R" {
		value = 5 + 0.05*waste
	} else if kind == "C" {
		value = 500 + 0.25*(waste-80)
	} else if kind == "I" {
		value = 800 + 0.04*(waste-100)
	} else {
		fmt.Println("Tipo de consumidor inválido.")
		return
	}
	fmt.Printf("Conta: %v\n", acc)
	fmt.Printf("Valor da conta: R$%v", toFixed(value, 2))
}
func Q06() {
	var n int
	fmt.Print("Número de temperaturas em ºF: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Temperatura %v: ", i+1)
		var t float64
		fmt.Scan(&t)
		fmt.Printf("%vºF equivalem a %vºC.\n", t, toFixed(5*(t-32)/9, 2))
	}
}
func Q07() {
	var t, r float64
	fmt.Print("Temperatura em ºF: ")
	fmt.Scan(&t)
	fmt.Print("Quantidade de chuva em polegadas: ")
	fmt.Scan(&r)
	fmt.Printf("O valor em ºC: %v\nA quantidade de chuva em mm: %v", toFixed(5*(t-32)/9, 2), toFixed(r*25.4, 2))
}
func Q08() {
	var r, h float64
	fmt.Print("Raio da lata: ")
	fmt.Scan(&r)
	fmt.Print("Altura da lata: ")
	fmt.Scan(&h)
	fmt.Printf("O valor do custo de fabricação é: R$%v.", toFixed((2*math.Pi*r*h+2*math.Pi*math.Pow(r, 2))*100, 2))
}
func Q09() {
	fmt.Print("Valores de A, B e C: ")
	coefs := inputToList()
	if len(coefs) != 3 {
		fmt.Println("Dígitos inválidos")
		return
	}
	fmt.Printf("∆ = %v.", math.Pow(coefs[1], 2)-4*coefs[0]*coefs[2])
}
func Q10() {
	fmt.Print("Escreva os quatro elementos da matriz (M11, M12, M21, M22): ")
	els := inputToList(4)
	fmt.Printf("O valor do determinante é de %v.", els[0]*els[3]-els[1]*els[2])
}

func inputToList(params ...int) []float64 {
	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')
	notes := toFloat64Elements(strings.Split(strings.TrimSpace(inp), " "))
	if len(params) != 0 {
		if len(notes) != params[0] {
			fmt.Println("O input está incorreto.")
			return []float64{}
		}
	}

	return notes
}
func toFloat64Elements(l1 []string) []float64 {
	var floats = []float64{}
	for i := range l1 {
		l1[i] = strings.TrimSuffix(l1[i], "\n")
		float, _ := strconv.ParseFloat(l1[i], 32)
		floats = append(floats, float)
	}
	return floats
}
func toFixed(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
