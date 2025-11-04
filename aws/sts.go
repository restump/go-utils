package awsutils

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func GetCallerIdentity(
	c aws.Config, timeout time.Duration, options ...func(*sts.Options),
) (*sts.GetCallerIdentityOutput, error) {
	client := sts.NewFromConfig(c, options...)
	input := &sts.GetCallerIdentityInput{}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := client.GetCallerIdentity(ctx, input)
	if err != nil {
		return nil, err
	}
	return res, nil
}
