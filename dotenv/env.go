package dotenv

import (
	"bitbucket.org/mexisme/get-secrets/errors"

	"github.com/mexisme/multiconfig"
	"github.com/mexisme/multiconfig/env"
	dotenv "github.com/mexisme/multiconfig/pathed"
)

//TODO EnvAddConfig
func EnvAddConfig(envs *multiconfig.Map) func(string, string) error {
	return func(path, body string) error {
		envs.AddItem(dotenv.New().SetPath(path).SetBody(body))
		return nil
	}
}

//nolint:godox TODO EnvMerge
func EnvMerge(envs *multiconfig.Map) *env.Config {
	merged, err := envs.Merge()
	if err != nil {
		errors.PanicOnErrors(err)
	}

	return env.New().SetBodyMap(merged)
}
