package log_level

import "github.com/project-flogo/core/data/coerce"

//
// Settings
//

type Settings struct {
	Path string `md:"path" default:"/app/log-level"` // Endpoint path
}

func (s *Settings) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"path": s.Path,
	}
}

func (s *Settings) fromMap(values map[string]interface{}) error {
	var err error

	s.Path, err = coerce.ToString(values["path"])
	if err != nil {
		return err
	}

	return nil
}

//
// Output
//

type Output struct {
	LogLevel string `md:"logLevel"` // The data received from path
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"logLevel": o.LogLevel,
	}
}

func (o *Output) fromMap(values map[string]interface{}) error {
	var err error

	o.LogLevel, err = coerce.ToString(values["logLevel"])
	if err != nil {
		return err
	}

	return nil
}

//
// Reply
//

type Reply struct {
	Reply string `md:"reply"` // The reply send back as confirmation for changed log level
}

func (r *Reply) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"reply": r.Reply,
	}
}

func (r *Reply) fromMap(values map[string]interface{}) error {
	var err error

	r.Reply, err = coerce.ToString(values["reply"])
	if err != nil {
		return err
	}

	return nil
}

//
// HandlerSettings
//

type HandlerSettings struct {
	FlowLogLevel string `md:"flowLogLevel"` // Log level per flow instance
}

func (hs *HandlerSettings) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"flowLogLevel": hs.FlowLogLevel,
	}
}

func (hs *HandlerSettings) fromMap(values map[string]interface{}) error {
	var err error

	hs.FlowLogLevel, err = coerce.ToString(values["flowLogLevel"])
	if err != nil {
		return err
	}

	return nil
}
