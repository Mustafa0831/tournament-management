package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	router := gin.New()
	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Настройка прокси для сервисов
	teamServiceURL, _ := url.Parse("http://team-service:8081")
	divisionServiceURL, _ := url.Parse("http://division-service:8082")
	matchServiceURL, _ := url.Parse("http://match-service:8083")
	playoffServiceURL, _ := url.Parse("http://playoff-service:8084")
	resultServiceURL, _ := url.Parse("http://result-generation-service:8085")

	router.Any("/teams/*proxyPath", reverseProxy(teamServiceURL))
	router.Any("/teams", reverseProxy(teamServiceURL))

	router.Any("/divisions/*proxyPath", reverseProxy(divisionServiceURL))
	router.Any("/divisions", reverseProxy(divisionServiceURL))

	router.Any("/matches/*proxyPath", reverseProxy(matchServiceURL))
	router.Any("/matches", reverseProxy(matchServiceURL))

	router.Any("/playoffs/*proxyPath", reverseProxy(playoffServiceURL))
	router.Any("/playoffs", reverseProxy(playoffServiceURL))

	router.Any("/generate-results/*proxyPath", reverseProxy(resultServiceURL))
	router.Any("/generate-results", reverseProxy(resultServiceURL))

	log.Println("API Gateway is running on port 8080")
	router.Run(":8080")
}

func reverseProxy(target *url.URL) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(target)
		proxy.Director = func(req *http.Request) {
			log.Printf("Proxying request to %s%s", target.Host, c.Request.URL.Path)
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = c.Request.URL.Path
			req.URL.RawQuery = c.Request.URL.RawQuery
			req.Header = c.Request.Header
			req.Host = target.Host
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
