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
	signOut := s.echo.Group("api/auth/sign-out", s.middleware.AuthMiddleware)
	updateUser := s.echo.Group("api/auth/update-user", s.middleware.AuthMiddleware)

	updateUser.POST("", s.authController.UpdateUser)
	s.echo.POST("api/auth/sign-up", s.authController.SignUp)
	s.echo.POST("api/auth/sign-in", s.authController.SignIn)
	signOut.POST("", s.authController.SignOut)

	s.echo.GET("/api/jobs/:jobId", s.jobController.GetJobById)
	s.echo.GET("/api/jobs/search", s.jobController.SearchJobs)
	s.echo.POST("/api/jobs/filter", s.jobController.GetJobsByFilters)

	jobsWithAuth := s.echo.Group("/api/jobs", s.middleware.AuthMiddleware)
	jobsWithAuth.POST("", s.jobController.CreateJob)
	jobsWithAuth.PUT("/:jobId", s.jobController.UpdateJob)

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
	applications := s.echo.Group("/api/applications", s.middleware.AuthMiddleware)
	applications.POST("", s.jobApplicationController.CreateApplication)
	applications.PUT("/:applicationId/status", s.jobApplicationController.UpdateApplicationStatus)
	applications.DELETE("/:applicationId", s.jobApplicationController.DeleteApplication)

	portfolios := s.echo.Group("/api/portfolios")
	portfolios.POST("", s.portfolioController.CreatePortfolio, s.middleware.AuthMiddleware)
	portfolios.POST("/:portfolioId/items", s.portfolioController.AddPortfolioItems, s.middleware.AuthMiddleware)
	portfolios.GET("/student/:studentId", s.portfolioController.GetPortfolioByStudent)

	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
