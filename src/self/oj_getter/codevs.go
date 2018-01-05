package oj_getter

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const (
	codevsUrl = "www.codevs.cn/problem"
)

type CodeVSGetter struct {
	BaseGetter
}

func (c *CodeVSGetter) getter() {

}

//处理对应的题目,例如:"http://www.codevs.cn/problem/1001/
func (c *CodeVSGetter) GetProblem(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get() error!")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll() error!")
		return
	}

	src := string(body)

	//将html标签全部转换成小写
	re, _ := regexp.Compile(`<[\S\s]+?>`)
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//匹配需要的数据,添加外层div防止非目标p标签的干扰
	re, _ = regexp.Compile(`<div class="panel-body">[\S\s]+?<p>[\S\s]+?</p>[\S\s]+?</div>`)
	temps := re.FindAllString(src, -1)

	for i := 0; i < len(temps); i++ {

		//读取p中的内容
		re, _ = regexp.Compile(`<p>[\S\s]+?</p>`)
		temps[i] = re.FindString(temps[i])
		re, _ = regexp.Compile(`<[\S\s]+?>`)
		temps[i] = re.ReplaceAllString(temps[i], "")

	}

	//c.Description = temps[0]
	//c.InputDescription = temps[1]
	//c.OutputDescription = temps[2]
	//c.SampleInput = temps[3]
	//c.SampleOutput = temps[4]
	//c.Hint = temps[5]
	fmt.Println(c)
}
