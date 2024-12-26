package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8

type Filelds map[string]interface{}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Filelds
	callers   []string
}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// 将传入的日志级别（Level类型的值）转换为字符串
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	//l := log.New(w, prefix, flag)
	//return &Logger{newLogger: l}
	return &Logger{newLogger: log.New(w, prefix, flag)} //用一行写完逻辑，比上面的代码更简洁
}

// 克隆 Logger 对象，返回一个新的 Logger 对象，它具有与当前 Logger 对象相同的配置，但可能有一些额外的字段或不同的上下文
func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

// 设置公共字段
func (l *Logger) WithFields(f Filelds) *Logger {
	ll := l.clone()       // 创建一个Logger的对象ll
	if ll.fields == nil { // 如果Logger对象的fields字段为空，则创建一个新的map
		ll.fields = make(Filelds)
	}
	for k, v := range f { // 遍历传入的字段，将其添加到Logger对象的fields字段中
		ll.fields[k] = v
	}
	return ll // 返回新的Logger对象
}

// 设置日志日志上下文
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone() //克隆一个Logger对象
	ll.ctx = ctx    //将传入的上下文赋值给Logger对象的ctx字段
	return ll       //返回新的Logger对象
}

// 设置当前某一层调用栈的信息
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return ll
}

// 设置当前所有调用栈的信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

// 将日志数据格式化为 JSON 格式，便于记录和传输
func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	// 创建一个新的 map，用于存储日志数据
	data := make(Filelds, len(l.fields)+4)
	// 将日志的级别转换为字符串，并存储在 map 中，键为 "level"
	data["level"] = level.String()
	// 将日志的上下文转换为字符串，并存储在 map 中，键为 "context"
	data["context"] = l.ctx
	// 将日志的时间转换为 Unix 时间戳，并存储在 map 中，键为 "time"
	data["time"] = time.Now().Local().UnixNano()
	// 将日志的消息存储在 map 中，键为 "message"
	data["message"] = message
	// 将日志的调用栈信息存储在 map 中，键为 "callers"
	data["callers"] = l.callers
	// 如果日志的公共字段不为空，则将其存储在 map 中，键为日志的公共字段的键
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}

func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message)) //将日志数据格式化为 JSON 格式
	content := string(body)                               //将 JSON 格式的日志数据转换为字符串
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Print(content)
	case LevelPanic:
		l.newLogger.Print(content)
		//default:
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.Output(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Output(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.Output(LevelWarn, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Output(LevelWarn, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.Output(LevelError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(LevelError, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v ...interface{}) {
	l.Output(LevelPanic, fmt.Sprint(v...))
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.Output(LevelPanic, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprintf(format, v...))
}
