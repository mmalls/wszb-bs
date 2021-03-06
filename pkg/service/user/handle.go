package user

import (
	"net/http"
	"strconv"

	"encoding/base64"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mmalls/wszb-bs/pkg/g"
	"github.com/mmalls/wszb-bs/pkg/model"
	"github.com/mmalls/wszb-bs/pkg/service/util"
	"github.com/xtfly/log4g"
	"golang.org/x/crypto/bcrypt"
)

var (
	mlog = log.GetLogger("user")
)

// HandleCreate api: Post /rest/v1/users
func HandleCreate(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	o := &model.User{}
	if err = c.Bind(o); err != nil {
		return
	}

	if err = o.GetByPhone(); err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if o.ID != 0 {
		c.Status(http.StatusConflict)
		return
	}

	hp, err1 := bcrypt.GenerateFromPassword([]byte(o.Password), 10)
	if err = err1; err != nil {
		return
	}
	o.Password = base64.StdEncoding.EncodeToString(hp)

	if err = o.Save(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}

//HandleGet api: Post /rest/v1/users/{userId}
func HandleGet(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	o := &model.User{ID: util.CvtDef(c.Param("userId"), 0)}

	if err = o.Get(); err != nil {
		return
	}
	o.Password = ""
	c.JSON(http.StatusOK, o)
}

// HandleLogin api: Post /rest/v1/users/auth
func HandleLogin(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	o := &model.Auth{}
	if err = c.Bind(o); err != nil {
		return
	}

	u := &model.User{Phone: o.Phone}
	if err = u.GetByPhone(); err != nil {
		err = g.UnauthErr(mlog, err)
		return
	}

	hp, err1 := base64.StdEncoding.DecodeString(u.Password)
	if err = err1; err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword(hp, []byte(o.Password)); err != nil {
		return
	}

	ll := &model.LoginLog{UserID: u.ID, IP: c.ClientIP()}
	if err = ll.Save(); err != nil {
		return
	}

	jwt, err1 := genJWT(u)
	if err = err1; err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": u, "token": jwt})
}

//openssl genrsa -out jwt_pri.pem 1024
//openssl pkcs8 -topk8 -inform PEM -in jwt_pri.pem -outform PEM –nocrypt
//openssl rsa -in jwt_pri.pem -pubout -out jwt_pub.pem
func genJWT(u *model.User) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims = &jwt.StandardClaims{
		// set the expire time
		// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		Audience:  strconv.Itoa(u.ID),
	}
	return t.SignedString(g.Cfg.JWT.SignKey)
}
