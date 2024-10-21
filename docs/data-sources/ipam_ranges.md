---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_ipam_ranges Data Source - terraform-provider-bloxone"
subcategory: "IPAM/DHCP"
description: |-
  Retrieves information about existing DHCP Ranges.
  A Range object represents a set of contiguous IP addresses in the same IP space with no gap, expressed as a (start, end) pair within a given subnet that are grouped together for administrative purpose and protocol management. The start and end values are not required to align with CIDR boundaries.
---

# bloxone_ipam_ranges (Data Source)

Retrieves information about existing DHCP Ranges.

A Range object represents a set of contiguous IP addresses in the same IP space with no gap, expressed as a (start, end) pair within a given subnet that are grouped together for administrative purpose and protocol management. The start and end values are not required to align with CIDR boundaries.

## Example Usage

```terraform
# Get all Ranges
data "bloxone_ipam_ranges" "example_all_ranges" {}

## Get specific Range by start and end values
data "bloxone_ipam_ranges" "example_range_by_start_end" {
  filters = {
    "start" = "192.168.1.15",
    "end"   = "192.168.1.30"
  }
}

## Get specific Range by name
data "bloxone_ipam_ranges" "example_range_by_name" {
  filters = {
    "name" = "example_range"
  }
}

# Get Range by tag
data "bloxone_ipam_ranges" "example_range_by_tag" {
  tag_filters = {
    location = "site1"
  }
}
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

Required:

- `end` (String) The end IP address of the range.
- `space` (String) The resource identifier.
- `start` (String) The start IP address of the range.

Optional:

- `comment` (String) The description for the range. May contain 0 to 1024 characters. Can include UTF-8.
- `dhcp_host` (String) The resource identifier.
- `dhcp_options` (Attributes List) The list of DHCP options. May be either a specific option or a group of options. (see [below for nested schema](#nestedatt--results--dhcp_options))
- `disable_dhcp` (Boolean) Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.  Defaults to _false_.
- `exclusion_ranges` (Attributes List) The list of all exclusion ranges in the scope of the range. (see [below for nested schema](#nestedatt--results--exclusion_ranges))
- `filters` (Attributes List) The list of all allow/deny filters of the range. (see [below for nested schema](#nestedatt--results--filters))
- `inheritance_sources` (Attributes) (see [below for nested schema](#nestedatt--results--inheritance_sources))
- `name` (String) The name of the range. May contain 1 to 256 characters. Can include UTF-8.
- `tags` (Map of String) The tags for the range in JSON format.

Read-Only:

- `compartment_id` (String) The compartment associated with the object. If no compartment is associated with the object, the value defaults to empty.
- `created_at` (String) Time when the object has been created.
- `id` (String) The resource identifier.
- `inheritance_assigned_hosts` (Attributes List) The list of the inheritance assigned hosts of the object. (see [below for nested schema](#nestedatt--results--inheritance_assigned_hosts))
- `inheritance_parent` (String) The resource identifier.
- `parent` (String) The resource identifier.
- `protocol` (String) The type of protocol (_ip4_ or _ip6_).
- `space_name` (String) The name of the IP Space the range belongs to.
- `tags_all` (Map of String) The tags for the range in JSON format including default tags.
- `threshold` (Attributes) The utilization threshold settings for the range. (see [below for nested schema](#nestedatt--results--threshold))
- `updated_at` (String) Time when the object has been updated. Equals to _created_at_ if not updated after creation.
- `utilization` (Attributes) The utilization statistics of IPV4 addresses for the range. (see [below for nested schema](#nestedatt--results--utilization))
- `utilization_v6` (Attributes) The utilization of IPV6 addresses in the range. (see [below for nested schema](#nestedatt--results--utilization_v6))

<a id="nestedatt--results--dhcp_options"></a>
### Nested Schema for `results.dhcp_options`

Optional:

- `group` (String) The resource identifier.
- `option_code` (String) The resource identifier.
- `option_value` (String) The option value.
- `type` (String) The type of item. Valid values are:
  * _group_
  * _option_


<a id="nestedatt--results--exclusion_ranges"></a>
### Nested Schema for `results.exclusion_ranges`

Required:

- `end` (String) The end address of the exclusion range.
- `start` (String) The start address of the exclusion range.

Optional:

- `comment` (String) The description for the exclusion range. May contain 0 to 1024 characters. Can include UTF-8.


<a id="nestedatt--results--filters"></a>
### Nested Schema for `results.filters`

Required:

- `access` (String) The access type of DHCP filter (_allow_ or _deny_).  Defaults to _allow_.

Optional:

- `hardware_filter_id` (String) The resource identifier.
- `option_filter_id` (String) The resource identifier.


<a id="nestedatt--results--inheritance_sources"></a>
### Nested Schema for `results.inheritance_sources`

Optional:

- `dhcp_options` (Attributes) The inheritance configuration for the _dhcp_options_ field. (see [below for nested schema](#nestedatt--results--inheritance_sources--dhcp_options))

<a id="nestedatt--results--inheritance_sources--dhcp_options"></a>
### Nested Schema for `results.inheritance_sources.dhcp_options`

Optional:

- `action` (String) The inheritance setting. Valid values are:
  * _inherit_: Use the inherited value.
  * _block_: Don't use the inherited value.

  Defaults to _inherit_.
- `value` (Attributes List) The inherited DHCP option values. (see [below for nested schema](#nestedatt--results--inheritance_sources--dhcp_options--value))

<a id="nestedatt--results--inheritance_sources--dhcp_options--value"></a>
### Nested Schema for `results.inheritance_sources.dhcp_options.value`

Optional:

- `action` (String) The inheritance setting. Valid values are:
  * _inherit_: Use the inherited value.
  * _block_: Don't use the inherited value.

  Defaults to _inherit_.

Read-Only:

- `display_name` (String) The human-readable display name for the object referred to by _source_.
- `source` (String) The resource identifier.
- `value` (Attributes) The inherited value for the DHCP option. (see [below for nested schema](#nestedatt--results--inheritance_sources--dhcp_options--value--value))

<a id="nestedatt--results--inheritance_sources--dhcp_options--value--value"></a>
### Nested Schema for `results.inheritance_sources.dhcp_options.value.value`

Read-Only:

- `option` (Attributes) Option inherited from the ancestor. (see [below for nested schema](#nestedatt--results--inheritance_sources--dhcp_options--value--value--option))
- `overriding_group` (String) The resource identifier.

<a id="nestedatt--results--inheritance_sources--dhcp_options--value--value--option"></a>
### Nested Schema for `results.inheritance_sources.dhcp_options.value.value.option`

Read-Only:

- `group` (String) The resource identifier.
- `option_code` (String) The resource identifier.
- `option_value` (String) The option value.
- `type` (String) The type of item. Valid values are:
  * _group_
  * _option_






<a id="nestedatt--results--inheritance_assigned_hosts"></a>
### Nested Schema for `results.inheritance_assigned_hosts`

Optional:

- `host` (String) The resource identifier.

Read-Only:

- `display_name` (String) The human-readable display name for the host referred to by _ophid_.
- `ophid` (String) The on-prem host ID.


<a id="nestedatt--results--threshold"></a>
### Nested Schema for `results.threshold`

Required:

- `enabled` (Boolean) Indicates whether the utilization threshold for IP addresses is enabled or not.
- `high` (Number) The high threshold value for the percentage of used IP addresses relative to the total IP addresses available in the scope of the object. Thresholds are inclusive in the comparison test.
- `low` (Number) The low threshold value for the percentage of used IP addresses relative to the total IP addresses available in the scope of the object. Thresholds are inclusive in the comparison test.


<a id="nestedatt--results--utilization"></a>
### Nested Schema for `results.utilization`

Read-Only:

- `abandon_utilization` (Number) The percentage of abandoned IP addresses relative to the total IP addresses available in the scope of the object.
- `abandoned` (String) The number of IP addresses in the scope of the object which are in the abandoned state (issued by a DHCP server and then declined by the client).
- `dynamic` (String) The number of IP addresses handed out by DHCP in the scope of the object. This includes all leased addresses, fixed addresses that are defined but not currently leased and abandoned leases.
- `free` (String) The number of IP addresses available in the scope of the object.
- `static` (String) The number of defined IP addresses such as reservations or DNS records. It can be computed as _static_ = _used_ - _dynamic_.
- `total` (String) The total number of IP addresses available in the scope of the object.
- `used` (String) The number of IP addresses used in the scope of the object.
- `utilization` (Number) The percentage of used IP addresses relative to the total IP addresses available in the scope of the object.


<a id="nestedatt--results--utilization_v6"></a>
### Nested Schema for `results.utilization_v6`

Optional:

- `abandoned` (String)
- `dynamic` (String)
- `static` (String)
- `total` (String)
- `used` (String)