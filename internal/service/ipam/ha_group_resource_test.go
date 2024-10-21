package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccHaGroupResource_basic(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test"
	var v ipam.HAGroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupBasicConfig("active", "active", "test-ha", "active-active"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "hosts", "HOSTS_REPLACE_ME"),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// func TestAccHaGroupResource_disappears(t *testing.T) {
// 	resourceName := "bloxone_ipam_ha_group.test"
// 	var v ipam.HAGroup

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		CheckDestroy:             testAccCheckHaGroupDestroy(context.Background(), &v),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccHaGroupBasicConfig("HOSTS_REPLACE_ME", "NAME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
// 					testAccCheckHaGroupDisappears(context.Background(), &v),
// 				),
// 				ExpectNonEmptyPlan: true,
// 			},
// 		},
// 	})
// }

func TestAccHaGroupResource_AnycastConfigId(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test_anycast_config_id"
	var v ipam.HAGroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupAnycastConfigId("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "ANYCAST_CONFIG_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "anycast_config_id", "ANYCAST_CONFIG_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHaGroupAnycastConfigId("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "ANYCAST_CONFIG_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "anycast_config_id", "ANYCAST_CONFIG_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHaGroupResource_Comment(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test_comment"
	var v ipam.HAGroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupComment("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHaGroupComment("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHaGroupResource_Hosts(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test_hosts"
	var v ipam.HAGroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupHosts("HOSTS_REPLACE_ME", "NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hosts", "HOSTS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHaGroupHosts("HOSTS_REPLACE_ME", "NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hosts", "HOSTS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHaGroupResource_IpSpace(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test_ip_space"
	var v ipam.HAGroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupIpSpace("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "IP_SPACE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ip_space", "IP_SPACE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHaGroupIpSpace("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "IP_SPACE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ip_space", "IP_SPACE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHaGroupResource_Mode(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test_mode"
	var v ipam.HAGroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupMode("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "MODE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mode", "MODE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHaGroupMode("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "MODE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mode", "MODE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHaGroupResource_Name(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test_name"
	var v ipam.HAGroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupName("HOSTS_REPLACE_ME", "NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHaGroupName("HOSTS_REPLACE_ME", "NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHaGroupResource_Status(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test_status"
	var v ipam.HAGroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupStatus("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "STATUS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "status", "STATUS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHaGroupStatus("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "STATUS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "status", "STATUS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHaGroupResource_StatusV6(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test_status_v6"
	var v ipam.HAGroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupStatusV6("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "STATUS_V6_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "status_v6", "STATUS_V6_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHaGroupStatusV6("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "STATUS_V6_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "status_v6", "STATUS_V6_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHaGroupResource_Tags(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test_tags"
	var v ipam.HAGroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupTags("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "TAGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHaGroupTags("HOSTS_REPLACE_ME", "NAME_REPLACE_ME", "TAGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckHaGroupExists(ctx context.Context, resourceName string, v *ipam.HAGroup) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			HaGroupAPI.
			Read(ctx, rs.Primary.ID).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetResult()
		return nil
	}
}

func testAccCheckHaGroupDestroy(ctx context.Context, v *ipam.HAGroup) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
			HaGroupAPI.
			Read(ctx, *v.Id).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckHaGroupDisappears(ctx context.Context, v *ipam.HAGroup) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
			HaGroupAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccHaGroupBasicConfig(role1, role2, name, mode string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_ha_group" "test" {
	hosts = [
		{
			host = data.bloxone_dhcp_hosts.test.results.0.id
			role = %q
		},
		{
			host = data.bloxone_dhcp_hosts.test.results.1.id
			role = %q
		}
	]
	name = %q
	mode = %q
}`, role1, role2, name, mode)

	return strings.Join([]string{acctest.TestAccBase_DhcpHosts(), config}, "")
}

func testAccHaGroupAnycastConfigId(hosts string, name string, anycastConfigId string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test_anycast_config_id" {
    hosts = %q
    name = %q
    anycast_config_id = %q
}
`, hosts, name, anycastConfigId)
}

func testAccHaGroupComment(hosts string, name string, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test_comment" {
    hosts = %q
    name = %q
    comment = %q
}
`, hosts, name, comment)
}

func testAccHaGroupHosts(hosts string, name string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test_hosts" {
    hosts = %q
    name = %q
}
`, hosts, name)
}

func testAccHaGroupIpSpace(hosts string, name string, ipSpace string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test_ip_space" {
    hosts = %q
    name = %q
    ip_space = %q
}
`, hosts, name, ipSpace)
}

func testAccHaGroupMode(hosts string, name string, mode string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test_mode" {
    hosts = %q
    name = %q
    mode = %q
}
`, hosts, name, mode)
}

func testAccHaGroupName(hosts string, name string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test_name" {
    hosts = %q
    name = %q
}
`, hosts, name)
}

func testAccHaGroupStatus(hosts string, name string, status string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test_status" {
    hosts = %q
    name = %q
    status = %q
}
`, hosts, name, status)
}

func testAccHaGroupStatusV6(hosts string, name string, statusV6 string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test_status_v6" {
    hosts = %q
    name = %q
    status_v6 = %q
}
`, hosts, name, statusV6)
}

func testAccHaGroupTags(hosts string, name string, tags string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test_tags" {
    hosts = %q
    name = %q
    tags = %q
}
`, hosts, name, tags)
}
