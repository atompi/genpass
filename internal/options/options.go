package options

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Version string = "v1.0.0"

type LogOptions struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}

type CoreOptions struct {
	Threads int        `yaml:"threads"`
	Log     LogOptions `yaml:"log"`
}

type InputOptions struct {
	Type            string `yaml:"type"`
	Path            string `yaml:"path"`
	AccessKeyId     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	RegionId        string `yaml:"region_id"`
}

type OutputOptions struct {
	Type string `yaml:"type"`
	Path string `yaml:"path"`
}

type ResourceOptions struct {
	Type string `yaml:"type"`
}

type Options struct {
	Core     CoreOptions     `yaml:"core"`
	Input    InputOptions    `yaml:"input"`
	Output   OutputOptions   `yaml:"output"`
	Resource ResourceOptions `yaml:"resource"`
}

func NewOptions() (opts Options) {
	optsSource := viper.AllSettings()
	err := createOptions(optsSource, &opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "create options failed:", err)
		os.Exit(1)
	}
	return
}
