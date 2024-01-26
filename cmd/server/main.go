package main

import (
	"fmt"

	transportHttp "github.com/Adi-ty/leetlearner-api/internal/transport/http"
)

func Run() error {
	fmt.Println("Starting up server")

    httpHandler := transportHttp.NewHandler()
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
    fmt.Println("Leet learner API")
    if err := Run(); err != nil {
		fmt.Println(err)
	}
}
