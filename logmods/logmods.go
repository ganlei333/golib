package logmods

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//SugarLogger 日志对象指针
var SugarLogger *zap.SugaredLogger

//LogConf 日志配置结构体
type LogConf struct {
	LogFilename   string
	LogMaxSize    int
	LogMaxBackups int
	LogMaxAge     int
	LogCompress   bool
	LogLevel      string
}

// func main() {
//     InitLogger()
//     defer sugarLogger.Sync()
//     simpleHttpGet("www.topgoer.com")
//     simpleHttpGet("http://www.topgoer.com")
// }int

//InitLogger 初始化日志对象
func InitLogger(logconf *LogConf) {
	writeSyncer := getLogWriter(logconf)
	encoder := getEncoder()
	//zapcore.DebugLevel
	var core zapcore.Core
	switch logconf.LogLevel {
	case "Debug":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	case "Info":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	case "Warn":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.WarnLevel)
	case "Error":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.ErrorLevel)
	case "DPanic":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.DPanicLevel)
	case "Panic":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.PanicLevel)
	case "Fatal":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.FatalLevel)
	default:
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.FatalLevel)
	}
	//core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logconf *LogConf) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logconf.LogFilename,
		MaxSize:    logconf.LogMaxSize,
		MaxBackups: logconf.LogMaxBackups,
		MaxAge:     logconf.LogMaxAge,
		Compress:   logconf.LogCompress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// func simpleHttpGet(url string) {
//     sugarLogger.Debugf("Trying to hit GET request for %s", url)
//     resp, err := http.Get(url)
//     if err != nil {
//         sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
//     } else {
//         sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
//         resp.Body.Close()
//     }
// }
