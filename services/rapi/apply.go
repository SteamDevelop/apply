package rapi

import (
"context"
"github.com/scryinfo/dot/dot"
"github.com/scryinfo/dot/dots/grpc/gserver"
)

const (
	RapiServerTypeId = "RapiServerId"
)

type RapiServerImp struct {
	ServerNobl gserver.ServerNobl `dot:""`
}

func (serv *RapiServerImp) Apply(context.Context, *ApplyReq) (*ApplyRes, error) {
	res := &ApplyRes{}
	return res, nil
}

func (serv *RapiServerImp) Sign(context.Context, *SignReq) (*SignRes, error) {
	res := &SignRes{}
	return res, nil
}

//func newHiServer(conf interface{}) (dot.Dot, error) {
//	var err error = nil
//	var bs []byte = nil
//	if bt, ok := conf.([]byte); ok {
//		bs = bt
//	} else {
//		return nil, dot.SError.Parameter
//	}
//	dconf := &config{}
//	err = dot.UnMarshalConfig(bs, dconf)
//	if err != nil {
//		return nil, err
//	}
//
//	d := &RapiServer{
//		conf: *dconf,
//	}
//
//	return d, err
//}



func (serv *RapiServerImp) Start(ignore bool) error {
	RegisterRapiServer(serv.ServerNobl.Server(), serv)
	return nil
}

//RapiServerTypeLives make all type lives
func RapiServerTypeLives() []*dot.TypeLives {
	tl := &dot.TypeLives{
		Meta: dot.Metadata{TypeId: RapiServerTypeId, NewDoter: func(conf interface{}) (dot.Dot, error) {
			return &RapiServerImp{}, nil
		}},
		Lives: []dot.Live{
			dot.Live{
				LiveId:    RapiServerTypeId,
				RelyLives: map[string]dot.LiveId{"ServerNobl": gserver.ServerNoblTypeId},
			},
		},
	}

	lives := []*dot.TypeLives{gserver.ServerNoblTypeLive(), tl}

	return lives
}
