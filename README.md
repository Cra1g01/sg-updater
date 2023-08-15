# Security Group Updater

sg-updater is an application that runs on local machines (or virtual ones) and
updates AWS security groups to include the IP address of the machine running the application.
If a rule in a security group that matches the description has a different IP address, a new
rule is added with the same protocol and description, and the matched rule is removed.

This project serves to replace an existing system. I will generally be considering this project as
completed as of version 1.0.0, however, feel free to create an issue or even a PR if you find any
critical bugs. I may get around to adding features/improving things in the future.

Usage
-----

To use sg-updater, first [build](#build) the application with [your config](#config).

Once built, place the executable in your PATH and run:
```sh
sg-updater
```

Build
-----

To build this project:
- [Install][go-install] Go
- Clone this repository
```sh
https://github.com/Cra1g01/sg-updater.git
```
- Modify the [config](#config)
- Build the project
```sh
go build -o sg-updater cmd/sg-updater/main.go
```

Config
------

There are a few steps you need to take before building and using the application.

### AWS

You should create an [IAM account][create-iam] with only the permissions required, for example:

A user called sg-updater with the following policy:
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": [
                "ec2:RevokeSecurityGroupIngress",
                "ec2:AuthorizeSecurityGroupIngress",
                "ec2:DescribeSecurityGroups",
                "ec2:CreateTags"
            ],
            "Resource": "*"
        }
    ]
}
```

Then [create access keys][create-access] for the user.

### Code

`pkg/config/config.go` is the config file. Add your AWS access key details and security groups there.

Account config example:
```go
// pkg/config/config.go
var Account AccountConfig = AccountConfig{
	AccessKey:    "AKASDHJKASD908ASDASD",
	AccessSecret: "dsDADklJKka/asdSAD897SDASdhajks+Hasd23AS",
}
```

Security group config example:
```go
// pkg/config/config.go
var SecurityGroups SecurityGroupsConfig = SecurityGroupsConfig{
	SecurityGroups: []sgconnection.SecurityGroup{
		{
			SgId:   "sg-0cd123b1231252154",
			Region: "eu-west-2",
		},
    {
			SgId:   "sg-0ca941a5432341435",
			Region: "eu-west-2",
		},
	},
}
```


[go-install]: https://go.dev/doc/install
[create-iam]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html
[create-access]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html
