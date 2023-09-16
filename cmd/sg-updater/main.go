package main

import (
	"context"
	"log"
	// "fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"

	ac "github.com/Cra1g01/sg-updater/pkg/config"
	"github.com/Cra1g01/sg-updater/pkg/sgconnection"
	"github.com/Cra1g01/sg-updater/pkg/sys"
)

func main() {
	username := sys.GetUsername()
	ipAddr := sys.GetIpAddr()
	update(username, ipAddr)
}

func update(username string, ipAddr string) {
	sgs := ac.SecurityGroups
	groups := sgs.GroupByRegion()
	for region, val := range groups {
		creds := credentials.NewStaticCredentialsProvider(ac.Account.AccessKey, ac.Account.AccessSecret, "")
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(region),
			config.WithCredentialsProvider(creds),
		)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		client := ec2.NewFromConfig(cfg)
		for _, sg := range val {
			rules := sg.GetIngressRules(client)
			for _, rule := range rules {
				if rule.Description == username && rule.Ip != ipAddr {
					newRule := sgconnection.IpRule{
						Ip:          ipAddr + "/32",
						Protocol:    rule.Protocol,
						Description: rule.Description,
					}
					sg.AuthorizeIngress(client, newRule)
					sg.RevokeIngress(client, rule)
					break
				}
			}
		}
	}
}
