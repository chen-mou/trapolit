package utils

import (
	"errors"
	"fmt"
	"runtime"
	"trapolit/lib/i18n"
)

type CustomError error

func NewError(lang i18n.Lang, msg string) CustomError {
	//TODO:这里要做到可配置 正式环境不要出现堆栈信息
	return errors.New(printCallers() + "\n" + i18n.Translate(lang, msg))
}

func printCallers() string {
	const depth = 10
	pc := make([]uintptr, depth)
	n := runtime.Callers(3, pc) // skip=2：跳过 runtime.Callers 和当前函数 还有NewError
	frames := runtime.CallersFrames(pc[:n])
	var res string
	for {
		frame, more := frames.Next()
		res = res + fmt.Sprintf("%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line)
		if !more {
			break
		}
	}
	return res
}
