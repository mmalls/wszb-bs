package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mmalls/wszb-bs/pkg/g"
	"github.com/xtfly/gokits/gcrypto"
	log "github.com/xtfly/log4g"
)

var (
	glog = log.GetLogger("model")
	db   *gorm.DB
)

// Init all glabol vars
func Init() error {
	if err := openDB(); err != nil {
		return err
	}

	return nil
}

// openDB db connections
func openDB() (err error) {
	crypto, err := gcrypto.NewCrypto(g.Cfg.Common.KeyFactor)
	if err != nil {
		return err
	}
	pwd, err := crypto.DecryptStr(g.Cfg.Database.Password)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s:%s@%s", g.Cfg.Database.UserName, pwd, g.Cfg.Database.URL)
	db, err = gorm.Open(g.Cfg.Database.Driver, url)
	if err != nil {
		return err
	}
	db.SetLogger(logger{})
	//if g.Cfg.Common.Debug {
	db = db.Debug()
	//}
	db = db.AutoMigrate(&User{},
		&Custom{},
		&Channel{},
		&Goods{},
		&Order{},
		&LoginLog{},
	)
	return nil
}

type logger struct {
}

func (l logger) Print(v ...interface{}) {
	glog.Debug(v...)
}
