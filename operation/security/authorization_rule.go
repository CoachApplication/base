package security

/**
 * Interfaces
 */

// AuthorizationRule is a single authorizer, that can authorize an operation
type AuthorizationRule interface {
	// AuthorizeOperation asks to authorize an operation
	AuthorizeOperation(op api_operation.Operation) AuthorizationRuleResult
}

// AuthorizationRuleResult is the result of an authorization request, with meta and UI data about the authorization 
type AuthorizationRuleResult interface {
	// Id of rule that triggered the result
	RuleId() string
	// Message about result that can be shown in UI
	Message() string 
	// Does the result explicitly allow the operation
	Allow() bool   
	// Does the result explicitly deny the opration  
	Deny() bool      
}

/**
 * Utilite RuleResult implementations
 */

// An empty Rule result for cases where no explicity rule has applied
type NoRuleResult string

// NewNoRuleResult creates a RuleResult for situations where no rule has applied, with an optional message
func NewNoRuleResult(message string) *NoRuleResult {
	return &(NoRuleResult(message))
}

// RuleResult converts this to a RuleResult inteface (for clarity and validation)
func (nrr *NoRuleResult) RuleResult() RuleResult {
	return RuleResult(nrr)
}

// Id of rule that triggered the result
func (nrr *NoRuleResult) RuleId() string {
	return "noruleapplied"
}
// Message about result that can be shown in UI
func (nrr *NoRuleResult) Message() string {
	if message := string(nrr); message == "" {
		return "No rule applied"
	} else {
		return message		
	}
}
// Does the result explicitly allow the operation
func (nrr *NoRuleResult) Allow() bool {
	return false
}   
// Does the result explicitly deny the opration  
func (nrr *NoRuleResult) Deny() bool {
	return false
}  

/**
 * Utility Rule implementations
 */

// AuthorizationOrderedListRule is an AuthorizationRule that actually runs a set of rules
type AuthorizationOrderedListRule []AuthorizationRule

// AuthorizationRule converts this struct to an AuthorizationRule (to validate that it implements the interface)
func (aolr *AuthorizeOperation) AuthorizationRule() AuthorizationRule {
	return AuthorizationRule(aolr)
}

// AuthorizeOperation authorizes an operation with the list of rules
func (aolr *AuthorizeOperation) AuthorizeOperation(op api_operation.Operation) RuleResult {
	for _, rule := range []AuthorizationRule(aolr) {
		res := rule.AuthorizeOperation(op)

		if res.Allow() || res.Deny() {
			return res
		}
	}
	return NewNoRuleResult("No rule applied").RuleResult()
}
