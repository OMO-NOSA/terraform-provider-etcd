---
page_title: "grant_user_role Resource - terraform-provider-etcd"
subcategory: ""
description: |-
  This resource grants a predefined role to a user.
---

# Resource `etcd_grant_user_role resource`

The grant_user_role resource grants a predefined role to the user, it takes as argument;

- role_name
- username

## Example Usage

```terraform

resource "etcd_grant_user_role" "example" {
  role_name = etcd_role.role.name
  username = etcd_user.user.username
}

```

## Schema

### Argument Reference

- **role_name** (String, Required) The role_name to grant user.
- **username** (String, Required) user to grant role.

