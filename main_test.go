package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomInt(t *testing.T) {
	x := randomInt(2, 5)
	assert.GreaterOrEqual(t, x, 2)
	assert.LessOrEqual(t, x, 5)
}

func TestParseInput(t *testing.T) {
	x := "0.1,	0.3,	0.5,	0,	1,	2,	3,	4,	5,	5.5,	6,	7,	8,	9,	9.1,	9.2,	10,	10.5,	11,	11.015,	11.02,	11.03,	11.04,	11.05,	11.06,	11.065,	11.07,	11.08,	11.09,	11.095,	11.15,	11.2,	11.25,	11.5,	11.75,	12,	12.1,	12.2,	13,	14,	14.1,	14.5,	14.6,	14.65,	14.7,	14.75,	17.75,	14.8,	15,	15.5,	16,	17,	17.5,	18,	19,	20,	21,	22,	23,	24,	25,	26,"
	r := parseInput(x)
	assert.Len(t, r, 62)
	x = "0.1,	0.3,	0.5,	0,	1,	2,	3,	4,	5,	5.5,	6,	7,	8,	9,	9.1,	9.2,	10,	10.5,	11,	11.015,	11.02,	11.03,	11.04,	11.05,	11.06,	11.065,	11.07,	11.08,	11.09,	11.095,	11.15,	11.2,	11.25,	11.5,	11.75,	12,	12.1,	12.2,	13,	14,	14.1,	14.5,	14.6,	14.65,	14.7,	14.75,	17.75,	14.8,	15,	15.5,	16,	17,	17.5,	18,	19,	20,	21,	22,	23,	24,	25,	26"
	r = parseInput(x)
	assert.Len(t, r, 62)
	x = "12"
	r = parseInput(x)
	assert.Len(t, r, 1)
}
