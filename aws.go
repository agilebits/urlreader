package urlreader

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

func getAWSConfig(region string, creds *credentials.Credentials) *aws.Config {
	conf := &aws.Config{}
	if creds == nil {
		// Grab the metadata URL
		metadataURL := os.Getenv("AWS_METADATA_URL")
		if metadataURL == "" {
			metadataURL = "http://169.254.169.254:80/latest"
		}

		creds = credentials.NewChainCredentials(
			[]credentials.Provider{
				&credentials.EnvProvider{},
				&credentials.SharedCredentialsProvider{Filename: "", Profile: ""},
				&ec2rolecreds.EC2RoleProvider{
					Client: ec2metadata.New(session.New(&aws.Config{
						Endpoint: aws.String(metadataURL),
					})),
				},
			})
	}

	conf.Credentials = creds
	if region == "" {
		region = os.Getenv("AWS_REGION")
	}

	if region != "" {
		conf.Region = aws.String(region)
	}

	return conf
}
