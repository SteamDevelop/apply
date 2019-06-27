package main

import (
	"github.com/scryinfo/apply/services"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/line"
	"github.com/scryinfo/scryg/sutils/ssignal"
	"go.uber.org/zap"
	"os"
)

func main() {
	l, err := line.BuildAndStart(services.Add)
	if err != nil {
		dot.Logger().Errorln("", zap.Error(err))
		return
	}
	defer line.StopAndDestroy(l, true)
	dot.Logger().Infoln("dot ok")

	ssignal.WaitCtrlC(func(s os.Signal) bool {
		return false
	})
	dot.Logger().Infoln("dot will stop")
}
