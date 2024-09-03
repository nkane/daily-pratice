package leetcode

import "sort"

/*
	1268: Search Suggestions Systems
*/

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	root := &Trie{}
	result := [][]string{}

	for _, product := range products {
		current := root
		for _, ch := range product {
			if current.children[ch-'a'] != nil {
				current = current.children[ch-'a']
			} else {
				current.children[ch-'a'] = &Trie{}
				current = current.children[ch-'a']
			}
			current.words = append(current.words, product)
		}
	}

	node := root
	for _, ch := range searchWord {
		if node != nil && node.children[ch-'a'] != nil {
			node = node.children[ch-'a']
			if len(node.words) <= 3 {
				result = append(result, node.words)
			} else {
				result = append(result, node.words[:3])
			}
		} else {
			node = nil
			result = append(result, []string{})
		}
	}

	return result
}

type Trie struct {
	children [26]*Trie
	words    []string
}
