package channel

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mmalls/wszb-bs/pkg/g"
	"github.com/mmalls/wszb-bs/pkg/model"
	log "github.com/xtfly/log4g"
)

var (
	mlog = log.GetLogger("channel")
)

// HandleList api: Get /rest/v1/users/{userId}/channels
func HandleList(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	userID := c.Param("userId")
	iuid, _ := strconv.Atoi(userID)
	o := &model.Channel{UserID: iuid}
	var row []model.Channel
	if row, err = o.ListByUserID(); err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"channles": &row})
}

// HandleCreate api: Post /rest/v1/users/{userId}/channels
func HandleCreate(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	userID := c.Param("userId")
	iuid, _ := strconv.Atoi(userID)
	o := &model.Channel{UserID: iuid}
	if err = c.Bind(o); err != nil {
		return
	}

	if err = o.Save(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}

//HandleDelete api: Post /rest/v1/users/{userId}/channels/{channelId}
func HandleDelete(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	cid := c.Param("channelId")
	icid, _ := strconv.Atoi(cid)
	o := &model.Channel{ID: icid}

	if err = o.Delete(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}

//HandleGet api: Post /rest/v1/users/{userId}/channels/{channelId}
func HandleGet(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	cid := c.Param("channelId")
	icid, _ := strconv.Atoi(cid)
	o := &model.Channel{ID: icid}

	if err = o.Get(); err != nil {
		return
	}
	c.JSON(http.StatusOK, o)
}

//HandleUpdate api: Post /rest/v1/users/{userId}/channels/{channelId}
func HandleUpdate(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	userID := c.Param("userId")
	iuid, _ := strconv.Atoi(userID)
	cid := c.Param("channelId")
	icid, _ := strconv.Atoi(cid)
	o := &model.Channel{ID: icid, UserID: iuid}

	if err = c.Bind(o); err != nil {
		return
	}

	if err = o.Save(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}