package registry

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/micro/v3/service/config"
	"go.uber.org/dig"

	"github.com/hb-chen/micro-starter/service/greeting/domain/repository/persistence/gorm"
	"github.com/hb-chen/micro-starter/service/greeting/domain/repository/persistence/memory"
	"github.com/hb-chen/micro-starter/service/greeting/domain/service"
	"github.com/hb-chen/micro-starter/service/greeting/usecase"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()

	err := buildGreetingUseCase(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func buildGreetingUseCase(c *dig.Container) error {
	conf, _ := config.Get("persistence")
	persistence := conf.String("")

	// ORM选择，gorm、xorm...
	switch persistence {
	case "gorm":
		// DB初始化
		gorm.InitDB()
		err := c.Provide(gorm.NewGreetingRepository)
		if err != nil {
			return err
		}
	default:
		// 默认memory作为mock
		err := c.Provide(memory.NewGreetingRepository)
		if err != nil {
			return err
		}
	}

	err := c.Provide(service.NewGreetingService)
	if err != nil {
		return err
	}
	err = c.Provide(usecase.NewGreetingUseCase)
	if err != nil {
		return err
	}

	return nil
}
