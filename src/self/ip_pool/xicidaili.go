package ip_pool

import (
	"log"
	"math/rand"
	"strconv"
	"strings"

	"self/proxy_pool/ua_pool"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

const (
	urlXicidaili = "http://www.xicidaili.com/nn/"
)

func getIpFromXicidaili() []string {
	ipList := make([]string, 0)

	ua := ua_pool.GetRandomUa()
	i := rand.Intn(1000) + 1

	request := gorequest.New()
	resp, _, errs := request.Get(urlXicidaili+strconv.FormatInt(int64(i), 10)).Set("User-Agent", ua).End()

	if errs != nil {
		log.Println(errs)
		return []string{}
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	resp.Body.Close()
	if err != nil {
		log.Println(err)
		return []string{}
	}

	doc.Find(".odd").Each(func(i int, selection *goquery.Selection) {
		var ip, ipType string
		ip = selection.Children().Eq(1).Text() + ":" + selection.Children().Eq(2).Text()
		ipType = selection.Children().Eq(5).Text()
		//log.Println(ipType + ip_pool)

		if strings.ToLower(ipType) == "http" {
			ipList = append(ipList, ip)
		}
	})
	log.Println("XiCi done")
	return ipList
}
