package main

func solution(lottos []int, winNums []int) []int {
	count := 0
	zeroCount := 0
	for _, lottoNum := range lottos {
		if lottoNum == 0 {
			zeroCount++
			continue
		}
		for _, winNum := range winNums {
			if lottoNum == winNum {
				count++
			}
		}
	}
	topCollect := count + zeroCount
	downCollect := count
	if topCollect < 2 {
		topCollect = 1
	}
	if downCollect < 2 {
		downCollect = 1
	}
	topRank := 7 - topCollect
	downRank := 7 - downCollect

	return []int{topRank, downRank}
}
