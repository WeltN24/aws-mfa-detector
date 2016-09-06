package main

import (
	"flag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func main() {
	var (
		awsRegion     = flag.String("aws.region", "eu-central-1", "AWS region")
		userToExclude = falg.String("exclude", "", "Users to exclude (Usernames)")
		err           error
	)

	flag.Parse()

	iamCli := iam.New(session.New(), aws.NewConfig().WithRegion(*awsRegion))
}
