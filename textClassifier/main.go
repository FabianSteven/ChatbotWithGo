package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// dataset devuelve un mapa de frases a sus clases desde un archivo
func dataset(file string) map[string]string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dataset := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		data := strings.Split(l, "\t")
		if len(data) != 2 {
			continue
		}
		sentence := data[0]
		if data[1] == "1" {
			dataset[sentence] = greeting
		} else if data[1] == "2" {
			dataset[sentence] = liked
		} else if data[1] == "3" {
			dataset[sentence] = disliked
		} else if data[1] == "4" {
			dataset[sentence] = orderPiza
		} else if data[1] == "5" {
			dataset[sentence] = orderHamburger
		} else if data[1] == "6" {
			dataset[sentence] = orderSalad
		} else if data[1] == "7" {
			dataset[sentence] = orderSoda
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return dataset
}

func main() {
	// Inicializar un nuevo clasificador
	nb := newClassifier()
	// Obtener el conjunto de datos de un archivo de texto
	dataset := dataset("./sentiment labelled sentences/chats2.txt")
	// Entrenar el clasificador con el conjunto de datos
	nb.train(dataset)

	// Solicita las entradas de la consola
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Escriba su comentario: ")
		sentence, _ := reader.ReadString('\n')
		// Clasificar la frase de entrada
		result := nb.classify(sentence)
		class := ""
		if result[greeting] > result[liked] && result[greeting] > result[disliked] && result[greeting] > result[orderPiza] && result[greeting] > result[orderHamburger] && result[greeting] > result[orderSalad] && result[greeting] > result[orderSoda] {
			class = greeting
			fmt.Printf("> Hola, en que podemos ayudarle - %s\n\n", class)
		} else if result[liked] > result[disliked] && result[liked] > result[orderPiza] && result[liked] > result[orderHamburger] && result[liked] > result[orderSalad] && result[liked] > result[orderSoda] {
			class = liked
			fmt.Printf("> Nos alegra que le gustara la comida - %s\n\n", class)
		} else if result[disliked] > result[orderPiza] && result[disliked] > result[orderHamburger] && result[disliked] > result[orderSalad] && result[disliked] > result[orderSoda] {
			class = disliked
			fmt.Printf("> Tendremos en cuenta su comentario y mejoraremos - %s\n\n", class)
		} else if result[orderPiza] > result[orderHamburger] && result[orderPiza] > result[orderSalad] && result[orderPiza] > result[orderSoda] {
			class = orderPiza
			fmt.Printf("> Una pizza en camino - %s\n\n", class)
		} else if result[orderHamburger] > result[orderSalad] && result[orderHamburger] > result[orderSoda] {
			class = orderHamburger
			fmt.Printf("> Una hamburguesa en camino - %s\n\n", class)
		} else if result[orderSalad] > result[orderSoda] {
			class = orderSalad
			fmt.Printf("> Ok, se esta preparando una ensalada - %s\n\n", class)
		} else {
			class = orderSoda
			fmt.Printf("> Ya le traemso su bebida - %s\n\n", class)
		}

	}
}
