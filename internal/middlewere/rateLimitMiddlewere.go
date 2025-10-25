package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type clientData struct {
	Requests int
	Reset    time.Time
}

var (
	clients = make(map[string]*clientData)
	lock    sync.Mutex
	limit   = 3                // 3 requests
	window  = 10 * time.Second // per 10 seconds
)

//Only Effective for Smalll scale

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Fetch the client Ip
		ip := c.ClientIP()

		lock.Lock()
		data, exists := clients[ip]

		if !exists || time.Now().After(data.Reset) {
			data = &clientData{
				Requests: 0,
				Reset:    time.Now().Add(window),
			}
			clients[ip] = data
		}

		if data.Requests >= limit {
			lock.Unlock()
			c.JSON(http.StatusTooManyRequests, gin.H{"erroe": "Too many request"})
			c.Abort()
			return
		}

		data.Requests++
		lock.Unlock()

		c.Next()

	}
}
