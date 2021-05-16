package log_level

import (
	"fmt"
	"github.com/project-flogo/core/trigger"
)

func init() {
	fmt.Println("Init log-level-trigger metatada")
	_ = trigger.Register(&Trigger{}, &Factory{})
}

type Settings struct {
	ASetting int `md:"Setting"`
}

type HandlerSettings struct {
	ASetting string `md:"aSetting,required"`
}

type Output struct {
	AnOutput string `md:"anOutput"`
}

func (o *Output) FromMap(values map[string]interface{}) error {

	//var err error
	//o.AnOutput, err = coerce.ToString(values["anOutput"])
	//if err != nil {
	//	return err
	//}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"anOutput": o.AnOutput,
	}
}

type Reply struct {
	AReply interface{} `md:"aReply"`
}

func (r *Reply) FromMap(values map[string]interface{}) error {
	r.AReply = values["aReply"]
	return nil
}

func (r *Reply) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"aReply": r.AReply,
	}
}
