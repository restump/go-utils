package awsutils

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	secrets "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecretValue(
	c aws.Config, 
	timeout time.Duration,
	secretId, versionId, versionStage *string,
	options ...func(*secrets.Options),
) (*secrets.GetSecretValueOutput, error) {
	client := secrets.NewFromConfig(c, options...)
	input := &secrets.GetSecretValueInput{
		SecretId:     secretId,
		VersionId:    versionId,
		VersionStage: versionStage,
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := client.GetSecretValue(ctx, input)
	if err != nil {
		return nil, err
	}
	return res, nil

}
