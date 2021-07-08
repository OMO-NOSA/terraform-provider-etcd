---
page_title: "Authentication Resource - terraform-provider-etcd"
subcategory: ""
description: |-
  This resource helps toggle the authentication setting for a etcd cluster -- Setting it either to enabled or disabled.
---

# Resource `etcd_auth resource`

This resource helps toggle the authentication setting for a etcd cluster -- Setting it either to enabled or disabled.


## Example Usage

```terraform

resource "etcd_auth" "auth" {
  enabled = false
}

```

## Schema

### Arguments Reference

- **enabled** (bool, Optional) Can either be set to true or false based on preference.


