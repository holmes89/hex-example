variable "app_version" {
}
variable "redis_url" {
}
variable "redis_pswd" {
}

provider "aws" {
  region = "us-east-1"
}

module "hex-api" {
  source = "../modules/api-lambda"

  name = "hex-example"
  display_name = "Hex Example"
  bucket = "hex-lambda"
  app_version = "${var.app_version}"
  env_vars = {
      REDIS_URL = "${var.redis_url}"
      REDIS_PASSWORD = "${var.redis_pswd}"
  }
}

output "url" {
  value = "${module.hex-api.base_url}"
}
