package sgconnection

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (sg *SecurityGroup) RevokeIngress(client *ec2.Client, ipRule IpRule) {
	input := &ec2.RevokeSecurityGroupIngressInput{
		GroupId: aws.String(sg.SgId),
		IpPermissions: []types.IpPermission{
			{
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

	result, err := client.RevokeSecurityGroupIngress(context.TODO(), input)
	if err != nil {
        log.Fatalln(err)
	}

	fmt.Println(result)
}
