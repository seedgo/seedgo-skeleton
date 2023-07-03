package middleware

import (
	"github.com/gin-gonic/gin"
	"seedgo-skeleton/util"
)

func TraceMiddware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tid := ctx.GetHeader("X-Tid")
		if len(tid) == 0 {
			tid = util.NextUid()
		}

		ctx.Set("tid", tid)

		ctx.Next()
	}

}
