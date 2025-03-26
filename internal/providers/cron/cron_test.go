package cron_test

import (
	"sync"
	"testing"
	"time"

	"github.com/mauriciofsnts/bot/internal/providers/cron"
	"github.com/stretchr/testify/require"
)

func TestCron(t *testing.T) {
	c := cron.New()

	var mu sync.Mutex
	var executed bool

	job := func() {
		mu.Lock()
		executed = true
		mu.Unlock()
	}

	err := c.AddJob("@every 1s", job)
	require.NoError(t, err)

	c.Start()
	defer c.Stop()

	jobs := c.List()
	require.Len(t, jobs, 1, "Should have 1 job")

	time.Sleep(3 * time.Second)

	mu.Lock()
	defer mu.Unlock()
	require.True(t, executed, "Should have been executed")
}
