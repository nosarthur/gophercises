package deck

const (
	spade = iota
	diamond
	club
	heart
)

type Card struct {
	Suit  Suit
	Value uint
}
