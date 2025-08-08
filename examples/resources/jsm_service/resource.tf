terraform {
  required_providers {
    atlassian-operations = {
      source = "atlassian/atlassian-operations"
    }
  }
}

# Basic JSM service example
resource "atlassian-operations_service" "basic" {
  name        = "JSM Service Test"
  description = "Description"
  tier        = 3
  type        = "SOFTWARE_SERVICES"
  owner = "438a6ccd-7daf-4afb-b0ee-9b440256a0a61"
  change_approvers = {
    groups = [
      "6e9fbdd1-f7c8-4066-b342-413ed4133a0"
    ]
  }
  responders = {
    users = [
      "61a11abf12351318cd0056fa95ff8"
    ]
    teams = [
      "123123"
    ]
  }
  stakeholders = {
    users = [
      "61a11abf3618cd004123ff8"
    ]
  }
  projects = {
    ids = [
      "10002"
    ]
  }
} 