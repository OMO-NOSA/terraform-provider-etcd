---
page_title: "Etcd Provider"
subcategory: ""
description: |-
  
---

# etcd Provider



## Example Usage

```terraform
terraform {
  required_providers {
    etcd = {
      version = "0.1"
      source  = "hashicorp.com/passbase/etcd"
    }
  }
}


provider "etcd" {
  endpoints = [ "localhost:2379" ]
  username = "root"
  password = "root"
  
}

resource "etcd_user" "user"{
  username = "passbase"
  password = "password"
 
}

```

## Schema


### Arguments Reference

- **username** (String, Required) The root username.
- **password** (String, Required) The root user password.
- **endpoints** (String, Required) Cluster endpoint.
