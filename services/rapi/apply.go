package rapi

import (
	"context"
	"github.com/scryinfo/apply/dots/auth2"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/grpc/gserver"
	"github.com/scryinfo/dp/dots/auth/stub"
)

const (
	RapiServerTypeId = "RapiServerId"
)

type rapiConfig struct {
	Bcurl   string
	ChainId string
}

type rapiServerImp struct {
	ServerNobl gserver.ServerNobl `dot:""`
	Auth2      *auth2.Auth2       `dot:""`
}

func (c *rapiServerImp) Apply(ctx context.Context, req *ApplyReq) (res *ApplyRes, err error) {

	res = &ApplyRes{}
	{
		ares, err2 := c.Auth2.Client().GenerateAddress(ctx, &stub.AddressParameter{
			Password: req.Pass,
		})
		if err2 != nil {
			err = err2
		} else {
			res.Addr = ares.Address

			err = nil
		}
	}
	return res, err
}

func (c *rapiServerImp) Sign(context.Context, *SignReq) (*SignRes, error) {
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

func (c *rapiServerImp) Start(ignore bool) error {
	RegisterRapiServer(c.ServerNobl.Server(), c)
	return nil
}

//RapiServerTypeLives make all type lives
func RapiServerTypeLives() []*dot.TypeLives {
	tl := &dot.TypeLives{
		Meta: dot.Metadata{TypeId: RapiServerTypeId, NewDoter: func(conf interface{}) (dot.Dot, error) {
			return &rapiServerImp{}, nil
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
