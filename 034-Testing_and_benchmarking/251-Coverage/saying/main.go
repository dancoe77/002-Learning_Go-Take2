package saying

import "fmt"

func Greet(s string) string {
	return fmt.Sprint("Welcome my dear ", s)
}

//go test -cover

//go test -coverprofile c.out

//go tool cover -html=c.out

//go tool cover -help