package g

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/xtfly/log4g"
)

var keyFunc = func(token *jwt.Token) (interface{}, error) {
	return Cfg.JWT.VerifyKey, nil
}

func Auth(mlog log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !Cfg.JWT.Enable {
			return
		}

		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			mlog.Warnf("invalid Authorization header: %v", auth)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token := auth[len("Bearer "):]
		t, err := jwt.Parse(token, keyFunc)
		if err != nil {
			mlog.Warnf("parser token failed, token=%v，err=%v", token, err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err := t.Claims.Valid(); err != nil {
			mlog.Warnf("validate token failed, token=%v，err=%v", token, err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		sc, ok := t.Claims.(*jwt.StandardClaims)
		if ok {
			c.Set("user_id", sc.Audience)
		}
	}
}
