---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_dns_forward_zones Data Source - terraform-provider-bloxone"
subcategory: "DNS"
description: |-
  Retrieves information about existing Forward DNS Zones.
---

# bloxone_dns_forward_zones (Data Source)

Retrieves information about existing Forward DNS Zones.

## Example Usage

```terraform
# Get forward zones filtered by an attribute
data "bloxone_dns_forward_zones" "example_by_attribute" {
  filters = {
    fqdn = "domain.com."
  }
}

# Get forward zones filtered by tag
data "bloxone_dns_forward_zones" "example_by_tag" {
  tag_filters = {
    region = "eu"
  }
}

# Get all forward zones
data "bloxone_dns_forward_zones" "example_all" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filters` (Map of String) Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes, e.g. name. If you specify multiple filters, the results returned will have only resources that match all the specified filters.
- `tag_filters` (Map of String) Tag Filters are used to return a more specific list of results filtered by tags. If you specify multiple filters, the results returned will have only resources that match all the specified filters.

### Read-Only

- `results` (Attributes List) (see [below for nested schema](#nestedatt--results))

<a id="nestedatt--results"></a>
### Nested Schema for `results`

Optional:

- `comment` (String) Optional. Comment for zone configuration.
- `disabled` (Boolean) Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.
- `external_forwarders` (Attributes List) Optional. External DNS servers to forward to. Order is not significant. (see [below for nested schema](#nestedatt--results--external_forwarders))
- `forward_only` (Boolean) Optional. _true_ to only forward.
- `fqdn` (String) Zone FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation.
- `hosts` (List of String) The resource identifier.
- `internal_forwarders` (List of String) The resource identifier.
- `nsgs` (List of String) The resource identifier.
- `tags` (Map of String) Tagging specifics.
- `view` (String) The resource identifier.

Read-Only:

- `created_at` (String) The timestamp when the object has been created.
- `id` (String) The resource identifier.
- `mapped_subnet` (String) Reverse zone network address in the following format: "ip-address/cidr". Defaults to empty.
- `mapping` (String) Read-only. Zone mapping type. Allowed values:  * _forward_,  * _ipv4_reverse_.  * _ipv6_reverse_.  Defaults to _forward_.
- `parent` (String) The resource identifier.
- `protocol_fqdn` (String) Zone FQDN in punycode.
- `tags_all` (Map of String) Tagging specifics includes default tags.
- `updated_at` (String) The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation.
- `warnings` (Attributes List) The list of a forward zone warnings. (see [below for nested schema](#nestedatt--results--warnings))

<a id="nestedatt--results--external_forwarders"></a>
### Nested Schema for `results.external_forwarders`

Required:

- `address` (String) Server IP address.

Optional:

- `fqdn` (String) Server FQDN.

Read-Only:

- `protocol_fqdn` (String) Server FQDN in punycode.


<a id="nestedatt--results--warnings"></a>
### Nested Schema for `results.warnings`

Optional:

- `message` (String) Warning message.
- `name` (String) Name of a warning.