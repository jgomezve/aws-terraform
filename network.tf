locals {
  aws_subnets = flatten([for net in var.networks : [
    for subnet in net.subnets : {
      key               = "${subnet.cidr}"
      network           = net.network
      cidr              = subnet.cidr
      availability_zone = subnet.az
    }
    ]
  ])
}

resource "aws_vpc" "vpc" {
  for_each   = { for net in var.networks : net.network => net }
  cidr_block = each.value.network
}


resource "aws_subnet" "subnet" {
  for_each          = { for cidr in local.aws_subnets : cidr.key => cidr }
  vpc_id            = aws_vpc.vpc[each.value.network].id
  cidr_block        = each.value.cidr
  availability_zone = each.value.availability_zone
}