package test

import (
	"crypto/tls"
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// func TestTerraforEc2AndS3Tags(t *testing.T) {
// 	t.Parallel()

// 	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
// 		TerraformDir: "../ec2-s3",
// 	})

// 	defer terraform.Destroy(t, terraformOptions)

// 	terraform.InitAndApply(t, terraformOptions)

// 	ec2InstanceIDs := terraform.OutputList(t, terraformOptions, "instances_ids")
// 	for _, ec2Id := range ec2InstanceIDs {
// 		instanceTags := aws.GetTagsForEc2Instance(t, "us-east-2", ec2Id)
// 		assert.Equal(t, "Flugel", instanceTags["Name"])
// 		assert.Equal(t, "InfraTeam", instanceTags["Owner"])
// 	}

// 	s3Names := terraform.OutputList(t, terraformOptions, "s3_names")
// 	for _, s3Name := range s3Names {
// 		instanceTags := aws.GetS3BucketTags(t, "us-east-2", s3Name)
// 		assert.Equal(t, "Flugel", instanceTags["Name"])
// 		assert.Equal(t, "InfraTeam", instanceTags["Owner"])
// 	}
// }

func TestTerraforAlbAndEc2(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../alb-ec2",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	lb_dns := terraform.Output(t, terraformOptions, "lb_dns")

	time.Sleep(60 * time.Second)
	tlsConfig := tls.Config{}
	statusCode, body := http_helper.HttpGet(t, fmt.Sprintf("http://%s", lb_dns), &tlsConfig)

	assert.Equal(t, 200, statusCode)
	assert.NotNil(t, body)
	assert.Equal(t, `{"name" : "Fluglel ALB"}`, body)
}
