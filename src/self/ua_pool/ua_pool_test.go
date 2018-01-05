package ua_pool

import (
	"fmt"
	"os"
	"testing"
)

func TestGetUaFromFile(t *testing.T) {
	os.Chdir("/Users/shiyi/Downloads/fightcoder-manager")
	InitUaPool()

	for i := 0; i < 10; i++ {
		ua := GetRandomUa()
		fmt.Println(ua)
	}
}
