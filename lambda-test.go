package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Username string
}

func handler(e Event) (string, error) {
	return fmt.Sprintf("<h1> %s from 3mbda </h1>", e.Username), nil
}

func lambdaMain() {
	lambda.Start(handler)
}
