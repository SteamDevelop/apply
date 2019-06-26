package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scryinfo/dot/dot"
)

const (
	BcTypeId = "7a54d9d5-a808-4057-9bef-86359d7fd8ee"
)

type bcConfig struct {
	Url          string `json:"url"`
	ChainId      string `json:"chainId"`
	TokenAddress string `json:"tokenAddress"`
	ApplyAddress string `json:"applyAddress"`
}
type Bc struct {
	conf      bcConfig
	ethclient *ethclient.Client
	apply     *Apply
	token     *Token
}

func (c *Bc) Token() *Token {
	return c.token
}

func (c *Bc) Apply() *Apply {
	return c.apply
}

func (c *Bc) Destroy(ignore bool) error {
	c.apply = nil
	c.token = nil
	if c.ethclient != nil {
		t := c.ethclient
		c.ethclient = nil

		t.Close()
	}
	return nil
}

func (c *Bc) Create(l dot.Line) error {
	var err error
	c.ethclient, err = ethclient.Dial(c.conf.Url)
	if err != nil {
		return err
	}

	c.token, err = NewToken(common.HexToAddress(c.conf.TokenAddress), c.ethclient)
	if err != nil {
		return err
	}

	c.apply, err = NewApply(common.HexToAddress(c.conf.ApplyAddress), c.ethclient)
	if err != nil {
		return err
	}

	return err
}

func newBc(conf interface{}) (dot.Dot, error) {
	var err error = nil
	var bs []byte = nil
	if bt, ok := conf.([]byte); ok {
		bs = bt
	} else {
		return nil, dot.SError.Parameter
	}
	dconf := &bcConfig{}
	err = dot.UnMarshalConfig(bs, dconf)
	if err != nil {
		return nil, err
	}

	d := &Bc{
		conf: *dconf,
	}

	return d, err
}

//BcTypeLives make all type lives
func BcTypeLives() []*dot.TypeLives {
	tl := &dot.TypeLives{
		Meta: dot.Metadata{TypeId: BcTypeId, NewDoter: func(conf interface{}) (dot.Dot, error) {
			return newBc(conf)
		}},
	}

	lives := make([]*dot.TypeLives, 0, 3)
	lives = append(lives, tl)

	return lives
}
