package order

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mmalls/wszb-bs/pkg/g"
	"github.com/mmalls/wszb-bs/pkg/model"
	log "github.com/xtfly/log4g"
)

var (
	mlog = log.GetLogger("order")
)

// HandleList api: Get /rest/v1/users/{userId}/orders
func HandleList(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	userID := c.Param("userId")
	iuid, _ := strconv.Atoi(userID)
	o := &model.Order{UserID: iuid}
	var row []model.OrderWithInfo

	if row, err = o.ListByUserID(); err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": &row})
}

// HandleCreate api: Post /rest/v1/users/{userId}/orders
func HandleCreate(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	userID := c.Param("userId")
	iuid, _ := strconv.Atoi(userID)
	o := &model.Order{UserID: iuid}
	if err = c.Bind(o); err != nil {
		return
	}

	if err = o.Save(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}

//HandleDelete api: Post /rest/v1/users/{userId}/orders/{orderId}
func HandleDelete(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	cid := c.Param("orderId")
	icid, _ := strconv.Atoi(cid)
	o := &model.Order{ID: icid}

	if err = o.Delete(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}

//HandleGet api: Post /rest/v1/users/{userId}/orders/{orderId}
func HandleGet(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	cid := c.Param("orderId")
	icid, _ := strconv.Atoi(cid)
	o := &model.Order{ID: icid}

	if err = o.Get(); err != nil {
		return
	}
	c.JSON(http.StatusOK, o)
}

//HandleUpdate api: Post /rest/v1/users/{userId}/orders/{orderId}
func HandleUpdate(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	userID := c.Param("userId")
	iuid, _ := strconv.Atoi(userID)
	cid := c.Param("orderId")
	icid, _ := strconv.Atoi(cid)
	o := &model.Order{ID: icid, UserID: iuid}

	if err = c.Bind(o); err != nil {
		return
	}

	if err := o.Save(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}
