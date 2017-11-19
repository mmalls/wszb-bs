package custom

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mmalls/wszb-bs/pkg/g"
	"github.com/mmalls/wszb-bs/pkg/model"
	"github.com/mmalls/wszb-bs/pkg/service/util"
	log "github.com/xtfly/log4g"
)

var (
	mlog = log.GetLogger("custom")
)

// HandleList api: Get /rest/v1/users/{userId}/customs
func HandleList(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	custom := &model.Custom{UserID: util.CvtDef(c.Param("userId"), 0)}
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

	custom := &model.Custom{UserID: util.CvtDef(c.Param("userId"), 0)}
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

	custom := &model.Custom{ID: util.CvtDef(c.Param("customId"), 0)}
	if err := custom.Delete(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}

//HandleGet api: Post /rest/v1/users/{userId}/customs/{customId}
func HandleGet(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	custom := &model.Custom{ID: util.CvtDef(c.Param("customId"), 0)}
	if err = custom.Get(); err != nil {
		return
	}
	c.JSON(http.StatusOK, custom)
}

//HandleUpdate api: Post /rest/v1/users/{userId}/customs/{customId}
func HandleUpdate(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	custom := &model.Custom{ID: util.CvtDef(c.Param("customId"), 0),
		UserID: util.CvtDef(c.Param("userId"), 0)}

	if err = c.Bind(custom); err != nil {
		return
	}

	if err = custom.Save(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}
