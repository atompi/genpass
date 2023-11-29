package execute

import (
	"github.com/atompi/metabot/internal/options"
	"go.uber.org/zap"
)

func Execute(opts options.Options) {
	zap.L().Sugar().Infof("options: ", opts)
}
