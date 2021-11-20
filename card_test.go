package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Club})
	fmt.Println(Card{Suit: Joker})
	fmt.Println(Card{Rank: Nine, Suit: Spade})
	fmt.Println(Card{Rank: King, Suit: Diamond})

	// Output:
	// Ace of Hearts
	// Two of Clubs
	// Joker
	// Nine of Spades
	// King of Diamonds
}

func TestNew(t *testing.T){
	cards := NewDeck()

	if len(cards) != 13*4 {
		t.Error("Length of cards is not equal to 52.")
	}
	
}
// Sorting
func TestDefaultSort(t *testing.T) {
	cards := NewDeck(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}

	if cards[0] != exp{
		t.Error("Expected Ace of Spades: Receiced:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := NewDeck(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}

	if cards[0] != exp{
		t.Error("Expected Ace of Spades: Receiced:", cards[0])
	}
}