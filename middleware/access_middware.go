package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"seedgo-skeleton/common"
	"strings"
	"time"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func sliceContain(sli []string, k string) bool {
	for _, i := range sli {
		if i == k {
			return true
		}
	}
	return false
}

func AccessLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//headerList := []string{"Authorization", "X-Tid"}
		headerList := []string{"X-Tid"}

		t := time.Now()
		body, _ := ctx.GetRawData()
		var header string
		for k, v := range ctx.Request.Header {
			if sliceContain(headerList, k) {
				header = header + k + ":" + strings.Join(v, ",") + ";"
			}
		}

		bodystr := ""
		// ingore file upload body, beacuse is too big
		contentType := ctx.Request.Header.Get("Content-Type")
		if !strings.Contains(contentType, "multipart/form-data") {
			bodystr = string(body)
		}
		common.Infof(ctx, "access log request, uri: %s, method: %s, header: %s, params: %s",
			ctx.Request.RequestURI,
			ctx.Request.Method,
			header,
			bodystr,
		)
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		ctx.Next()
		// after request

		costtime := time.Since(t).Microseconds()
		common.Infof(ctx, "access log response, costtime: %dms, result: %s", costtime, blw.body.String())

	}
}
