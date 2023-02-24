package registry

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hb-chen/micro-starter/service/greeting/repo/gorm"
	"github.com/micro/micro/v3/service/config"
	"go.uber.org/dig"

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
	persistence := conf.String("gorm")

	// ORM选择，gorm、xorm...
	switch persistence {
	case "xorm":
	case "gorm":
	default:
		// DB初始化
		c.Provide(gorm.NewDB)
		err := c.Provide(gorm.NewGreetingRepository)
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
