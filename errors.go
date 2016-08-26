package errors

import (
	"fmt"
	"runtime"
	"strconv"
)

func New(params ...string) *errors {
	newerr := &errors{}

	if len(params) == 0 {
		return &errors{Msg: params[0], calldepth: 0}
	}

	if len(params) > 2 {
		return &errors{Msg: "Ops too many params", calldepth: 0}
	}

	for idx, value := range params {
		switch {
		case idx == 0:
			newerr.Msg = value
		case idx == 1:
			depth, e := strconv.Atoi(value)
			if e != nil || depth < 0 {
				newerr.calldepth = 1
			}
		}
	}

	var (
		ok   bool
		line int
		file string
	)

	_, file, line, ok = runtime.Caller(newerr.calldepth)
	if !ok {
		file = "Ops no file name"
		line = -1
	}

	if len(newerr.Msg) > 79 {
		msg := []byte(newerr.Msg)
		newerr.Msg = fmt.Sprintf("%s\n%s: %s: %v\n",
			msg[:79], msg[79:], file, line)

	}
	newerr.Msg = fmt.Sprintf("%s: %s: %v\n", newerr.Msg, file, line)

	return newerr
}

type errors struct {
	Msg       string
	calldepth int
}

func (e *errors) Error() string {
	return e.Msg
}
