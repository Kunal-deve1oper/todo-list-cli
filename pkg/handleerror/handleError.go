package handleerror

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
