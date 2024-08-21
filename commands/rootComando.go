package commands

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bep/simplecobra"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type rootComando struct {
	logger  zap.Logger
	Printf  func(format string, v ...interface{})
	Println func(a ...interface{})
	Out     io.Writer
	silence bool
	start   time.Time

	comandos []simplecobra.Commander
}

func (r *rootComando) Commands() []simplecobra.Commander {
	return r.comandos
}

func (r *rootComando) Name() string {
	return viper.GetString("app")
}

func (r *rootComando) Run(ctx context.Context, cmder *simplecobra.Commandeer, args []string) error {
	r.Println("Press Ctrl + C to stop texcoco")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	return nil
}

func (r *rootComando) PreRun(cmd, runnder *simplecobra.Commandeer) error {
	r.Out = os.Stdout

	if r.silence {
		r.Out = io.Discard
	}

	r.Printf = func(format string, v ...interface{}) {
		if !r.silence {
			fmt.Fprintf(r.Out, format, v...)
		}
	}

	r.Println = func(a ...interface{}) {
		if !r.silence {
			fmt.Fprintln(r.Out, a...)
		}
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	lumberjackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("logs/%s.log", viper.GetString("app")),
		MaxSize:    10, // megabtyes
		MaxBackups: 3,
		MaxAge:     7, // days
		Compress:   true,
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "@timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(lumberjackLogger),
		zap.InfoLevel,
	)

	r.logger = *zap.New(core)
	defer r.logger.Sync()

	r.start = time.Now()

	return nil
}

func (r *rootComando) Init(cmder *simplecobra.Commandeer) error {
	cmd := cmder.CobraCommand
	cmd.Use = "texcoco [flags]"
	cmd.Short = "texcoco get sport data"
	cmd.Long = `texcoco is the main command, used to get sport data`

	cmd.PersistentFlags().BoolVarP(&r.silence, "silence", "s", false, "silence mode")
	_ = cmd.RegisterFlagCompletionFunc("silence", cobra.NoFileCompletions)

	return nil
}

func (r *rootComando) executionTime(command string) {
	elapsed := time.Since(r.start)
	r.Printf("execution time for %s:%s (%s)\n", viper.GetString("app"), command, elapsed)
}
