package rapi

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/scryinfo/apply/dots/auth/stub"
	"github.com/scryinfo/apply/dots/auth2"
	"github.com/scryinfo/apply/services/contract"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/gindot"
	"github.com/scryinfo/dot/dots/grpc/gserver"
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
	Bc         *contract.Bc       `dot:""`
	GinRouter  *gindot.Router `dot:""`
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

			_, err = c.Bc.Apply().Apply(c.Bc.TransactOpts(), common.HexToAddress(res.Addr), req.Finger) //todo sign
		}
	}
	return res, err
}

func (c *rapiServerImp) Sign(ctx context.Context, req *SignReq) (res *SignRes, err error) {
	res = &SignRes{}

	if len(req.Finger) < 0 || len(req.Addr) < 0 {
		err = errors.New("finger or addr is empty")
	} else {
		_, err = c.Bc.Apply().Sign(c.Bc.TransactOpts(), common.HexToAddress(req.Addr), req.Finger)
		if err == nil {

		}
	}

	return res, err
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

func (c *rapiServerImp) AfterAllStart(l dot.Line)  {
	c.GinRouter.Router().Any("/rapi.Rapi/apply", func(i *gin.Context) {

	})
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
				RelyLives: map[string]dot.LiveId{"ServerNobl": gserver.ServerNoblTypeId, "Bc": contract.BcTypeId, "GinRouter":gindot.RouterTypeId},
			},
		},
	}

	lives := []*dot.TypeLives{gserver.ServerNoblTypeLive(), tl}
	lives = append(lives, contract.BcTypeLives()...)

	return lives
}
