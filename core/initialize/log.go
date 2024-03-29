package initialize

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"myapp/core/common"

	tool "github.com/anu1097/golang-masking-tool"
	"github.com/anu1097/golang-masking-tool/filter"
	lg "github.com/sirupsen/logrus"
)

type Logging interface {
	//init new entry log
	NewEntry() *lg.Entry
	//log with level
	Error(ctx context.Context, format string, v ...interface{})
	Warning(ctx context.Context, format string, v ...interface{})
	Debug(ctx context.Context, format string, v ...interface{})
	Info(ctx context.Context, format string, v ...interface{})
	Fatal(ctx context.Context, format string, v ...interface{})
	//mask sensitive data on entry log
	FieldMask(data interface{}) string
	// set fields to entry log
	SetFields(entry *lg.Entry, fields map[string]interface{}) *lg.Entry
	//set entry to parent context
	SetEntry(parent context.Context, entry *lg.Entry) context.Context
	// Retrieve log entry from ctx, if there is no entry then return a completely new entry
	GetEntry(ctx context.Context) *lg.Entry
}

type logRus struct {
	env    string
	id     string
	level  string
	logger *lg.Logger
}

type ctxKey string

const (
	ctxKeyEntry ctxKey = "logger_entry"
)

func NewLogRus(level, id, env string) Logging {
	logger := lg.New()
	switch strings.ToUpper(level) {
	case "DEBUG":
		logger.SetLevel(lg.DebugLevel)
	case "ERROR":
		logger.SetLevel(lg.ErrorLevel)
	}

	fm := &lg.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.999Z07:00",
	}
	if env == common.EnvLocal {
		fm.PrettyPrint = true
	}
	logger.SetFormatter(fm)

	return &logRus{
		env:    env,
		id:     id,
		level:  level,
		logger: logger,
	}
}

func (k *logRus) NewEntry() *lg.Entry {
	return lg.NewEntry(k.logger)
}

func (k *logRus) SetFields(entry *lg.Entry, fields map[string]interface{}) *lg.Entry {
	return entry.WithFields(fields)
}

func (k *logRus) SetEntry(parent context.Context, entry *lg.Entry) context.Context {
	return context.WithValue(parent, ctxKeyEntry, entry)
}

func (k *logRus) GetEntry(ctx context.Context) *lg.Entry {
	entry, ok := ctx.Value(ctxKeyEntry).(*lg.Entry)
	if !ok {
		return lg.NewEntry(k.logger)
	}

	return entry
}

func (k *logRus) Error(ctx context.Context, format string, v ...interface{}) {
	k.GetEntry(ctx).Errorf(fmt.Sprintf(format, v...))
}
func (k *logRus) Warning(ctx context.Context, format string, v ...interface{}) {
	k.GetEntry(ctx).Warnf(fmt.Sprintf(format, v...))
}

func (k *logRus) Debug(ctx context.Context, format string, v ...interface{}) {
	k.GetEntry(ctx).Debugf(fmt.Sprintf(format, v...))
}

func (k *logRus) Info(ctx context.Context, format string, v ...interface{}) {
	k.GetEntry(ctx).Infof(fmt.Sprintf(format, v...))
}

func (k *logRus) Fatal(ctx context.Context, format string, v ...interface{}) {
	k.GetEntry(ctx).Fatalf(fmt.Sprintf(format, v...))
}

// just use for pointer of model
// @TBD
func (k *logRus) FieldMask(data interface{}) string {
	var rs interface{}
	maskTool := tool.NewMaskTool(filter.TagFilter())
	if k.env != common.EnvProd {
		rs = data
	} else {
		rs = maskTool.MaskDetails(data)
	}

	rb, err := json.Marshal(rs)
	if err != nil {
		return ""
	}
	return strings.ReplaceAll(string(rb), "[filtered]", "*****")
}
