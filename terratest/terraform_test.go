package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraforEc2AndS3Tags(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../ec2-s3",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	ec2InstanceIDs := terraform.OutputList(t, terraformOptions, "instances_ids")
	for _, ec2Id := range ec2InstanceIDs {
		instanceTags := aws.GetTagsForEc2Instance(t, "us-east-2", ec2Id)
		assert.Equal(t, "Flugel", instanceTags["Name"])
		assert.Equal(t, "InfraTeam", instanceTags["Owner"])
	}

	s3Names := terraform.OutputList(t, terraformOptions, "s3_names")
	for _, s3Name := range s3Names {
		instanceTags := aws.GetS3BucketTags(t, "us-east-2", s3Name)
		assert.Equal(t, "Flugel", instanceTags["Name"])
		assert.Equal(t, "InfraTeam", instanceTags["Owner"])
	}
}
