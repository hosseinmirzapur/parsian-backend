package config

import (
	"context"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	smithyendpoints "github.com/aws/smithy-go/endpoints"
)

type resolverV2 struct{}

func (*resolverV2) ResolveEndpoint(ctx context.Context, params s3.EndpointParameters) (
	smithyendpoints.Endpoint, error,
) {
	endpoint, err := url.Parse("https://parsian.storage.iran.liara.space")
	return smithyendpoints.Endpoint{
		URI: *endpoint,
	}, err
}

func LoadAWSConfig() aws.Config {
	cfg, err := awsCfg.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

func GetClient() *s3.Client {
	return s3.NewFromConfig(LoadAWSConfig(), func(o *s3.Options) {
		o.EndpointResolverV2 = &resolverV2{}
	})
}
