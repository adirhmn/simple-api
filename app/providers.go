package app

import (
	"simpleapi/repository"
	"simpleapi/tools/dbpool"
	"simpleapi/usecase"
)

type ServiceContainers struct {
	UserUseCase usecase.UserUseCaseInt
	Close       func() error
}

func NewServiceContainers(cfg *Config, dbPools *dbpool.DBPools) (*ServiceContainers, error) {

	svCont := &ServiceContainers{}

	svCont.UserUseCase = usecase.NewUserUseCase(repository.NewUserRepo(dbPools))

	svCont.Close = func() error {
		return nil
	}

	return svCont, nil
}
