---
page_title: "etcd_resource Resource - terraform-provider-etcd"
subcategory: ""
description: |-
  resource to set a key/value pair on the Etcd Cluster
---

# Resource `key_value resource`

resource to set a key/value pair on the Etcd Cluster

## Example Usage

```terraform

resource "etcd_key_value" "example" {
  key = "Passbase"
  value = "Awesome"

}

```

## Schema

### Argument Reference

- **key** (String, Required) Key name.
- **value** (String, Required) value of key.


