/*
container module is dependency container.
this dependency container is recommended to use to all of app.
*/
package container

import (
	"github.com/drakejin/fiber-aws-serverless/config"
	"github.com/drakejin/fiber-aws-serverless/db"
)

type Container struct {
	Config    *config.Config
	ServiceDB *db.Client
}

func New(cfg *config.Config) (*Container, error) {
	c := &Container{}
	c.Config = cfg
	serviceDB, err := CreateServiceDBClient(cfg)
	if err != nil {
		return nil, err
	}
	c.ServiceDB = serviceDB

	return c, nil
}

func CreateServiceDBClient(cfg *config.Config) (*db.Client, error) {
	client, err := db.New(cfg)
	if err != nil {
		return nil, err
	}
	return client, err
}
