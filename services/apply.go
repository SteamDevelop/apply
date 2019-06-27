package services

import (
	"github.com/scryinfo/apply/dots/auth2"
	"github.com/scryinfo/apply/services/rapi"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/grpc/gserver"
)

//AddNewers add newers of api
func Add(l dot.Line) error {
	//add the controller
	lives := gserver.HttpNoblTypeLives()
	lives = append(lives, auth2.Auth2TypeLives()...)
	lives = append(lives, rapi.RapiServerTypeLives()...)
	err := l.PreAdd(lives...)

	return err
}
