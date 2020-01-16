package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(12)
}

func main() {
	i := numeroAleatorio()
	fmt.Print("Almoço será no: ")
	switch i {
	case 0:
		fmt.Println("Açãizão de patrício !")
	case 1:
		fmt.Println("Buteco da linguiça !")
	case 2:
		fmt.Println("Zurich !")
	case 3:
		fmt.Println("Buteco Bradesco !")
	case 4:
		fmt.Println("BKzão !")
	case 5:
		fmt.Println("Shopping !")
	case 6:
		fmt.Println("Violetas Bar !")
	case 7:
		fmt.Println("Hotel mania !")
	case 8:
		fmt.Println("Boteco toyShow !")
	case 9:
		fmt.Println("Vegano escolhe !")
	case 10:
		fmt.Println("Japa !")
	case 11:
		fmt.Println("Qualquer lugar mas, com cerveja")
	case 12:
		fmt.Println("Pastel !")
	}
}
