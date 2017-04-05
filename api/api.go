package api

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Server store the stats / data of every deployment
type Server struct {
	Config *viper.Viper

	Apps  map[string]*AppMetrics
	mutex sync.Mutex
}

type AppMetrics struct {
	Name       string             `json:"name"`
	QoSMetrics map[string]float64 `json:"metrics"`
}

func (app *AppMetrics) setMetric(name string, value float64) {
	app.QoSMetrics[name] = value
}

// NewServer return an instance of Server struct.
func NewServer(config *viper.Viper) *Server {
	return &Server{
		Config: config,
		Apps:   make(map[string]*AppMetrics),
	}
}

// StartServer start a web server
func (s *Server) StartServer() error {
	router := gin.New()

	// Global middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	daemonsGroup := router.Group("/v1/apps/")
	{
		daemonsGroup.GET("/metrics", s.getAllMetrics)
		daemonsGroup.POST("/:app/metrics/:metric", s.setMetric)
	}

	return router.Run(":" + s.Config.GetString("port"))
}

func (s *Server) getAllMetrics(c *gin.Context) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  s.Apps,
	})
}

func (s *Server) setMetric(c *gin.Context) {
	app := c.Param("app")
	metric := c.Param("metric")

	var json struct {
		Value float64 `json:"value" binding:"required"`
	}

	if err := c.Bind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Sprintf("Unable to parse request: %s", err.Error()),
		})
		return
	}

	s.mutex.Lock()
	s.mutex.Unlock()

	_, ok := s.Apps[app]
	if !ok {
		s.Apps[app] = &AppMetrics{
			Name:       app,
			QoSMetrics: make(map[string]float64),
		}
	}

	s.Apps[app].setMetric(metric, json.Value)

	c.JSON(http.StatusOK, gin.H{
		"error": false,
	})
}
