package sgconnection

type SecurityGroup struct {
	SgId   string
	Region string
}

type IpRule struct {
	Ip          string
	Description string
	Protocol    string
	FromPort    int32
	ToPort      int32
}
