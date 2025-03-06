//go:build integration
// +build integration

package notify_test

import (
	"testing"

	"github.com/ZeroBl21/z-timer/notify"
)

func TestSend(t *testing.T) {
	n := notify.New("test title", "test msg", notify.SeverityNormal)

	if err := n.Send(); err != nil {
		t.Error(err)
	}
}
