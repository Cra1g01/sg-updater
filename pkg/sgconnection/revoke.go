package sgconnection

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (sg *SecurityGroup) RevokeIngress(svc *ec2.EC2, ipRule IpRule) {
	input := &ec2.RevokeSecurityGroupIngressInput{
		GroupId: aws.String(sg.SgId),
		IpPermissions: []*ec2.IpPermission{
			{
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

	result, err := svc.RevokeSecurityGroupIngress(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Fatalln(err.Error())
			}
		} else {
			log.Fatalln(err.Error())
		}
		return
	}

	fmt.Println(result)
}
