type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

type Rank uint8 

const (
	_ Rank = iota
	Ace,
	Two,
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit
	Rank 
}

