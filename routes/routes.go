package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // Get a list of available events
	server.GET("/events/:id", getEvent) // Get a list of available single event

	authenticated := server.Group("/")                               // use a GROUP to group certain routes together that will use the middleware to auth
	authenticated.Use(middlewares.Authenticate)                      // use the middleware to authenticate BEFORE user uses the route
	authenticated.POST("/events", createEvent)                       // creating a bookable new event (Auth Required)
	authenticated.PUT("/events/:id", updateEvent)                    // updating an event (Auth required, only by creator)
	authenticated.DELETE("/events/:id", deleteEvent)                 // deleting an event (Auth required, only by creator)
	authenticated.POST("/events/:id/register", registerForEvent)     // registering a user for an event (Auth Required)
	authenticated.DELETE("/events/:id/register", cancelRegistration) // deleting a registration for a user (Auth Required)

	server.POST("/signup", signup) // create a new user
	server.POST("/login", login)   // Authenticate user login (AUTH TOKEN (JWT))
}
