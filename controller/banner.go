package controller

import (
	"gopkg.in/macaron.v1"
	"model"
)

// {get} /banner 列出所有资源
// * @apiParam {String} [page=1] 指定第几页
// * @apiParam {String} [limit=10] 指定每页的记录数
// * @apiParam {Boolean} [is_show] 指定is_show过滤
func BannerListAllGetHandler(ctx *macaron.Context) {
	isShow := ctx.ParamsEscape("is_show")
	// page := ctx.ParamsInt("page")
	// limit := ctx.ParamsInt("limit")
	var isShowBool bool
	if isShow == "true" {
		isShowBool = true
	} else {
		isShowBool = false
	}
	b := new(model.Banner)
	b.IsShow = isShowBool
	result, err := b.GetAll()
	if err != nil {
		logger.Error(err)
		ctx.Status(500)
		return
	}
	ctx.JSON(200, result)
}
