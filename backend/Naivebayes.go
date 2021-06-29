package main

//Valores de las frases
const (
	greeting       = "greeting"
	liked          = "liked"
	disliked       = "disliked"
	orderPizza     = "orderPizza"
	orderHamburger = "orderHamburger"
	orderSalad     = "orderSalad"
	orderSoda      = "orderSoda"
)

/*
 * Clasificador
 */

type wordFrequency struct {
	word    string
	counter map[string]int
}

type classifier struct {
	dataset map[string][]string
	words   map[string]wordFrequency
}

// newClassifier devuelve un nuevo clasificador con el conjunto de datos y las palabras vacías
func newClassifier() *classifier {
	c := new(classifier)
	c.dataset = map[string][]string{
		greeting:       []string{},
		liked:          []string{},
		disliked:       []string{},
		orderPizza:     []string{},
		orderHamburger: []string{},
		orderSalad:     []string{},
		orderSoda:      []string{},
	}
	c.words = map[string]wordFrequency{}
	return c
}

// train rellena el conjunto de datos y palabras de un clasificador con el conjunto de datos de entrada map

func (c *classifier) train(dataset map[string]string) {
	for sentence, class := range dataset {
		c.addSentence(sentence, class)
		words := tokenize(sentence)
		for _, w := range words {
			c.addWord(w, class)
		}
	}
}

// classify devuelve las probabilidades de que una frase sea de cada clase
func (c classifier) classify(sentence string) map[string]float64 {
	words := tokenize(sentence)
	greetProb := c.probability(words, greeting)
	likedProb := c.probability(words, liked)
	dislikedProb := c.probability(words, disliked)
	pizaProb := c.probability(words, orderPizza)
	hamburgerProb := c.probability(words, orderHamburger)
	saladProb := c.probability(words, orderSalad)
	sodaProb := c.probability(words, orderSoda)
	return map[string]float64{
		greeting:       greetProb,
		liked:          likedProb,
		disliked:       dislikedProb,
		orderPizza:     pizaProb,
		orderHamburger: hamburgerProb,
		orderSalad:     saladProb,
		orderSoda:      sodaProb,
	}
}

// addSentence añade una frase y su clase al mapa del conjunto de datos de un clasificador
func (c *classifier) addSentence(sentence, class string) {
	c.dataset[class] = append(c.dataset[class], sentence)
}

// addSentence añade una palabra al mapa de palabras de un clasificador y actualiza su frecuencia
func (c *classifier) addWord(word, class string) {
	wf, ok := c.words[word]
	if !ok {
		wf = wordFrequency{word: word, counter: map[string]int{
			greeting:       0,
			liked:          0,
			disliked:       0,
			orderPizza:     0,
			orderHamburger: 0,
			orderSalad:     0,
			orderSoda:      0,
		}}
	}
	wf.counter[class]++
	c.words[word] = wf
}

// priorProb devuelve la probabilidad previa de cada clase del clasificador
func (c classifier) priorProb(class string) float64 {
	return float64(len(c.dataset[class])) / float64(len(c.dataset[greeting])+len(c.dataset[liked])+len(c.dataset[disliked])+len(c.dataset[orderPizza])+len(c.dataset[orderHamburger])+len(c.dataset[orderSalad])+len(c.dataset[orderSoda]))
}

// totalWordCount devuelve el recuento de palabras de una clase (duplicada también cuenta)
func (c classifier) totalWordCount(class string) int {
	greetCount := 0
	likedCount := 0
	dislikedCount := 0
	pizaCount := 0
	hamburgerCount := 0
	saladCount := 0
	sodaCount := 0
	for _, wf := range c.words {
		greetCount += wf.counter[greeting]
		likedCount += wf.counter[liked]
		dislikedCount += wf.counter[disliked]
		pizaCount += wf.counter[orderPizza]
		hamburgerCount += wf.counter[orderHamburger]
		saladCount += wf.counter[orderSalad]
		sodaCount += wf.counter[orderSoda]
	}
	if class == greeting {
		return greetCount
	} else if class == liked {
		return likedCount
	} else if class == disliked {
		return dislikedCount
	} else if class == orderPizza {
		return pizaCount
	} else if class == orderHamburger {
		return hamburgerCount
	} else if class == orderSalad {
		return saladCount
	} else if class == orderSoda {
		return sodaCount
	} else {
		return greetCount + likedCount + dislikedCount + pizaCount + hamburgerCount + saladCount + sodaCount
	}
}

// totalDistinctWordCount devuelve el número de palabras distintas en el conjunto de datos
func (c classifier) totalDistinctWordCount() int {
	greetCount := 0
	likedCount := 0
	dislikedCount := 0
	pizaCount := 0
	hamburgerCount := 0
	saladCount := 0
	sodaCount := 0
	for _, wf := range c.words {
		greetCount += zeroOneTransform(wf.counter[greeting])
		likedCount += zeroOneTransform(wf.counter[liked])
		dislikedCount += zeroOneTransform(wf.counter[disliked])
		pizaCount += zeroOneTransform(wf.counter[orderPizza])
		hamburgerCount += zeroOneTransform(wf.counter[orderHamburger])
		saladCount += zeroOneTransform(wf.counter[orderSalad])
		sodaCount += zeroOneTransform(wf.counter[orderSoda])
	}
	return greetCount + likedCount + dislikedCount + pizaCount + hamburgerCount + saladCount + sodaCount
}

// la probabilidad retoma la probabilidad de que una lista de palabras esté en una clase
func (c classifier) probability(words []string, class string) float64 {
	prob := c.priorProb(class)
	for _, w := range words {
		count := 0
		if wf, ok := c.words[w]; ok {
			count = wf.counter[class]
		}
		prob *= (float64((count + 1)) / float64((c.totalWordCount(class) + c.totalDistinctWordCount())))
	}
	for _, w := range words {
		count := 0
		if wf, ok := c.words[w]; ok {
			count += (wf.counter[greeting] + wf.counter[liked] + wf.counter[disliked] + wf.counter[orderPizza] + wf.counter[orderHamburger] + wf.counter[orderSalad] + wf.counter[orderSoda])
		}
		prob /= (float64((count + 1)) / float64((c.totalWordCount("") + c.totalDistinctWordCount())))
	}
	return prob
}
