package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
)

func LoadAWSConfig() aws.Config {

	cfg, err := awsCfg.LoadDefaultConfig(context.TODO(), func(options *awsCfg.LoadOptions) error {
		options.Region = "us-east-1"
		options.EC2IMDSEndpoint = "https://parsian.storage.iran.liara.space"

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
