package twiliserver

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Gorynychdo/twilisip/internal/app/model"
    "github.com/levigross/grequests"
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

func (s *server) serve() error {
    return s.engine.Run(s.config.BindAddr)
}

func (s *server) configureRouter() {
    s.engine.POST("/register", s.register)
}

func (s *server) register(c *gin.Context) {
    var json model.Phone
    if err := c.ShouldBindJSON(&json); err != nil {
        s.error(c, http.StatusBadRequest, err)
        return
    }

    ro := &grequests.RequestOptions{
        Auth: []string{s.config.TwilioAccountSID, s.config.TwilioAuthToken},
        Data: map[string]string{"PhoneNumber": json.Number},
    }

    res, err := grequests.Post(s.config.TwilioCallerIDURL, ro)
    if err != nil {
        s.error(c, http.StatusBadRequest, err)
        return
    }

    c.JSON(200, gin.H{"status": "OK"})
}

func (s *server) error(c *gin.Context, status int, err error) {
    c.JSON(status, gin.H{
        "status": status,
        "error":  err.Error(),
    })
}
