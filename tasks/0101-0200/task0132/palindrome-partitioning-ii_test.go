package task0132

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCut(t *testing.T) {
	assert.Equal(t, 1, minCut("aab"))
}

func Test_minCut2(t *testing.T) {
	assert.Equal(t, 0, minCut("a"))
}

func Test_minCut3(t *testing.T) {
	assert.Equal(t, 1, minCut("ab"))
}

func Test_minCut4(t *testing.T) {
	assert.Equal(t, 452, minCut("apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp"))
}

func Test_minCut5(t *testing.T) {
	assert.Equal(t, 1, minCut("abbab"))
}
