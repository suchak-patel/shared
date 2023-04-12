# Using module from TF registry

module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = "${var.Environment}-vpc"
  cidr = var.vpc_cidr_block

  azs             = var.vpc_azs
  private_subnets = var.vpc_private_subnets
  public_subnets  = var.vpc_public_subnets

  enable_nat_gateway = true

  tags = {
    Terraform = "true"
    Environment = var.Environment
  }
}
