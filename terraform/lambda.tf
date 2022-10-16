variable "bucket" {
  type = string
}

variable "artifact" {
  type = string
}

module "aws" {
  source         = "./lambda"
  artifactBucket = var.bucket
  artifactName   = var.artifact
  aws_region     = "eu-west-1"
  smtp_host      = "test"
  smtp_port      = 587
}