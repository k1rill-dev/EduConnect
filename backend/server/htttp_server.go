package server

import (
	_ "EduConnect/docs"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *server) runHttpServer() error {
	s.echo.Use(s.middleware.CORS())

	s.mapRoutes()

	return s.echo.Start(s.cfg.Http.Port)
}

func (s *server) mapRoutes() {
	s.echo.POST("api/auth/sign-up", s.authController.SignUp)
	s.echo.POST("api/auth/sign-in", s.authController.SignIn)
	s.echo.POST("api/auth/sign-out", s.authController.SignOut)

	s3Group := s.echo.Group("api/s3/")
	s3Group.Use(middleware.Logger())
	s3Group.Use(middleware.Recover())

	s3Group.POST("/upload", s.s3.UploadFile)
	s3Group.GET("/files/:id", s.s3.DownloadFile)
	s3Group.GET("/link/:id", s.s3.GetFileLink)
	s3Group.DELETE("/files/:id", s.s3.DeleteFile)
	// s.echo.POST("api/auth/sign-in", s.authController.SignInWithWallet)
	// s.echo.POST("api/auth/verify-signature", s.authController.VerifySignature)
	// s.echo.POST("api/auth/refresh-tokens", s.authController.RefreshTokens)

	// authGroup := s.echo.Group("api/auth", s.middleware.AuthMiddleware)
	// authGroup.POST("/sign-out", s.authController.SignOut)
	// authGroup.POST("/change-role", s.authController.ChangeRole)

	// s.echo.POST("api/auth/verify-token", s.authController.VerifyToken)

	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
