package task0682

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calPoints(t *testing.T) {
	// "5" - Add 5 to the record, record is now [5].
	// "2" - Add 2 to the record, record is now [5, 2].
	// "C" - Invalidate and remove the previous score, record is now [5].
	// "D" - Add 2 * 5 = 10 to the record, record is now [5, 10].
	// "+" - Add 5 + 10 = 15 to the record, record is now [5, 10, 15].
	// The total sum is 5 + 10 + 15 = 30.
	assert.Equal(t, 30, calPoints([]string{"5", "2", "C", "D", "+"}))
}

func Test_calPoints2(t *testing.T) {
	// "5" - Add 5 to the record, record is now [5].
	// "-2" - Add -2 to the record, record is now [5, -2].
	// "4" - Add 4 to the record, record is now [5, -2, 4].
	// "C" - Invalidate and remove the previous score, record is now [5, -2].
	// "D" - Add 2 * -2 = -4 to the record, record is now [5, -2, -4].
	// "9" - Add 9 to the record, record is now [5, -2, -4, 9].
	// "+" - Add -4 + 9 = 5 to the record, record is now [5, -2, -4, 9, 5].
	// "+" - Add 9 + 5 = 14 to the record, record is now [5, -2, -4, 9, 5, 14].
	// The total sum is 5 + -2 + -4 + 9 + 5 + 14 = 27.
	assert.Equal(t, 27, calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}
