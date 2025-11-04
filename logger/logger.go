package loggerutils

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/mdobak/go-xerrors"
)

type StackFrame struct {
	Func   string `json:"func"`
	Source string `json:"source"`
	Line   int    `json:"line"`
}

func MarshalStack(err error) []StackFrame {
	trace := xerrors.StackTrace(err)

	if len(trace) == 0 {
		return nil
	}

	frames := trace.Frames()
	s := make([]StackFrame, len(frames))

	for i, v := range frames {
		f := StackFrame{
			Source: filepath.Join(
				filepath.Base(filepath.Dir(v.File)),
				filepath.Base(v.File),
			),
			Func: filepath.Base(v.Function),
			Line: v.Line,
		}
		s[i] = f
	}

	return s
}

func FormatError(err error) slog.Value {
	var groupValues []slog.Attr

	groupValues = append(groupValues, slog.String("msg", err.Error()))

	frames := MarshalStack(err)

	if frames != nil {
		groupValues = append(groupValues,
			slog.Any("trace", frames),
		)
	}
	return slog.GroupValue(groupValues...)
}

func ReplaceAttr(_ []string, a slog.Attr) slog.Attr {
	switch a.Value.Kind() {
	case slog.KindAny:
		switch v := a.Value.Any().(type) {
		case error:
			a.Value = FormatError(v)
		}
	}

	if a.Key == "msg" {
		return slog.Attr{Key: "message", Value: a.Value}
	}

	return a
}

func New(level slog.Level) *slog.Logger {
	opts := &slog.HandlerOptions{Level: level, ReplaceAttr: ReplaceAttr}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	// detect when in an AWS lambda runtime environment; add 'requestId' attribute
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if lambdaContext, ok := lambdacontext.FromContext(ctx); ok {
		logger.With(slog.String("requestId", lambdaContext.AwsRequestID))
	}
	return logger
}
