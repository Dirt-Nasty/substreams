package block

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRange_Split(t *testing.T) {
	og := &Range{
		StartBlock:        706,
		ExclusiveEndBlock: 1250,
	}

	expected := []*Range{
		{706, 800},
		{800, 1000},
		{1000, 1200},
		{1200, 1250},
	}

	actual := og.Split(200)

	require.Equal(t, expected, actual)
}

func TestRangeMerged(t *testing.T) {
	assert.Equal(t, ParseRanges("10-40,50-70").String(), ParseRanges("10-20,20-30,30-40,50-60,60-70").Merged().String())
	assert.Equal(t, ParseRanges("10-40,60-70").String(), ParseRanges("10-20,20-30,30-40,60-70").Merged().String())
	assert.Equal(t, ParseRanges("10-40").String(), ParseRanges("10-20,20-30,30-40").Merged().String())
	assert.Equal(t, ParseRanges("1-5,10-12,13-14").String(), ParseRanges("1-2,2-3,3-4,4-5,10-12,13-14").Merged().String())
}

func TestRangeMergedChunked(t *testing.T) {
	assert.Equal(t,
		ParseRanges("10-30,30-40,50-70").String(),
		ParseRanges("10-20,20-30,30-40,50-60,60-70").MergedChunked(20).String(),
	)
	assert.Equal(t,
		ParseRanges("10-30,30-50,50-60,80-100").String(),
		ParseRanges("10-20,20-30,30-40,40-50,50-60,80-90,90-100").MergedChunked(20).String(),
	)
	assert.Equal(t,
		ParseRanges("10-20,20-30,30-40").String(),
		ParseRanges("10-20,20-30,30-40").MergedChunked(5).String(),
	)
	assert.Equal(t,
		ParseRanges("1-4,4-5,10-12,13-14").String(),
		ParseRanges("1-2,2-3,3-4,4-5,10-12,13-14").MergedChunked(3).String(),
	)
}
