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
	s.echo.POST("api/auth/sign-up", s.authController.SignUp)
	s.echo.POST("api/auth/sign-in", s.authController.SignIn)
	s.echo.POST("api/auth/sign-out", s.authController.SignOut)

	s.echo.POST("api/course/", s.courseController.CreateCourse)

	s.echo.Static("api/files/", "storage")

	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
