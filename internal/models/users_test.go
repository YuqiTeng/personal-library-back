package userdata

import (
	"testing"
)

func TestGetUserById(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		testGetUserById(t, "123", "0")
	})
}