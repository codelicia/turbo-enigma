package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewEnv(t *testing.T) {
	t.Run("Env present", func(t *testing.T) {
		defer os.Unsetenv("MY_AWESOME_ENV")
		os.Setenv("MY_AWESOME_ENV", "some-content")

		_, err := NewEnv([]string{"MY_AWESOME_ENV"})
		assert.Empty(t, err)
	})

	t.Run("Env missing", func(t *testing.T) {
		_, err := NewEnv([]string{"MY_AWESOME_ENV"})
		assert.Errorf(t, err, "missing %s in environment variable", "MY_AWESOME_ENV")
	})
}

func TestEnv_Get(t *testing.T) {
	t.Run("Env present", func(t *testing.T) {
		defer os.Unsetenv("MY_AWESOME_ENV")
		os.Setenv("MY_AWESOME_ENV", "some-content")

		env, err := NewEnv([]string{"MY_AWESOME_ENV"})
		assert.Empty(t, err)

		assert.Equal(t, "some-content", env.Get("MY_AWESOME_ENV"))
	})

	t.Run("Env missing", func(t *testing.T) {
		defer os.Unsetenv("MY_AWESOME_ENV")
		os.Setenv("MY_AWESOME_ENV", "some-content")

		env, err := NewEnv([]string{"MY_AWESOME_ENV"})
		assert.Empty(t, err)

		assert.Panics(t, func() {
			_ = env.Get("ANOTHER_AWESOME_ENV")
		})
	})
}
