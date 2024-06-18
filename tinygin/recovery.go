package tinygin

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func trace(msg string) string {
	// pcs means program counters
	pcs := make([]uintptr, 10)
	n := runtime.Callers(3, pcs[:])

	var strBui strings.Builder
	strBui.WriteString(msg + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		strBui.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return strBui.String()
}

func Recovery() HandlerFunc {
	return func(ctx *Context) {
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("%s", err)
				log.Printf("%s\n", trace(msg))
				ctx.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		ctx.Next()
	}
}
