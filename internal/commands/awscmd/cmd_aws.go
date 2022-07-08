// Package awscmd contains all commands for the aws section of the fusion cli
//
// Commands
//
// AWS commands can be executed with
// 	fusion new aws
package awscmd

type Globals struct {
	Region string `enum:"us-west-1,us-west-2,us-east-1,us-east-2,af-south-1,ap-east-1,ap-southeast-3,ap-south-1,ap-northeast-3,ap-northeast-2,ap-southeast-1,ap-southeast-2,ap-northeast-1,ca-central-1,eu-central-1,eu-west-1,eu-west-2,eu-south-1,eu-west-3,eu-north-1,me-south-1,sa-east-1,us-gov-east-1,us-gov-west-1" default:"us-east-1"`
}

type AWS struct {
	Stack struct {
		ServerlessWebsite NewAWSServerlessWebsiteCmd `cmd:"" help:"Create new AWS static website w/ serverless backend"`
	} `cmd:"" help:"Create new stacks of resources"`

	Resource struct {
		SecurityGroup NewSecurityGroupCmd `cmd:"" help:"Create new AWS Security Group"`
		S3Bucket      NewS3BucketCmd      `cmd:"" help:"Create new AWS S3 bucket" name:"s3_bucket"`
		Cloudfront    NewCloudfrontCmd    `cmd:"" help:"Create new AWS cloudfront distribution"`
		Route53       NewRoute53ZoneCmd   `cmd:"" help:"Create new AWS route53 hosted zone"`
		Lambda        NewLambdaCommand    `cmd:"" help:"Create new AWS lambda function"`
		IamPolicy     NewIamPolicyCmd     `cmd:"" help:"Create new AWS IAM policy"`
		ApiGateway    NewAPIGatewayCmd    `cmd:"" help:"Create new AWS API gateway v2"`
		VpcStack      NewVpcStackCmd      `cmd:"" help:"Create new VPC with public & private subnets, igws, & NAT gateways"`
	} `cmd:"" help:"Create new single resource"`
}
