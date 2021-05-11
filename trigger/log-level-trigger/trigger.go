package log_level

import (
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

var triggerMetadata = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{})

func init() {
	_ = trigger.Register(&Trigger{}, &Factory{})
}

type Trigger struct {
	settings *Settings
	logger   log.Logger
	handlers map[string]*clientHandler
}

type clientHandler struct {
	handler  trigger.Handler
	settings *HandlerSettings
}

type Factory struct {
}

func (t Trigger) Start() error {
	t.logger.Info("Starting trigger logic: log-level")

	return nil
}

func (t Trigger) Stop() error {
	t.logger.Info("Stopping trigger: log-level")

	return nil
}

func (t Trigger) Initialize(ctx trigger.InitContext) error {
	t.logger = ctx.Logger()

	t.logger.Info("Initialising log-trigger")
	return nil
}

func (f Factory) Metadata() *trigger.Metadata {
	return triggerMetadata
}

func (f Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}

	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}

	return &Trigger{settings: s}, nil
}
