package card

type CardSet struct {
	Cards []Card
}

func (c *CardSet) FromString(s string) {
	for i := 0; i < len(s); i += 2 {
		c.Cards = append(c.Cards, FromString(s[i:i+2]))
	}
}

func NewCardSet(n int) CardSet {
	return CardSet{Cards: make([]Card, n)}
}

func (c *CardSet) Print() {
	for i := 0; i < len(c.Cards); i++ {
		c.Cards[i].Print()
	}
}

func (c *CardSet) Length() int {
	return len(c.Cards)
}

func (c *CardSet) Get(i int) Card {
	return c.Cards[i]
}

func (c *CardSet) Set(i int, card Card) {
	c.Cards[i] = card
}

func (c *CardSet) swap(i int, j int) {
	c.Cards[i], c.Cards[j] = c.Cards[j], c.Cards[i]
}

func (c *CardSet) AddCard(card Card) {
	c.Cards = append(c.Cards, card)
}

func (c *CardSet) AddCards(cards CardSet) {
	c.Cards = append(c.Cards, cards.Cards...)
}
