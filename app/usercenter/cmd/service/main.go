package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"go.opentelemetry.io/otel/attribute"
	"kratos-admin/pkg/tracex"
	"os"
	"strings"

	"kratos-admin/app/usercenter/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = ""
	// Version is the version of the compiled software.
	Version string = ""
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, r registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(r),
	)
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	Name = bc.Service.Name
	Version = bc.Service.Version
	addr := strings.SplitN(bc.Server.Grpc.Addr, ":", 2)
	if len(addr) == 2 {
		id += ":" + addr[1]
	}
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	if err := tracex.SetTracerProvider(bc.Trace.Endpoint, func() attribute.KeyValue {
		return semconv.ServiceNameKey.String(Name)
	}, func() attribute.KeyValue {
		return semconv.ServiceVersionKey.String(Version)
	}, func() attribute.KeyValue {
		return semconv.ServiceInstanceIDKey.String(id)
	}); err != nil {
		log.Error(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Registry, bc.Data, bc.JwtAuth, logger)
	if err != nil {
		panic(err)
	}
	if cleanup != nil {
		defer cleanup()
	}
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
