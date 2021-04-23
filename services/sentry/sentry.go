package sentry

import (
	"cliTest/config"
	"github.com/getsentry/sentry-go"
	"log"
)

type sentryService struct{}

var (
	SentryService = sentryService{}
	configData    = config.GetInstance()
)

// Init initialization sentry service
func (s *sentryService) Init() {
	sentry.Logger.SetPrefix("[sentry sdk] ")
	sentry.Logger.SetFlags(log.Ldate | log.Ltime)

	err := sentry.Init(sentry.ClientOptions{
		Dsn:   configData.Sentry.Dsn,
		Debug: configData.DevMode,
	})

	if err != nil {
		log.Fatalf("sentry initialization failed: %s", err)
	}
}

func (s *sentryService) Message(extras *map[string]interface{}, tag, message string) {
	sentry.WithScope(func(scope *sentry.Scope) {
		setScopeDetails(tag, extras, scope)
		scope.SetLevel(sentry.LevelInfo)
		sentry.CaptureMessage(message)
	})
}

func (s *sentryService) Error(extras *map[string]interface{}, tag string, err error) {
	sentry.WithScope(func(scope *sentry.Scope) {
		setScopeDetails(tag, extras, scope)
		scope.SetLevel(sentry.LevelError)
		sentry.CaptureException(err)
	})
}

// setScopeDetails for configuration sentry scope
func setScopeDetails(tag string, extras *map[string]interface{}, scope *sentry.Scope) {
	scope.SetTag("tag", tag)
	if extras != nil {
		scope.SetExtras(*extras)
	}
}
