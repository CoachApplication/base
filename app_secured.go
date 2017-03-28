package standard

import (
	api "github.com/CoachApplication/coach-api"
)

type SecuredApplication struct {
	Application

	authOpDecFactory AuthorizedOperationDecorationFactory
}

type AuthorizedOperationDecorationFactory interface {
	DecorateOperation(api.Operation) api.Operation
}

// NewSecuredApplication Constructor for SecuredApplication
func NewSecuredApplication(authOpDecFactory AuthorizedOperationDecorationFactory) *SecuredApplication {
	return &SecuredApplication{
		Application:      *NewApplication(nil),
		authOpDecFactory: authOpDecFactory,
	}
}

func (sa *SecuredApplication) Operations() api.Operations {
	ops := NewOperations()

	targetOps := sa.Application.Operations()
	for _, opId := range targetOps.Order() {
		targetOp, _ := targetOps.Get(opId)
		ops.Add(sa.authOpDecFactory.DecorateOperation(targetOp))
	}

	return ops.Operations()
}
