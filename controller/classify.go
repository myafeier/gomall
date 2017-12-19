package controller

import (
	"gopkg.in/macaron.v1"
	"model"
)

//  {get} /classify/?:level 列出所有资源
// ?:rank 分类级别
func ClassifyGetListHandler(ctx *macaron.Context) {
	level := ctx.ParamsInt("level")
	logger.Debug(level)

	if level != 0 {
		v, ok := model.CacheClassify.LevelMap[level]
		if ok {
			ctx.JSON(200, v)
		}
	} else {
		ctx.JSON(200, model.CacheClassify.LevelStruct)
	}
	logger.Debugf("%#v", model.CacheClassify)
}
