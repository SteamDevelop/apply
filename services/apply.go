package services

import (
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/grpc/gserver"
)

//AddNewers add newers of api
func Add(l dot.Line) error {
	//add the controller
	typeLives := gserver.GinNoblTypeLives()
	err := l.PreAdd(typeLives...)
	return err
}
