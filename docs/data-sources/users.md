---
page_title: "etcd_Users_data_source Data Source - terraform-provider-etcd"
subcategory: ""
description: |-
  Sample data source in the Terraform provider etcd.
---

# Data Source `etcd_users data_source`

Sample data source in the Terraform provider scaffolding.

## Example Usage

```terraform

data "etcd_users" "edu" {
  depends_on = [
    etcd_user.user,
  ]
}

```

## Schema

### Required


### Optional




