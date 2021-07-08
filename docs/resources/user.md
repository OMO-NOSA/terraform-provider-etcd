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

resource "etcd_user" "user"{
  username = "passbase"
  password = "password"
 
}

```

## Schema

### Arguments Reference

- **username** (String, Required) The username to be created.
- **password** (String, Required) Password for the user, password length should be > 9 characters.


