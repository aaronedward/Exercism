package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine":
		return 9
	case card == "ten" || card == "jack" || card == "queen" || card == "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var handSum = ParseCard(card1) + ParseCard(card2)
	var dealerVal = ParseCard(dealerCard)
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case handSum == 21 && (dealerVal != 11 && dealerVal != 10):
		return "W"
	case handSum == 21 && (dealerVal == 11 || dealerVal == 10):
		return "S"
	case handSum >= 17 && handSum <= 20:
		return "S"
	case handSum >= 12 && handSum <= 16 && dealerVal < 7:
		return "S"
	case handSum >= 12 && handSum <= 16 && dealerVal >= 7:
		return "H"
	case handSum <= 11:
		return "H"
	default:
		return "S"
	}
}
