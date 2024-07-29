package leetcode

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestPredictPartyVictory(t *testing.T) {
	senate := "RD"
	expected := "Radiant"
	got := predictPartyVictory(senate)
	assert.Assert(t, expected == got)

	senate = "RDD"
	expected = "Dire"
	got = predictPartyVictory(senate)
	assert.Assert(t, expected == got)

	senate = "RDDRRD"
	expected = "Radiant"
	got = predictPartyVictory(senate)
	assert.Assert(t, expected == got)

	senate = "DDR"
	expected = "Dire"
	got = predictPartyVictory(senate)
	assert.Assert(t, expected == got)

	senate = "RDR"
	expected = "Radiant"
	got = predictPartyVictory(senate)
	assert.Assert(t, expected == got)
}
