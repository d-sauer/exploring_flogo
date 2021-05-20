package log_level

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"time"
)

func createEndpoint(t *Trigger) func() {
	t.logger.Debug("Start creating endpoint")
	path := t.settings.Path

	restoreLogger := initLogger(path, 1065)

	return restoreLogger
}

func initLogger(path string, port int) func() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		nanos := t.UnixNano()
		millis := nanos / int64(time.Millisecond)
		enc.AppendInt64(millis)
	}
	//encoder := zapcore.NewJSONEncoder(config)
	encoder := zapcore.NewConsoleEncoder(config)
	atom := zap.NewAtomicLevel()
	logr := zap.New(zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), atom))

	mux := http.NewServeMux()
	mux.Handle(path, atom)
	go http.ListenAndServe(fmt.Sprintf(":%s", port), mux)

	return zap.ReplaceGlobals(logr)
}
