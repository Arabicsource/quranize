package quranize

import (
	"bytes"
	"strings"
)

var memo = make(map[string][]string)

func inTree(harfs []rune, node *Node) bool {
	for _, harf := range harfs {
		node = getChild(node.Children, harf)
		if node == nil {
			return false
		}
	}
	return true
}

func combine(heads, tails []string) []string {
	combinations := []string{}
	for _, head := range heads {
		for _, tail := range tails {
			if head == "" {
				combinations = append(combinations, tail)
			} else {
				combinations = append(combinations, head+tail)
				combinations = append(combinations, head+" "+tail)
				combinations = append(combinations, head+"ال"+tail)
				combinations = append(combinations, head+" ال"+tail)
				combinations = append(combinations, head+"ا"+tail)
				combinations = append(combinations, head+"ا "+tail)
			}
		}
	}
	return combinations
}

func quranize(text string) []string {
	if text == "" {
		return []string{""}
	}

	if cache, ok := memo[text]; ok {
		return cache
	}

	kalimas := []string{}
	l := len(text)
	for width := 1; width <= maxWidth && width <= l; width++ {
		if tails, ok := hijaiyas[text[l-width:]]; ok {
			heads := quranize(text[:l-width])
			for _, combination := range combine(heads, tails) {
				if inTree([]rune(combination), root) {
					kalimas = append(kalimas, combination)
				}
			}
		}
	}

	memo[text] = kalimas
	return kalimas
}

func removeDoubleChar(text string) string {
	buffer := bytes.NewBuffer(nil)
	for i := range text {
		if i == 0 || text[i-1] != text[i] {
			buffer.WriteByte(text[i])
		}
	}
	return buffer.String()
}

func insertResult(results []string, newResult string) []string {
	for _, result := range results {
		if result == newResult {
			return results
		}
	}
	return append(results, newResult)
}

func Encode(text string) []string {
	text = strings.Replace(text, " ", "", -1)
	results := []string{}
	for _, result := range quranize(text) {
		results = insertResult(results, result)
	}
	for _, result := range quranize(removeDoubleChar(text)) {
		results = insertResult(results, result)
	}
	return results
}
