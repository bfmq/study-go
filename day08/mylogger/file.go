package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

var (
	MaxSize = 50000
)

type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	Level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

func NewFileLog(levelStr string, fp, fn string, maxSize int64) *FileLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       level,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, MaxSize),
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open error log file failed,err:%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	// 开goroutine执行写日志任务
	go f.writeLogBackground()
	// // 开5个goroutine执行写日志任务
	// for i:=0;i<5;i++{
	// 	go f.writeLogBackground()
	// }
	return err
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file failed,err:%v\n", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	now := time.Now()
	nowStr := now.Format("20060102150405000")
	fileInfo, err := file.Stat()
	file.Close()
	if err != nil {
		fmt.Printf("get file failed,err:%v\n", err)
		return nil, err
	}

	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s/%s.bak%s", f.filePath, fileInfo.Name(), nowStr)
	os.Rename(logName, newLogName)
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed,err:%v\n", err)
		return nil, err
	}
	return fileObj, nil
}

func (f *FileLogger) writeLogBackground() {
	for {
		// 先判断是否需要切割日志
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		// 从通道中取日志
		select {
		case logmsg := <-f.logChan:
			log := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logmsg.timestamp, getLogString(logmsg.Level), logmsg.funcName, logmsg.fileName, logmsg.line, logmsg.msg)
			fmt.Fprint(f.fileObj, log)

			if logmsg.Level >= ERROR {
				if f.checkSize(f.fileObj) {
					newFile, err := f.splitFile(f.errFileObj)
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				fmt.Fprint(f.errFileObj, log)
			}
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)

		// 将logMsg对象扔进处理通道
		logmsg := &logMsg{
			Level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now.Format("2006-01-02 15:04:05"),
			line:      lineNo,
		}
		// 保证日志记录不会阻塞
		select {
		case f.logChan <- logmsg:
		default:
		}

	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
