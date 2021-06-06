package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LogInfo(message string, a... interface{}) {
	_, _ = fmt.Fprintf(gin.DefaultWriter, message+"\n", a)

}

func LogError(message string, a... interface{}) {
	_, _ = fmt.Fprintf(gin.DefaultErrorWriter, message+"\n", a)
}
