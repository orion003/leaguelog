package logging

import (
	log15 "leaguelog/Godeps/_workspace/src/gopkg.in/inconshreveable/log15.v2"
)

type Logger interface {
	Debug(msg string, ctx ...interface{})
	Info(msg string, ctx ...interface{})
	Warn(msg string, ctx ...interface{})
	Error(msg string, ctx ...interface{})
	Critical(msg string, ctx ...interface{})
}

type Log15 struct {
	log log15.Logger
}

func NewLog15() Log15 {
	log := Log15{
		log: log15.Root(),
	}

	return log
}

func (l Log15) Debug(msg string, ctx ...interface{}) {
	l.log.Debug(msg, ctx...)
}

func (l Log15) Info(msg string, ctx ...interface{}) {
	l.log.Info(msg, ctx...)
}

func (l Log15) Warn(msg string, ctx ...interface{}) {
	l.log.Warn(msg, ctx...)
}

func (l Log15) Error(msg string, ctx ...interface{}) {
	l.log.Error(msg, ctx...)
}

func (l Log15) Critical(msg string, ctx ...interface{}) {
	l.log.Crit(msg, ctx...)
}
