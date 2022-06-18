package page_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToken_ForIndex(t *testing.T) {
	n := NewTokenGenerate()

	i, err := n.GetIndex(n.ForIndex(1))
	assert.NoError(t, err)
	assert.Equal(t, i, 1)

	n2 := NewTokenGenerate(SetTokenSalt("test1"))
	s2 := n2.ForIndex(1)
	_, err = n.GetIndex(s2)
	assert.Error(t, err)
	assert.Equal(t, err, ErrInvalidToken)

	n3 := NewTokenGenerate(SetTokenMaxIndex(10))

	_, err = n3.GetIndex(n3.ForIndex(11))
	assert.Error(t, err)
	assert.Equal(t, err, ErrOverMaxPageSizeToken)

	n4 := NewTokenGenerate(SetTokenTimeLimitation(time.Second * 5))
	s4 := n4.ForIndex(1)
	time.Sleep(6 * time.Second)
	_, err = n4.GetIndex(s4)
	assert.Error(t, err)
	assert.Equal(t, err, ErrOverdueToken)
}
