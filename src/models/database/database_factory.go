package database

import "errors"

type DatabaseFactory struct{}

func (f *DatabaseFactory) CreateDatabase(dbType string) (Database, error) {
	switch dbType {
	case "mysql":
		return &MySQL{}, nil
	default:
		return nil, errors.New("database type not supported")
	}
}
