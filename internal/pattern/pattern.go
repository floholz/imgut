package pattern

import (
	"fmt"
	"strconv"
	"strings"
)

type UrlSplit struct {
	split     string
	isPattern bool
}

func ResolveUrl(url string) ([]string, error) {
	resolvedUrl := make([]UrlSplit, 1)
	// resolvedUrlIsPattern := []bool{false}
	patternIndexes := []int{-1}
	for i, c := range url {
		switch c {
		case '{':
			if len(patternIndexes)%2 != 1 {
				return nil, fmt.Errorf("invalid pattern start")
			}
			patternIndexes = append(patternIndexes, i)
			resolvedUrl = append(resolvedUrl, UrlSplit{isPattern: true})
		case '}':
			if len(patternIndexes)%2 != 0 {
				return nil, fmt.Errorf("invalid pattern end")
			}
			patternIndexes = append(patternIndexes, i)
			resolvedUrl = append(resolvedUrl, UrlSplit{isPattern: false})
		default:
			resolvedUrl[len(resolvedUrl)-1].split += string(c)
		}
	}
	if resolvedUrl[0].split == "" && len(resolvedUrl) > 1 {
		resolvedUrl = resolvedUrl[1:]
	}

	patternIndexes = append(patternIndexes, len(url)-1)
	if len(patternIndexes)%2 != 0 {
		patternIndexes = patternIndexes[:len(patternIndexes)-1]
	}

	urls := make([]string, 1)
	for _, urlSplit := range resolvedUrl {
		if urlSplit.isPattern {
			permutations := ResolvePattern(urlSplit.split)
			urlsLen := len(urls)
			for urlIdx := 0; urlIdx < urlsLen; urlIdx++ {
				tmpUrl := urls[urlIdx]
				urls[urlIdx] += permutations[0]
				for permIdx := 1; permIdx < len(permutations); permIdx++ {
					urls = append(urls, tmpUrl+permutations[permIdx])
				}
			}
		} else {
			for urlIdx := 0; urlIdx < len(urls); urlIdx++ {
				urls[urlIdx] += urlSplit.split
			}
		}
	}

	return urls, nil
}

func ResolvePattern(pattern string) []string {
	var resolveList []string
	orPatternList := strings.Split(pattern, "|")
	for _, orPattern := range orPatternList {
		resolveList = append(resolveList, resolveSinglePattern(orPattern)...)
	}
	return resolveList
}

func resolveSinglePattern(pattern string) []string {
	resolveList := make([]string, 1)
	bufferIdx := 0
	for pChar := 0; pChar < len(pattern); pChar++ {
		switch pattern[pChar] {
		case '-':
			handleHyphen(pattern, &pChar, &bufferIdx, &resolveList)
		case '?':
			handleQuestionMark(pattern, &pChar, &bufferIdx, &resolveList)
		case '.':
			handleDot(pattern, &pChar, &bufferIdx, &resolveList)

		}
	}
	if len(resolveList) == 1 && resolveList[0] == "" {
		resolveList = make([]string, 0)
	}
	if bufferIdx < len(pattern)-1 {
		resolveList = append(resolveList, pattern[bufferIdx:])
	}
	return resolveList
}

func handleHyphen(pattern string, patternIdx, bufferIdx *int, resolveList *[]string) {
	bufferEnd := max(*bufferIdx, *patternIdx-1)
	for c := 0; c < len(*resolveList); c++ {
		(*resolveList)[c] += pattern[*bufferIdx:bufferEnd]
	}
	*bufferIdx = bufferEnd + 2

	start := pattern[*patternIdx-1]
	end := pattern[*patternIdx+1]
	var tmpBuildStr []string
	for c := start; c <= end; c++ {
		for _, str := range *resolveList {
			tmpBuildStr = append(tmpBuildStr, str+string(c))
		}
	}
	*resolveList = tmpBuildStr
	*patternIdx++
}

func handleQuestionMark(pattern string, patternIdx, bufferIdx *int, resolveList *[]string) {
	bufferEnd := max(*bufferIdx, *patternIdx-1)
	bsLen := len(*resolveList)
	for i := 0; i < bsLen; i++ {
		str := pattern[*bufferIdx:bufferEnd]
		(*resolveList)[i] += str
		*resolveList = append(*resolveList, str+string(pattern[bufferEnd]))
	}
	*bufferIdx = bufferEnd + 2
}

func handleDot(pattern string, patternIdx, bufferIdx *int, resolveList *[]string) {
	if pattern[*patternIdx-1] != '.' {
		return
	}
	if *patternIdx+1 >= len(pattern)-1 {
		return
	}
	bufferEnd := max(*bufferIdx, *patternIdx-1)
	startValStr := pattern[*bufferIdx:bufferEnd]
	startVal, errStart := strconv.ParseInt(startValStr, 10, 64)
	if errStart != nil {
		fmt.Printf("Error parsing range start value %s\n", startValStr)
		*patternIdx = len(pattern)
		*bufferIdx = *patternIdx
		return
	}
	endValStr := pattern[*patternIdx+1:]
	endVal, errEnd := strconv.ParseInt(endValStr, 10, 64)
	if errEnd != nil {
		fmt.Printf("Error parsing range end value %s\n", endValStr)
		*patternIdx = len(pattern)
		*bufferIdx = *patternIdx
		return
	}
	var tmpBuildStr []string
	for i := startVal; i < endVal; i++ {
		for _, str := range *resolveList {
			tmpBuildStr = append(tmpBuildStr, str+strconv.FormatInt(i, 10))
		}
	}
	*resolveList = tmpBuildStr
	*patternIdx = len(pattern)
	*bufferIdx = *patternIdx
}
