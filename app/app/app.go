package main

import (
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/gindot"
	"github.com/scryinfo/dot/dots/line"
	"github.com/scryinfo/scryg/sutils/sfile"
	"github.com/scryinfo/scryg/sutils/ssignal"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

func main() {
	l, err := line.BuildAndStart(Add)
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

//AddNewers add newers of api
func Add(l dot.Line) error {
	//add the controller
	lives := gindot.TypeLiveGinDot()
	err := l.PreAdd(lives)

	l.ToDotEventer().ReSetLiveEvents(dot.LiveId(gindot.EngineLiveId), &dot.LiveEvents{
		AfterCreate: func(live *dot.Live, l dot.Line) {
			if engine, ok := live.Dot.(*gindot.Engine); ok {
				ex := "."
				{
					et, err := os.Executable()
					if err == nil {
						ex = filepath.Dir(et)
						if !sfile.ExistFile(filepath.Join(ex, "dist-apply")) {
							temp := filepath.Dir(ex)
							temp = filepath.Dir(temp)
							if sfile.ExistFile(filepath.Join(temp, "app/dist")) {
								ex = temp
							}
						}
					}
				}
				engine.GinEngine().StaticFile("/apply", filepath.Join(ex,"dist-apply/index.html"))
				engine.GinEngine().Static("/static", filepath.Join(ex,"dist-apply/static"))
				engine.GinEngine().StaticFile("/sign", filepath.Join(ex,"dist-sign/index.html"))
			}
		},
	})
	return err
}
