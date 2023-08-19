package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
)

func LoadAWSConfig() aws.Config {

	cfg, err := awsCfg.LoadDefaultConfig(context.TODO())
	cfg.Region = "us-east-1"

	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
