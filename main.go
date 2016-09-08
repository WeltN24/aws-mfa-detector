package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

type Result struct {
	Data interface{} `json:"data"`
}

func main() {
	var (
		awsRegion     = flag.String("aws.region", "eu-central-1", "AWS region")
		userToExclude = flag.String("exclude", "", "Users to exclude from detection (, separated)")
		err           error
	)

	flag.Parse()

	iamCli := iam.New(session.New(), aws.NewConfig().WithRegion(*awsRegion))

	users, err := getUsers(iamCli)

	if err != nil {
		log.Fatal(err)
	}

	usersToExclude := strings.Split(*userToExclude, ",")
	users = excludeUsers(users, usersToExclude)

	usersWithoutMfa := make([]map[string]string, 0, len(users))
	for _, userName := range users {
		has, err := hasMfa(iamCli, userName)

		if err != nil {
			log.Fatal(err)
		}

		if !has {
			usersWithoutMfa = append(usersWithoutMfa, map[string]string{
				"{#USERNAME}": userName,
			})
		}
	}

	err = json.NewEncoder(os.Stdout).Encode(Result{Data: usersWithoutMfa})

	if err != nil {
		log.Fatal(err)
	}
}

func getUsers(iamCli *iam.IAM) ([]string, error) {
	resp, err := iamCli.ListUsers(&iam.ListUsersInput{})

	if err != nil {
		return []string{}, fmt.Errorf("getting users: %v", err)
	}

	users := make([]string, 0, len(resp.Users))
	for _, user := range resp.Users {
		users = append(users, *user.UserName)
	}

	return users, nil
}

func hasMfa(iamCli *iam.IAM, userName string) (bool, error) {
	resp, err := iamCli.ListMFADevices(&iam.ListMFADevicesInput{
		UserName: aws.String(userName),
	})

	if err != nil {
		return false, fmt.Errorf("getting mfa devices for user %v: %v", userName, err)
	}

	return len(resp.MFADevices) > 0, nil
}

func excludeUsers(originalList []string, usersToExclude []string) []string {
	usersToCheck := make([]string, 0, len(originalList))
	for _, user := range originalList {
		if !stringInSlice(user, usersToExclude) {
			usersToCheck = append(usersToCheck, user)
		}
	}
	return usersToCheck
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
