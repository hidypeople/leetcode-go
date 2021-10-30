package task1044

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestDupSubstring(t *testing.T) {
	assert.Equal(t, "ana", longestDupSubstring("banana"))
}

func Test_longestDupSubstring2(t *testing.T) {
	assert.Equal(t, "", longestDupSubstring("abcd"))
}

func Test_longestDupSubstring3(t *testing.T) {
	assert.Equal(t, "pppppppppppppppppppppppppppppppppppp", longestDupSubstring("ppppppppppppppppppppppppppppppppppppp"))
}
