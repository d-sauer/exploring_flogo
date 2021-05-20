package log_level

import (
	"errors"
	"fmt"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
	"time"
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
	settings      *Settings
	handlers      []trigger.Handler
	logger        log.Logger
	restoreLogger func()
}

// Initialize initializes the trigger
func (t *Trigger) Initialize(ctx trigger.InitContext) error {
	fmt.Println("Initialize trigger")

	t.handlers = ctx.GetHandlers()
	t.logger = ctx.Logger()
	path := t.settings.Path

	t.logger.Infof("Initialize endpoint path: `%s`", path)

	restoreLoggerFunc := createEndpoint(t)
	if restoreLoggerFunc == nil {
		return errors.New("Can't create custom logger")
	}

	t.restoreLogger = restoreLoggerFunc

	return nil
}

func (t Trigger) Start() error {
	t.logger.Info("Starting trigger logic: log-level")
	const MAX = 100

	for i := 0; i < MAX; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Count: ", i, " of ", MAX)

		t.logger.Info("Test info message i:", i)
		t.logger.Debug("Test debug message i:", i)
	}
	return nil
}

func (t Trigger) Stop() error {
	t.logger.Info("Stopping trigger")
	defer t.restoreLogger()

	return nil
}
