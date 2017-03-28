package null

import (
	api "github.com/CoachApplication/coach-api"
	base "github.com/CoachApplication/coach-base"
)

type BaseOperation struct{}

func (gop *BaseOperation) Validate() api.Result {
	return base.MakeSuccessfulResult()
}
func (gop *BaseOperation) Exec(props api.Properties) api.Result {
	return base.MakeFailedResult()
}
