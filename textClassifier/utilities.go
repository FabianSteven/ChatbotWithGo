package main

import (
	"math"
	"regexp"
	"strings"
)

/*
 * Utilidades
 */

// Las stopwords son palabras que tienen muy poco significado
var stopwords = map[string]struct{}{
	"a": struct{}{}, "aca": struct{}{}, "ahi": struct{}{}, "al": struct{}{}, "algo": struct{}{}, "algun": struct{}{}, "alguno": struct{}{},
	"algunos": struct{}{}, "alguna": struct{}{}, "algunas": struct{}{}, "alla": struct{}{}, "alli": struct{}{}, "ambos": struct{}{},
	"ante": struct{}{}, "antes": struct{}{}, "aquel": struct{}{}, "aquella": struct{}{}, "aquellas": struct{}{}, "aquello": struct{}{}, "aquellos": struct{}{},
	"aqui": struct{}{}, "arriba": struct{}{}, "asi": struct{}{}, "atras": struct{}{}, "aun": struct{}{}, "aunque": struct{}{}, "bien": struct{}{},
	"cada": struct{}{}, "casi": struct{}{}, "como": struct{}{}, "con": struct{}{}, "cual": struct{}{}, "cuales": struct{}{}, "cualquier": struct{}{},
	"cualquiera": struct{}{}, "cuando": struct{}{}, "cuanto": struct{}{}, "cuan": struct{}{}, "cuantos": struct{}{}, "cuantas": struct{}{}, "de": struct{}{},
	"del": struct{}{}, "demas": struct{}{}, "desde": struct{}{}, "donde": struct{}{}, "dos": struct{}{}, "el": struct{}{}, "ella": struct{}{},
	"ellas": struct{}{}, "ellos": struct{}{}, "en": struct{}{}, "eres": struct{}{}, "esas": struct{}{}, "y": struct{}{}, "eso": struct{}{},
	"ese": struct{}{}, "esos": struct{}{}, "after": struct{}{}, "estas": struct{}{}, "esta": struct{}{}, "este": struct{}{}, "esto": struct{}{},
	"estos": struct{}{}, "etc": struct{}{}, "ha": struct{}{}, "hasta": struct{}{}, "la": struct{}{}, "las": struct{}{}, "los": struct{}{},
	"lo": struct{}{}, "mis": struct{}{}, "mi": struct{}{}, "mientras": struct{}{}, "muy": struct{}{}, "nosotros": struct{}{},
	"nosotras": struct{}{}, "nuestra": struct{}{}, "nuestro": struct{}{}, "nuestras": struct{}{}, "nuestros": struct{}{}, "otra": struct{}{}, "otros": struct{}{},
	"otro": struct{}{}, "para": struct{}{}, "pero": struct{}{}, "pues": struct{}{}, "que": struct{}{}, "si": struct{}{}, "siempre": struct{}{},
	"siendo": struct{}{}, "sin": struct{}{}, "sino": struct{}{}, "sobre": struct{}{}, "sr": struct{}{}, "sra": struct{}{}, "su": struct{}{},
	"sus": struct{}{}, "te": struct{}{}, "tu": struct{}{}, "tus": struct{}{}, "un": struct{}{}, "una": struct{}{}, "uno": struct{}{},
	"unas": struct{}{}, "unos": struct{}{}, "usted": struct{}{}, "ustedes": struct{}{}, "vosotros": struct{}{}, "vosotras": struct{}{}, "vuestra": struct{}{},
	"vuestro": struct{}{}, "vuestras": struct{}{}, "vuestros": struct{}{}, "ya": struct{}{}, "esa": struct{}{}, "yo": struct{}{},
}

func isStopword(w string) bool {
	_, ok := stopwords[w]
	return ok
}

// limpiar los caracteres que no son números y reducir el tamaño de los mismos
func cleanup(sentence string) string {
	re := regexp.MustCompile("[^a-zA-Z 0-9]+")
	return re.ReplaceAllString(strings.ToLower(sentence), "")
}

// tokenize crear un array de palabras a partir de una frase
func tokenize(sentence string) []string {
	s := cleanup(sentence)
	words := strings.Fields(s)
	var tokens []string
	for _, w := range words {
		if !isStopword(w) {
			tokens = append(tokens, w)
		}
	}
	return tokens
}

// zeroOneTransform devuelve
// 0 si el argumento x = 0
// 1 en caso contrario
func zeroOneTransform(x int) int {
	return int(math.Ceil(float64(x) / (float64(x) + 1.0)))
}
