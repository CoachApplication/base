package configprovider

import (
	api "github.com/james-nesbitt/coach-api"
	base "github.com/james-nesbitt/coach-base"
	base_errors "github.com/james-nesbitt/coach-base/errors"
	base_config "github.com/james-nesbitt/coach-base/operation/configuration"
)

type GetOperation struct {
	base_config.GetOperationBase
	provider Provider
}

func NewGetOperation(provider Provider) *GetOperation{
	return &GetOperation{
		provider: provider,
	}
}

func (gon GetOperation) Validate(props api.Properties) api.Result {
	return base.MakeSuccessfulResult()
}

func (gon GetOperation) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	key := ""
	if keyProp, err := props.Get(base_config.PROPERTY_KEY_KEY); err == nil {
		key = keyProp.Get().(string)
	} else {
		res.AddError(err)
		res.AddError(base_errors.RequiredPropertyWasEmptyError{Key: base_config.PROPERTY_KEY_KEY})
	}

	scopedConfig := base_config.NewStandardScopedConfig()
	for _, scope := range gon.provider.Scopes() {
		if config, err := gon.provider.Get(key, scope); err == nil {
			scopedConfig.Set(scope, config)
		} else {
			res.AddError(err)
		}
	}

	scopedConfigProp := base_config.ScopedConfigProperty{}
	scopedConfigProp.Set(scopedConfig)
	res.AddProperty(api.Property(&scopedConfigProp))

	return res.Result()
}
