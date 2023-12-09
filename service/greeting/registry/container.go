package registry

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hb-chen/micro-starter/service/greeting/repo/gorm"
	"github.com/hb-chen/micro-starter/service/greeting/repo/memory"
	"go.uber.org/dig"
	"micro.dev/v4/service/config"

	"github.com/hb-chen/micro-starter/service/greeting/domain/usecase"
	"github.com/hb-chen/micro-starter/service/greeting/service"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()

	err := build(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func build(c *dig.Container) error {
	conf, _ := config.Get("persistence")
	persistence := conf.String("")

	// ORM选择，gorm、xorm...
	switch persistence {
	case "xorm":
	case "gorm":
		// DB初始化
		err := c.Provide(gorm.NewDB)
		if err != nil {
			return err
		}
		err = c.Provide(gorm.NewGreetingRepository)
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

	err := c.Provide(usecase.NewGreetingUsecase)
	if err != nil {
		return err
	}
	err = c.Provide(service.NewGreetingService)
	if err != nil {
		return err
	}

	return nil
}
