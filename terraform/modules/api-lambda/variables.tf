variable "name" {
    description = "Name of the Function"
}

variable "display_name" {
    description = "Display name for module"
}

variable "bucket" {
    description = "Name of the bucket"
}

variable "app_version" {
    description = "Version of the Function"
}

variable "env_vars" {
    type = "map"
    description ="Environmental variables for Function"
}

variable "stage" {
    default = "api"
    description = "Stage for deployment"
}
