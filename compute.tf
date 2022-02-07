resource "aws_network_interface" "nic" {
  for_each    = { for nic in var.instances : nic.name => nic }
  subnet_id   = aws_subnet.subnet[each.value.subnet].id
  private_ips = each.value.ips
}

resource "aws_instance" "my_ec2" {
  for_each      = { for vm in var.instances : vm.name => vm }
  ami           = each.value.ami
  instance_type = each.value.type

  network_interface {
    network_interface_id = aws_network_interface.nic[each.value.name].id
    device_index         = 0
  }

  credit_specification {
    cpu_credits = "unlimited"
  }

  tags = each.value.tags
}