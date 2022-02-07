networks = [
  {
    network = "172.16.0.0/16"
    subnets = [
      {
        cidr = "172.16.0.0/24"
        az   = "us-east-2a"
      },
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
  },
  {
    name   = "VM2"
    ami    = "ami-089c6f2e3866f0f14"
    type   = "t2.micro"
    subnet = "172.16.1.0/24"
    ips    = ["172.16.1.20"]
    tags = {
      Name  = "Flugel"
      Owner = "InfraTeam"
    }
  }
]

storage = [
  {
    name = "jgomezve-aws-bucket"
    acl  = "private"
    tags = {
      name  = "flugel"
      owner = "infraTeam"
    }
  }
]