package log_level

import (
	"fmt"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

var triggerMetadata = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{}, &Reply{})

func init() {
	fmt.Println("Init log-level-trigger")
	_ = trigger.Register(&Trigger{}, &Factory{})
}

// Factory is a kafka trigger factory
type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMetadata
}

// New implements trigger.Factory.New
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}

	return &Trigger{settings: s}, nil
}

// Trigger ...
type Trigger struct {
	settings *Settings
	handlers []trigger.Handler
	logger   log.Logger
}

// Initialize initializes the trigger
func (t *Trigger) Initialize(ctx trigger.InitContext) error {
	fmt.Println("Initialize trigger")

	t.handlers = ctx.GetHandlers()
	t.logger = ctx.Logger()
	path := t.settings.Path
	t.logger.Infof("Initialize endpoint path: `%s`", path)

	err := createEndpoint(t)

	return err
}

func (t Trigger) Start() error {
	t.logger.Info("Starting trigger logic: log-level")

	return nil
}

func (t Trigger) Stop() error {
	t.logger.Info("Stopping trigger")
	return nil
}

//
//type Trigger struct {
//	settings *Settings
//	logger   log.Logger
//	handlers map[string]*clientHandler
//}
//
//type clientHandler struct {
//	handler  trigger.Handler
//	settings *HandlerSettings
//}
//
//type Factory struct {
//}
//
//func (t Trigger) Start() error {
//	t.logger.Info("Starting trigger logic: log-level")
//
//	return nil
//}
//
//func (t Trigger) Stop() error {
//	t.logger.Info("Stopping trigger: log-level")
//
//	return nil
//}
//
//func (t Trigger) Initialize(ctx trigger.InitContext) error {
//	t.logger = ctx.Logger()
//	t.logger.Info("Initialising log-trigger")
//
//	t.handlers = make(map[string]*clientHandler)
//
//	return nil
//}
//
//func (f Factory) Metadata() *trigger.Metadata {
//	return triggerMetadata
//}
//
//func (f Factory) New(config *trigger.Config) (trigger.Trigger, error) {
//	s := &Settings{}
//
//	err := metadata.MapToStruct(config.Settings, s, true)
//	if err != nil {
//		return nil, err
//	}
//
//	return &Trigger{settings: s}, nil
//}
