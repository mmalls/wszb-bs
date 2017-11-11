package custom

import (
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mmalls/wszb-bs/pkg/g"
	"github.com/mmalls/wszb-bs/pkg/model"
	log "github.com/xtfly/log4g"
)

var (
	mlog = log.GetLogger("custom")
)

// HandleList api: Get /rest/v1/users/{userId}/customs
func HandleList(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	userID := c.Param("userId")
	iuid, _ := strconv.Atoi(userID)
	custom := &model.Custom{UserID: iuid}
	var row []model.Custom
	if row, err = custom.ListByUserID(); err != nil {

		return
	}
	c.JSON(http.StatusOK, gin.H{"customs": &row})
}

// HandleCreate api: Post /rest/v1/users/{userId}/customs
func HandleCreate(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	userID := c.Param("userId")
	iuid, _ := strconv.Atoi(userID)
	custom := &model.Custom{UserID: iuid}
	if err = c.Bind(custom); err != nil {
		return
	}

	if err = custom.Save(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}

//HandleDelete api: Post /rest/v1/users/{userId}/customs/{customId}
func HandleDelete(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	cid := c.Param("customId")
	icid, _ := strconv.Atoi(cid)
	custom := &model.Custom{ID: icid}

	if err := custom.Delete(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}

//HandleGet api: Post /rest/v1/users/{userId}/customs/{customId}
func HandleGet(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	cid := c.Param("customId")
	icid, _ := strconv.Atoi(cid)
	custom := &model.Custom{ID: icid}

	if err = custom.Get(); err != nil {
		return
	}
	c.JSON(http.StatusOK, custom)
}

//HandleUpdate api: Post /rest/v1/users/{userId}/customs/{customId}
func HandleUpdate(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	userID := c.Param("userId")
	iuid, _ := strconv.Atoi(userID)
	cid := c.Param("customId")
	icid, _ := strconv.Atoi(cid)
	custom := &model.Custom{ID: icid, UserID: iuid}

	if err = c.Bind(custom); err != nil {
		return
	}

	if err = custom.Save(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}
