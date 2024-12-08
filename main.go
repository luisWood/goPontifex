package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func setupDeck(deckSize int) []int {
	initialDeck := make([]int, deckSize)

	for i := range initialDeck {
		initialDeck[i] = i + 1
	}

	rand.Shuffle(deckSize, func(i, j int) {
		initialDeck[i], initialDeck[j] = initialDeck[j], initialDeck[i]
	})

	return initialDeck
}

func findJokerIndex(deck []int, jokerValue int) int {
	jokerIndex := -1
	for index, value := range deck {
		if value == jokerValue {
			jokerIndex = index
		}
	}

	return jokerIndex
}

func validateJokerPosition(index int, step int, deckSize int) int {
	if step == 1 {
		if index == deckSize-1 {
			return step
		}
	} else if step == 2 {
		if index == deckSize-2 {
			return step
		} else if index == deckSize-1 {
			return step - 1
		}
	}
	return index + step
}

func moveJoker(jokerValue int, step int, deck []int) []int {
	// 1. find joker index
	jokerIndex := findJokerIndex(deck, jokerValue)

	// 2. Validate that the final position is not the last or first position in deck
	newJokerIndex := validateJokerPosition(jokerIndex, step, len(deck))

	// 3. Move joker {STEP} steps FAILING, NEED TO FIX TODO

	if jokerIndex < newJokerIndex {
		copy(deck[jokerIndex:newJokerIndex], deck[jokerIndex+1:newJokerIndex+1])
	} else if jokerIndex > newJokerIndex {
		copy(deck[newJokerIndex+1:jokerIndex+1], deck[newJokerIndex:jokerIndex])
	}

	deck[newJokerIndex] = jokerValue
	return deck
}

func tripleCut(deck []int, jokerAValue int, jokerBValue int) []int {
	jokerAIndex := findJokerIndex(deck, jokerAValue)
	jokerBIndex := findJokerIndex(deck, jokerBValue)

	if jokerAIndex > jokerBIndex {
		tempJokerAIndex := jokerBIndex
		jokerBIndex = jokerAIndex
		jokerAIndex = tempJokerAIndex
	}

	partOne := deck[:jokerAIndex]
	partTwo := deck[jokerAIndex : jokerBIndex+1]
	partThree := deck[jokerBIndex+1:]

	copy(deck[:], append(append(partThree, partTwo...), partOne...))

	return deck
}

func countCut(deck []int, deckSize int) []int {
	refValue := deck[deckSize-1]
	if refValue >= deckSize-1 {
		refValue = deckSize - 1
	}

	partialDeck := deck[:refValue]
	partOne := deck[refValue : deckSize-1]
	copy(deck[:], append(append(partOne, partialDeck...), deck[deckSize-1]))
	return deck
}

func getKeystream(deck []int) int {
	keystream := deck[deck[0]-1]
	return keystream
}

func numericToAlphabetic(numericMessage []int) []string {
	results := []string{}

	for _, num := range numericMessage {
		if num < 1 {
			continue // Skip invalid numbers
		}
		letter := string(rune(((num-1)%26 + 65)))
		results = append(results, strings.ToLower(letter))
	}
	return results
}

func alphabeticToNumeric(message string) []int {
	results := []int{}
	result := 0
	for _, char := range message {
		if char < 'a' || char > 'z' {
			continue // Skip non-alphabetic characters
		}
		result = ((int(char-'a') + 1) % 27)
		results = append(results, result)
	}

	return results
}

func encrypt(numericMessage []int) []int {
	DECK_SIZE := 10
	JOKER_A_VALUE := DECK_SIZE - 1
	JOKER_B_VALUE := DECK_SIZE
	initialDeck := setupDeck(DECK_SIZE)
	deck := initialDeck

	encryptedNumericMessage := []int{}
	for value := range numericMessage {
		deck = moveJoker(JOKER_A_VALUE, 1, deck)
		deck = moveJoker(JOKER_B_VALUE, 2, deck)
		deck = tripleCut(deck, JOKER_A_VALUE, JOKER_B_VALUE)
		deck = countCut(deck, DECK_SIZE)

		keystream := getKeystream(deck)
		encryptedNumericMessage = append(encryptedNumericMessage, value+keystream)
	}

	return encryptedNumericMessage
}

func main() {
	example := "Hello World"
	numericMessage := alphabeticToNumeric(example)
	encryptedNumericMessage := encrypt(numericMessage)
	fmt.Println(encryptedNumericMessage)
}
