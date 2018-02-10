package database

import "github.com/lordking/blaster/common"

type Database interface {
	NewConfig() interface{}
	ValidateBefore() error
	Connect() error
	GetConnection() interface{}
	Close() error
}

func Configure(configKey string, db Database) error {

	config := db.NewConfig()
	if err := common.ReadConfigKey(configKey, config); err != nil {
		return err
	}

	if err := db.ValidateBefore(); err != nil {
		return err
	}

	return nil
}
