package services

import (
	"github.com/scryinfo/apply/dots/auth2"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/gindot"
	"github.com/scryinfo/dot/dots/grpc/gserver"
	"github.com/scryinfo/scryg/sutils/sfile"
	"os"
	"path/filepath"
)

//AddNewers add newers of api
func Add(l dot.Line) error {
	//add the controller
	lives := gserver.GinNoblTypeLives()
	lives = append(lives, auth2.Auth2TypeLives()...)
	err := l.PreAdd(lives...)

	l.ToDotEventer().ReSetLiveEvents(dot.LiveId(gindot.EngineLiveId), &dot.LiveEvents{
		AfterCreate: func(live *dot.Live, l dot.Line) {
			if engine, ok := live.Dot.(*gindot.Engine); ok {
				ex := "."
				{
					et, err := os.Executable()
					if err == nil {
						ex = filepath.Dir(et)
						if !sfile.ExistFile(filepath.Join(ex, "dist")) {
							temp := filepath.Dir(ex)
							temp = filepath.Dir(temp)
							if sfile.ExistFile(filepath.Join(temp, "app/dist")) {
								ex = temp
							}
						}
					}
				}
				engine.GinEngine().Static("/app", filepath.Join(ex, "app"))
			}
		},
	})
	return err
}
