package controller

import (
	"filter"
	slogger "github.com/myafeier/logger"
	"gopkg.in/macaron.v1"
	"log"
	"os"
)

var logger slogger.ILogger

func init() {
	logger = slogger.NewSimpleLogger2(os.Stdout, "[Controller]", log.Lshortfile|log.Ldate|log.Lmicroseconds)

}

func GetUserSession(ctx *macaron.Context) (session *filter.Session) {
	logger.Debug("userInfo:", ctx.Data["userInfo"])
	session, _ = ctx.Data["userInfo"].(*filter.Session)
	return
}
