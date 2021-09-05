package task0899

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_orderlyQueue(t *testing.T) {
	assert.Equal(t, "acb", orderlyQueue("cba", 1))
	assert.Equal(t, "acb", orderlyQueue("bac", 1))
}

func Test_orderlyQueue2(t *testing.T) {
	assert.Equal(t, "aaabc", orderlyQueue("baaca", 3))
}

func Test_orderlyQueue3(t *testing.T) {
	assert.Equal(t, "imvxz", orderlyQueue("xmvzi", 2))
}

func Test_orderlyQueue4(t *testing.T) {
	assert.Equal(t,
		"aagtkuqxitavoyjqiupzadbdyymyvuteolyeerecnuptghlzsynozeuuvteryojyokpufanyrqqmtgxhyycltlnusyeyyqygwupc",
		orderlyQueue("xitavoyjqiupzadbdyymyvuteolyeerecnuptghlzsynozeuuvteryojyokpufanyrqqmtgxhyycltlnusyeyyqygwupcaagtkuq", 1),
	)
}
