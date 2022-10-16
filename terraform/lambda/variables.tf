variable "artifactBucket" {
  type        = string
  description = "S3 Bucket to read lambda zip file from"
}

variable "artifactName" {
  type        = string
  description = "Lambda compiled code zip file "
}

variable "handler" {
  type        = string
  default     = "main.handler"
  description = "Lambda function main handler"
}

variable "runtime" {
  type        = string
  default     = "go1.x"
  description = "Lambda runtime"
}

variable "functionName" {
  type        = string
  default     = "email-service"
  description = "Lambda function name"
}