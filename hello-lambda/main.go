// Copyright 2017 budougumi0617 All Rights Reserved.

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// HelloLambdaHandler uses lambdaContext, events.
// https://docs.aws.amazon.com/lambda/latest/dg/go-programming-model-handler-types.html
func HelloLambdaHandler(ctx context.Context, config events.ConfigEvent) (string, error) {
	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return "", errors.New("Cannot find LambdaContext")
	}

	// https://docs.aws.amazon.com/lambda/latest/dg/go-programming-model-logging.html
	// Use the log module
	log.Printf("log:InvokedFunctionArn = %s\n", lc.InvokedFunctionArn)
	// Use the fmt module
	fmt.Printf("fmt:InvokedFunctionArn = %s\n", lc.InvokedFunctionArn)

	// Print event
	log.Printf("config version is %v\n", config.Version)

	// https://docs.aws.amazon.com/lambda/latest/dg/go-programming-model-env-variables.html
	log.Printf("\"%s\" executes on \"%s\".\n", os.Getenv("NAME"), os.Getenv("AWS_EXECUTION_ENV"))

	return "function finished", nil
}

func main() {
	lambda.Start(HelloLambdaHandler)
}
