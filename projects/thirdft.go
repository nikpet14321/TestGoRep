package main

import (
	"fmt"
	"regexp"
	"strings"
	"sort"
)

// подсмотрел
type wordValue struct {
	Word string
	Value int
}

func AnalyzeText(text string) {
	var maxikey string
	t := map[string]int{}
	maxinum := -111111
	sumwords := 0
	sumword := 0
	
	re := regexp.MustCompile(`[^.,!? ]+`)
	words := re.FindAllString(text, -1)
	for _, element := range words {
		t[strings.ToLower(element)]++
	}
	for key, elemen := range t {
		sumwords += elemen
		sumword = len(t)
		if elemen > maxinum {
			maxinum = elemen
			maxikey = key
		}
	}
	p := make([]int, 0, len(t))
	for _, elem := range t {
		p = append(p, elem)
	}
	sort.Slice(p, func(i, j int) bool{
		return p[i] > p[j]
	})
	r := getTopWords(t, 5)

	fmt.Printf(`Количество слов: %d
Количество уникальных слов: %d
Самое часто встречающееся слово: "%s" (встречается %d раз)
Топ-5 самых часто встречающихся слов:
"%s": %d раз
"%s": %d раз
"%s": %d раз
"%s": %d раз
"%s": %d раз` + "\n", sumwords, sumword, maxikey, maxinum, r[0], p[0], r[1], p[1], r[2], p[2], r[3], p[3], r[4], p[4])
}

func getTopWords(wordMap map[string]int, n int) []string {
	t := make([]wordValue, 0, len(wordMap))
	res := make([]string, n)

	for key, element := range wordMap {
		t = append(t, wordValue{key, element})
	}
	sort.Slice(t, func(i, j int) bool{
		return t[i].Value > t[j].Value
	})
	for i := 0;i < n;i++ {
		res[i] = t[i].Word
	}
	return res
}