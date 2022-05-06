package e

import (
	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"

	oerrors "errors"
)

func Is(err error, target error) bool {
	return oerrors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return oerrors.As(err, target)
}

func IfErr(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func Ws(err error) error {
	return errors.WithStack(err)
}

func Wm(err error, msg string) error {
	return errors.WithMessage(err, msg)
}

// Something very low level.
func Traces(err error) bool {
	if err != nil {
		log.Tracef("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// Useful debugging information.
func Debugs(err error) bool {
	if err != nil {
		log.Debugf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// omething noteworthy happened!
func Infos(err error) bool {
	if err != nil {
		log.Infof("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// ou should probably take a look at this.
func Warns(err error) bool {
	if err != nil {
		log.Warnf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// Something failed but I'm not quitting. Calls os.Exit(1) after logging
func Errors(err error) bool {
	if err != nil {
		log.Errorf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// Bye. Calls panic() after logging
func Fatals(err error) bool {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// I'm bailing.
func Panics(err error) bool {
	if err != nil {
		log.Panicf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// Something very low level.
func IsTraces(err error, target error) bool {
	if Is(err, target) {
		log.Tracef("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// Useful debugging information.
func IsDebugs(err error, target error) bool {
	if Is(err, target) {
		log.Debugf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// omething noteworthy happened!
func IsInfos(err error, target error) bool {
	if Is(err, target) {
		log.Infof("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// ou should probably take a look at this.
func IsWarns(err error, target error) bool {
	if Is(err, target) {
		log.Warnf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// Something failed but I'm not quitting. Calls os.Exit(1) after logging
func IsErrors(err error, target error) bool {
	if Is(err, target) {
		log.Errorf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// Bye. Calls panic() after logging
func IsFatals(err error, target error) bool {
	if Is(err, target) {
		log.Fatalf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// I'm bailing.
func IsPanics(err error, target error) bool {
	if Is(err, target) {
		log.Panicf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// Something very low level.
func AsTraces(err error, target error) bool {
	if As(err, target) {
		log.Tracef("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// Useful debugging information.
func AsDebugs(err error, target error) bool {
	if As(err, target) {
		log.Debugf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// omething noteworthy happened!
func AsInfos(err error, target error) bool {
	if As(err, target) {
		log.Infof("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// ou should probably take a look at this.
func AsWarns(err error, target error) bool {
	if As(err, target) {
		log.Warnf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// Something failed but I'm not quitting. Calls os.Exit(1) after logging
func AsErrors(err error, target error) bool {
	if As(err, target) {
		log.Errorf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// Bye. Calls panic() after logging
func AsFatals(err error, target error) bool {
	if As(err, target) {
		log.Fatalf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// I'm bailing.
func AsPanics(err error, target error) bool {
	if As(err, target) {
		log.Panicf("%+v\n", errors.WithStack(err))
		return true
	}
	return false
}

// Something very low level.
func Trace(err error) bool {
	if err != nil {
		log.Trace(errors.WithStack(err))
		return true
	}
	return false
}

// Useful debugging information.
func Debug(err error) bool {
	if err != nil {
		log.Debug(errors.WithStack(err))
		return true
	}
	return false
}

// omething noteworthy happened!
func Info(err error) bool {
	if err != nil {
		log.Info(errors.WithStack(err))
		return true
	}
	return false
}

// ou should probably take a look at this.
func Warn(err error) bool {
	if err != nil {
		log.Warn(errors.WithStack(err))
		return true
	}
	return false
}

// Something failed but I'm not quitting. Calls os.Exit(1) after logging
func Error(err error) bool {
	if err != nil {
		log.Error(errors.WithStack(err))
		return true
	}
	return false
}

// Bye. Calls panic() after logging
func Fatal(err error) bool {
	if err != nil {
		log.Fatal(errors.WithStack(err))
		return true
	}
	return false
}

// I'm bailing.
func Panic(err error) bool {
	if err != nil {
		log.Panic(errors.WithStack(err))
		return true
	}
	return false
}

// Something very low level.
func IsTrace(err error, target error) bool {
	if Is(err, target) {
		log.Trace(errors.WithStack(err))
		return true
	}
	return false
}

// Useful debugging information.
func IsDebug(err error, target error) bool {
	if Is(err, target) {
		log.Debug(errors.WithStack(err))
		return true
	}
	return false
}

// omething noteworthy happened!
func IsInfo(err error, target error) bool {
	if Is(err, target) {
		log.Info(errors.WithStack(err))
		return true
	}
	return false
}

// ou should probably take a look at this.
func IsWarn(err error, target error) bool {
	if Is(err, target) {
		log.Warn(errors.WithStack(err))
		return true
	}
	return false
}

// Something failed but I'm not quitting. Calls os.Exit(1) after logging
func IsError(err error, target error) bool {
	if Is(err, target) {
		log.Error(errors.WithStack(err))
		return true
	}
	return false
}

// Bye. Calls panic() after logging
func IsFatal(err error, target error) bool {
	if Is(err, target) {
		log.Fatal(errors.WithStack(err))
		return true
	}
	return false
}

// I'm bailing.
func IsPanic(err error, target error) bool {
	if Is(err, target) {
		log.Panic(errors.WithStack(err))
		return true
	}
	return false
}

// Something very low level.
func AsTrace(err error, target interface{}) bool {
	if As(err, target) {
		log.Trace(errors.WithStack(err))
		return true
	}
	return false
}

// Useful debugging information.
func AsDebug(err error, target interface{}) bool {
	if As(err, target) {
		log.Debug(errors.WithStack(err))
		return true
	}
	return false
}

// omething noteworthy happened!
func AsInfo(err error, target interface{}) bool {
	if As(err, target) {
		log.Info(errors.WithStack(err))
		return true
	}
	return false
}

// ou should probably take a look at this.
func AsWarn(err error, target interface{}) bool {
	if As(err, target) {
		log.Warn(errors.WithStack(err))
		return true
	}
	return false
}

// Something failed but I'm not quitting. Calls os.Exit(1) after logging
func AsError(err error, target interface{}) bool {
	if As(err, target) {
		log.Error(errors.WithStack(err))
		return true
	}
	return false
}

// Bye. Calls panic() after logging
func AsFatal(err error, target interface{}) bool {
	if As(err, target) {
		log.Fatal(errors.WithStack(err))
		return true
	}
	return false
}

// I'm bailing.
func AsPanic(err error, target interface{}) bool {
	if As(err, target) {
		log.Panic(errors.WithStack(err))
		return true
	}
	return false
}
