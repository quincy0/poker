package poker

import "sort"

const (
	// Diamond 方块
	Diamond = 1
	// Heart 红心
	Heart = 2
	// Club 梅花
	Club = 3
	// Spade 黑桃
	Spade = 4

	// RANK1 同花顺
	RANK1 = 1
	// RANK2 同花
	RANK2 = 2
	// RANK3 顺子
	RANK3 = 3
	// RANK4 豹子
	RANK4 = 4
	// RANK5 对子
	RANK5 = 5
	// RANK6 单张
	RANK6 = 6
)

// Pokers 扑克牌
var Pokers = [][]int{
	{13, 1}, {12, 1}, {11, 1}, {10, 1}, {9, 1}, {8, 1}, {7, 1}, {6, 1}, {5, 1}, {4, 1}, {3, 1}, {2, 1}, {1, 1},
	{13, 2}, {12, 2}, {11, 2}, {10, 2}, {9, 2}, {8, 2}, {7, 2}, {6, 2}, {5, 2}, {4, 2}, {3, 2}, {2, 2}, {1, 2},
	{13, 3}, {12, 3}, {11, 3}, {10, 3}, {9, 3}, {8, 3}, {7, 3}, {6, 3}, {5, 3}, {4, 3}, {3, 3}, {2, 3}, {1, 3},
	{13, 4}, {12, 4}, {11, 4}, {10, 4}, {9, 4}, {8, 4}, {7, 4}, {6, 4}, {5, 4}, {4, 4}, {3, 4}, {2, 4}, {1, 4},
}

func getRank(cards [][]int) (int, bool) {
	if len(cards) != 3 {
		return 0, false
	}
	scores, suits := classifyCards(cards)
	sort.Ints(scores)
	if isSameSuit(suits) && isStraight(scores) {
		return RANK1, true
	}
	if isSameSuit(suits) {
		return RANK2, true
	}
	if isStraight(scores) {
		return RANK3, true
	}
	if isSameScore(scores) {
		return RANK4, true
	}
	if isPair(scores) {
		return RANK5, true
	}
	return RANK6, true
}

// 是否是同花
func isSameSuit(suits []int) bool {
	return suits[0] == suits[1] && suits[0] == suits[2]
}

// 是否是顺子
func isStraight(sortedScores []int) bool {
	// A K Q
	if sortedScores[0] == 1 && sortedScores[1] == 12 && sortedScores[2] == 13 {
		return true
	}
	return (sortedScores[0]+1 == sortedScores[1]) && (sortedScores[1]+1 == sortedScores[2])
}

// 是否是豹子
func isSameScore(sortedScores []int) bool {
	return sortedScores[0] == sortedScores[1] && sortedScores[1] == sortedScores[2]
}

// 是否是对子
func isPair(sortedScores []int) bool {
	return sortedScores[0] == sortedScores[1] || sortedScores[1] == sortedScores[2]
}

func classifyCards(cards [][]int) ([]int, []int) {
	scores, suits := make([]int, 3), make([]int, 3)
	for i := 0; i < len(cards); i++ {
		scores[i] = cards[i][0]
		suits[i] = cards[i][1]
	}
	return scores, suits
}
