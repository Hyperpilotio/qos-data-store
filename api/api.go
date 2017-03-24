package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Server store the stats / data of every deployment
type Server struct {
	Config *viper.Viper
	// Maps name of metrics to its value
	Counter map[string]int64
	mutex   sync.Mutex
}

// NewServer return an instance of Server struct.
func NewServer(config *viper.Viper) *Server {
	return &Server{
		Config:  config,
		Counter: make(map[string]int64),
	}
}

// StartServer start a web server
func (s *Server) StartServer() error {
	if s.Config.GetString("filesPath") == "" {
		return errors.New("filesPath is not specified in the configuration file")
	}

	if err := os.Mkdir(s.Config.GetString("filesPath"), 0755); err != nil {
		if !os.IsExist(err) {
			return errors.New("Unable to create filesPath directory: " + err.Error())
		}
	}

	//gin.SetMode("release")
	router := gin.New()

	// Global middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	daemonsGroup := router.Group("/v1/metrics")
	{
		daemonsGroup.GET("", s.getMetricsList)
		daemonsGroup.GET("/:metric", s.getMetric)
		daemonsGroup.POST("/:metric", s.setMetric)

	}

	return router.Run(":" + s.Config.GetString("port"))
}

func (s *Server) getMetricsList(c *gin.Context) {
	if len(s.Counter) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": true,
			"data":  "no metrics at all",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  s.Counter,
	})
}

func (s *Server) getMetric(c *gin.Context) {
	metric := c.Param("metric")
	if _, ok := s.Counter[metric]; !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": true,
			"data":  fmt.Sprintf("metric %s not found", metric),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  s.Counter[metric],
	})
}

func (s *Server) setMetric(c *gin.Context) {
	var metric string

	metric = c.Param("metric")
	if c.Query("val") == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": true,
			"data":  fmt.Sprintf("value for metric %s not exist", metric),
		})
		return
	}

	v, err := strconv.ParseInt(c.Query("val"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": true,
			"data":  fmt.Sprintf("value %s for metric %s is not valid", c.Query("val"), metric),
		})
		return
	}

	s.Counter[metric] = v

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  s.Counter[metric],
	})
}
