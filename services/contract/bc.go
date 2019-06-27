package contract

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scryinfo/dot/dot"
	"math/big"
)

const (
	BcTypeId = "7a54d9d5-a808-4057-9bef-86359d7fd8ee"
)

type bcConfig struct {
	Url          string `json:"url"`
	ChainId      int64  `json:"chainId"`
	TokenAddress string `json:"tokenAddress"`
	ApplyAddress string `json:"applyAddress"`

	OwnerAddress string `json:"ownerAddress"`
	OwnerKey     string `json:"ownerKey"`
}
type Bc struct {
	conf      bcConfig
	ethclient *ethclient.Client
	apply     *Apply
	token     *Token
	signer    types.Signer
	pkey      *ecdsa.PrivateKey
}

func (c *Bc) TransactOpts() *bind.TransactOpts {
	re := &bind.TransactOpts{
		From: common.HexToAddress(c.conf.OwnerAddress),
		Signer: func(signer types.Signer, addresses common.Address, tx *types.Transaction) (transaction *types.Transaction, e error) {
			stx, err := types.SignTx(tx, c.signer, c.pkey)
			return stx, err
		},
	}
	return re
}

func (c *Bc) Token() *Token {
	return c.token
}

func (c *Bc) Apply() *Apply {
	return c.apply
}

func (c *Bc) EthClient() *ethclient.Client {
	return c.ethclient
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

	c.pkey, err = crypto.HexToECDSA(c.conf.OwnerKey)
	if err != nil {
		return err
	}
	c.signer = types.NewEIP155Signer(big.NewInt(c.conf.ChainId))

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
