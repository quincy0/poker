package poker

import (
	"fmt"
	"github.com/quincy0/poker/trie"
	"sort"
	"sync"
)

// Input 输入牌型
type Input struct {
	root *trie.Node
}

// InitMethod 初始化
func InitMethod() *Input {
	input := &Input{
		root: trie.NewTrieNode(0),
	}
	build(input.root, [][]int{}, map[int]bool{})
	return input
}

// Build 构建字典树
func build(node *trie.Node, cards [][]int, cardsIndex map[int]bool) {
	var wg sync.WaitGroup
	for i := 0; i < len(Pokers); i++ {
		if cardsIndex[i] {
			continue
		}
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			tempCards := make([][]int, len(cards))
			copy(tempCards, cards)
			cardsMap := make(map[int]bool)
			for k, v := range cardsIndex {
				cardsMap[k] = v
			}
			cardsMap[index] = true
			tempCards = append(tempCards, Pokers[index])
			rank := 0
			if len(tempCards) == 3 {
				rank, _ = getRank(tempCards)
			}
			key := fmt.Sprintf("%d%d", Pokers[index][0], Pokers[index][1])
			node.Add(key, rank)
			child, _ := node.Child(key)
			if rank == 0 {
				build(child, tempCards, cardsMap)
			}
		}(i)
	}
	wg.Wait()
}

// Compare 比较两副牌大小
func (in *Input) Compare(p1 [][]int, p2 [][]int) int {
	p1Rank := in.search(p1)
	p2Rank := in.search(p2)
	res := compare(p2Rank, p1Rank)
	if res != 0 {
		return res
	}

	score1, _ := classifyCards(p1)
	score2, _ := classifyCards(p2)
	return compareScore(p1Rank, score1, score2)
}

func (in *Input) search(p [][]int) int {
	node := in.root
	for i := 0; i < len(p); i++ {
		key := fmt.Sprintf("%d%d", p[i][0], p[i][1])
		node, _ = node.Child(key)
	}
	return node.Rank
}

/**
比较排面大小
score1 > score2 => 1
score1 == score2 => 0
score1 < score2 => -1
*/
func compareScore(rank int, score1, score2 []int) int {
	// 处理A
	handleA(rank, score1)
	handleA(rank, score2)

	// 豹子只比较一次
	if rank == RANK4 {
		return compare(score1[0], score2[0])
	}

	sort.Ints(score1)
	sort.Ints(score2)

	// 一对先比较对子
	if rank == RANK5 {
		pairNum1, single1 := getPairNum(score1)
		pairNum2, single2 := getPairNum(score2)
		res := compare(pairNum1, pairNum2)
		if res != 0 {
			return res
		}
		return compare(single1, single2)
	}

	res := compare(score1[2], score2[2])
	if res != 0 {
		return res
	}
	res = compare(score1[1], score2[1])
	if res != 0 {
		return res
	}
	return compare(score1[0], score2[0])
}

func compare(num1, num2 int) int {
	if num1 > num2 {
		return 1
	} else if num1 == num2 {
		return 0
	} else {
		return -1
	}
}

// 获取对子值及单值
func getPairNum(sortedScore []int) (int, int) {
	if sortedScore[0] == sortedScore[1] {
		if sortedScore[0] == 1 {
			sortedScore[0] = 14
		}
		return sortedScore[0], sortedScore[2]
	}
	return sortedScore[2], sortedScore[0]
}

func handleA(rank int, scores []int) {
	// A 2 3 不处理
	if rank == RANK1 || rank == RANK3 {
		if scores[0] == 1 && (scores[1] == 2 || scores[2] == 2) {
			return
		}
	}

	for i:=0; i<len(scores); i++ {
		if scores[i] == 1 {
			scores[i] =14
		}
	}
}
