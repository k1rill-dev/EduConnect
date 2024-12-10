package server

import (
	_ "EduConnect/docs"

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
	applications := s.echo.Group("/api/applications", s.middleware.AuthMiddleware)
	applications.POST("", s.jobApplicationController.CreateApplication)
	applications.PUT("/:applicationId/status", s.jobApplicationController.UpdateApplicationStatus)
	applications.DELETE("/:applicationId", s.jobApplicationController.DeleteApplication)

	// s.echo.POST("api/auth/sign-in", s.authController.SignInWithWallet)
	// s.echo.POST("api/auth/verify-signature", s.authController.VerifySignature)
	// s.echo.POST("api/auth/refresh-tokens", s.authController.RefreshTokens)

	// authGroup := s.echo.Group("api/auth", s.middleware.AuthMiddleware)
	// authGroup.POST("/sign-out", s.authController.SignOut)
	// authGroup.POST("/change-role", s.authController.ChangeRole)

	// s.echo.POST("api/auth/verify-token", s.authController.VerifyToken)

	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
