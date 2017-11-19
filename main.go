package main

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mmalls/comm"
	"github.com/mmalls/comm/ginmw"
	"github.com/mmalls/wszb-bs/pkg/g"
	"github.com/mmalls/wszb-bs/pkg/model"
	"github.com/mmalls/wszb-bs/pkg/service"
	"github.com/xtfly/log4g"
)

var (
	mlog = log.GetLogger("main")
	hlog = log.GetLogger("http")
)

func main() {
	// parse the command flags
	comm.ParseCmd()

	// parse config file
	if err := g.ParseCfg(); err != nil {
		mlog.Errorf("Parse config file, %v", err)
		os.Exit(2)
	}
	//mlog.Infof("%v", g.Cfg)

	// connect to databse
	if err := model.Init(); err != nil {
		mlog.Errorf("Init variable, %v", err)
		os.Exit(2)
	}

	// start http server
	r := gin.New()
	if !g.Cfg.Common.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Use(ginmw.Logger(hlog), gin.Recovery(), cors.Default())
	pwd, _ := os.Getwd()
	r.StaticFile("/", filepath.Join(pwd, "static", "index.html"))
	r.Static("/static", filepath.Join(pwd, "static", "static"))
	if err := service.Init(r); err != nil {
		os.Exit(2)
	}
	comm.StartHTTP(hlog, &g.Cfg.Common, r, func() {
		// stopped callback
	})
}
