package utils

import (
	"github.com/sirupsen/logrus"
	"go-project/global"
)

//打印info级别日志
func LogInfo(tag string, msg ...interface{}) {
	global.MyLogger.WithFields(logrus.Fields{"tag": tag}).Info(msg)
}

//打印info级别日志 格式后
func LogInfoF(tag string, format string, msg ...interface{}) {
	global.MyLogger.WithFields(logrus.Fields{"tag": tag}).Infof(format, msg)
}

//打印Warn级别日志
func LogWarn(tag string, msg ...interface{}) {
	global.MyLogger.WithFields(logrus.Fields{"tag": tag}).Warn(msg)
}

//打印Warn级别日志 格式后
func LogWarnF(tag string, format string, msg ...interface{}) {
	global.MyLogger.WithFields(logrus.Fields{"tag": tag}).Warnf(format, msg)
}

//打印error级别日志
func LogError(tag string, msg ...interface{}) {
	global.MyLogger.WithFields(logrus.Fields{"tag": tag}).Error(msg)
}

//打印error级别日志 格式后
func LogErrorF(tag string, format string, msg ...interface{}) {
	global.MyLogger.WithFields(logrus.Fields{"tag": tag}).Errorf(format, msg)
}
