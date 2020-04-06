package twiliserver

import (
    "github.com/gin-gonic/gin"
)

type server struct {
    config *Config
    engine *gin.Engine
}

func newServer(config *Config) *server {
    s := &server{
        config: config,
        engine: gin.Default(),
    }

    s.configureRouter()

    return s
}

func (s *server) configureRouter() {
    s.engine.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
}

func (s *server) serve() error {
    return s.engine.Run(s.config.BindAddr)
}
