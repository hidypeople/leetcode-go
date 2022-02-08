package task0165

import (
	"strconv"
	"strings"
)

// Given two version numbers, version1 and version2, compare them.
//   Version numbers consist of one or more revisions joined by a dot '.'.
//   Each revision consists of digits and may contain leading zeros.
//   Every revision contains at least one character.
//   Revisions are 0-indexed from left to right, with the leftmost revision being revision 0,
//     the next revision being revision 1, and so on. For example 2.5.33 and 0.1 are valid version numbers.
// To compare version numbers, compare their revisions in left-to-right order.
// Revisions are compared using their integer value ignoring any leading zeros.
// This means that revisions 1 and 001 are considered equal.
// If a version number does not specify a revision at an index, then treat the revision as 0.
// For example, version 1.0 is less than version 1.1 because their revision 0s are the same,
// but their revision 1s are 0 and 1 respectively, and 0 < 1.
// Return the following:
//   If version1 < version2, return -1.
//   If version1 > version2, return 1.
//   Otherwise, return 0.
//
// Constraints:
//   1 <= version1.length, version2.length <= 500
//   version1 and version2 only contain digits and '.'.
//   version1 and version2 are valid version numbers.
//   All the given revisions in version1 and version2 can be stored in a 32-bit integer.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Compare Version Numbers.
//   Memory Usage: 2.1 MB, less than 13.43% of Go online submissions for Compare Version Numbers.
func compareVersion(version1 string, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")
	i := 0
	for i < len(v1) && i < len(v2) {
		v1part, v2part := strings.TrimLeft(v1[i], "0"), strings.TrimLeft(v2[i], "0")
		v1num, _ := strconv.Atoi(v1part)
		v2num, _ := strconv.Atoi(v2part)
		if v1num < v2num {
			return -1
		} else if v1num > v2num {
			return 1
		}
		i++
	}
	// Process tail
	var tail []string
	res := 0
	if len(v1) < len(v2) {
		res = -1
		tail = v2
	} else if len(v1) > len(v2) {
		res = 1
		tail = v1
	} else {
		return 0
	}
	for j := i; j < len(tail); j++ {
		if strings.Trim(tail[j], "0") != "" {
			return res
		}
	}
	return 0
}
