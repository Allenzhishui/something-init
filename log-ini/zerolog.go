package log_ini

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"strings"
)

func SetLogLevel() {
	env := viper.GetString("LOG_LEVEL")
	level := zerolog.InfoLevel
	if env != "" {

		switch strings.ToLower(env) {
		case "trace":
			level = zerolog.TraceLevel
		case "debug":
			level = zerolog.DebugLevel
		case "info":
			level = zerolog.InfoLevel
		case "warn":
			level = zerolog.WarnLevel
		case "error":
			level = zerolog.ErrorLevel
		case "fatal":
			level = zerolog.FatalLevel
		case "panic":
			level = zerolog.PanicLevel
		}
	}
	zerolog.SetGlobalLevel(level)
	log.Debug().Str("log_level", level.String()).Msgf("the log level will be set as %s", env) //不会打印出来此条日志
	log.Info().Str("log_level", level.String()).Msgf("the log level will be set as %s", env)
	log.Error().Str("log_level", level.String()).Msgf("the log level will be set as %s", env)
}
