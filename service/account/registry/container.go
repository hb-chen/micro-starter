package registry

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/micro/v3/service/config"
	"go.uber.org/dig"

	"github.com/hb-chen/micro-starter/service/account/domain/repository/persistence/gorm"
	"github.com/hb-chen/micro-starter/service/account/domain/repository/persistence/memory"
	"github.com/hb-chen/micro-starter/service/account/domain/repository/persistence/xorm"
	"github.com/hb-chen/micro-starter/service/account/domain/service"
	"github.com/hb-chen/micro-starter/service/account/usecase"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()

	buildUserUseCase(c)

	return c, nil
}

func buildUserUseCase(c *dig.Container) {
	conf, _ := config.Get("persistence")
	persistence := conf.String("")

	// ORM选择，gorm、xorm...
	switch persistence {
	case "xorm":
		// DB初始化
		xorm.InitDB()
		c.Provide(xorm.NewUserRepository)
	case "gorm":
		// DB初始化
		gorm.InitDB()
		c.Provide(gorm.NewUserRepository)
	default:
		// 默认memory作为mock
		c.Provide(memory.NewUserRepository)
	}

	c.Provide(service.NewUserService)
	c.Provide(usecase.NewUserUseCase)
}
