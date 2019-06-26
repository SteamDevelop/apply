package auth2

import (
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/grpc/conns"
	"github.com/scryinfo/dp/dots/auth/stub"
)

const (
	Auth2TypeId = "b573aab1-4bea-46d5-b9d3-6c8da5ca690e"
)

type Auth2 struct {
	Conn   *conns.ConnName `dot:""`
	client stub.KeyServiceClient
}

func (c *Auth2) Client() stub.KeyServiceClient {
	return c.client
}

func (c *Auth2) AfterAllInject(l dot.Line) {
	c.client = stub.NewKeyServiceClient(c.Conn.Conn())
}

//Auth2TypeLives make all type lives
func Auth2TypeLives() []*dot.TypeLives {
	tl := &dot.TypeLives{
		Meta: dot.Metadata{TypeId: Auth2TypeId, NewDoter: func(conf interface{}) (dot.Dot, error) {
			return &Auth2{}, nil
		}},
		Lives: []dot.Live{
			dot.Live{
				LiveId:    Auth2TypeId,
				RelyLives: map[string]dot.LiveId{"Conn": conns.ConnNameTypeId},
			},
		},
	}

	lives := make([]*dot.TypeLives, 0, 3)
	lives = append(lives, conns.ConnNameTypeLives()...)
	lives = append(lives, tl)

	return lives
}
