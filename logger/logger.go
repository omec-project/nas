// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package logger

import (
	"os"
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/omec-project/util/logger"
	"github.com/omec-project/util/logger_conf"
	"github.com/sirupsen/logrus"
)

var (
	log         *logrus.Logger
	NasLog      *logrus.Entry
	NasMsgLog   *logrus.Entry
	ConvertLog  *logrus.Entry
	SecurityLog *logrus.Entry
)

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}

	free5gcLogHook, err := logger.NewFileHook(logger_conf.Free5gcLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err == nil {
		log.Hooks.Add(free5gcLogHook)
	}

	selfLogHook, err := logger.NewFileHook(logger_conf.LibLogDir+"nas.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err == nil {
		log.Hooks.Add(selfLogHook)
	}

	NasLog = log.WithFields(logrus.Fields{"component": "LIB", "category": "NAS"})
	NasMsgLog = log.WithFields(logrus.Fields{"component": "NAS", "category": "Message"})
	ConvertLog = log.WithFields(logrus.Fields{"component": "NAS", "category": "Convert"})
	SecurityLog = log.WithFields(logrus.Fields{"component": "NAS", "category": "Security"})
}

func SetLogLevel(level logrus.Level) {
	NasLog.Infoln("set log level :", level)
	log.SetLevel(level)
}

func SetReportCaller(set bool) {
	NasLog.Infoln("set report call :", set)
	log.SetReportCaller(set)
}
