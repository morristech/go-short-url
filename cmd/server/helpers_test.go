package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateHash(t *testing.T) {
	data := map[string]string{
		"":    "1",
		"8":   "9",
		"9":   "A",
		"A":   "B",
		"X":   "Y",
		"Z":   "a",
		"y":   "z",
		"z":   "z1",
		"z1":  "z2",
		"z8":  "z9",
		"z9":  "zA",
		"zX":  "zY",
		"zY":  "zZ",
		"za":  "zb",
		"zy":  "zz",
		"zz":  "zz1",
		"zz1": "zz2",
		"zzA": "zzB",
		"zzY": "zzZ",
		"zzy": "zzz",
	}

	for seed, expected := range data {
		actual := GenerateHash(seed)
		assert.Equal(t, expected, actual)
	}
}
