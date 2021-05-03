package app

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap/zapcore"
	"net/url"
	"regexp"
	"time"
)

type Args struct {
	ListenAddress        string
	PrimaryDatabaseDSN   string
	SecondaryDatabaseDSN string
	ImageStorageURI      string
	LogLevel             string
	ShutdownTimeout      time.Duration
}

func NewArgs(opts *docopt.Opts) (*Args, error) {
	listenAddress, _ := opts.String("--listen")
	if checkListenAddress(listenAddress) == false {
		return nil, fmt.Errorf("invalid option --listen; format: 127.0.0.1:80")
	}

	primaryDatabaseDSN, _ := opts.String("<primary-dsn>")
	if checkPostgresDSN(primaryDatabaseDSN) == false {
		return nil, fmt.Errorf("invalid argument <primary-dsn>; format: postgres://user:pass@host:port/db?option=value")
	}

	secondaryDatabaseDSN, _ := opts.String("--secondary-db")
	if secondaryDatabaseDSN != "" && checkPostgresDSN(secondaryDatabaseDSN) == false {
		return nil, fmt.Errorf("invalid option --secondary-db; format: postgres://user:pass@host:port/db?option=value")
	}

	imageStorageURI, _ := opts.String("--image-storage")
	if checkImageStorageURI(imageStorageURI) == false {
		return nil, fmt.Errorf("invalid option --image-storage; format: http(s)://id:secret@host:port/bucket")
	}

	logLevel, _ := opts.String("--log-level")
	if checkLogLevel(logLevel) == false {
		return nil, fmt.Errorf("invalid option --log-level; allowed values: %v", getAllowedLogLevels())
	}

	shutdownTimeout, _ := opts.Int("--shutdown-timeout")
	if shutdownTimeout < 0 {
		return nil, fmt.Errorf("invalid option --shutdown-timeout; must be greater or equal zero")
	}

	return &Args{
		ListenAddress:        listenAddress,
		PrimaryDatabaseDSN:   primaryDatabaseDSN,
		SecondaryDatabaseDSN: secondaryDatabaseDSN,
		ImageStorageURI:      imageStorageURI,
		LogLevel:             logLevel,
		ShutdownTimeout:      time.Duration(shutdownTimeout) * time.Second,
	}, nil
}

func checkListenAddress(addr string) bool {
	pattern := regexp.MustCompile(`^(?P<ip>\d+\.\d+\.\d+\.\d+):(?P<port>\d+)$`)

	return pattern.MatchString(addr)
}

func checkPostgresDSN(dsn string) bool {
	_, err := pg.ParseURL(dsn)

	return err == nil
}

func checkImageStorageURI(uri string) bool {
	parsed, err := url.Parse(uri)
	if err != nil {
		return false
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return false
	}

	if parsed.User.Username() == "" {
		return false
	}

	_, isPasswordSet := parsed.User.Password()
	if !isPasswordSet {
		return false
	}

	if parsed.Path == "" {
		return false
	}

	return true
}

func checkLogLevel(level string) bool {
	for _, l := range getAllowedLogLevels() {
		if level == l {
			return true
		}
	}

	return false
}

func getAllowedLogLevels() []string {
	return []string{
		zapcore.DebugLevel.String(),
		zapcore.InfoLevel.String(),
		zapcore.WarnLevel.String(),
		zapcore.ErrorLevel.String(),
	}
}