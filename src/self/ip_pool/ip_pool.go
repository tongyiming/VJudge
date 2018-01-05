package ip_pool

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
)

const (
	checkIpUrl    = "http://httpbin.org/get"
	timeout       = 30 * time.Second
	checkTimeout  = 2 * time.Second
	checkInterval = 1 * time.Second
	rePullNum     = 10
)

var (
	rwMutex *sync.RWMutex
	ipPool  []string
)

func InitIpPoll() {
	rwMutex = new(sync.RWMutex)

	ipPool = pullIpInNetWork()
	go startCheck()
}

func GetProxyAgent() *gorequest.SuperAgent {
	ip := getRandomIp()
	if ip == "" {
		return gorequest.New().Timeout(timeout)
	}

	return gorequest.New().Proxy("http://" + ip).Timeout(timeout)
}

func getRandomIp() string {
	length := len(ipPool)
	if length == 0 {
		return ""
	}

	index := rand.Intn(len(ipPool))
	return ipPool[index]
}

func startCheck() {
	for {
		if len(ipPool) != 0 {
			index := rand.Intn(len(ipPool))
			ip := ipPool[index]
			if !checkIp(ip) {
				rwMutex.Lock()
				ipPool = append(ipPool[:index], ipPool[index+1:]...)
				rwMutex.Unlock()
			}
		}

		if len(ipPool) < rePullNum {
			addIpForPool()
		}

		time.Sleep(checkInterval)
	}
}

func addIpForPool() {
	ipList := pullIpInNetWork()
	ipList = append(ipList, ipPool...)

	rwMutex.Lock()
	ipPool = ipList
	rwMutex.Unlock()
}

//ip筛选
func pullIpInNetWork() []string {
	ipL := make([]string, 0)

	ipL = append(ipL, getIpFromXicidaili()...)
	ipL = append(ipL, getIpFromData5u()...)

	ipList := make([]string, 0)
	for _, ip := range ipL {
		isExist := false
		for _, rip := range ipPool {
			if rip == ip {
				isExist = true
				break
			}
		}
		if !isExist {
			ipList = append(ipList, ip)
		}
	}

	fmt.Println(len(ipList), " ", len(ipL))

	return ipList
}

func checkIp(ip string) bool {
	resp, _, err := gorequest.New().Proxy("http://" + ip).Get(checkIpUrl).Timeout(checkTimeout).End()
	if err != nil {
		return false
	}
	if resp.StatusCode == 200 {
		return true
	}
	return false
}
