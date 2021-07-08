---
page_title: "etcd_key_value_data_source Data Source - terraform-provider-etcd"
subcategory: ""
description: |-
  Sample data source in the Terraform provider etcd.
---

# Data Source `etcd_data_source`

Sample data source in the Terraform provider scaffolding.

## Example Usage

```terraform

data "etcd_key_value" "edu" {
   key = "Nosa"
}
 
```

## Schema

### Required

- **key** (String, Required) Sample attribute.

### Optional



