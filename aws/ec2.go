package awsutils

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/restump/go-utils"
)

func GetEnabledRegions(
	c aws.Config,
	timeout time.Duration,
	options ...func(*ec2.Options),
) (*ec2.DescribeRegionsOutput, error) {
	client := ec2.NewFromConfig(c)
	input := &ec2.DescribeRegionsInput{
		Filters: []ec2types.Filter{
			{
				Name: utils.ToPtr("opt-in-status"),
				Values: []string{
					"opt-in-not-required",
					"opted-in",
				},
			},

		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := client.DescribeRegions(ctx, input)
	if err != nil {
		return nil, err
	}
	return res, nil
}


