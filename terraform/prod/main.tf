module "hex-api" {
  source "../modules/api-lambda"

  name = "hex-example"
  bucket = "hex-lambda"
  app_version = "${var.app_version}"
  env_vars = {
    variables = {
      REDIS_URL = "${var.redis_url}"
      REDIS_PASSWORD = "${var.redis_pswd}"
    }
  }
}
