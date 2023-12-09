// Package profile is for specific profiles
// @todo this package is the definition of cruft and
// should be rewritten in a more elegant way
package profile

import (
	"path/filepath"

	"github.com/hb-chen/micro-starter/service/auth/noop"
	"github.com/urfave/cli/v2"
	microAuth "micro.dev/v4/service/auth"
	"micro.dev/v4/service/broker"
	memBroker "micro.dev/v4/service/broker/memory"
	"micro.dev/v4/service/config"
	storeConfig "micro.dev/v4/service/config/store"
	microEvents "micro.dev/v4/service/events"
	evStore "micro.dev/v4/service/events/store"
	memStream "micro.dev/v4/service/events/stream/memory"
	"micro.dev/v4/service/logger"
	"micro.dev/v4/service/model"
	"micro.dev/v4/service/model/sql"
	"micro.dev/v4/service/profile"
	"micro.dev/v4/service/registry"
	memRegistry "micro.dev/v4/service/registry/memory"
	microRuntime "micro.dev/v4/service/runtime"
	"micro.dev/v4/service/runtime/local"
	microStore "micro.dev/v4/service/store"
	"micro.dev/v4/service/store/file"
	mem "micro.dev/v4/service/store/memory"
	"micro.dev/v4/util/user"
)

func init() {
	if err := profile.Register("starter-local", Local); err != nil {
		logger.Fatalf("Error profile register: %v", err)
	}

	if err := profile.Register("starter-kubernetes", Kubernetes); err != nil {
		logger.Fatalf("Error profile register: %v", err)
	}

	if err := profile.Register("starter-test", Kubernetes); err != nil {
		logger.Fatalf("Error profile register: %v", err)
	}
}

// Local profile to run locally
var Local = &profile.Profile{
	Name: "starter-local",
	Setup: func(ctx *cli.Context) error {
		// catch all
		profile.SetupDefaults()

		microAuth.DefaultAuth = noop.NewAuth()
		microStore.DefaultStore = file.NewStore(file.WithDir(filepath.Join(user.Dir, "server", "store")))
		profile.SetupConfigSecretKey()
		config.DefaultConfig, _ = storeConfig.NewConfig(microStore.DefaultStore, "")
		profile.SetupJWT()

		// the registry service uses the memory registry, the other core services will use the default
		// rpc client and call the registry service
		if ctx.Args().Get(1) == "registry" {
			profile.SetupRegistry(memRegistry.NewRegistry())
		} else {
			// set the registry address
			registry.DefaultRegistry.Init(
				registry.Addrs("localhost:8000"),
			)

			profile.SetupRegistry(registry.DefaultRegistry)
		}

		// the broker service uses the memory broker, the other core services will use the default
		// rpc client and call the broker service
		if ctx.Args().Get(1) == "broker" {
			profile.SetupBroker(memBroker.NewBroker())
		} else {
			broker.DefaultBroker.Init(
				broker.Addrs("localhost:8003"),
			)
			profile.SetupBroker(broker.DefaultBroker)
		}

		// set the store in the model
		// TODO sql model
		model.DefaultModel = sql.NewModel()

		// use the local runtime, note: the local runtime is designed to run source code directly so
		// the runtime builder should NOT be set when using this implementation
		microRuntime.DefaultRuntime = local.NewRuntime()

		var err error
		microEvents.DefaultStream, err = memStream.NewStream()
		if err != nil {
			logger.Fatalf("Error configuring stream: %v", err)
		}
		microEvents.DefaultStore = evStore.NewStore(
			evStore.WithStore(microStore.DefaultStore),
		)

		microStore.DefaultBlobStore, err = file.NewBlobStore()
		if err != nil {
			logger.Fatalf("Error configuring file blob store: %v", err)
		}

		// Configure tracing with Jaeger (forced tracing):
		tracingServiceName := ctx.Args().Get(1)
		if len(tracingServiceName) == 0 {
			tracingServiceName = "Micro"
		}
		// openTracer, _, err := jaeger.New(
		// 	opentelemetry.WithServiceName(tracingServiceName),
		// 	opentelemetry.WithSamplingRate(1),
		// )
		// if err != nil {
		// 	logger.Fatalf("Error configuring opentracing: %v", err)
		// }
		// opentelemetry.DefaultOpenTracer = openTracer

		return nil
	},
}

// Kubernetes profile to run on kubernetes with zero deps. Designed for use with the micro helm chart
var Kubernetes = &profile.Profile{
	Name: "starter-kubernetes",
	Setup: func(ctx *cli.Context) (err error) {
		// catch all
		profile.SetupDefaults()

		microAuth.DefaultAuth = noop.NewAuth()

		microRuntime.DefaultRuntime = local.NewRuntime()

		microEvents.DefaultStream, err = memStream.NewStream()
		if err != nil {
			logger.Fatalf("Error configuring stream: %v", err)
		}

		microStore.DefaultStore = file.NewStore(file.WithDir("/store"))
		microStore.DefaultBlobStore, err = file.NewBlobStore(file.WithDir("/store/blob"))
		if err != nil {
			logger.Fatalf("Error configuring file blob store: %v", err)
		}

		// set the store in the model
		// TODO sql model
		model.DefaultModel = sql.NewModel()

		// the registry service uses the memory registry, the other core services will use the default
		// rpc client and call the registry service
		if ctx.Args().Get(1) == "registry" {
			profile.SetupRegistry(memRegistry.NewRegistry())
		} else {
			// set the registry address
			registry.DefaultRegistry.Init(
				registry.Addrs("micro-server.micro.svc.cluster.local:8000"),
			)

			profile.SetupRegistry(registry.DefaultRegistry)
		}

		// the broker service uses the memory broker, the other core services will use the default
		// rpc client and call the broker service
		if ctx.Args().Get(1) == "broker" {
			profile.SetupBroker(memBroker.NewBroker())
		} else {
			broker.DefaultBroker.Init(
				broker.Addrs("micro-server.micro.svc.cluster.local:8003"),
			)
			profile.SetupBroker(broker.DefaultBroker)
		}

		config.DefaultConfig, err = storeConfig.NewConfig(microStore.DefaultStore, "")
		if err != nil {
			logger.Fatalf("Error configuring config: %v", err)
		}
		profile.SetupConfigSecretKey()

		// Configure tracing with Jaeger:
		tracingServiceName := ctx.Args().Get(1)
		if len(tracingServiceName) == 0 {
			tracingServiceName = "Micro"
		}
		// openTracer, _, err := jaeger.New(
		// 	opentelemetry.WithServiceName(tracingServiceName),
		// 	opentelemetry.WithTraceReporterAddress("localhost:6831"),
		// )
		// if err != nil {
		// 	logger.Fatalf("Error configuring opentracing: %v", err)
		// }
		// opentelemetry.DefaultOpenTracer = openTracer

		return nil
	},
}

// Test profile is used for the go test suite
var Test = &profile.Profile{
	Name: "starter-test",
	Setup: func(ctx *cli.Context) error {
		// catch all
		profile.SetupDefaults()

		microAuth.DefaultAuth = noop.NewAuth()
		microStore.DefaultStore = mem.NewStore()
		microStore.DefaultBlobStore, _ = file.NewBlobStore()
		config.DefaultConfig, _ = storeConfig.NewConfig(microStore.DefaultStore, "")
		profile.SetupRegistry(memRegistry.NewRegistry())
		// set the store in the model
		// TODO sql model
		model.DefaultModel = sql.NewModel()
		return nil
	},
}
