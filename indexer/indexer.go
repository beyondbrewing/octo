package indexer

import (
	"fmt"
	"sync"

	ers "github.com/beyondbrewing/octo/errors"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

type Config struct {
	//configs
	ChainParams *chaincfg.Params
	MaxPeers    uint8
	Peers       []string

	mu      sync.Mutex
	tipHash *chainhash.Hash

	synced bool
}

type Option func(*Config)

func DefaultConfig() *Config {
	return &Config{
		ChainParams: &chaincfg.MainNetParams,
		MaxPeers:    10,
	}
}

func (c *Config) validate() error {
	if c.ChainParams == nil {
		return fmt.Errorf("%w: chain params must not be nil", ers.ErrIndexerUnknownNetwork)
	}
	if c.MaxPeers <= 0 {
		return fmt.Errorf("indexer: max peers must be positive, got %d", c.MaxPeers)
	}
	return nil
}

func New(opts ...Option) (*Config, error) {
	cfg := DefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	// log := cfg.Logger
	// if log == nil {
	// 	log = logger.Default()
	// }
	// log = log.With("component", "indexer")

	// // client, err := btcclient.NewBtcClient(
	// // 	btcclient.WithChainParams(cfg.ChainParams),
	// // 	btcclient.WithMaxPeers(cfg.MaxPeers),
	// // 	btcclient.WithLogger(log),
	// // 	btcclient.WithListeners(cfg.Listeners),
	// // )
	// if err != nil {
	// 	return nil, fmt.Errorf("indexer: failed to create btc client: %w", err)
	// }

	// return &Indexer{
	// 	cfg:     cfg,
	// 	client:  client,
	// 	logger:  log,
	// 	tipHash: cfg.ChainParams.GenesisHash, // read from database needed to be added here
	// }, nil

	return &Config{}, nil
}

func WithLogger() Option {
	return func(c *Config) {}
}

func WithNetwork(name string) Option {
	return func(c *Config) {
		c.ChainParams = resolveChainParams(name)
	}
}

func WithEnodePeers(addrs ...string) Option {
	return func(c *Config) { c.Peers = addrs }
}
func WithMaxPeers(n int) Option {
	return func(c *Config) { c.MaxPeers = uint8(n) }

}

func resolveChainParams(name string) *chaincfg.Params {
	switch name {
	case "mainnet":
		return &chaincfg.MainNetParams
	case "signet":
		return &chaincfg.SigNetParams
	case "testnet3":
		return &chaincfg.TestNet3Params
	case "regtest":
		return &chaincfg.RegressionNetParams
	case "simnet":
		return &chaincfg.SimNetParams
	default:
		return &chaincfg.MainNetParams
	}
}
