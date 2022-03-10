package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type inputData struct {
	ID *int `json:"id"`
}

func (i *inputData) Validate() error {
	return nil
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, in *inputData) error {
	deadline, _ := ctx.Deadline()
	childCtx, cancel := context.WithDeadline(ctx, deadline.Add(-100*time.Millisecond))
	defer cancel()

	if err := in.Validate(); err != nil {
		return fmt.Errorf("input validate failure:  %+v", err)
	}

	return run(childCtx)
}

func run(ctx context.Context) error {
	return nil
}
