resource "aws_s3_bucket" "s3_bucket" {
  for_each = { for s3 in var.storage : s3.name => s3 }
  bucket   = each.value.name
  acl      = each.value.acl
  tags     = each.value.tags
}