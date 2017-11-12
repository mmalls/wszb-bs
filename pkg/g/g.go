package g

import (
	"net/http"

	"crypto/rsa"

	"io/ioutil"

	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mmalls/comm"
	"github.com/xtfly/log4g"
)

// TopConfig top config node
type TopConfig struct {
	Common   comm.Common   `yaml:"common"`
	Database comm.Database `yaml:"database"`
	JWT      JWT           `yaml:"jwt"`
}

type JWT struct {
	PublicKey  string          `yaml:"public_key"`
	PrivateKey string          `yaml:"private_key"`
	Enable     bool            `yaml:"enable"`
	VerifyKey  *rsa.PublicKey  `yaml:"-"`
	SignKey    *rsa.PrivateKey `yaml:"-"`
}

// glabol vars
var (
	Cfg = &TopConfig{}
)

func ParseCfg() error {
	err := comm.ParseCfg(Cfg)
	if err != nil {
		return err
	}

	var bs []byte
	bs, err = ioutil.ReadFile(Cfg.JWT.PrivateKey)
	if err != nil {
		return err
	}
	Cfg.JWT.SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(bs)
	if err != nil {
		return err
	}

	bs, err = ioutil.ReadFile(Cfg.JWT.PublicKey)
	if err != nil {
		return err
	}

	Cfg.JWT.VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(bs)
	return nil
}

// HandleErr ...
func HandleErr(c *gin.Context, mlog log.Logger, err *error) {
	if *err == nil {
		return
	}

	e := *err
	if strings.Contains(e.Error(), "json") {
		c.Status(http.StatusBadRequest)
	} else if strings.Contains(e.Error(), "crypto/bcrypt: hashedPassword") {
		c.Status(http.StatusUnauthorized)
	} else if e == gorm.ErrRecordNotFound {
		switch c.Request.Method {
		case http.MethodGet, http.MethodPut:
			c.Status(http.StatusNotFound)
		case http.MethodDelete:
			c.Status(http.StatusNoContent)
		default:
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.Status(http.StatusInternalServerError)
	}

	mlog.Error(*err)
}
