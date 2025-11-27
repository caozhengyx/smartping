package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/caozhengyx/smartping/src/funcs"
	"github.com/caozhengyx/smartping/src/g"
	"github.com/caozhengyx/smartping/src/http"
	"github.com/jakecoffman/cron"
	//"sync"
)

// Init config
var Version = "0.9.0"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	version := flag.Bool("v", false, "show version")
	flag.Parse()
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}
	g.ParseConfig(Version)
	go funcs.ClearArchive()
	c := cron.New()
	c.AddFunc("*/60 * * * * *", func() {
		go funcs.Ping()
		go funcs.Mapping()
		if g.Cfg.Mode["Type"] == "cloud" {
			go funcs.StartCloudMonitor()
		}
	}, "ping")
	c.AddFunc("0 0 * * * *", func() {
		go funcs.ClearArchive()
	}, "mtc")
	c.Start()
	http.StartHttp()
}
