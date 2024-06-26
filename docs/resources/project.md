---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "todoist_project Resource - terraform-provider-todoist"
subcategory: ""
description: |-
  Project Resource
---

# todoist_project (Resource)

Project Resource

## Example Usage

```terraform
resource "todoist_project" "my_project" {
  color       = "...my_color..."
  is_favorite = true
  name        = "Andres Rogahn"
  parent_id   = "...my_parent_id..."
  view_style  = "...my_view_style..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the project

### Optional

- `color` (String) The color of the project icon. Refer to the name column in the Colors guide for more info. https://developer.todoist.com/guides/#colors
- `is_favorite` (Boolean) Whether the project is a favorite (a true or false value).
- `parent_id` (String) Parent project ID. Requires replacement if changed.
- `view_style` (String) A string value (either list or board, default is list). This determines the way the project is displayed within the Todoist clients.

### Read-Only

- `comment_count` (Number) Number of project comments.
- `id` (String) Project ID.
- `is_inbox_project` (Boolean) Whether the project is the user's Inbox (read-only).
- `is_shared` (Boolean) Whether the project is shared (read-only, a true or false value).
- `is_team_inbox` (Boolean) Whether the project is the Team Inbox (read-only).
- `order` (Number) Project position under the same parent (read-only, will be 0 for inbox and team inbox projects).
- `url` (String) URL to access this project in the Todoist web or mobile applications.

## Import

Import is supported using the following syntax:

```shell
terraform import todoist_project.my_todoist_project ""
```
