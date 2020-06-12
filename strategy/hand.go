package strategy

type Hand struct {
	handValue int
}

const (
	HANDVALUE_GUU = iota
	HANDVALUE_CHO
	HANDVALUE_PAA
)

var hands = []*Hand{
	{HANDVALUE_GUU},
	{HANDVALUE_CHO},
	{HANDVALUE_PAA},
}

func hand(handValue int) *Hand {
	return hands[handValue]
}

func (h *Hand) IsStrongerThan(hand *Hand) bool {
	return h.fight(hand) == 1
}

func (h *Hand) ISWeakerThan(hand *Hand) bool {
	return h.fight(hand) == -1
}

func (h *Hand) fight(hand *Hand) int {
	if h == hand {
		return 0
	} else if (h.handValue+1)%3 == hand.handValue {
		return 1
	} else {
		return -1
	}
}
