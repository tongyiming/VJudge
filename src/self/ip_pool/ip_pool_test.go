package ip_pool

import (
	"fmt"
	"os"
	"testing"

	"self/proxy_pool/ua_pool"
)

func TestXicidaili(t *testing.T) {
	os.Chdir("/Users/shiyi/Downloads/fightcoder-manager")

	ua_pool.InitUa()
	getIpFromXicidaili()
}

func TestData5u(t *testing.T) {
	os.Chdir("/Users/shiyi/Downloads/fightcoder-manager")

	ua_pool.InitUa()
	getIpFromData5u()
}

func TestCheckIp(t *testing.T) {
	os.Chdir("/Users/shiyi/Downloads/fightcoder-manager")

	ua_pool.InitUa()
	pullIpInNetWork()
}

func TestGetRandomIp(t *testing.T) {
	InitIpPoll()

	for i := 0; i < 10; i++ {
		fmt.Println(i, getRandomIp())
	}
}
