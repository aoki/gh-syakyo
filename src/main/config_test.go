package main

import (
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	home := os.Getenv("HOME")
	config := loadConfig(home + "/.config/gh_test")
	assert.Equal(t, "ringohub", config.User)
}
