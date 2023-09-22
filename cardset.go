package twoplustwogo

type CardSet struct {
	Cards []Card
}

func (cs *CardSet) NewCardSet(s string) {
	for i := 0; i < len(s); i += 2 {
		cs.Cards = append(cs.Cards, NewCard(s[i:i+2]))
	}
}

func EmptyCardSet() CardSet {
	return CardSet{Cards: make([]Card, 0)}
}
func NewCardSet() CardSet {
	return CardSet{Cards: make([]Card, 0)}
}

func (cs *CardSet) Print() {
	for i := 0; i < len(cs.Cards); i++ {
		cs.Cards[i].Print()
	}
}

func (cs *CardSet) Length() int {
	return len(cs.Cards)
}

func (cs *CardSet) Get(i int) Card {
	return cs.Cards[i]
}

func (cs *CardSet) Set(i int, card Card) {
	cs.Cards[i] = card
}

func (cs *CardSet) swap(i int, j int) {
	cs.Cards[i], cs.Cards[j] = cs.Cards[j], cs.Cards[i]
}

func (cs *CardSet) AddCard(card Card) {
	cs.Cards = append(cs.Cards, card)
}

func (cs *CardSet) AddCards(cards CardSet) {
	cs.Cards = append(cs.Cards, cards.Cards...)
}
func NewHandFromCodes(codes []int) CardSet {
	var hand CardSet
	for i := 0; i < len(codes); i++ {
		hand.AddCard(Card(FromCode(codes[i])))
	}
	return hand
}

//	func FromCards(cards []Card) CardSet {
//		var cardSet CardSet
//		cardSet.Cards = cards
//		return cardSet
//	}
func NewHand(cardString string) CardSet {
	var hand CardSet
	hand.NewCardSet(cardString)
	return hand
}
func NewBoard(cardString string) CardSet {
	var hand CardSet
	hand.NewCardSet(cardString)
	return hand
}

func (cs *CardSet) ToString() string {
	var s string
	for i := 0; i < len(cs.Cards); i++ {
		s += cs.Cards[i].ToString()
	}
	return s
}

func (cs *CardSet) RemoveCard(card Card) {
	for i := 0; i < cs.Length(); i++ {
		if cs.Get(i) == card {
			cs.Cards = append(cs.Cards[:i], cs.Cards[i+1:]...)
			return
		}
	}
}
