output "vpc_ids" {
  value = [
    for vpc in aws_vpc.vpc : vpc.id
  ]
}

output "instances_ids" {
  value = [
    for ec2 in aws_instance.ec2 : ec2.id
  ]
}

output "s3_ids" {
  value = [
    for s3 in aws_s3_bucket.s3_bucket : s3.id
  ]
}

output "s3_names" {
  value = [
    for s3 in aws_s3_bucket.s3_bucket : s3.bucket
  ]
}