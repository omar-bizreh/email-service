variable "smtp_host" {
  type        = string
  description = "SMTP host address"
}

variable "smtp_port" {
  type        = number
  description = "SMTP host port"
}

variable "aws_region" {
  type        = string
  description = "AWS Region for SSM"
}