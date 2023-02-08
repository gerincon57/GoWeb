package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handlers"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
	//"golang.org/x/tools/cmd/getgo/server"
	//"golang.org/x/tools/cmd/getgo/server"
)

type Router struct {
	db []domain.Ticket

	en *gin.Engine
}

func NewRouter(en *gin.Engine, storage *[]domain.Ticket) *Router {
	return &Router{en: en, db: *storage}
}

func (r *Router) SetRoutes() {
	r.MapRoutes()
}

func (r *Router) MapRoutes() {
	rp := tickets.NewRepository(r.db)
	sv := tickets.NewService(rp) //&repo
	hd := handlers.NewTicket(sv) //&sv

	r.en.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	t := r.en.Group("ticket")
	t.Use(middlewares.AuthMiddleware())
	t.GET("getByCountry/:dest", hd.GetTicketsByCountry())
	t.GET("getAverage/:dest", hd.AverageDestination())

}
