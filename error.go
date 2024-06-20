package errors

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Error struct {
	msg string

	file string
	line int
	fn   string
}

func (err *Error) Error() string {
	return err.String()
}

func (err *Error) String() string {
	return fmt.Sprintf("%s:%d @ fn:%s # %s", err.file, err.line, err.fn, err.msg)
}

func (err *Error) As(target any) bool {
	return errors.As(err, target)
}

func (err *Error) Is(target error) bool {
	return errors.Is(err, target)
}

func (err *Error) Unwrap() error {
	return errors.Unwrap(err)
}

func (err *Error) Join(errs ...error) error {
	return errors.Join(append([]error{err}, errs...)...)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func Join(errs ...error) error {
	return errors.Join(errs...)
}

func New(args ...any) error {
	msg := fmt.Sprintln(args...)
	name, file, line := getCallerInfo(0)

	return &Error{msg: msg[:len(msg)-1], file: file, line: line, fn: name}
}

func NewCustom(skip int, args ...any) error {
	msg := fmt.Sprintln(args...)
	name, file, line := getCallerInfo(skip)

	return &Error{msg: msg[:len(msg)-1], file: file, line: line, fn: name}
}

func getCallerInfo(skip int) (string, string, int) {
	const caller = 2

	pc, file, line, ok := runtime.Caller(caller + skip)
	details := runtime.FuncForPC(pc)
	if !ok || details == nil {
		return "", "", 0
	}

	funcName := details.Name()
	lastSlash := strings.LastIndexByte(funcName, os.PathSeparator)
	if lastSlash < 0 {
		lastSlash = 0
	}

	lastDot := strings.LastIndexByte(funcName[lastSlash:], '.') + lastSlash

	cwd, err := os.Getwd()
	if err != nil {
		return "", "", 0
	}

	file = strings.TrimPrefix(file, cwd+"/")

	return funcName[lastDot+1:], file, line
}

func GetMessage(err error) string {
	if err == nil {
		return ""
	}

	if e, ok := err.(*Error); ok {
		return e.msg
	}

	return ""
}
