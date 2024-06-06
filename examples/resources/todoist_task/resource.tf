resource "todoist_task" "my_task" {
  assignee_id     = "...my_assignee_id..."
  content         = "...my_content..."
  description     = "...my_description..."
  due_date        = "...my_due_date..."
  due_datetime    = "...my_due_datetime..."
  due_lang        = "...my_due_lang..."
  due_string      = "...my_due_string..."
  duration_amount = 7
  duration_unit   = "...my_duration_unit..."
  order           = 0
  parent_id       = "...my_parent_id..."
  priority        = 3
  project_id      = "...my_project_id..."
  section_id      = "...my_section_id..."
}