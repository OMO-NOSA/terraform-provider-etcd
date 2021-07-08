---
page_title: "etcd_resource Resource - terraform-provider-etcd"
subcategory: ""
description: |-
  Sample resource in the Terraform provider etcd.
---

# Resource `etcd_resource`

Sample resource in the Terraform provider scaffolding.

## Example Usage

```terraform

resource "etcd_auth" "auth" {
  enabled = false
}

resource "etcd_key_value" "edu" {
  key = "Passbase"
  value = "Awesome"

}

resource "etcd_user" "user"{
  username = "passbase"
  password = var.user_password
 
}

resource "etcd_role" "role" {
  name = "developer"
  
}

resource "etcd_grant_role_permission" "perm" {
  role_name = "developer"
  key = etcd_key_value.edu.key
  permission = "WRITE"
  range = "test"

}

resource "etcd_grant_user_role" "gmt" {
  role_name = etcd_role.role.name
  username = etcd_user.user.username
}

```

## Schema

### Optional

- **id** (String, Optional) The ID of this resource.
- **sample_attribute** (String, Optional) Sample attribute.


