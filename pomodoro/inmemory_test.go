//go:build inmemory
// +build inmemory

package pomodoro_test

import (
	"testing"

	"github.com/ZeroBl21/go-ztimer/pomodoro"
	"github.com/ZeroBl21/go-ztimer/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()

	return repository.NewInMemoryRepo(), func() {}
}
