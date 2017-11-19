package goods

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

// HandleList api: Get /rest/v1/users/{userId}/goods
func HandleList(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	o := &model.Goods{UserID: util.CvtDef(c.Param("userId"), 0)}
	var row []model.GoodsWitchChl
	if row, err = o.ListByUserID(); err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"goods": &row})
}

// HandleCreate api: Post /rest/v1/users/{userId}/goods
func HandleCreate(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	o := &model.Goods{UserID: util.CvtDef(c.Param("userId"), 0)}
	if err = c.Bind(o); err != nil {
		return
	}

	if err := o.Save(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}

//HandleDelete api: Post /rest/v1/users/{userId}/goods/{goodsId}
func HandleDelete(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	o := &model.Goods{ID: util.CvtDef(c.Param("goodsId"), 0)}
	if err := o.Delete(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}

//HandleGet api: Post /rest/v1/users/{userId}/goods/{goodsId}
func HandleGet(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	o := &model.Goods{ID: util.CvtDef(c.Param("goodsId"), 0)}
	if err = o.Get(); err != nil {
		return
	}
	c.JSON(http.StatusOK, o)
}

//HandleUpdate api: Post /rest/v1/users/{userId}/goods/{goodsId}
func HandleUpdate(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	o := &model.Goods{ID: util.CvtDef(c.Param("goodsId"), 0),
		UserID: util.CvtDef(c.Param("userId"), 0)}

	if err = c.Bind(o); err != nil {
		return
	}

	if err = o.Save(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}
