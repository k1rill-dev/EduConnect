package server

import (
	"EduConnect/pkg/config"
	"EduConnect/pkg/controllers"
	"EduConnect/pkg/jwt"
	"EduConnect/pkg/logger"
	"EduConnect/pkg/middlewares"
	"EduConnect/pkg/mongodb"
	"EduConnect/pkg/redis"
	"EduConnect/pkg/repo"
	"EduConnect/pkg/s3"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {
	log                      logger.Logger
	cfg                      *config.Config
	echo                     *echo.Echo
	authController           *controllers.AuthController
	jobController            *controllers.JobController
	jobApplicationController *controllers.ApplicationController
	portfolioController      *controllers.PortfolioController
	mongoClient              *mongo.Client
	middleware               *middlewares.MiddlewareManager
  s3             *s3.S3Storage
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{
		log:  log,
		cfg:  cfg,
		echo: echo.New(),
		//middleware: middlewares.NewMiddlewareManager(log, cfg),
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

	userRepo := repo.NewMongoAccountRepo(s.log, s.cfg, s.mongoClient)
	jobRepo := repo.NewJobRepo(s.log, s.cfg, s.mongoClient)
	jobApplicationRepo := repo.NewJobApplicationRepo(s.log, s.cfg, s.mongoClient)
	jwtRepo := repo.NewJwtRepo(s.log, s.cfg, redisClient, mongo)
	portfolioRepo := repo.NewPortfolioRepositoryMongo(s.log, s.cfg, s.mongoClient)
	jwtManager := jwt.NewJwtManager(s.log, s.cfg, jwtRepo)
	validate := s.setupValidator()
	s.middleware = middlewares.NewMiddlewareManager(s.log, s.cfg, jwtManager)
	s.authController = controllers.NewAuthController(s.log, s.cfg, userRepo, validate, jwtManager)
	s.jobController = controllers.NewJobController(s.log, s.cfg, jobRepo, validate, jwtManager, userRepo)
	s.jobApplicationController = controllers.NewApplicationController(s.log, s.cfg, jobApplicationRepo, validate, jwtManager, userRepo)
	s.portfolioController = controllers.NewPortfolioController(s.log, s.cfg, portfolioRepo, validate, jwtManager, userRepo)

	s3Storage, err := s3.NewS3Storage(s.log, s.cfg, s.mongoClient)
	if err != nil {
		s.log.Fatalf("S3Storage failed to start: %v", err)
	}
	s.s3 = s3Storage

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

func (s *server) setupValidator() *validator.Validate {
	validate := validator.New()

	// validate.RegisterValidation("eth_addr", func(fl validator.FieldLevel) bool {
	// 	addr := fl.Field().String()
	// 	return common.IsHexAddress(addr)
	// })

	return validate
}
