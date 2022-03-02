package main

import (
	"goproject/AnaloGo/process"
)

func main() {
	m := make(map[string]string)
	process.Run("운영계1", "C10", "file", m)

}
