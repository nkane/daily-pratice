package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestSuggestedProducts(t *testing.T) {
	tests := []struct {
		Name       string
		Product    []string
		SearchWord string
		Output     [][]string
	}{
		{
			Name: "Example 1",
			Product: []string{
				"mobile", "mouse", "moneypot", "monitor", "mousepad",
			},
			SearchWord: "mouse",
			Output: [][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := suggestedProducts(test.Product, test.SearchWord)
			assert.DeepEqual(t, test.Output, got)
		})
	}
}
