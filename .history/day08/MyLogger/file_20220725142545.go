package myLogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type LogMsg struct {
	Level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timeStamp string
	line      int
}

type FileLogger struct {
	Level       LogLevel
	filePath    string // 文件路径
	fileName    string // 文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *LogMsg
}

func NewFileLogger(levelstr, filepath, filename string, max int64) *FileLogger {
	logLevel, err := parseLogLevel(levelstr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    filepath,
		fileName:    filename,
		maxFileSize: max,
		logChan:     make(chan *LogMsg, 50000),
	}
	err = fl.initFile() // 按照文件路径与文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

// 初始化
func (f *FileLogger) initFile() error {
	fp := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed,%v\n", err)
		return err
	}
	errfileObj, err := os.OpenFile(fp+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err file failed,%v\n", err)
		return err
	}
	//日志文件都已经打开
	f.fileObj = fileObj
	f.errFileObj = errfileObj
	return nil
}

// 关闭日志
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

// 判断是否需要记录该日志
func (l *FileLogger) enable(level LogLevel) bool {
	return l.Level <= level
}


func (f *FileLogger) writeLogBackGround(){
	
}

// 记录日志的方法
func (l *FileLogger) log(lv LogLevel, msg string, arg ...interface{}) {
	if l.enable(lv) {
		fotmat := fmt.Sprintf(msg, arg...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		if l.checkSize(l.fileObj) {
			newFile, err := l.splitFile(l.fileObj)
			if err != nil {
				fmt.Printf("split file failed ,%v\n", err)
				return
			}
			l.fileObj = newFile
		}
		fmt.Fprintf(l.fileObj, "[%s] [%s] [%s: %s: %d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, fotmat)
		if lv >= ERROR {
			//如果要记录的文件大于等于ERROR级别，还要在err日志文件中再记录一遍
			if l.checkSize(l.errFileObj) {
				newFile, err := l.splitFile(l.errFileObj)
				if err != nil {
					fmt.Printf("split file failed ,%v\n", err)
					return
				}
				l.errFileObj = newFile
			}
			fmt.Fprintf(l.errFileObj, "[%s] [%s] [%s: %s: %d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, fotmat)
		}
	}
}

// 判断日志是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}
	//如果当前文件大小 大于等于 日志文件的最大值就返回true
	return fileinfo.Size() >= f.maxFileSize
}

// 分割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要进行切割

	// 备份 rename
	nowStr := time.Now().Format("20060102030405000")
	// 获取文件对象的详细信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return nil, err
	}
	LogName := path.Join(f.filePath, fileInfo.Name())                        //拿到当前日志文件的完整路径
	nowLogName := fmt.Sprintf("%s/%s.bak%s", f.filePath, f.fileName, nowStr) // 拼接一个日志文件备份的名字
	// 关闭当前日志文件
	file.Close()

	os.Rename(LogName, nowLogName)
	// 打开一个新的日志文件
	fileObj, err := os.OpenFile(LogName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Open new log file ,err %v\n", err)
		return nil, err
	}

	// 将打开的新日志文件对象赋值给 f.fileObj
	return fileObj, nil
}

func (l *FileLogger) Debug(msg string, arg ...interface{}) {
	l.log(DEBUG, msg, arg...)
}

func (l *FileLogger) Info(msg string, arg ...interface{}) {
	l.log(INFO, msg, arg...)
}

func (l *FileLogger) Warning(msg string, arg ...interface{}) {
	l.log(WARNING, msg, arg...)
}

func (l *FileLogger) Error(msg string, arg ...interface{}) {
	l.log(ERROR, msg, arg...)
}

func (l *FileLogger) Fatal(msg string, arg ...interface{}) {
	l.log(FATAL, msg, arg...)
}
