package order

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mmalls/wszb-bs/pkg/g"
	"github.com/mmalls/wszb-bs/pkg/model"
	"github.com/mmalls/wszb-bs/pkg/service/util"
	log "github.com/xtfly/log4g"
)

var (
	mlog = log.GetLogger("order")
)

// HandleList api: Get /rest/v1/users/{userId}/orders
func HandleList(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	o := &model.Order{UserID: util.CvtDef(c.Param("userId"), 0)}
	var rows []model.OrderWithCustom

	offset := util.CvtDef(c.Query("offset"), 0)
	limit := util.CvtDef(c.Query("limit"), 100)

	if rows, err = o.ListByUserID(offset, limit); err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": &rows})
}

// HandleCreate api: Post /rest/v1/users/{userId}/orders
func HandleCreate(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	o := &model.OrderWithCustom{}
	if err = c.Bind(o); err != nil {
		return
	}
	o.UserID = util.CvtDef(c.Param("userId"), 0)
	o.CustomID = o.Custom.ID
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
	o := &model.OrderWithCustom{}
	o.ID = icid

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
	o := &model.OrderWithCustom{}
	o.ID = icid
	o.UserID = iuid

	if err = c.Bind(o); err != nil {
		return
	}

	if err := o.Update(); err != nil {
		return
	}
	c.Status(http.StatusOK)
}

//HandleStatsQuery api: Post /rest/v1/users/{userId}/stats/orders?by=month
func HandleStatsQuery(c *gin.Context) {
	var err error
	defer g.HandleErr(c, mlog, &err)

	o := &model.Order{UserID: util.CvtDef(c.Param("userId"), 0)}

	//
	end := time.Now()
	begin := end
	qbegin := c.Query("begin")
	qend := c.Query("end")
	if qend != "" {
		if iend, err := time.ParseInLocation("2006-01-02", qend, time.Local); err != nil {
			end = iend
		}
	}
	if qbegin != "" {
		if ibegin, err := time.ParseInLocation("2006-01-02", qbegin, time.Local); err != nil {
			begin = ibegin
		}
	}

	//
	by := c.Query("by")
	format := "2006-01-02"

	switch by {
	case "month":
		begin = end.AddDate(0, -1, 0)
	case "week":
		begin = end.AddDate(0, 0, -7)
	case "year":
		begin = end.AddDate(-1, 0, 0)
		format = "2006-01"
	default:
		begin = end.AddDate(0, 0, -7)
	}

	var rows []model.OrderWithCustom
	if rows, err = o.Query(begin, end); err != nil {
		return
	}

	out := &model.OrderStats{UserID: o.UserID}
	out.TotalOrder = len(rows)
	customMap := make(map[int]int)
	goodsMap := make(map[int]int)
	out.Items = make([]*model.OrderItem, 0, out.TotalOrder)
	last := &model.OrderItem{}
	for _, r := range rows {
		out.TotalIncoming += r.TotalSellPrice
		for _, g := range r.Goods {
			out.TotalQuantity += g.Quantity
			goodsMap[g.GoodsID] = g.GoodsID
		}
		customMap[r.CustomID] = r.CustomID

		// calc by day or month
		f := r.CreatedAt.Format(format)
		if f == last.Key {
			last.Incoming += r.TotalSellPrice
		} else {
			last = &model.OrderItem{Key: f, Incoming: r.TotalSellPrice}
			out.Items = append(out.Items, last)
		}
	}
	out.TotalCustom = len(customMap)
	out.TotalGoods = len(goodsMap)

	c.JSON(http.StatusOK, out)
}
