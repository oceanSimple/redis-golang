package test

import (
	"github.com/spf13/viper"
	"testing"
)

func TestViper(t *testing.T) {
	tactic := viper.GetString("aof.tactic")
	t.Logf("tactic: %s", tactic)
}
