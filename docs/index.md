---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "todoist Provider"
subcategory: ""
description: |-
  Todoist REST API v2 for ChatGPT Actions
---

# todoist Provider

Todoist REST API v2 for ChatGPT Actions

## Example Usage

```terraform
terraform {
  required_providers {
    todoist = {
      source  = "speakeasy/todoist"
      version = "0.1.0"
    }
  }
}

provider "todoist" {
  # Configuration options
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `api_key` (String, Sensitive)
- `server_url` (String) Server URL (defaults to https://api.todoist.com/rest/v2)
