// SPDX-FileCopyrightText: 2022 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package logger

import (
	"os"
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	logger_util "github.com/omec-project/util/logger"
	"github.com/omec-project/util/logger_conf"
	"github.com/sirupsen/logrus"
)

var (
	log             *logrus.Logger
	CommTestCommLog *logrus.Entry
	CommTestAmfLog  *logrus.Entry
	CommTestSmfLog  *logrus.Entry
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

	free5gcLogHook, err := logger_util.NewFileHook(logger_conf.Free5gcLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err == nil {
		log.Hooks.Add(free5gcLogHook)
	}

	selfLogHook, err := logger_util.NewFileHook(logger_conf.LibLogDir+"common_consumer_test_data.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err == nil {
		log.Hooks.Add(selfLogHook)
	}

	CommTestCommLog = log.WithFields(logrus.Fields{"component": "CommonTest", "category": "Comm"})
	CommTestAmfLog = log.WithFields(logrus.Fields{"component": "CommonTest", "category": "Amf"})
	CommTestSmfLog = log.WithFields(logrus.Fields{"component": "CommonTest", "category": "Smf"})
}

func SetLogLevel(level logrus.Level) {
	CommTestCommLog.Infoln("set log level :", level)
	log.SetLevel(level)
}

func SetReportCaller(reportCaller bool) {
	CommTestCommLog.Infoln("set report call :", reportCaller)
	log.SetReportCaller(reportCaller)
}
