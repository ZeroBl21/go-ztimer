//go:build inmemory
// +build inmemory

package cmd

import (
	"github.com/ZeroBl21/go-ztimer/pomodoro"
	"github.com/ZeroBl21/go-ztimer/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
