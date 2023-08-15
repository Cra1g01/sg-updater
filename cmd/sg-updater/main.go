package main

import (
	// "fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/Cra1g01/sg-updater/pkg/config"
	"github.com/Cra1g01/sg-updater/pkg/sgconnection"
	"github.com/Cra1g01/sg-updater/pkg/sys"
)

func main() {
	username := sys.GetUsername()
	ipAddr := sys.GetIpAddr()
	update(username, ipAddr)
}

func update(username string, ipAddr string) {
	sgs := config.SecurityGroups
	groups := sgs.GroupByRegion()
	for region, val := range groups {
		sess := session.Must(session.NewSession(&aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials(config.Account.AccessKey, config.Account.AccessSecret, ""),
		}))
		svc := ec2.New(sess)
		for _, sg := range val {
			rules := sg.GetIngressRules(svc)
			for _, rule := range rules {
				if rule.Description == username && rule.Ip != ipAddr {
					newRule := sgconnection.IpRule{
						Ip:          ipAddr + "/32",
						Protocol:    rule.Protocol,
						Description: rule.Description,
					}
					sg.AuthorizeIngress(svc, newRule)
					sg.RevokeIngress(svc, rule)
					break
				}
			}
		}
	}
}

