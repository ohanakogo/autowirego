package autowirego

import (
	"github.com/gookit/slog"
	"runtime"
)

func HandleRecover(do func(err any)) {
	if r := recover(); r != nil {
		do(r)
	}
}

func HandleNonError(err error, postprocess func()) {
	if err == nil {
		if postprocess != nil {
			postprocess()
		}
	}
}

func HandleError(err error, postprocess func()) {
	_, file, line, _ := runtime.Caller(1)
	if err != nil {
		slog.Errorf("error encountered[%s:%d]", file, line)
		slog.Errorf("output: %s", err.Error())
		if postprocess != nil {
			postprocess()
		}
	}
}
