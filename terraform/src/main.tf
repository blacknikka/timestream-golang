terraform {
  required_version = ">= 0.12.0"
  backend "s3" {
    region  = "us-east-1"
    encrypt = true

    bucket = "terraform-bucket-fortfstate-us-east-1"
    key    = "terraform-timestream-aws.tfstate"
  }
}

provider "aws" {
  region = var.region
}

module "network" {
  source = "./modules/network"

  base_name = var.base_name
}

module "ec2" {
  source = "./modules/ec2"

  base_name      = var.base_name
  vpc_main       = module.network.vpc_main
  subnet_for_app = module.network.subnet_for_app
}
