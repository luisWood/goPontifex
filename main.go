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
			return 1
		}
	} else if step == 2 {
		if index == deckSize-2 {
			return 1
		} else if index == deckSize-1 {
			return 2
		}
	}
	return index + step
}

func moveJoker(jokerValue int, step int, deck []int) []int {
	// 1. find joker index
	jokerIndex := findJokerIndex(deck, jokerValue)

	// 2. Validate that the final position is not the last or first position in deck
	newJokerIndex := validateJokerPosition(jokerIndex, step, len(deck))

	// 3. Move joker {STEP} steps

	if jokerIndex < newJokerIndex {
		partA := deck[:jokerIndex]
		partB := deck[jokerIndex+1 : newJokerIndex+1]
		copy(deck[:], append(partA, partB...))
		deck[newJokerIndex] = jokerValue

	} else if jokerIndex > newJokerIndex {
		// deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 9}
		//  0  1  2  3  4  5  6  7  8  9
		//   1, 10, 2, 3, 4, 5, 6, 7, 8, 9
		//jokerIndex = 7
		//
		partA := deck[:newJokerIndex]
		partB := deck[newJokerIndex-1 : jokerIndex]
		copy(deck[:], append(partA, partB...))
		deck[newJokerIndex] = jokerValue

	}

	// deck[newJokerIndex] = jokerValue
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
	// keystream := -1

	if deck[0] == len(deck) && deck[len(deck)-1] <= len(deck)-1 {
		return deck[len(deck)-1]
	} else if deck[0] == len(deck) && deck[len(deck)-1] >= len(deck)-1 {
		return -1
	} else {
		return deck[deck[0]]
	}
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
	message = strings.ToLower(message)
	for _, char := range message {
		if char < 'a' || char > 'z' {
			continue // Skip non-alphabetic characters
		}
		result = ((int(char-'a') + 1) % 27)
		results = append(results, result)
	}

	return results
}

func encrypt(numericMessage []int) struct {
	message   []int
	deck      []int
	keystream []int
} {
	DECK_SIZE := 10
	JOKER_A_VALUE := DECK_SIZE - 1
	JOKER_B_VALUE := DECK_SIZE
	initialDeck := setupDeck(DECK_SIZE)

	deck := make([]int, len(initialDeck))
	copy(deck, initialDeck)

	encryptedNumericMessage := []int{}
	keystreamSlice := []int{}

	for _, value := range numericMessage {
		deck = moveJoker(JOKER_A_VALUE, 1, deck)
		deck = moveJoker(JOKER_B_VALUE, 2, deck)
		deck = tripleCut(deck, JOKER_A_VALUE, JOKER_B_VALUE)
		deck = countCut(deck, DECK_SIZE)

		keystream := -1
		keystream = getKeystream(deck)
		if keystream != -1 {
			keystreamSlice = append(keystreamSlice, keystream)
			encryptedNumericMessage = append(encryptedNumericMessage, value+keystream)
		} else {
			deck = moveJoker(JOKER_A_VALUE, 1, deck)
			deck = moveJoker(JOKER_B_VALUE, 2, deck)
			deck = tripleCut(deck, JOKER_A_VALUE, JOKER_B_VALUE)
			deck = countCut(deck, DECK_SIZE)

			keystream = getKeystream(deck)
			keystreamSlice = append(keystreamSlice, keystream)
			encryptedNumericMessage = append(encryptedNumericMessage, value+keystream)
		}

	}

	result := struct {
		message   []int
		deck      []int
		keystream []int
	}{
		encryptedNumericMessage,
		initialDeck,
		keystreamSlice,
	}

	return result
}

func decrypt(message []int, deck []int) struct {
	message   []int
	keystream []int
} {
	DECK_SIZE := len(deck)
	JOKER_A_VALUE := DECK_SIZE - 1
	JOKER_B_VALUE := DECK_SIZE

	keystreamSlice := []int{}
	decryptedNumericMessage := []int{}
	for _, value := range message {
		deck = moveJoker(JOKER_A_VALUE, 1, deck)
		deck = moveJoker(JOKER_B_VALUE, 2, deck)
		deck = tripleCut(deck, JOKER_A_VALUE, JOKER_B_VALUE)
		deck = countCut(deck, DECK_SIZE)

		// keystream := getKeystream(deck)

		keystream := -1

		keystream = getKeystream(deck)

		if keystream != -1 {
			keystreamSlice = append(keystreamSlice, keystream)
			decryptedNumericMessage = append(decryptedNumericMessage, value-keystream)
		} else {
			deck = moveJoker(JOKER_A_VALUE, 1, deck)
			deck = moveJoker(JOKER_B_VALUE, 2, deck)
			deck = tripleCut(deck, JOKER_A_VALUE, JOKER_B_VALUE)
			deck = countCut(deck, DECK_SIZE)

			keystream = getKeystream(deck)
			keystreamSlice = append(keystreamSlice, keystream)
			decryptedNumericMessage = append(decryptedNumericMessage, value-keystream)
		}

		// keystreamSlice = append(keystreamSlice, keystream)
	}

	result := struct {
		message   []int
		keystream []int
	}{
		decryptedNumericMessage,
		keystreamSlice,
	}

	return result
}

func main() {
	example := "this is a test"
	numericMessage := alphabeticToNumeric(example)

	result := encrypt(numericMessage)
	deck := result.deck

	encryptedMessage := result.message
	decryptedResult := decrypt(encryptedMessage, deck)
	fmt.Println("keystream: ", decryptedResult)

	decryptedNumericMessage := decryptedResult.message
	decryptedMessage := numericToAlphabetic(decryptedNumericMessage)

	fmt.Println("Original:", example)
	fmt.Println("Encrypted:", numericToAlphabetic(encryptedMessage))
	fmt.Println("Decrypted:", (decryptedMessage))
	fmt.Println("Encryption Keystream:", result.keystream)
	fmt.Println("Decryption Keystream:", decryptedResult.keystream)
	// fmt.Println(numericMessage)
	fmt.Println(result.keystream)
	fmt.Println(decryptedResult.keystream)
}
