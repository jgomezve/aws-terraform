# AWS Terraform & Terratest
[![Tests](https://github.com/jgomezve/terraform-sandbox/actions/workflows/test.yaml/badge.svg)](https://github.com/jgomezve/terraform-sandbox/actions/workflows/test.yaml)

## Prerequisites

* Make sure to have Go 1.17 installed on your computer

    * [Install Go](https://go.dev/doc/install)

* Set your AWS credentials as environmental variables 


        export AWS_ACCESS_KEY_ID="anaccesskey"
        export AWS_SECRET_ACCESS_KEY="asecretkey"

> **_NOTE:_**:  More information here [AWS Terraform provider](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)

## Execution - Use Case EC2 Intance & S3 Bucket

### Terraform

* Go to the **EC2 Instance & S3 Bucket** folder

        cd ec2-s3

* Initialize Terraform and the AWS provider

        terraform init

* Setup the EC2 Intances, Network and S3 Bucket configuration in the file `ec2-s3/variables.auto.tfvars`. Here an example:

```hcl
networks = [
  {
    network = "172.16.0.0/16"
    subnets = [
      {
        cidr = "172.16.1.0/24"
        az   = "us-east-2a"
      }
    ]
  }
]

instances = [
  {
    name   = "VM1"
    ami    = "ami-089c6f2e3866f0f14"
    type   = "t2.micro"
    subnet = "172.16.1.0/24"
    ips    = ["172.16.1.10"]
    tags = {
      Name  = "Flugel"
      Owner = "InfraTeam"
    }
  }
]

storage = [
  {
    name = "my-aws-bucket"
    acl  = "private"
    tags = {
      Name  = "Flugel"
      Owner = "InfraTeam"
    }
  }
]
```

> **_NOTE:_**: See the `ec2-s3/variables.tf` file to check the supported schema

* Execute Terraform

        terraform apply -auto-approve


* Destroy Terraform Resources (Optional)

        terraform destroy -auto-approve


### Terratest

* Go to the `terratest` folder and download the Go dependencies

        cd terratest
        go get -v -d

* Execute the test file `terraform_test.go`

        go test -v