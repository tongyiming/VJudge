package vjudge

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	codevsUrl = "http://codevs.cn/submission/?id="
)

var (
	codevsUserList = []string{"tyming@fightcoder.com"}
	codevsPassList = []string{"tyming"}
	codevsMutexMap map[string]*sync.Mutex
)

var CodeVSRes = map[string]int{

	"等待测试 Pending":                0,
	"测试通过 Accepted":               1,
	"编译错误 Compile Error":          2,
	"测试失败 Rejected":               3,
	"正在评测 Running":                4,
	"答案错误 Wrong Answer":           5,
	"题目无效 Invalid Problem":        6,
	"非法调用 Restricted Call":        7,
	"运行错误 Runtime Error":          8,
	"暂无数据 Data Missed":            9,
	"超出时间 Time Limit Exceeded":    10,
	"超出空间 Memory Limit Exceeded":  11,
	"过多输出 Output Limit Exceeded":  12,
	"等待重测 Rejudge Pending":        13,
	"运行错误(内存访问非法) Runtime Error":  14,
	"运行错误(浮点错误)    Runtime Error": 15,
	"正在编译 COMPILING":              16}

var CodeVSLang = map[string]string{
	"C":      "c",
	"C++":    "cpp",
	"Pascal": "pas"}

type CodeVSJudger struct {
}

func (this *CodeVSJudger) Submit(problemId, language, code string) string {

	//init jar
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar, Timeout: time.Second * 10}
	a, _ := client.Get("http://www.codevs.cn/")
	fmt.Println(a.StatusCode)
	//fmt.Println(a.Cookies())
	//b, _ := ioutil.ReadAll(a.Body)
	//html := string(b)
	//fmt.Println(html)

	index := rand.Intn(5)
	index = index % len(userList)
	//login data
	uvLogin := url.Values{}
	uvLogin.Add("username", codevsUserList[index])
	uvLogin.Add("password", codevsPassList[index])
	fmt.Println(codevsUserList[index], codevsPassList[index], uvLogin)

	req, err := http.NewRequest("POST", "https://login.codevs.com/api/auth/login", strings.NewReader(uvLogin.Encode()))
	if err != nil {
		return "POST https://login.codevs.com/api/auth/login error"
	}
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Origin", "https://login.codevs.com")
	req.Header.Add("Host", "login.codevs.com")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Add("Authorization", "No login")
	req.Header.Add("Content-Length", "56")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Referer", "https://login.codevs.com/auth/login")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.86 Safari/537.36")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Transport = tr
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return "1"
	}
	fmt.Println(req, "login:", resp.StatusCode)
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	html := string(b)
	if strings.Index(html, "登录失败 ，请仔细核对你的帐号与密码") >= 0 {
		return "2"
	}
	return ""
	////获取csrfmiddlewaretoken
	//req, err = http.NewRequest("GET", "http://codevs.cn/problem/1000/", nil)
	//if err != nil {
	//	return "3"
	//}
	//
	//resp, err = client.Do(req)
	//if err != nil {
	//	return "4"
	//}
	//defer resp.Body.Close()
	//
	//b, _ = ioutil.ReadAll(resp.Body)
	//html = string(b)
	//
	//re, _ := regexp.Compile(`<input type="hidden" name="csrfmiddlewaretoken" value="(.*?)">`)
	//temp := re.FindStringSubmatch(html)
	//fmt.Println("csrfmiddlewaretoken", temp)

	//uv := url.Values{}
	//
	//uv.Add("code", code)
	//uv.Add("id", problemId)
	//uv.Add("format", CodeVSLang[language])
	//fmt.Println(CodeVSLang[language])
	//uv.Add("csrfmiddlewaretoken", "")
	//
	//req, err = http.NewRequest("POST", "http://codevs.cn/judge/", strings.NewReader(uv.Encode()))
	//if err != nil {
	//	return ""
	//}
	//
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	//req.Header.Add("Host", "codevs.cn")
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.91 Safari/537.36")
	//
	//resp, err = client.Do(req)
	//if err != nil {
	//	return ""
	//}
	//fmt.Println(resp.StatusCode)
	//fmt.Println("111111111111111")
	//defer resp.Body.Close()
	//
	//return ""
	/*
		uv := url.Values{}
		uv.Add("problemid", problemId)
		uv.Add("format", CodeVSLang[language])
		uv.Add("code", code)
		uv.Add("csrfmiddlewaretoken",temp[1])
		//uv.Add("problemid", "1000")
		//uv.Add("language", "2")
		//uv.Add("usercode", "#include<iostream>using namespace std;int main(){int a,b;while(cin>>a>>b){cout<<a+b<<endl;}return 0;}")

		req, err = http.NewRequest("POST", "http://acm.hdu.edu.cn/submit.php?action=submit", strings.NewReader(uv.Encode()))
		if err != nil {
			return ""
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Host", "acm.hdu.edu.cn")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.91 Safari/537.36")

		resp, err = client.Do(req)
		if err != nil {
			return ""
		}
		defer resp.Body.Close()
		fmt.Println(time.Now())

		b, _ = ioutil.ReadAll(resp.Body)
		html = string(b)

		if strings.Index(html, "One or more following ERROR(s) occurred.") >= 0 {
			return ""
		}

		resp, _ = client.Get("http://acm.hdu.edu.cn/status.php?first=&pid=1000&user=mysake&lang=3&status=0")
		b, _ = ioutil.ReadAll(resp.Body)
		html = string(b)

		re = regexp.MustCompile(`(\d+)</td><td>(.*?)</td><td>(?s:.*?)<font color=.*?>(.*?)</font>.*?<td>(\d+)MS</td>`)
		temps := re.FindAllStringSubmatch(html, -1)
		//fmt.Println("=============", temps[0][1])
		return temps[0][1]

	*/

}
