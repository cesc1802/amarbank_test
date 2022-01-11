package httpserver

import (
	"amarbank/pkg/httpserver/middleware"
	"amarbank/pkg/logger"
	"amarbank/pkg/validation"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net"
	"net/http"
	"reflect"
	"strings"
)

type GinOption struct {
	name string
	port string
	host string
	mode string
}

type GinService struct {
	isRunning bool
	engine    *gin.Engine
	*http.Server
	handlers []func(engine *gin.Engine)
	logger   logger.Logger
	*GinOption
}

func NewGinService(log logger.Logger, mode, host, port string) *GinService {
	return &GinService{
		isRunning: false,
		handlers:  []func(*gin.Engine){},
		logger:    log,
		GinOption: &GinOption{
			name: "GIN-HTTP-SERVICE",
			mode: mode,
			host: host,
			port: port,
		},
	}
}

func jsonTagNameFunc(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

func (s *GinService) Configure() error {
	if s.isRunning {
		return nil
	}
	if s.mode == "" {
		s.mode = "debug"
	}

	gin.SetMode(s.mode)
	s.engine = gin.New()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(jsonTagNameFunc)
		v.RegisterValidation("space-between", validation.StringContainWhiteSpace)
	}

	s.engine.RedirectTrailingSlash = true
	s.engine.RedirectFixedPath = true

	// Recovery
	s.engine.Use(middleware.Recovery())

	s.isRunning = true
	return nil
}

func (s *GinService) Start() error {
	if err := s.Configure(); err != nil {
		return err
	}

	// Setup handlers
	for _, hdl := range s.handlers {
		hdl(s.engine)
	}

	s.Server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.host, s.port),
		Handler: s.engine,
	}
	s.logger.Info("Listening and serving HTTP on ", s.host, s.port)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {

		s.logger.Info("Listening error: %v", err)
		return err
	}

	err = s.Serve(lis)

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *GinService) Stop() error {
	if s.Server != nil {
		return s.Server.Shutdown(context.Background())
	}

	return nil
}

func (s *GinService) AddHandler(hdl func(engine *gin.Engine)) {
	if s.isRunning {
		return
	}
	s.handlers = append(s.handlers, hdl)
}
