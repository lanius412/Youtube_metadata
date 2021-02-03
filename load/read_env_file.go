package load

import (
	"github.com/joho/godotenv"

	"fmt"
	"os"
)

//read developerKey from .env file
func Read_env_file() (developerKey string) {
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Println("Can't Read .env File Correctly")
	}
	developerKey = os.Getenv("DEVELOPERKEY")
	//log.Println(developerKey)
	return developerKey
}