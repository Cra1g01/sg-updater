package sgconnection

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (sg *SecurityGroup) AuthorizeIngress(svc *ec2.EC2, ipRule IpRule) {
	input := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(sg.SgId),
		IpPermissions: []*ec2.IpPermission{
			{
				// FromPort:   aws.Int64(22),
				// ToPort: aws.Int64(22),
				IpProtocol: aws.String(ipRule.Protocol),
				IpRanges: []*ec2.IpRange{
					{
						CidrIp:      aws.String(ipRule.Ip),
						Description: aws.String(ipRule.Description),
					},
				},
			},
		},
	}

	result, err := svc.AuthorizeSecurityGroupIngress(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Fatalln(aerr.Error())
			}
		} else {
			log.Fatalln(aerr.Error())
		}
		return
	}

	fmt.Println(result)
}
