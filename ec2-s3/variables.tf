variable "networks" {
  type = list(object({
    network = string
    subnets = list(object({
      cidr = string
      az   = string
      test = test
    }))
  }))
}

variable "instances" {
  type = list(object({
    name   = string
    ami    = string
    type   = string
    subnet = string
    ips    = list(string)
    tags   = map(string)
  }))
}

variable "storage" {
  type = list(object({
    name = string
    acl  = string
    tags = map(string)
  }))
}
