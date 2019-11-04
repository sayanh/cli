package cli

import (
	"github.com/kyma-project/cli/pkg/step"
)

//Options defines available options for the command
type Options struct {
	Verbose bool
	step.Factory
	KubeconfigPath string
}

//NewOptions creates options with default values
func NewOptions() *Options {
	return &Options{}
}
