package static

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserProxy_Login(t *testing.T) {
	err := NewUserProxy(&User{}).Login("123", "eq")
	require.Nil(t, err)
}
