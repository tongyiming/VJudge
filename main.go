/**
 * Created by shiyi on 2017/10/1.
 * Email: shiyi@fightcoder.com
 */

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"self/commons"
	"self/oj_getter"

	"self/commons/g"

	"github.com/TV4/graceful"
	"github.com/gin-gonic/gin"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//接收命令行参数
	version := flag.Bool("v", false, "show version")
	cfgfile := flag.String("c", "cfg/cfg.toml.debug", "set config file")
	flag.Parse()

	if *version {
		fmt.Println("version:", g.GitVer)
		fmt.Println("build time:", g.BuildTime)
		os.Exit(0)
	}

	// 初始化框架组件
	commons.InitAll(*cfgfile)

	gin.SetMode(g.Conf().Run.Mode)

	oj_getter.StartGetter()

	//初始化路由
	//router := gin.Default()

	//controllers.Register(router)

	//go dispatch.StartConsume()

	//优雅退出
	graceful.LogListenAndServe(&http.Server{
		Addr: ":8888",
		//Handler: router,
	})

	//关闭框架组件
	commons.CloseAll()
}
