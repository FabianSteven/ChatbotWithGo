package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
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
			dataset[sentence] = orderPizza
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

func process(messData messages) messResponseData {

	var messres messResponseData
	// Inicializar un nuevo clasificador
	nb := newClassifier()
	// Obtener el conjunto de datos de un archivo de texto
	dataset := dataset("./sentiment labelled sentences/chats2.txt")
	// Entrenar el clasificador con el conjunto de datos
	nb.train(dataset)

	// Solicita las entradas de la consola
	//reader := bufio.NewReader(os.Stdin)
	fmt.Print("Escriba su comentario: ")
	//var prueba string = "hola como esta"
	//var prueba2 string = "como"
	//var prueba3 string = "esta"
	sentence, _ := messData.Num1, "" //reader.ReadString('\n')
	fmt.Print(reflect.TypeOf(sentence))
	// Clasificar la frase de entrada
	result := nb.classify(sentence)
	class := ""
	if result[greeting] > result[liked] && result[greeting] > result[disliked] && result[greeting] > result[orderPizza] && result[greeting] > result[orderHamburger] && result[greeting] > result[orderSalad] && result[greeting] > result[orderSoda] {
		class = greeting
		messres.Ans = "> Hola, en que podemos ayudarle - %s\n\n" //, class)
		messres.Class = class
	} else if result[liked] > result[disliked] && result[liked] > result[orderPizza] && result[liked] > result[orderHamburger] && result[liked] > result[orderSalad] && result[liked] > result[orderSoda] {
		class = liked
		messres.Ans = "> Nos alegra que le gustara la comida - %s\n\n" //, class)
		messres.Class = class
	} else if result[disliked] > result[orderPizza] && result[disliked] > result[orderHamburger] && result[disliked] > result[orderSalad] && result[disliked] > result[orderSoda] {
		class = disliked
		messres.Ans = "> Tendremos en cuenta su comentario y mejoraremos - %s\n\n" //, class)
		messres.Class = class
	} else if result[orderPizza] > result[orderHamburger] && result[orderPizza] > result[orderSalad] && result[orderPizza] > result[orderSoda] {
		class = orderPizza
		messres.Ans = "> Una pizza en camino - %s\n\n" //, class)
		messres.Class = class
	} else if result[orderHamburger] > result[orderSalad] && result[orderHamburger] > result[orderSoda] {
		class = orderHamburger
		messres.Ans = "> Una hamburguesa en camino - %s\n\n" //, class)
		messres.Class = class
	} else if result[orderSalad] > result[orderSoda] {
		class = orderSalad
		messres.Ans = "> Ok, se esta preparando una ensalada - %s\n\n" //, class)
		messres.Class = class
	} else {
		class = orderSoda
		messres.Ans = "> Ya le traemso su bebida - %s\n\n" //, class
		messres.Class = class
	}
	return messres
}
