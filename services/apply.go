package services

import (
	"github.com/scryinfo/apply/dots/auth2"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/grpc/gserver"
)

//AddNewers add newers of api
func Add(l dot.Line) error {
	//add the controller
	lives := gserver.GinNoblTypeLives()
	lives = append(lives, auth2.Auth2TypeLives()...)
	err := l.PreAdd(lives...)
	return err
}
