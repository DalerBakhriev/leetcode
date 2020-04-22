package binarywatch

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func backTrack(
	allPossibleMinutes map[int]int,
	usedMinutesSum map[int]int,
	currIterNum int,
	currMinutesSum int,
	minutesCandidates []int,
	startIndex int,
	maxIters int,
	freqMap map[int]int,
) {
	if currIterNum == maxIters {
		_, alreadyUsed := allPossibleMinutes[currMinutesSum]
		_, usedByLowerNum := usedMinutesSum[currMinutesSum]
		if !alreadyUsed && !usedByLowerNum && currMinutesSum < 720 {
			allPossibleMinutes[currMinutesSum] = 1
			usedMinutesSum[currMinutesSum] = maxIters
		}
		return
	}

	for ind := startIndex; ind < len(minutesCandidates); ind++ {
		minuteCandidate := minutesCandidates[ind]
		if freqMap[minuteCandidate] > 0 {
			freqMap[minuteCandidate]--
			backTrack(
				allPossibleMinutes,
				usedMinutesSum,
				currIterNum+1,
				currMinutesSum+minuteCandidate,
				minutesCandidates,
				startIndex+1,
				maxIters,
				freqMap,
			)
			freqMap[minuteCandidate]++
		}

	}
}

func convertMinutesToTime(timeInMinutes map[int]int) []string {

	uniqueTimeAsString := make(map[string]int)
	timeAsStringArr := make([]string, 0)

	for totalMinutes := range timeInMinutes {
		hours, minutes := totalMinutes/60, totalMinutes%60

		hoursAsString := strconv.Itoa(hours)
		minutesAsString := strconv.Itoa(minutes)

		if len(minutesAsString) == 1 {
			minutesAsString = strings.Join([]string{"0", minutesAsString}, "")
		}
		time := strings.Join([]string{hoursAsString, ":", minutesAsString}, "")
		if _, alreadySaved := uniqueTimeAsString[time]; !alreadySaved {
			uniqueTimeAsString[time] = 1
		}
	}

	for time := range uniqueTimeAsString {
		timeAsStringArr = append(timeAsStringArr, time)
	}

	return timeAsStringArr
}

// ReadBinaryWatch ...
func ReadBinaryWatch(num int) []string {

	if num == 0 {
		return []string{"0:00"}
	}

	minutes := []int{1, 2, 4, 8, 16, 32, 60, 120, 240, 480}
	freqMap := map[int]int{
		1:   1,
		2:   1,
		4:   1,
		8:   1,
		16:  1,
		32:  1,
		60:  1,
		120: 1,
		240: 1,
		480: 1,
	}

	usedByLowerNum := make(map[int]int)
	for i := 1; i < num; i++ {
		allPossibleMinutes := make(map[int]int)
		backTrack(
			allPossibleMinutes,
			usedByLowerNum,
			0,
			0,
			minutes,
			0,
			i,
			freqMap,
		)
	}

	finalAllPossibleMinutes := make(map[int]int)
	backTrack(
		finalAllPossibleMinutes,
		usedByLowerNum,
		0,
		0,
		minutes,
		0,
		num,
		freqMap,
	)
	allPossibleTimes := convertMinutesToTime(finalAllPossibleMinutes)
	sort.Strings(allPossibleTimes)

	return allPossibleTimes
}

// Correct solution

// ReadBinaryWatchCorrect ...
func ReadBinaryWatchCorrect(num int) []string {
	if num == 0 {
		return []string{"0:00"}
	}
	hourMap := map[int][]int{
		0: []int{0},
		1: []int{1, 2, 4, 8},
		2: []int{3, 5, 6, 9, 10},
		3: []int{7, 11},
	}
	res := []string{}
	oneMap := map[int][]int{}
	for i := 0; i < 60; i++ {
		count := cntOfOne(i)
		if _, ok := oneMap[count]; !ok {
			oneMap[count] = []int{}
		}
		oneMap[count] = append(oneMap[count], i)
	}
	for length := 0; length <= num && length < 4; length++ {
		merge(hourMap[length], oneMap[num-length], &res)
	}
	return res
}

func cntOfOne(num int) int {
	cnt := 0
	for num > 0 {
		num = num & (num - 1)
		cnt++
	}
	return cnt
}

func merge(hour, minute []int, res *[]string) {
	formate := "%d:%d"
	ltTenFormate := "%d:0%d"
	var str string
	for _, h := range hour {
		for _, m := range minute {
			str = fmt.Sprintf(formate, h, m)
			if m < 10 {
				str = fmt.Sprintf(ltTenFormate, h, m)
			}
			*res = append(*res, str)
		}
	}
}
