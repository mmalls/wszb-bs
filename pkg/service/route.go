package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mmalls/wszb-bs/pkg/g"
	"github.com/mmalls/wszb-bs/pkg/service/channel"
	"github.com/mmalls/wszb-bs/pkg/service/custom"
	"github.com/mmalls/wszb-bs/pkg/service/goods"
	"github.com/mmalls/wszb-bs/pkg/service/order"
	"github.com/mmalls/wszb-bs/pkg/service/user"
	log "github.com/xtfly/log4g"
)

var (
	hlog = log.GetLogger("http")
)

// Init ...
func Init(r gin.IRouter) error {
	apis(r)
	return nil
}

// apis ...
func apis(r gin.IRouter) {
	v1 := r.Group("/rest/v1")

	v1.POST("/users", user.HandleCreate)
	v1.POST("/auth", user.HandleLogin)

	v1auth := v1.Use(g.Auth(hlog))

	v1auth.GET("/users/:userId", user.HandleGet)

	v1auth.GET("/users/:userId/customs", custom.HandleList)
	v1auth.GET("/users/:userId/customs/:customId", custom.HandleGet)
	v1auth.POST("/users/:userId/customs", custom.HandleCreate)
	v1auth.PUT("/users/:userId/customs/:customId", custom.HandleUpdate)
	v1auth.DELETE("/users/:userId/customs/:customId", custom.HandleDelete)

	v1auth.GET("/users/:userId/channels", channel.HandleList)
	v1auth.GET("/users/:userId/channels/:channelId", channel.HandleGet)
	v1auth.POST("/users/:userId/channels", channel.HandleCreate)
	v1auth.PUT("/users/:userId/channels/:channelId", channel.HandleUpdate)
	v1auth.DELETE("/users/:userId/channels/:channelId", channel.HandleDelete)

	v1auth.GET("/users/:userId/goods", goods.HandleList)
	v1auth.GET("/users/:userId/goods/:goodsId", goods.HandleGet)
	v1auth.POST("/users/:userId/goods", goods.HandleCreate)
	v1auth.PUT("/users/:userId/goods/:goodsId", goods.HandleUpdate)
	v1auth.DELETE("/users/:userId/goods/:goodsId", goods.HandleDelete)

	v1auth.GET("/users/:userId/orders", order.HandleList)
	v1auth.GET("/users/:userId/orders/:orderId", order.HandleGet)
	v1auth.POST("/users/:userId/orders", order.HandleCreate)
	v1auth.PUT("/users/:userId/orders/:orderId", order.HandleUpdate)
	v1auth.DELETE("/users/:userId/orders/:orderId", order.HandleDelete)
	v1auth.GET("/users/:userId/stats/orders", order.HandleStatsQuery)

}
