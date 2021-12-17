package types

type ServerInfo struct {
	Environment string  `mapstructure:"env"`
	Name        *string `mapstructure:"name"`
	IsPrimary   bool    `mapstructure:"is-primary"`
}
