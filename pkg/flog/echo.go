package flog

import (
	"bytes"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/labstack/gommon/log"
	"github.com/valyala/fasttemplate"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"unsafe"
)

func NewEchoLogger(level string) (e *EchoLogger) {
	e = &EchoLogger{
		level:    uint32(0),
		skip:     0,
		prefix:   "",
		template: e.newTemplate(""),
		bufferPool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 256))
			},
		},
	}
	e.SetOutput(os.Stdout)
	return e
}

type EchoLogger struct {
	prefix     string
	level      uint32
	skip       int
	output     io.Writer
	template   *fasttemplate.Template
	levels     []string
	bufferPool sync.Pool
	mutex      sync.Mutex
}

func (e *EchoLogger) Output() io.Writer {
	return e.output
}

func (e *EchoLogger) SetOutput(w io.Writer) {
	e.output = w
}

func (e *EchoLogger) Prefix() string {
	return e.prefix
}

func (e *EchoLogger) SetPrefix(p string) {
	e.prefix = p
}

func (e *EchoLogger) Level() log.Lvl {
	return log.Lvl(atomic.LoadUint32(&e.level))
}

func (e *EchoLogger) SetLevel(v log.Lvl) {
	atomic.StoreUint32(&e.level, uint32(v))
}

func (e *EchoLogger) SetHeader(h string) {
	e.template = e.newTemplate(h)
}

func (e *EchoLogger) newTemplate(format string) *fasttemplate.Template {
	return fasttemplate.New(format, "${", "}")
}

func (e *EchoLogger) Print(i ...interface{}) {
	l.Print(i...)
}

func (e *EchoLogger) Printf(format string, args ...interface{}) {
	l.Printf(format, args...)
}

func (e *EchoLogger) Printj(j log.JSON) {
	l.Print(jsonString(j))
}

func (e *EchoLogger) Debug(i ...interface{}) {
	l.Debug().Caller(1).Msg(fmt.Sprint(i...))
}

func (e *EchoLogger) Debugf(format string, args ...interface{}) {
	l.Debug().Caller(1).Msg(fmt.Sprintf(format, args...))
}

func (e *EchoLogger) Debugj(j log.JSON) {
	l.Debug().Caller(1).Msg(jsonString(j))
}

func (e *EchoLogger) Info(i ...interface{}) {
	l.Info().Caller(1).Msg(fmt.Sprint(i...))
}

func (e *EchoLogger) Infof(format string, args ...interface{}) {
	l.Info().Caller(1).Msg(fmt.Sprintf(format, args...))
}

func (e *EchoLogger) Infoj(j log.JSON) {
	l.Info().Caller(1).Msg(jsonString(j))
}

func (e *EchoLogger) Warn(i ...interface{}) {
	l.Warn().Caller(1).Msg(fmt.Sprint(i...))
}

func (e *EchoLogger) Warnf(format string, args ...interface{}) {
	l.Warn().Caller(1).Msg(fmt.Sprintf(format, args...))
}

func (e *EchoLogger) Warnj(j log.JSON) {
	l.Warn().Caller(1).Msg(jsonString(j))
}

func (e *EchoLogger) Error(i ...interface{}) {
	l.Error().Caller(1).Msg(fmt.Sprint(i...))
}

func (e *EchoLogger) Errorf(format string, args ...interface{}) {
	l.Error().Caller(1).Msg(fmt.Sprintf(format, args...))
}

func (e *EchoLogger) Errorj(j log.JSON) {
	l.Error().Caller(1).Msg(jsonString(j))
}

func (e *EchoLogger) Fatal(i ...interface{}) {
	l.Fatal().Caller(1).Msg(fmt.Sprint(i...))
}

func (e *EchoLogger) Fatalj(j log.JSON) {
	l.Fatal().Caller(1).Msg(jsonString(j))
}

func (e *EchoLogger) Fatalf(format string, args ...interface{}) {
	l.Fatal().Caller(1).Msg(fmt.Sprintf(format, args...))
}

func (e *EchoLogger) Panic(i ...interface{}) {
	l.Panic().Caller(1).Msg(fmt.Sprint(i...))
}

func (e *EchoLogger) Panicj(j log.JSON) {
	l.Panic().Caller(1).Msg(jsonString(j))
}

func (e *EchoLogger) Panicf(format string, args ...interface{}) {
	l.Panic().Caller(1).Msg(fmt.Sprintf(format, args...))
}

func jsonString(j log.JSON) string {
	b, _ := sonic.Marshal(j)
	return *(*string)(unsafe.Pointer(&b))
}
