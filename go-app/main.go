package main
import (
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)
var (
	numbers []int
	lock    sync.Mutex
)
func getSign(n int) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	}
	return 0
}
func addNumber(c *gin.Context) {
	log.Info("POST /add called")
	type Input struct {
		Number int `json:"number"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		log.WithError(err).Error("Invalid JSON input")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please send a number."})
		return
	}
	num := input.Number
	log.WithField("number", num).Info("Received number")
	lock.Lock()
	defer lock.Unlock()
	log.WithField("list", numbers).Debug("Current list before processing")
	if len(numbers) == 0 || getSign(numbers[0]) == getSign(num) {
		numbers = append(numbers, num)
		log.WithField("list", numbers).Info("Added number to list")
		c.JSON(http.StatusOK, gin.H{"list": numbers})
		return
	}
	toRemove := -num
	log.WithField("toRemove", toRemove).Warn("Opposite sign detected, applying FIFO logic")
	newList := []int{}
	for _, n := range numbers {
		if toRemove == 0 {
			newList = append(newList, n)
			continue
		}
		if n <= toRemove {
			toRemove -= n
			log.WithField("removed", n).Warn("Removed")
		} else {
			log.WithFields(log.Fields{
				"original": n,
				"reduceBy": toRemove,
			}).Info("Reduced value")
			newList = append(newList, n-toRemove)
			toRemove = 0
		}
	}
	numbers = newList
	log.WithField("list", numbers).Info("Final updated list")
	c.JSON(http.StatusOK, gin.H{"list": numbers})
}
func main() {
	// Set log level to Debug to see more details
	log.SetLevel(log.DebugLevel)
	r := gin.Default()
	r.POST("/add", addNumber)
	r.GET("/list", func(c *gin.Context) {
		log.Info("GET /list called")
		lock.Lock()
		defer lock.Unlock()
		log.WithField("list", numbers).Info("Returning list")
		c.JSON(http.StatusOK, gin.H{"list": numbers})
	})
	log.Info("Server starting at http://localhost:8080")
	r.Run(":8080")
}
