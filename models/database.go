package models

// SQLDataBase struct
type SQLDataBase struct {
	Server          string `toml:"Server"`
	Port            int    `toml:"Port"`
	Database        string `toml:"Database"`
	ApplicationName string `toml:"ApplicationName"`
	PoolSize        int    `toml:"PoolSize"`
	UserID          string
	Password        string
}
