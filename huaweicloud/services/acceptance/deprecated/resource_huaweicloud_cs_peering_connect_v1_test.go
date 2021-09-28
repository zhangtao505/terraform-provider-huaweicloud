// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://www.github.com/huaweicloud/magic-modules
//
// ----------------------------------------------------------------------------

package deprecated

import (
	"fmt"
	"testing"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils/fmtp"

	"github.com/chnsz/golangsdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
)

func TestAccCsPeeringConnectV1_basic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAccPreCheckDeprecated(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckCsPeeringConnectV1Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsPeeringConnectV1_basic(acctest.RandString(10)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsPeeringConnectV1Exists(),
				),
			},
		},
	})
}

func testAccCsPeeringConnectV1_basic(val string) string {
	return fmt.Sprintf(`
resource "huaweicloud_cs_cluster_v1" "cluster" {
  name = "terraform_cs_cluster_v1_test%s"
}

resource "huaweicloud_vpc_v1" "vpc" {
  name = "terraform_vpc_v1_test%s"
  cidr = "192.168.0.0/16"
}

resource "huaweicloud_vpc_subnet_v1" "subnet" {
  name = "terraform_vpc_subnet_v1_test%s"
  cidr = "192.168.0.0/16"
  gateway_ip = "192.168.0.1"
  vpc_id = "${huaweicloud_vpc_v1.vpc.id}"
}

resource "huaweicloud_cs_peering_connect_v1" "peering" {
  name = "terraform_cs_peering_connect_v1_test%s"
  target_vpc_info {
    vpc_id = "${huaweicloud_vpc_v1.vpc.id}"
  }
  cluster_id = "${huaweicloud_cs_cluster_v1.cluster.id}"
}
	`, val, val, val, val)
}

func testAccCheckCsPeeringConnectV1Destroy(s *terraform.State) error {
	config := acceptance.TestAccProvider.Meta().(*config.Config)
	client, err := config.CloudStreamV1Client(acceptance.HW_REGION_NAME)
	if err != nil {
		return fmtp.Errorf("Error creating sdk client, err=%s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "huaweicloud_cs_peering_connect_v1" {
			continue
		}

		url, err := replaceVarsForTest(rs, "reserved_cluster/{cluster_id}/peering/{id}")
		if err != nil {
			return err
		}
		url = client.ServiceURL(url)

		_, err = client.Get(url, nil, &golangsdk.RequestOpts{
			MoreHeaders: map[string]string{"Content-Type": "application/json"}})
		if err == nil {
			return fmtp.Errorf("huaweicloud_cs_peering_connect_v1 still exists at %s", url)
		}
	}

	return nil
}

func testAccCheckCsPeeringConnectV1Exists() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := acceptance.TestAccProvider.Meta().(*config.Config)
		client, err := config.CloudStreamV1Client(acceptance.HW_REGION_NAME)
		if err != nil {
			return fmtp.Errorf("Error creating sdk client, err=%s", err)
		}

		rs, ok := s.RootModule().Resources["huaweicloud_cs_peering_connect_v1.peering"]
		if !ok {
			return fmtp.Errorf("Error checking huaweicloud_cs_peering_connect_v1.peering exist, err=not found this resource")
		}

		url, err := replaceVarsForTest(rs, "reserved_cluster/{cluster_id}/peering/{id}")
		if err != nil {
			return fmtp.Errorf("Error checking huaweicloud_cs_peering_connect_v1.peering exist, err=building url failed: %s", err)
		}
		url = client.ServiceURL(url)

		_, err = client.Get(url, nil, &golangsdk.RequestOpts{
			MoreHeaders: map[string]string{"Content-Type": "application/json"}})
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return fmtp.Errorf("huaweicloud_cs_peering_connect_v1.peering is not exist")
			}
			return fmtp.Errorf("Error checking huaweicloud_cs_peering_connect_v1.peering exist, err=send request failed: %s", err)
		}
		return nil
	}
}