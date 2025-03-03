package config

import "biyard.co/common/base"

type AppConfig struct {
	base.Config

	Feepayer FeepayerConfig `yaml:"feepayer"`
	Owner    OwnerConfig    `yaml:"owner"`

	Contracts []string `yaml:"contracts"`

	ContractAddresses ContractConfig `yaml:"contractAddresses"`

	Klaytn   KlaytnConfig      `yaml:"klaytn"`
	Services FeeEndpointConfig `yaml:"services"`

	OAuth OAuthConfig `yaml:"oauth"`
	Hint  HintConfig  `yaml:"hint"`

	ABI string `yaml:"abi"`

	Endpoint      string              `yaml:"endpoint"`
	Discord       DiscordConfig       `yaml:"discord"`
	ICP           ICPConfig           `yaml:"icp"`
	Authorization AuthorizationConfig `yaml:"authorization"`
}

type AuthorizationConfig struct {
	Code string `yaml:"code"`

	Discord DiscordConfig `yaml:"discord"`
	ICP     ICPConfig     `yaml:"icp"`
}

type ICPConfig struct {
	Endpoint string `yaml:"endpoint"`
}

type DiscordConfig struct {
	BotToken string               `yaml:"botToken"`
	Channels DiscordChannelConfig `yaml:"channels"`
}

type DiscordChannelConfig struct {
	DAO string `yaml:"dao"`
}

type ContractConfig struct {
	AccountProfile   string `yaml:"accountProfile"`
	ClosedTokenClaim string `yaml:"closedTokenClaim"`
	Shop             string `yaml:"shop"`
	NFT              string `yaml:"nft"`
	EXP              string `yaml:"exp"`
}

type KlaytnConfig struct {
	Endpoint string `yaml:"endpoint"`
	ChainID  int    `yaml:"chainId"`
}

type FeeEndpointConfig struct {
	Endpoint string `yaml:"feeEndpoint"`
}

type FeepayerConfig struct {
	Address    string `yaml:"address"`
	PrivateKey string `yaml:"privateKey"`
}

type OwnerConfig struct {
	Address    string `yaml:"address"`
	PrivateKey string `yaml:"privateKey"`
}

type OAuthConfig struct {
	RedirectURI []string       `yaml:"redirectUri"`
	Kakao       ProviderConfig `yaml:"kakao"`
	Google      ProviderConfig `yaml:"google"`
}

type ProviderConfig struct {
	ClientID       string `yaml:"clientId"`
	CredentialFile string `yaml:"credentialFile"`
	Endpoint       string `yaml:"endpoint"`
}

type HintConfig struct {
	Secret  string        `yaml:"secret"`
	Account AccountConfig `yaml:"account"`
}

type AccountConfig struct {
	Address    string `yaml:"address"`
	PublicKey  string `yaml:"publicKey"`
	PrivateKey string `yaml:"privateKey"`
}

func DefaultAppConfig() AppConfig {
	conf := base.DefaultConfig()

	conf.Logging.Level = "debug"
	return AppConfig{
		Config: conf,
	}
}
