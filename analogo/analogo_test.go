package main

import (
	"goproject/AnaloGo/process"
	"testing"
)

func TestMain(t *testing.T) {
	m := make(map[string]string)
	process.Run("운영계1", "C10", "file", m)

}
