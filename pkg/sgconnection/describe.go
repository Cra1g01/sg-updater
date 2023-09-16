package sgconnection

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (sg *SecurityGroup) Describe(client *ec2.Client) *types.SecurityGroup {
	res, err := client.DescribeSecurityGroups(context.TODO(), &ec2.DescribeSecurityGroupsInput{
		GroupIds: []string{
			sg.SgId,
		},
	})
	if err != nil {
		log.Fatalf("Failed to describe security group, %v", err)
	}
	return &res.SecurityGroups[0]
}

func (sg *SecurityGroup) GetIngressRules(client *ec2.Client) []IpRule {
	res := sg.Describe(client)
	rules := []IpRule{}
	for _, ipPermission := range res.IpPermissions {
		ipRange := ipPermission.IpRanges[0]
		fromPort := ipPermission.FromPort
		if fromPort == nil {
			port := int32(-1)
			fromPort = &port
		}
		toPort := ipPermission.ToPort
		if toPort == nil {
			port := int32(-1)
			toPort = &port
		}
		ip := IpRule{
			Ip:          *ipRange.CidrIp,
			Description: *ipRange.Description,
			Protocol:    *ipPermission.IpProtocol,
			FromPort:    *fromPort,
			ToPort:      *toPort,
		}
		rules = append(rules, ip)
	}
	return rules
}
