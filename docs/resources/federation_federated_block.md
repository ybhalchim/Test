---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_federation_federated_block Resource - terraform-provider-bloxone"
subcategory: "IPAM Federation""
description: |-
  Manages a Federated Block.
  The Federated Block object allows a uniform representation of the address space segmentation, supporting functions such as administrative grouping, routing aggregation, delegation etc.
---

# bloxone_federation_federated_block (Resource)

Manages a Federated Block.

The Federated Block object allows a uniform representation of the address space segmentation, supporting functions such as administrative grouping, routing aggregation, delegation etc.

## Example Usage

```terraform
resource "bloxone_federation_federated_realm" "example" {
  name = "example_federation_federated_realm"
}

resource "bloxone_federation_federated_block" "example" {
  name            = "example_federation_federated_block"
  federated_realm = bloxone_federation_federated_realm.example.id
  cidr            = 24
  address         = "10.10.0.0"

  //tags
  tags = {
    site = "Site A"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cidr` (Number) The CIDR of the federated block. This is required, if _address_ does not specify it in its input.
- `federated_realm` (String) The resource identifier.

### Optional

- `address` (String) The address of the subnet in the form “a.b.c.d”
- `comment` (String) The description for the federated block. May contain 0 to 1024 characters. Can include UTF-8.
- `name` (String) The name of the federated block. May contain 1 to 256 characters. Can include UTF-8.
- `tags` (Map of String) The tags for the federated block in JSON format.

### Read-Only

- `allocation_v4` (Attributes) The percentage of the Federated Block’s total address space that is consumed by Leaf Terminals. (see [below for nested schema](#nestedatt--allocation_v4))
- `created_at` (String) Time when the object has been created.
- `id` (String) The resource identifier.
- `parent` (String) The resource identifier.
- `protocol` (String) The type of protocol of federated block (_ip4_ or _ip6_).
- `tags_all` (Map of String) The tags of the federation block in JSON format including default tags.
- `updated_at` (String) Time when the object has been updated. Equals to _created_at_ if not updated after creation.

<a id="nestedatt--allocation_v4"></a>
### Nested Schema for `allocation_v4`

Read-Only:

- `allocated` (Number) Percent of total space allocated.
- `delegated` (Number) Percent of total space delegated.
- `overlapping` (Number) Percent of total space in overlapping blocks.
- `reserved` (Number) Percent of total space reserved.