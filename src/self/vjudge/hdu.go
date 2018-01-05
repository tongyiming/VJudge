package vjudge

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"

	"fmt"

	"github.com/pkg/errors"
)

type HDUJudger struct {
	client   *http.Client
	token    string
	pat      *regexp.Regexp
	username string
	userpass string
}

const HDUToken = "HDU"

var HDURes = map[string]int{"Queuing": 0,
	"Compiling":                                 1,
	"Running":                                   1,
	"Accepted":                                  2,
	"Compilation Error":                         3,
	"Runtime Error<br>(STACK_OVERFLOW)":         4,
	"Runtime Error<br>(ACCESS_VIOLATION)":       4,
	"Runtime Error<br>(ARRAY_BOUNDS_EXCEEDED)":  4,
	"Runtime Error<br>(FLOAT_DENORMAL_OPERAND)": 4,
	"Runtime Error<br>(FLOAT_DIVIDE_BY_ZERO)":   4,
	"Runtime Error<br>(FLOAT_OVERFLOW)":         4,
	"Runtime Error<br>(FLOAT_UNDERFLOW )":       4,
	"Runtime Error<br>(INTEGER_OVERFLOW)":       4,
	"Runtime Error<br>(INTEGER_DIVIDE_BY_ZERO)": 4,
	"Wrong Answer":                              5,
	"Time Limit Exceeded":                       6,
	"Memory Limit Exceeded":                     7,
	"Output Limit Exceeded":                     8,
	"Presentation Error":                        9,
	"System Error":                              10}

//var HDULang = map[int]int{
//	LanguageC:    3,
//	LanguageCPP:  2,
//	LanguageJAVA: 5}

var hduLangPush map[int]int = map[int]int{
	C:    1,
	CPP:  2,
	JAVA: 3}

var hduLangPull map[int]int = map[int]int{
	C:    3,
	CPP:  2,
	JAVA: 5}

func (h *HDUJudger) Init() error {
	jar, _ := cookiejar.New(nil)
	h.client = &http.Client{Jar: jar, Timeout: time.Second * 10}
	h.token = HDUToken
	pattern := `(\d+)</td><td>(.*?)</td><td>(?s:.*?)<font color=.*?>(.*?)</font>.*?<td>(\d+)MS</td><td>(\d+)K</td><td><a href="/viewcode.php\?rid=\d+"  target=_blank>(\d+) B</td><td>.*?</td>`

	h.pat = regexp.MustCompile(pattern)
	h.username = "mysake"
	h.userpass = "123456"
	return nil
}

func (h *HDUJudger) Match(token string) bool {
	if token == HDUToken {
		return true
	}
	return false
}
func (h *HDUJudger) Login() error {

	h.client.Get("http://acm.hdu.edu.cn")

	uv := url.Values{}
	uv.Add("username", h.username)
	uv.Add("userpass", h.userpass)
	uv.Add("login", "Sign In")

	req, err := http.NewRequest("POST", "http://acm.hdu.edu.cn/userloginex.php?action=login", strings.NewReader(uv.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "acm.hdu.edu.cn")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.91 Safari/537.36")

	resp, err := h.client.Do(req)
	if err != nil {
		log.Println("err", err)
		return err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	html := string(b)
	if strings.Index(html, "No such user or wrong password.") >= 0 {
		return errors.New("Login Failed")
	}

	return nil
}

func (h *HDUJudger) Submit() error {

	uv := url.Values{}
	uv.Add("check", "0")
	uv.Add("problemid", "1000")
	uv.Add("language", "2")
	uv.Add("usercode", "#include<iostream>using namespace std;int main(){int a,b;while(cin>>a>>b){cout<<a+b<<endl;}return 0;}")

	req, err := http.NewRequest("POST", "http://acm.hdu.edu.cn/submit.php?action=submit", strings.NewReader(uv.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "acm.hdu.edu.cn")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.91 Safari/537.36")

	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(time.Now())

	b, _ := ioutil.ReadAll(resp.Body)
	html := string(b)

	if strings.Index(html, "One or more following ERROR(s) occurred.") >= 0 {
		return errors.New("Submit Failed")
	}
	return nil
}
