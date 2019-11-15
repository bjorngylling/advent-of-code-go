package main

import (
	"testing"
)

func TestParseLn(t *testing.T) {
	playerCount, lastMarbleValue := parseLn("9 players; last marble is worth 25 points")
	if playerCount != 9 {
		t.Errorf("Expected playerCount to be 9 but was %+v\n", playerCount)
	}
	if lastMarbleValue != 25 {
		t.Errorf("Expected lastMableValue to be 25 but was %+v\n", lastMarbleValue)
	}
}

func TestScore(t *testing.T) {
	playerCount, lastMarbleValue, expectedScore := 9, 25, 32
	result := score(playerCount, lastMarbleValue)
	if result != expectedScore {
		t.Errorf("Expected score to be %+v but was %+v [playerCount=%d, lastMarbleValue=%d]\n",
			expectedScore, result, playerCount, lastMarbleValue)
	}
	playerCount, lastMarbleValue, expectedScore = 10, 1618, 8317
	result = score(playerCount, lastMarbleValue)
	if result != expectedScore {
		t.Errorf("Expected score to be %+v but was %+v [playerCount=%d, lastMarbleValue=%d]\n",
			expectedScore, result, playerCount, lastMarbleValue)
	}
	playerCount, lastMarbleValue, expectedScore = 13, 7999, 146373
	result = score(playerCount, lastMarbleValue)
	if result != expectedScore {
		t.Errorf("Expected score to be %+v but was %+v [playerCount=%d, lastMarbleValue=%d]\n",
			expectedScore, result, playerCount, lastMarbleValue)
	}
	playerCount, lastMarbleValue, expectedScore = 17, 1104, 2764
	result = score(playerCount, lastMarbleValue)
	if result != expectedScore {
		t.Errorf("Expected score to be %+v but was %+v [playerCount=%d, lastMarbleValue=%d]\n",
			expectedScore, result, playerCount, lastMarbleValue)
	}
	playerCount, lastMarbleValue, expectedScore = 21, 6111, 54718
	result = score(playerCount, lastMarbleValue)
	if result != expectedScore {
		t.Errorf("Expected score to be %+v but was %+v [playerCount=%d, lastMarbleValue=%d]\n",
			expectedScore, result, playerCount, lastMarbleValue)
	}
	playerCount, lastMarbleValue, expectedScore = 30, 5807, 37305
	result = score(playerCount, lastMarbleValue)
	if result != expectedScore {
		t.Errorf("Expected score to be %+v but was %+v [playerCount=%d, lastMarbleValue=%d]\n",
			expectedScore, result, playerCount, lastMarbleValue)
	}
}
