package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

// 커스텀 에러 타입 (스택 정보 보관)
type StackError struct {
	Msg   string
	Err   error
	Stack []string
}

func (e *StackError) Error() string {
	return fmt.Sprintf("%s: %v", e.Msg, e.Err)
}

// Logrus 커스텀 포매터
type PrettyFormatter struct {
	logrus.TextFormatter
}

func (f *PrettyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	bytes, err := f.TextFormatter.Format(entry)
	if err != nil {
		return nil, err
	}

	// 스택 에러가 있으면 멀티라인 처리
	if err, ok := entry.Data[logrus.ErrorKey].(*StackError); ok {
		stackStr := "Stack Trace:\n  " + strings.Join(err.Stack, "\n  ")
		bytes = append(bytes, []byte(stackStr)...)
	}
	return bytes, nil
}
