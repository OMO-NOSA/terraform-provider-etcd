---
page_title: "etcd_grant_role_permission Resource - terraform-provider-etcd"
subcategory: ""
description: |-
  This resource grants some set of permissions to an existing role with the etcd cluster
---

# Resource `grant_role_permission`

grant_role_permission resource in the Terraform Etcd Provider.

## Example Usage

```terraform
resource "etcd_grant_role_permission" "example" {
  role_name = "developer"
  key = etcd_key_value.edu.key
  permission = "WRITE"
  range = "test"

}

```

## Schema

### Arguments Reference

- **role_name** (String, required) Name of an already created role.
- **key** (String, required) Name of key to grant permission On.
- **permission** (String, required) Permission to grant to role -- READ | WRITE | READWRITE.
- **range** (String, required) Range of keys to grant permissions 


