package main

import (
	"fmt"
	"math/rand"
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

func moveJoker(jokerValue int, step int, deck []int) []int {
	// 1. find joker index
	jokerIndex := findJokerIndex(deck, jokerValue)
	newJokerIndex := jokerIndex + step

	// 2. Validate that the final position is not the last or first position in deck
	if newJokerIndex+step >= len(deck) {
		newJokerIndex = 1
	}

	// 3. Move joker {STEP} steps
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
	partialDeck := deck[:refValue]

	partOne := deck[refValue : deckSize-1]
	fmt.Println(deck)

	fmt.Printf("Partial: %v\n", partialDeck)
	fmt.Printf("PartOne: %v\n", partOne)
	copy(deck[:], append(append(partOne, partialDeck...), refValue))

	return deck
}

func main() {
	DECK_SIZE := 10
	JOKER_A_VALUE := DECK_SIZE - 1
	JOKER_B_VALUE := DECK_SIZE

	// 1. Setup Deck
	deck := setupDeck(DECK_SIZE)

	// 2. Rotate A Joker
	deck = moveJoker(JOKER_A_VALUE, 1, deck)
	// 3. Rotate B Joker
	deck = moveJoker(JOKER_B_VALUE, 2, deck)

	// 4. Triple Cut
	deck = tripleCut(deck, JOKER_A_VALUE, JOKER_B_VALUE)
	// 5. Count Cut
	deck = countCut(deck, DECK_SIZE)

}
