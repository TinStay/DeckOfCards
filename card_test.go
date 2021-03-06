package deck

import (
	"fmt"
	"math/rand"
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

func TestShuffle(t *testing.T) {
	// First call to shuffleRand.Perm(52) should be: [40 35 ...]
	// make shuffleRand deterministic
	shuffleRand = rand.New(rand.NewSource(0))

	// Create decks
	cards := NewDeck()
	shuffled := Shuffle(cards)

	first := cards[40]
	second := cards[35]

	if first != shuffled[0]{
		t.Errorf("Expected %s. Received: %s", first, shuffled[0])
	} 
	if second != shuffled[1]{
		t.Errorf("Expected %s. Received: %s", second, shuffled[1])
	} 
}

func TestJokers(t *testing.T){
	cards := NewDeck(Jokers(3))
	count := 0

	for _, card := range cards {
		if(card.Suit == Joker){
			count++
		}
	}

	if(count != 3){
		t.Error("Expected 3 Jokers. Received:", count)
	}

}

func TestFilter(t *testing.T){
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}

	cards := NewDeck(Filter(filter))

	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three{
			t.Error("Expected all 2's and 3's to be filtered out")
		}
	}
}

func TestDeckNumbers(t *testing.T){
	cards := NewDeck(DeckNumber(3))

	// 13 ranks * 3 suits * 3 decks
	if len(cards) != 13*4*3 {
		t.Error("Expected 52*3. Received:", len(cards))
	}
}