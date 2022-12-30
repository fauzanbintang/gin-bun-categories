package utils

import (
	"fmt"
	"runtime"
)

func WrapError(err error) error {
	if err == nil {
		return nil
	}

	pc, file, line, _ := runtime.Caller(1)
	fmt.Println(pc, "=", file, "=", line, "dari wrap error")
	return wrapError(err, file, line, pc)
}

func wrapError(err error, file string, line int, pc uintptr) error {
	details := runtime.FuncForPC(pc)
	fmt.Println(details.Name())
	return fmt.Errorf("[%s][%d] %s: %w", file, line, details.Name(), err)
}
