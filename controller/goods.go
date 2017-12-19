package controller

import (
	"gopkg.in/macaron.v1"
	"model"
)

// {get} /api/goods 列出所有资源
// limit,pageNo
// classify  int
// name string ,模糊查询，临时不做，之后考虑用sphinx
// stat 状态
func GoodsListByPagination(ctx *macaron.Context) {
	limit := ctx.QueryInt("limit")
	pageNo := ctx.QueryInt("pageNo")
	classify := ctx.QueryInt64("classify")
	// name := ctx.ParamsEscape("name")
	stat := ctx.QueryInt("stat")

	if limit == 0 {
		limit = 10
	}
	if pageNo == 0 {
		pageNo = 1
	}
	logger.Debug("classify", classify, "|", limit, "|", pageNo, "|", stat)
	var classifySlice []int64
	if classify != 0 {
		classifySlice = model.FindAllClassifyChildIdSlice(classify)
		logger.Debug(classifySlice)
	}

	g := new(model.Goods)
	g.Stat = stat
	result, err := g.GetByPM(pageNo, limit, "", classifySlice)
	if err != nil {
		logger.Error(err)
		ctx.Status(500)
		return
	}
	ctx.JSON(200, result)

}
