terraform {
  required_providers {
    todoist = {
      source  = "speakeasy/todoist"
      version = "0.0.2"
    }
  }
}

variable "todoist_api_key" {}

provider "todoist" {
  api_key = var.todoist_api_key
}

resource "todoist_project" "my_project" {
  name        = "Terraform"
}

resource "todoist_task" "my_task" {
  content     = "does it change?"
  priority        = 3
  project_id      = todoist_project.my_project.id
}
