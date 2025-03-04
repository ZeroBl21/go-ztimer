//go:build !inmemory
// +build !inmemory

package cmd

import (
	"github.com/ZeroBl21/go-ztimer/pomodoro"
	"github.com/ZeroBl21/go-ztimer/pomodoro/repository"
	"github.com/spf13/viper"
)

func getRepo() (pomodoro.Repository, error) {
	repo, err := repository.NewSQLiteRepo(viper.GetString("db"))
	if err != nil {
		return nil, err
	}

	return repo, nil
}
