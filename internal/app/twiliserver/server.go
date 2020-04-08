package twiliserver

import (
    "bytes"
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Gorynychdo/twilisip/internal/app/model"
    "github.com/levigross/grequests"
)

const (
    ctxKeyReqBody = "request_body"
    ctxKeyResult  = "result"
)

var (
    errWrongParameters = errors.New("wrong parameters")
)

type server struct {
    config *Config
    engine *gin.Engine
}

func newServer(config *Config) *server {
    s := &server{
        config: config,
        engine: gin.New(),
    }

    s.configureRouter()

    return s
}

func (s *server) serve() error {
    return s.engine.Run(s.config.BindAddr)
}

func (s *server) configureRouter() {
    s.engine.Use(s.logRequest())
    s.engine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
        return fmt.Sprintf("[%v] |%s %3d %s|%s %s %s| %s %s %s\n",
            param.TimeStamp.Format("2006/01/02 - 15:04:05"),
            param.StatusCodeColor(), param.StatusCode, param.ResetColor(),
            param.MethodColor(), param.Method, param.ResetColor(),
            param.Path,
            param.Keys[ctxKeyReqBody],
            param.Keys[ctxKeyResult],
        )
    }))
    s.engine.Use(gin.Recovery())

    s.engine.POST("/register", s.register)
}

func (s *server) logRequest() gin.HandlerFunc {
    return func(c *gin.Context) {
        var buf bytes.Buffer
        tee := io.TeeReader(c.Request.Body, &buf)

        body, err := ioutil.ReadAll(tee)
        if err == nil {
            c.Request.Body = ioutil.NopCloser(&buf)
        }

        c.Set(ctxKeyReqBody, string(body))
        c.Next()
    }
}

func (s *server) register(c *gin.Context) {
    var json model.Phone
    if err := c.ShouldBindJSON(&json); err != nil {
        c.Set(ctxKeyResult, err)
        s.error(c, http.StatusBadRequest, errWrongParameters)
        return
    }

    res, err := grequests.Post(s.config.TwilioCallerIDURL, &grequests.RequestOptions{
        Auth: []string{s.config.TwilioAccountSID, s.config.TwilioAuthToken},
        Data: map[string]string{"PhoneNumber": json.Number},
    })
    if err != nil {
        c.Set(ctxKeyResult, err)
        s.error(c, http.StatusBadRequest, errWrongParameters)
        return
    }

    defer res.Close()

    response := gin.H{"res": res.String()}
    c.Set(ctxKeyResult, mustJson(response))
    c.JSON(200, response)
}

func (s *server) error(c *gin.Context, status int, err error) {
    c.JSON(status, gin.H{
        "status": status,
        "error":  err.Error(),
    })
}
