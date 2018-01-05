package ip_pool

import (
	"log"
	"strings"

	"self/proxy_pool/ua_pool"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

const (
	urlData5u = "http://www.data5u.com/free/gngn/index.shtml"
)

func getIpFromData5u() []string {
	ipList := make([]string, 0)

	ua := ua_pool.GetRandomUa()

	request := gorequest.New()
	resp, _, errs := request.Get(urlData5u).Set("User-Agent", ua).End()
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

	doc.Find(".l2").Each(func(i int, selection *goquery.Selection) {
		var ip, ipType string
		ip = selection.Children().Eq(0).Text() + ":" + selection.Children().Eq(1).Text()
		ipType = selection.Children().Eq(3).Text()
		//log.Println(ipType + ip_pool)

		if strings.ToLower(ipType) == "http" {
			ipList = append(ipList, ip)
		}
	})
	log.Println("Data5u done")
	return ipList
}
