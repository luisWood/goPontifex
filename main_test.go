package main

import (
	"reflect"
	"testing"
)

func TestFindJoker(t *testing.T) {
	dummyDeck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	jokerValue := 10
	expectedIndex := 9

	actualIndex := findJokerIndex(dummyDeck, jokerValue)
	if expectedIndex != actualIndex {
		t.Errorf("\nTest: FIND_JOKER_INDEX \n want: %d;\n Actual: %d;\n", expectedIndex, actualIndex)
	}
}

func TestValidateJokerPosition(t *testing.T) {
	deckSize := 10
	// test_case_#1
	step := 2
	jokerInitialIndex := 2
	expectedFinalIndex := 4
	actualFinalIndex := validateJokerPosition(jokerInitialIndex, step, deckSize)

	if expectedFinalIndex != actualFinalIndex {
		t.Errorf("\nTest: VALIDATE_JOKER_POSITION\n TEST-CASE#1\n want: %d;\n Actual: %d;\n", expectedFinalIndex, actualFinalIndex)
	}

	// test_case_#2
	step = 2
	jokerInitialIndex = 8
	expectedFinalIndex = 1
	actualFinalIndex = validateJokerPosition(jokerInitialIndex, step, deckSize)

	if expectedFinalIndex != actualFinalIndex {
		t.Errorf("\nTest: VALIDATE_JOKER_POSITION\n TEST-CASE#2\n  want: %d;\n Actual: %d;\n", expectedFinalIndex, actualFinalIndex)
	}

	// test_case_#3
	step = 2
	jokerInitialIndex = 9
	expectedFinalIndex = 2
	actualFinalIndex = validateJokerPosition(jokerInitialIndex, step, deckSize)

	// test_case_#4
	step = 1
	jokerInitialIndex = 9
	expectedFinalIndex = 1
	actualFinalIndex = validateJokerPosition(jokerInitialIndex, step, deckSize)

	// test_case_#5
	step = 1
	jokerInitialIndex = 2
	expectedFinalIndex = 3
	actualFinalIndex = validateJokerPosition(jokerInitialIndex, step, deckSize)
	if expectedFinalIndex != actualFinalIndex {
		t.Errorf("\nTest: VALIDATE_JOKER_POSITION\n TEST-CASE#5\n want: %d;\n Actual: %d;\n", expectedFinalIndex, actualFinalIndex)
	}
}

func TestMoveJoker(t *testing.T) {
	// test_case_#1
	step := 2
	jokerValue := 10
	deck := []int{1, 2, 3, 4, 10, 5, 6, 7, 8, 9}
	expectedDeck := []int{1, 2, 3, 4, 5, 6, 10, 7, 8, 9}
	actualDeck := moveJoker(jokerValue, step, deck)

	if !reflect.DeepEqual(actualDeck, expectedDeck) {
		t.Errorf("\nTest: MOVE_JOKER\n TEST-CASE#1\n want: %d;\n Actual: %d;\n", expectedDeck[:10], actualDeck[:10])
	}

	// test_case_#2
	deck = []int{1, 10, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedDeck = []int{1, 2, 3, 10, 4, 5, 6, 7, 8, 9}
	actualDeck = moveJoker(jokerValue, step, deck)

	if !reflect.DeepEqual(actualDeck, expectedDeck) {
		t.Errorf("\nTest: MOVE_JOKER\n TEST-CASE#2\n want: %d;\n Actual: %d;\n", expectedDeck[:10], actualDeck[:10])
	}

	// test_case_#4
	deck = []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 9}
	expectedDeck = []int{1, 10, 2, 3, 4, 5, 6, 7, 8, 9}
	actualDeck = moveJoker(jokerValue, step, deck)

	if !reflect.DeepEqual(actualDeck, expectedDeck) {
		t.Errorf("\nTest: MOVE_JOKER\n TEST-CASE#4\n want: %d;\n Actual: %d;\n", expectedDeck[:10], actualDeck[:10])
	}

	// test_case_#5
	deck = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedDeck = []int{1, 2, 10, 3, 4, 5, 6, 7, 8, 9}
	actualDeck = moveJoker(jokerValue, step, deck)

	if !reflect.DeepEqual(actualDeck, expectedDeck) {
		t.Errorf("\nTest: MOVE_JOKER\n TEST-CASE#4\n want: %d;\n Actual: %d;\n", expectedDeck[:10], actualDeck[:10])
	}

}

func TestTripleCut(t *testing.T) {
	jokerAValue := 9
	jokerBvalue := 10

	dummyDeck := []int{1, 2, 9, 3, 4, 5, 6, 10, 7, 8}
	expectedDeck := []int{7, 8, 9, 3, 4, 5, 6, 10, 1, 2}

	actualDeck := tripleCut(dummyDeck, jokerAValue, jokerBvalue)
	if !reflect.DeepEqual(actualDeck, expectedDeck) {
		t.Errorf("\nTest: TRIPLE_CUT\n TEST-CASE#1\n want: %d;\n Actual: %d;\n", expectedDeck[:10], actualDeck[:10])
	}

	dummyDeck = []int{1, 2, 3, 4, 5, 6, 10, 7, 8, 9}
	expectedDeck = []int{10, 7, 8, 9, 1, 2, 3, 4, 5, 6}

	actualDeck = tripleCut(dummyDeck, jokerAValue, jokerBvalue)
	if !reflect.DeepEqual(actualDeck, expectedDeck) {
		t.Errorf("\nTest: TRIPLE_CUT\n TEST-CASE#1\n want: %d;\n Actual: %d;\n", expectedDeck[:10], actualDeck[:10])
	}

	dummyDeck = []int{1, 2, 10, 3, 4, 5, 6, 9, 7, 8}
	expectedDeck = []int{7, 8, 10, 3, 4, 5, 6, 9, 1, 2}

	actualDeck = tripleCut(dummyDeck, jokerAValue, jokerBvalue)
	if !reflect.DeepEqual(actualDeck, expectedDeck) {
		t.Errorf("\nTest: TRIPLE_CUT\n TEST-CASE#3\n want: %d;\n Actual: %d;\n", expectedDeck[:10], actualDeck[:10])
	}
}

func TestCountCut(t *testing.T) {
	deckSize := 10
	dummyDeck := []int{1, 2, 9, 3, 4, 5, 6, 10, 7, 8}
	expectedDeck := []int{7, 1, 2, 9, 3, 4, 5, 6, 10, 8}

	actualDeck := countCut(dummyDeck, deckSize)

	if !reflect.DeepEqual(actualDeck, expectedDeck) {
		t.Errorf("\nTest: COUNT_CUT\n TEST-CASE#1\n want: %d;\n Actual: %d;\n", expectedDeck[:10], actualDeck[:10])
	}

	dummyDeck = []int{1, 2, 9, 3, 4, 5, 6, 7, 8, 10}
	expectedDeck = []int{1, 2, 9, 3, 4, 5, 6, 7, 8, 10}

	actualDeck = countCut(dummyDeck, deckSize)

	if !reflect.DeepEqual(actualDeck, expectedDeck) {
		t.Errorf("\nTest: COUNT_CUT\n TEST-CASE#2\n want: %d;\n Actual: %d;\n", expectedDeck[:10], actualDeck[:10])
	}
}

func TestGetKeyStream(t *testing.T) {
	dummyDeck := []int{2, 9, 1, 3, 4, 5, 6, 10, 7, 8}

	expectedKeystream := 1
	actualKeystream := getKeystream(dummyDeck)

	if expectedKeystream != actualKeystream {
		t.Errorf("\nTest: GET_KEY_STREAM\n TEST-CASE#1\n want: %d;\n Actual: %d;\n", expectedKeystream, actualKeystream)
	}

	dummyDeck = []int{7, 2, 1, 3, 9, 4, 5, 6, 10, 8}

	expectedKeystream = 6
	actualKeystream = getKeystream(dummyDeck)

	if expectedKeystream != actualKeystream {
		t.Errorf("\nTest: GET_KEY_STREAM\n TEST-CASE#2\n want: %d;\n Actual: %d;\n", expectedKeystream, actualKeystream)
	}
}

func TestNumericToAlphabet(t *testing.T) {
	dummyMessage := []int{1, 2, 3, 4}
	expectedDecryptedMessage := []string{"a", "b", "c", "d"}
	actualDecryptedMessage := numericToAlphabetic(dummyMessage)

	if !reflect.DeepEqual(actualDecryptedMessage, expectedDecryptedMessage) {
		t.Errorf("\nTest: NUMERIC_TO_ALPHABET\n TEST-CASE#1\n want: %s;\n Actual: %s;\n", expectedDecryptedMessage[:4], actualDecryptedMessage[:4])
	}

	dummyMessage = []int{26, 52, 78}
	expectedDecryptedMessage = []string{"z", "z", "z"}
	actualDecryptedMessage = numericToAlphabetic(dummyMessage)
	if !reflect.DeepEqual(actualDecryptedMessage, expectedDecryptedMessage) {
		t.Errorf("\nTest: NUMERIC_TO_ALPHABET\n TEST-CASE#2\n want: %s;\n Actual: %s;\n", expectedDecryptedMessage[:3], actualDecryptedMessage[:3])
	}
}

func TestAlphabetToNumeric(t *testing.T) {
	dummyMessage := "test"
	expectedResult := []int{20, 5, 19, 20}

	actualResult := alphabeticToNumeric(dummyMessage)

	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Errorf("\nTest: ALPHABET_TO_NUMERIC\n TEST-CASE#1\n want: %d;\n Actual: %d;\n", expectedResult[:4], actualResult[:4])
	}

	dummyMessage = "abcdefghijklmnopqrstuvwxyz"
	expectedResult = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}
	actualResult = alphabeticToNumeric(dummyMessage)

	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Errorf("\nTest: ALPHABET_TO_NUMERIC\n TEST-CASE#1\n want: %d;\n Actual: %d;\n", expectedResult[:25], actualResult[:25])
	}
}
