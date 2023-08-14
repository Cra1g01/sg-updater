package config

type AccountConfig struct {
	AccountId    string
	AccessKey    string
	AccessSecret string
}

type SecurityGroup struct {
	SgId   string
	Region string
}

type SecurityGroupsConfig struct {
	SecurityGroups []SecurityGroup
}

var Account AccountConfig = AccountConfig{
	AccountId:    "AccountID",
	AccessKey:    "AccountAccessKey",
	AccessSecret: "AccountAccessSecret",
}

var SecurityGroups SecurityGroupsConfig = SecurityGroupsConfig{
	SecurityGroups: []SecurityGroup{
		{
			SgId:   "SgID",
			Region: "Region",
		},
	},
}
