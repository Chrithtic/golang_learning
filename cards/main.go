// How many times can you draw a King of Hearts from a 52 card deck, shuffled 20 times?

package main

import (
	"fmt"
	"math/rand"
)

type Card struct {
	Type  string
	Suite string
}

type Deck []Card

func main() {
	for i := 1; i <= 20; i++ {
		deck := new()
		shuffle(deck)
		prob(deck)
	}
}

func shuffle(deck Deck) Deck {
	for i := 0; i < len(deck); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			deck[r], deck[i] = deck[i], deck[r]
		}
	}
	return deck
}

func new() (deck Deck) {
	// List out all the card types in the slice
	types := []string{"two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack",
		"queen", "king", "ace"}
	// List of suites in the deck.
	suites := []string{"spades", "heart", "diamond", "club"}
	// Loop thorugh types and suites, append to the deck.
	for i := 0; i < len(types); i++ {
		for s := 0; s < len(suites); s++ {
			card := Card{
				Type:  types[i],
				Suite: suites[s],
			}

			deck = append(deck, card)
		}
	}
	return
}

func prob(deck Deck) {
	var prob float64
	for i, card := range deck {

		if card.Type == "king" && card.Suite == "heart" {
			// TODO:implement logic for prob i divided by 52
			prob = float64(i) / float64(len(deck))
			fmt.Printf("index; %v, card: %v\n", i, card)
			fmt.Printf("prob; %v\n", prob)
		}
	}
}
