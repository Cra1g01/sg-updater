package sgconnection

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (sg *SecurityGroup) AuthorizeIngress(client *ec2.Client, ipRule IpRule) {
	input := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(sg.SgId),
		IpPermissions: []types.IpPermission{
			{
				// FromPort:   aws.Int64(22),
				// ToPort: aws.Int64(22),
				IpProtocol: aws.String(ipRule.Protocol),
				IpRanges: []types.IpRange{
					{
						CidrIp:      aws.String(ipRule.Ip),
						Description: aws.String(ipRule.Description),
					},
				},
			},
		},
	}

	result, err := client.AuthorizeSecurityGroupIngress(context.TODO(), input)
	if err != nil {
        log.Fatalln(err)
	}

	fmt.Println(result)
}
