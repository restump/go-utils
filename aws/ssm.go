package awsutils

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func GetParameter(
	c aws.Config,
	timeout time.Duration,
	name *string,
	withDecryption *bool,
	options ...func(*ssm.Options),
) (*ssm.GetParameterOutput, error) {
	client := ssm.NewFromConfig(c, options...)
	input := &ssm.GetParameterInput{
		Name:           name,
		WithDecryption: withDecryption,
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := client.GetParameter(ctx, input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetParameters(
	c aws.Config,
	timeout time.Duration,
	names []string,
	withDecryption *bool,
	options ...func(*ssm.Options),
) (*ssm.GetParametersOutput, error) {
	client := ssm.NewFromConfig(c, options...)
	input := &ssm.GetParametersInput{
		Names:          names,
		WithDecryption: withDecryption,
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := client.GetParameters(ctx, input)
	if err != nil {
		return nil, err
	}
	return res, nil
}
