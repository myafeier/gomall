package model

import (
	"fmt"
	slogger "github.com/myafeier/logger"
	"log"
	"os"
)

var logger slogger.ILogger

func init() {
	logger = slogger.NewSimpleLogger2(os.Stdout, "[Model]", log.Lshortfile|log.Ldate|log.Lmicroseconds)

}

func formatMoney(t float32) string {
	return fmt.Sprintf("¥%.2f", t)
}
