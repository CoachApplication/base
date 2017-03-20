package standard

import (
	"errors"
	api "github.com/james-nesbitt/coach-api"
)

type SecuredApplication struct {
	Application

	authOp api.Operation
	authOpPropId string
	authOpSuccessPropId string
}

// NewSecuredApplication Constructor for SecuredApplication
func NewSecuredApplication(authOp api.Operation, authOpPropId string, authOpSuccessPropId string) *SecuredApplication {
	return &SecuredApplication{
		authOp: authOp,
		authOpPropId: authOpPropId,
		authOpSuccessPropId: authOpSuccessPropId,
	}
}

func (sa *SecuredApplication) Operations() api.Operations {
	ops := NewOperations()

	targetOps := sa.Application.Operations()
	for _, opId := range targetOps.Order() {
		targetOp, _ := targetOps.Get(opId)
		authorizedOp := NewAuthorizingDecoratorOperation(targetOp, sa.authOp, sa.authOpPropId, sa.authOpSuccessPropId)
		ops.Add(api.Operation(authorizedOp))
	}

	return ops.Operations()
}

/**
 * SecuredApplication Decorating Operation for authentication
 */

// AuthorizingDecoratorOperation An operation which decorates another operation and Authorizes before running the Exec()
type AuthorizingDecoratorOperation struct {
	targetOp api.Operation
	authOp api.Operation

	authOpPropertyId string
	authSuccessPropertyId string
}

// NewAuthorizingDecoratorOperation Create a new authorizer decorator from an op, and an ops list
// The ops list is expected to contain any Operation objects needed to authorize.
func NewAuthorizingDecoratorOperation(op api.Operation, authOp api.Operation, authOpPropertyId string, authSuccessPropertyId string) *AuthorizingDecoratorOperation {
	return &AuthorizingDecoratorOperation{
		targetOp: op,
		authOp: authOp,
		authOpPropertyId: authOpPropertyId,
		authSuccessPropertyId: authSuccessPropertyId,
	}
}

// Id Provide a unique machine string identifier
func (ado *AuthorizingDecoratorOperation) Id() string {

}

// Ui Define UI elements for the Operation
func (ado *AuthorizingDecoratorOperation) Ui() api.Ui {
	return ado.targetOp.Ui()
}

// Usage Define how the operation is intended to be used
func (ado *AuthorizingDecoratorOperation) Usage() api.Usage {
	return ado.targetOp.Usage()
}

// Properties Provide Exec properties with default values
func (ado *AuthorizingDecoratorOperation) Properties() api.Properties {
	props := NewProperties()

	// auth Operation props
	props.Merge(ado.authOp.Properties())
	// target Operation Properties
	props.Merge(ado.targetOp.Properties())

	return props.Properties()
}

// Validate That the Operation can properly Execute and Authorize
func (ado *AuthorizingDecoratorOperation) Validate(props api.Properties) api.Result {
	res := NewResult()

	successProp, successPropFound := props.Get(ado.authSuccessPropertyId)
	operationProp, operationPropFound := props.Get(ado.authOpPropertyId)

	if successPropFound != nil || successProp.Type() != "bool" {
		// this authorization op is not valid, it is either missing its op or success property
		res.MarkFailed()

		if operationPropFound != nil {
			res.AddError(errors.New("Secure API Authorize operation is invalid.  It is missing the operation authorization property."))
		}
		if successPropFound != nil {
			res.AddError(errors.New("Secure API Authorize operation is invalid.  It is missing the authorization success property."))
		} else if successProp.Type() != "bool" {
			res.AddError(errors.New("Secure API Authorize operation is invalid.  The authorization success property is not a bool."))
		}

		res.AddError(errors.New("Secure Builder API could not execute authorized Operation."))
	}
	if operationPropFound != nil || operationProp.Type() != "coach.api.Operation" {
		res.MarkFailed()

		if op, success := operationProp.Get().(api.Operation); !success {
			res.AddError(errors.New("Secure API Authorize operation is invalud. No target operation was provided"))
		} else {
			targetRes := op.Validate()
			res.Merge(targetRes)
		}
	}

	res.MarkFinished()
	return res.Result()
}

// Exec Execute the Operation using passed Properties
func (ado *AuthorizingDecoratorOperation) Exec(props api.Properties) api.Result {
	res := NewResult()

	successProp, successPropFound := props.Get(ado.authSuccessPropertyId)
	operationProp, operationPropFound := props.Get(ado.authOpPropertyId)

	if successPropFound != nil || operationPropFound != nil || successProp.Type() != "bool" {
		// this authorization op is not valid, it is either missing its op or success property

		res.MarkFailed()

		if operationPropFound != nil {
			res.AddError(errors.New("Secure Builder API Authorize operation is invalid.  It is missing the operation authorization property."))
		}
		if successPropFound != nil {
			res.AddError(errors.New("Secure Builder API Authorize operation is invalid.  It is missing the authorization success property."))
		} else if successProp.Type() != "bool" {
			res.AddError(errors.New("Secure Builder API Authorize operation is invalid.  The authorization success property is not a bool."))
		}

		res.AddError(errors.New("Secure Builder API could not execute authorized Operation."))
	} else {

		operationProp.Set(ado.targetOp)

		authResult := ado.authOp.Exec(props)
		<-authResult.Finished()

		res.Merge(authResult)

		if !res.Success() {
			res.MarkFailed()
			res.AddError(errors.New("Operation authorization failed to execute."))
		} else {
			if successProp.Get().(bool) {
				// The Auth op returned a TRUE success value, so run the target Exec

				execResult := ado.targetOp.Exec(props)
				res.Merge(execResult)

				return res.Result()
			} else {
				// The Auth op returned a FALSE success value
				res.MarkFailed()
				res.AddError(errors.New("Authorization failed.  You are not permitted to execute the requested operation: " + op.authorized.Id()))
			}
		}

	}

	res.MarkFinished()
	return res.Result()
}