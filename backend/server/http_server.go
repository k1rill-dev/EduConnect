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
	s.echo.Use(s.middleware.CORS())
	s.echo.POST("api/auth/sign-up", s.authController.SignUp)
	s.echo.POST("api/auth/sign-in", s.authController.SignIn)
	s.echo.POST("api/auth/sign-out", s.authController.SignOut)

	updateUser := s.echo.Group("api/auth/update-user", s.middleware.AuthMiddleware)

	updateUser.POST("", s.authController.UpdateUser)

	signOut := s.echo.Group("api/auth/sign-out", s.middleware.AuthMiddleware)
	signOut.POST("", s.authController.SignOut)

	// courseGroup := s.echo.Group("api/course", s.middleware.AuthMiddleware)
	s.echo.POST("api/course/", s.courseController.CreateCourse, s.middleware.AuthMiddleware)
	s.echo.POST("api/course/submit-assignment", s.courseController.SubmitAssignment, s.middleware.AuthMiddleware)

	s.echo.GET("api/course/", s.courseController.GetCourses)

	s.echo.Static("api/photo/", "storage/photo")
	s.echo.Static("api/file/", "storage/assignments")

	applications := s.echo.Group("/api/applications", s.middleware.AuthMiddleware)
	applications.POST("", s.jobApplicationController.CreateApplication)
	applications.PUT("/:applicationId/status", s.jobApplicationController.UpdateApplicationStatus)
	applications.DELETE("/:applicationId", s.jobApplicationController.DeleteApplication)

	s.echo.GET("/api/jobs/:jobId", s.jobController.GetJobById)
	s.echo.GET("/api/jobs/search", s.jobController.SearchJobs)
	s.echo.POST("/api/jobs/filter", s.jobController.GetJobsByFilters)

	jobsWithAuth := s.echo.Group("/api/jobs", s.middleware.AuthMiddleware)
	jobsWithAuth.POST("", s.jobController.CreateJob)
	jobsWithAuth.PUT("/:jobId", s.jobController.UpdateJob)

	portfolios := s.echo.Group("/api/portfolios")
	portfolios.POST("", s.portfolioController.CreatePortfolio, s.middleware.AuthMiddleware)
	portfolios.POST("/:portfolioId/items", s.portfolioController.AddPortfolioItems, s.middleware.AuthMiddleware)
	portfolios.GET("/student/:studentId", s.portfolioController.GetPortfolioByStudent)

	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
