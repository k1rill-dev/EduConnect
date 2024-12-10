package server

import (
	"EduConnect/internal/repo"
	"EduConnect/pkg/config"
	"EduConnect/pkg/jwt"
	"EduConnect/pkg/logger"
	"EduConnect/pkg/middlewares"
	"EduConnect/pkg/mongodb"
	"EduConnect/pkg/redis"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {
	log  logger.Logger
	cfg  *config.Config
	echo *echo.Echo
	// authController *controllers.AuthController
	mongoClient *mongo.Client
	middleware  *middlewares.MiddlewareManager
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{
		log:  log,
		cfg:  cfg,
		echo: echo.New(),
		// middleware: *middlewares.NewMiddlewareManager(log, cfg),
	}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	mongo, err := mongodb.NewMongoDbConn(ctx, s.cfg.Mongo)
	if err != nil {
		return err
	}
	s.mongoClient = mongo
	s.initMongoDBCollections(ctx)

	redisConnector := redis.NewRedisConnector(s.log, s.cfg.Redis)
	redisClient, err := redisConnector.NewRedisConn(ctx)
	if err != nil {
		return err
	}

	// accountRepo := repo.NewMongoAccountRepo(s.log, s.cfg, mongo)
	// nonceRepo := repo.NewNonceRedisRepo(s.log, redisClient)
	jwtRepo := repo.NewJwtRepo(s.log, s.cfg, redisClient, mongo)
	jwtManager := jwt.NewJwtManager(s.log, s.cfg, jwtRepo)
	// nonceManager := nonce.NewNonceManager(s.log, s.cfg.Nonce, nonceRepo)
	// validate := s.setupValidator()
	s.middleware = middlewares.NewMiddlewareManager(s.log, s.cfg, jwtManager)
	// s.authController = controllers.NewAuthController(s.log, s.cfg, accountRepo, nonceManager, validate, jwtManager)

	go func() {
		if err := s.runHttpServer(); err != nil {
			s.log.Error("(HttpServer) err: %v", err)
			cancel()
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancelShutdownCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShutdownCtx()

	s.log.Infof("Server shutdown...")

	if err := s.echo.Shutdown(shutdownCtx); err != nil {
		s.log.Infof("Shutdown server with error: %v", err)
		return err
	}

	s.log.Infof("Server shutdown succesfuly")

	return nil
}

// func (s *server) setupValidator() *validator.Validate {
// 	validate := validator.New()

// 	validate.RegisterValidation("eth_addr", func(fl validator.FieldLevel) bool {
// 		addr := fl.Field().String()
// 		return common.IsHexAddress(addr)
// 	})

// 	en := en.New()
// 	uni := ut.New(en, en)

// 	trans, _ := uni.GetTranslator("en")
// 	en_translations.RegisterDefaultTranslations(validate, trans)
// 	validate.RegisterTranslation("eth_addr", trans, func(ut ut.Translator) error {
// 		return ut.Add("eth_addr", "{0} must be a valid Ethereum address", true)
// 	}, func(ut ut.Translator, fe validator.FieldError) string {
// 		t, _ := ut.T("eth_addr", fe.Field())
// 		return t
// 	}) // Убрать потом все равно не сильно помогает

// 	return validate
// }
