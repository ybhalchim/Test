package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccHaGroupDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_ha_groups.test"
	resourceName := "bloxone_ipam_ha_group.test"
	var v ipam.HAGroup

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHaGroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHaGroupDataSourceConfigFilters("HOSTS_REPLACE_ME", "NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					}, testAccCheckHaGroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccHaGroupDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_ha_groups.test"
	resourceName := "bloxone_ipam_ha_group.test"
	var v ipam.HAGroup
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHaGroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHaGroupDataSourceConfigTagFilters("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					}, testAccCheckHaGroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckHaGroupResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "anycast_config_id", dataSourceName, "results.0.anycast_config_id"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "hosts", dataSourceName, "results.0.hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "ip_space", dataSourceName, "results.0.ip_space"),
		resource.TestCheckResourceAttrPair(resourceName, "mode", dataSourceName, "results.0.mode"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "status", dataSourceName, "results.0.status"),
		resource.TestCheckResourceAttrPair(resourceName, "status_v6", dataSourceName, "results.0.status_v6"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccHaGroupDataSourceConfigFilters(hosts, name string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test" {
  hosts = %q
  name = %q
}

data "bloxone_ipam_ha_groups" "test" {
  filters = {
	hosts = bloxone_ipam_ha_group.test.hosts
  }
}
`, hosts, name)
}

func testAccHaGroupDataSourceConfigTagFilters(hosts, name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test" {
  hosts = %q
  name = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_ipam_ha_groups" "test" {
  tag_filters = {
	tag1 = bloxone_ipam_ha_group.test.tags.tag1
  }
}
`, hosts, name, tagValue)
}
