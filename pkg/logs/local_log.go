package logs

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
)

func loggerFiberConfig(pathFile string) logger.Config {
	if pathFile == "" {
		pathFile = "./app_log.log"
	}
	// Define file to logs
	file, err := os.OpenFile(pathFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	// Set config for logger
	loggerConfig := logger.Config{
		Output: file, // add file to save output
	}

	return loggerConfig

}
