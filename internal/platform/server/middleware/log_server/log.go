package log_server

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Middleware is a gin.HandlerFunc that middleware for logging HTTP requests and responses.
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the incoming request
		log.Printf("Incoming request: %s %s", c.Request.Method, c.Request.URL)

		// Call the next middleware/handler in the chain
		c.Next()

		// Log the outgoing response
		log.Printf("Outgoing response: %d", c.Writer.Status())

		// Log any errors that occurred during the request
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				log.Printf("Error occurred: %v", err)
			}
		}
	}
}
