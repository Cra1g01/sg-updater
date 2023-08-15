package sgconnection

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (sg *SecurityGroup) Describe(svc *ec2.EC2) *ec2.SecurityGroup {
	res, err := svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
		GroupIds: []*string{
			aws.String(sg.SgId),
		},
	})
	if err != nil {
		log.Fatalf("Failed to describe security group, %v", err)
	}
    var sgs []*ec2.SecurityGroup = res.SecurityGroups
    return sgs[0]
}

func (sg *SecurityGroup) GetIngressRules(svc *ec2.EC2) []IpRule {
    res := sg.Describe(svc)
    rules := []IpRule{}
    for _, ipPermission := range res.IpPermissions {
        ipRange := ipPermission.IpRanges[0]
        fromPort := ipPermission.FromPort
        if fromPort == nil {
            fromPort = aws.Int64(-1)
        }
        toPort := ipPermission.ToPort
        if toPort == nil {
            toPort = aws.Int64(-1)
        }
        ip := IpRule{
            Ip: *ipRange.CidrIp,
            Description: *ipRange.Description,
            Protocol: *ipPermission.IpProtocol,
            FromPort: *fromPort,
            ToPort: *toPort,
        }
        rules = append(rules, ip)
    }
    return rules
}
