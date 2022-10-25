package configs

import (
	"fmt"
	"os"
)

var (
	PORT = fmt.Sprintf(":%v", os.Getenv("PORT"))
)

const (
	API_VERSION = "/api/v1"
)
