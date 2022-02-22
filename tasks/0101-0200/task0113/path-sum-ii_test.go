package task0113

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pathSum(t *testing.T) {
	tree := BSTFromArrayInts([]int{5, 4, 8, 11, NULL, 13, 4, 7, 2, NULL, NULL, 5, 1})
	assert.Equal(t, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}, pathSum(tree, 22))
}

func Test_pathSum2(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, 3})
	assert.Equal(t, [][]int{}, pathSum(tree, 5))
}

func Test_pathSum3(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2})
	assert.Equal(t, [][]int{}, pathSum(tree, 0))
}

func Test_pathSum4(t *testing.T) {
	tree := BSTFromArrayInts([]int{-260, -202, -903, -980, -570, -858, 218, 764, -300, 205, NULL, -35, NULL, NULL, -204, 950, -769, 258, -652, 614, -584, 76, 817, -192, NULL, NULL, -114, 880, NULL, -200, 71, 671, 344, 801, 193, -18, 876, -920, -730, 222, 679, NULL, -680, NULL, NULL, NULL, -859, 744, -261, 692, NULL, -341, -163, NULL, NULL, 482, -979, 205, NULL, 146, 165, 801, 100, -656, 714, -629, 995, 474, 307, -581, -150, -941, NULL, NULL, NULL, -937, -69, -23, 82, NULL, -139, -591, NULL, -453, -861, -370, NULL, NULL, NULL, 216, 233, NULL, 430, NULL, 5, -110, NULL, NULL, -660, 624, -510, -588, NULL, NULL, 381, NULL, 368, 559, NULL, 521, -301, NULL, 522, 379, NULL, NULL, NULL, NULL, 456, 519, NULL, NULL, 482, 349, NULL, NULL, 19, NULL, NULL, 288, -811, NULL, -372, NULL, NULL, -536, NULL, -404, -457, -740, 860, NULL, NULL, -636, NULL, NULL, 342, -874, -462, -504, 781, 855, -392, NULL, NULL, NULL, 406, NULL, -758, 541, NULL, -947, NULL, NULL, NULL, NULL, NULL, -964, NULL, 600, -45, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, -194, NULL, NULL, NULL, -802, NULL, NULL, NULL, -3, NULL, -792, 672, 643, NULL, 14, NULL, NULL, 489, 457, NULL, NULL, NULL, NULL, 412, NULL, 558, NULL, NULL, NULL, NULL, -846, 158, -146, NULL, NULL, -76, -650, NULL, -782, NULL, -127, NULL, -678, NULL, NULL, NULL, NULL, NULL, NULL, -464, -426, NULL, -366, NULL, NULL, NULL, NULL, NULL, 81, -607, 716, NULL, NULL, -213, NULL, 379, NULL, NULL, NULL, NULL, 644, 445, NULL, NULL, -419, -845, -720, NULL, NULL, -915, NULL, NULL, NULL, NULL, NULL, NULL, -686, 594, -243, NULL, 496, NULL, 907, NULL, NULL, NULL, NULL, NULL, NULL, 579, 873, 702, NULL, NULL, NULL, -834, NULL, NULL, NULL, NULL, NULL, -300, -214, -466, NULL, NULL, 972, NULL, NULL, NULL, 814, NULL, -940, NULL, 763, NULL, NULL, NULL, NULL, -449, -844, NULL, NULL, NULL, NULL, -47})
	assert.Equal(t, [][]int{{-260, -903, -858, -35, 817, 222, 307, -301, -947, -76, 445, 579, 814, -47}}, pathSum(tree, -243))
}
