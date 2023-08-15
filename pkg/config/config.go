package config

import "github.com/Cra1g01/sg-updater/pkg/sgconnection"

type AccountConfig struct {
	AccountId    string
	AccessKey    string
	AccessSecret string
}

type SecurityGroupsConfig struct {
	SecurityGroups []sgconnection.SecurityGroup
}

func (sgs *SecurityGroupsConfig) GroupByRegion() map[string][]*sgconnection.SecurityGroup {
	m := make(map[string][]*sgconnection.SecurityGroup)
	for _, sg := range sgs.SecurityGroups {
		sg := sg
		_, exists := m[sg.Region]
		if exists {
			m[sg.Region] = append(m[sg.Region], &sg)
		} else {
			m[sg.Region] = []*sgconnection.SecurityGroup{
				&sg,
			}
		}
	}
	return m
}

var Account AccountConfig = AccountConfig{
	AccountId:    "AccountID",
	AccessKey:    "AccountAccessKey",
	AccessSecret: "AccountAccessSecret",
}

var IpURL string = "https://checkip.amazonaws.com"

var SecurityGroups SecurityGroupsConfig = SecurityGroupsConfig{
	SecurityGroups: []sgconnection.SecurityGroup{
		{
			SgId:   "SgID",
			Region: "Region",
		},
	},
}
