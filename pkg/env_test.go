package pkg_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"turboenigma/pkg"
)

func TestNewEnv(t *testing.T) {
	t.Run("Env present", func(t *testing.T) {
		defer os.Unsetenv("MY_AWESOME_ENV")
		os.Setenv("MY_AWESOME_ENV", "some-content")

		_, err := pkg.NewEnv([]string{"MY_AWESOME_ENV"})
		assert.Empty(t, err)
	})

	t.Run("Env missing", func(t *testing.T) {
		_, err := pkg.NewEnv([]string{"MY_AWESOME_ENV"})
		assert.Errorf(t, err, "missing %s in environment variable", "MY_AWESOME_ENV")
	})
}

func TestEnv_Get(t *testing.T) {
	t.Run("Env present", func(t *testing.T) {
		defer os.Unsetenv("MY_AWESOME_ENV")
		os.Setenv("MY_AWESOME_ENV", "some-content")

		env, err := pkg.NewEnv([]string{"MY_AWESOME_ENV"})
		assert.Empty(t, err)

		assert.Equal(t, "some-content", env.Get("MY_AWESOME_ENV"))
	})

	t.Run("Env missing", func(t *testing.T) {
		defer os.Unsetenv("MY_AWESOME_ENV")
		os.Setenv("MY_AWESOME_ENV", "some-content")

		env, err := pkg.NewEnv([]string{"MY_AWESOME_ENV"})
		assert.Empty(t, err)

		assert.Panics(t, func() {
			_ = env.Get("ANOTHER_AWESOME_ENV")
		})
	})
}