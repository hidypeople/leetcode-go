package task0049

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_groupAnagrams(t *testing.T) {
	assert.True(t, compareResults(
		[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}),
	))
}

func Test_groupAnagrams2(t *testing.T) {
	assert.True(t, compareResults(
		[][]string{},
		groupAnagrams([]string{}),
	))
}

func Test_groupAnagrams3(t *testing.T) {
	assert.True(t, compareResults(
		[][]string{{"a"}},
		groupAnagrams([]string{"a"}),
	))
}

func Test_groupAnagrams4(t *testing.T) {
	assert.True(t, compareResults(
		[][]string{{"abbbbbbbbbbb"}, {"aaaaaaaaaaab"}},
		groupAnagrams([]string{"abbbbbbbbbbb", "aaaaaaaaaaab"}),
	))
}

func Test_groupAnagrams5(t *testing.T) {
	assert.True(t, compareResults(
		[][]string{{"zoo"}, {"hag"}, {"pup"}, {"ids"}, {"tic"}, {"flo"}, {"red"}, {"fee"}, {"hus"}, {"fed"}, {"sky"}, {"yup", "yup"}, {"ben"}, {"pop"}, {"you"}, {"abe"}, {"den", "ned"}, {"cod"}, {"jul"}, {"jar"}, {"ray"}, {"paw"}, {"owl"}, {"oct"}, {"woo"}, {"ado"}, {"dem"}, {"eve"}, {"why"}, {"sui"}, {"ito"}, {"ago"}, {"arc"}, {"mys"}, {"fat"}, {"eon"}, {"asp"}, {"lea"}, {"sow"}},
		groupAnagrams([]string{"ray", "cod", "abe", "ned", "arc", "jar", "owl", "pop", "paw", "sky", "yup", "fed", "jul", "woo", "ado", "why", "ben", "mys", "den", "dem", "fat", "you", "eon", "sui", "oct", "asp", "ago", "lea", "sow", "hus", "fee", "yup", "eve", "red", "flo", "ids", "tic", "pup", "hag", "ito", "zoo"}),
	))
}

func compareResults(r1, r2 [][]string) bool {
	for _, r1strings := range r1 {
		isFound := false
		for _, r2strings := range r2 {
			if isStringsEqual(r1strings, r2strings) {
				isFound = true
				break
			}
		}
		if !isFound {
			return false
		}
	}
	return true
}

func isStringsEqual(strs1, strs2 []string) bool {
	for _, s1 := range strs1 {
		if !findString(strs2, s1) {
			return false
		}
	}
	return true
}

func findString(strs []string, s string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}
