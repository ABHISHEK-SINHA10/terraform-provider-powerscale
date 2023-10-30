/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"context"
	powerscale "dell/powerscale-go-client"
	"fmt"
	"regexp"
	"strings"
	"terraform-provider-powerscale/client"
	"terraform-provider-powerscale/powerscale/helper"
	"testing"

	. "github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// TestAccSmartPoolSettingsDatasource UT for SmartPoolSettingsDatasource, currently the test is against PowerScale 9.4.
func TestAccSmartPoolSettingsDatasource(t *testing.T) {
	var data = "data.powerscale_smartpool_settings.settings"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + SmartPoolSettingsDatasourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(data, "settings.global_namespace_acceleration_enabled", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.global_namespace_acceleration_state", regexp.MustCompile(`^\w+$`)),
					resource.TestMatchResourceAttr(data, "settings.manage_io_optimization", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.manage_io_optimization_apply_to_files", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.manage_protection", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.manage_protection_apply_to_files", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.protect_directories_one_level_higher", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.spillover_enabled", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.spillover_target.id", regexp.MustCompile(`^\d+$`)),
					resource.TestMatchResourceAttr(data, "settings.spillover_target.name", regexp.MustCompile(`^$|\w+`)),
					resource.TestMatchResourceAttr(data, "settings.spillover_target.type", regexp.MustCompile(`^\w+$`)),
					resource.TestMatchResourceAttr(data, "settings.ssd_l3_cache_default_enabled", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.ssd_qab_mirrors", regexp.MustCompile(`^\w+$`)),
					resource.TestMatchResourceAttr(data, "settings.ssd_system_btree_mirrors", regexp.MustCompile(`^\w+$`)),
					resource.TestMatchResourceAttr(data, "settings.ssd_system_delta_mirrors", regexp.MustCompile(`^\w+$`)),
					resource.TestMatchResourceAttr(data, "settings.virtual_hot_spare_deny_writes", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.virtual_hot_spare_hide_spare", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.virtual_hot_spare_limit_drives", regexp.MustCompile(`^\d+$`)),
					resource.TestMatchResourceAttr(data, "settings.virtual_hot_spare_limit_percent", regexp.MustCompile(`^\d+$`)),
				),
			},
		},
	})
}

func TestAccSmartPoolSettingsDatasourceNone(t *testing.T) {
	var data = "data.powerscale_smartpool_settings.settings"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetSmartPoolSettings).Return(mockV5StoragepoolSettingsNone, nil).Build()
				},
				Config: ProviderConfig + SmartPoolSettingsDatasourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(data, "settings.manage_io_optimization", "false"),
					resource.TestCheckResourceAttr(data, "settings.manage_io_optimization_apply_to_files", "false"),
					resource.TestCheckResourceAttr(data, "settings.manage_protection", "false"),
					resource.TestCheckResourceAttr(data, "settings.manage_protection_apply_to_files", "false"),
				),
			},
		},
		CheckDestroy: func(_ *terraform.State) error {
			if FunctionMocker != nil {
				FunctionMocker.UnPatch()
			}
			return nil
		},
	})
}

func TestAccSmartPoolSettingsDatasourceAll(t *testing.T) {
	var data = "data.powerscale_smartpool_settings.settings"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetSmartPoolSettings).Return(mockV5StoragepoolSettingsAll, nil).Build()
				},
				Config: ProviderConfig + SmartPoolSettingsDatasourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(data, "settings.manage_io_optimization", "true"),
					resource.TestCheckResourceAttr(data, "settings.manage_io_optimization_apply_to_files", "true"),
					resource.TestCheckResourceAttr(data, "settings.manage_protection", "true"),
					resource.TestCheckResourceAttr(data, "settings.manage_protection_apply_to_files", "true"),
				),
			},
		},
		CheckDestroy: func(_ *terraform.State) error {
			if FunctionMocker != nil {
				FunctionMocker.UnPatch()
			}
			return nil
		},
	})
}

func TestAccSmartPoolSettingsDatasourceErrorRequest(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetSmartPoolSettings).Return(nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfig + SmartPoolSettingsDatasourceConfig,
				ExpectError: regexp.MustCompile(`.*Error reading SmartPool settings*.`),
			},
		},
		CheckDestroy: func(_ *terraform.State) error {
			if FunctionMocker != nil {
				FunctionMocker.UnPatch()
			}
			return nil
		},
	})
}

func TestAccSmartPoolSettingsDatasourceErrorType(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetSmartPoolSettings).Return("invalid_type", nil).Build()
				},
				Config:      ProviderConfig + SmartPoolSettingsDatasourceConfig,
				ExpectError: regexp.MustCompile(`.*Failed to parse StoragePool Settings*.`),
			},
		},
		CheckDestroy: func(_ *terraform.State) error {
			if FunctionMocker != nil {
				FunctionMocker.UnPatch()
			}
			return nil
		},
	})
}

// TestAccSmartPoolSettingsDatasourceV16 UT for SmartPoolSettingsDatasource using mock response from PowerScale 9.5.
func TestAccSmartPoolSettingsDatasourceV16(t *testing.T) {
	var data = "data.powerscale_smartpool_settings.settings"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetSmartPoolSettings).Build().
						When(func(ctx context.Context, powerscaleClient *client.Client) bool {
							if strings.Contains(powerscaleClient.PscaleOpenAPIClient.GetConfig().Servers[0].URL, "localhost") {
								// enforce 9.5 (i.e. v16) endpoint in mock server
								powerscaleClient.OnefsVersion = client.OnefsVersion{Major: 9, Minor: 5, Patch: 0}
								return false
							} else if !strings.Contains(powerscaleClient.PscaleOpenAPIClient.GetConfig().Servers[0].URL, "localhost") &&
								powerscaleClient.OnefsVersion.IsLessThan("9.5.0") {
								// if running against an actual PowerScale 9.4, v16 is invalid, use mock data
								powerscaleClient.OnefsVersion = client.OnefsVersion{Major: 9, Minor: 5, Patch: 0}
								return true
							}
							return false
						}).Return(mockV16StoragepoolSettings, nil)
				},
				Config: ProviderConfig + SmartPoolSettingsDatasourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(data, "settings.default_transfer_limit_state", regexp.MustCompile(`^$|\w+`)),
					resource.TestMatchResourceAttr(data, "settings.default_transfer_limit_pct", regexp.MustCompile(`^$|^\d+$`)),
					resource.TestMatchResourceAttr(data, "settings.global_namespace_acceleration_enabled", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.global_namespace_acceleration_state", regexp.MustCompile(`^\w+$`)),
					resource.TestMatchResourceAttr(data, "settings.manage_io_optimization", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.manage_io_optimization_apply_to_files", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.manage_protection", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.manage_protection_apply_to_files", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.protect_directories_one_level_higher", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.spillover_enabled", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.spillover_target.id", regexp.MustCompile(`^\d+$`)),
					resource.TestMatchResourceAttr(data, "settings.spillover_target.name", regexp.MustCompile(`^$|\w+`)),
					resource.TestMatchResourceAttr(data, "settings.spillover_target.type", regexp.MustCompile(`^\w+$`)),
					resource.TestMatchResourceAttr(data, "settings.ssd_l3_cache_default_enabled", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.ssd_qab_mirrors", regexp.MustCompile(`^\w+$`)),
					resource.TestMatchResourceAttr(data, "settings.ssd_system_btree_mirrors", regexp.MustCompile(`^\w+$`)),
					resource.TestMatchResourceAttr(data, "settings.ssd_system_delta_mirrors", regexp.MustCompile(`^\w+$`)),
					resource.TestMatchResourceAttr(data, "settings.virtual_hot_spare_deny_writes", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.virtual_hot_spare_hide_spare", regexp.MustCompile(`^false|true$`)),
					resource.TestMatchResourceAttr(data, "settings.virtual_hot_spare_limit_drives", regexp.MustCompile(`^\d+$`)),
					resource.TestMatchResourceAttr(data, "settings.virtual_hot_spare_limit_percent", regexp.MustCompile(`^\d+$`)),
				),
			},
		},
		CheckDestroy: func(_ *terraform.State) error {
			if FunctionMocker != nil {
				FunctionMocker.UnPatch()
			}
			return nil
		},
	})
}

// TestAccSmartPoolSettingsDatasourceV16_files_at_default UT for SmartPoolSettingsDatasource using mock response from PowerScale 9.5.
func TestAccSmartPoolSettingsDatasourceV16FilesAtDefault(t *testing.T) {
	var data = "data.powerscale_smartpool_settings.settings"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetSmartPoolSettings).Return(mockV16StoragepoolSettingsFilesAtDefault, nil).Build()
				},
				Config: ProviderConfig + SmartPoolSettingsDatasourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(data, "settings.manage_io_optimization", "true"),
					resource.TestCheckResourceAttr(data, "settings.manage_io_optimization_apply_to_files", "false"),
					resource.TestCheckResourceAttr(data, "settings.manage_protection", "true"),
					resource.TestCheckResourceAttr(data, "settings.manage_protection_apply_to_files", "false"),
				),
			},
		},
		CheckDestroy: func(_ *terraform.State) error {
			if FunctionMocker != nil {
				FunctionMocker.UnPatch()
			}
			return nil
		},
	})
}

// TestAccSmartPoolSettingsDatasourceV16_none UT for SmartPoolSettingsDatasource using mock response from PowerScale 9.5.
func TestAccSmartPoolSettingsDatasourceV16None(t *testing.T) {
	var data = "data.powerscale_smartpool_settings.settings"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetSmartPoolSettings).Return(mockV16StoragepoolSettingsNone, nil).Build()
				},
				Config: ProviderConfig + SmartPoolSettingsDatasourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(data, "settings.manage_io_optimization", "false"),
					resource.TestCheckResourceAttr(data, "settings.manage_io_optimization_apply_to_files", "false"),
					resource.TestCheckResourceAttr(data, "settings.manage_protection", "false"),
					resource.TestCheckResourceAttr(data, "settings.manage_protection_apply_to_files", "false"),
				),
			},
		},
		CheckDestroy: func(_ *terraform.State) error {
			if FunctionMocker != nil {
				FunctionMocker.UnPatch()
			}
			return nil
		},
	})
}

// TestAccSmartPoolSettingsDatasourceV16_all UT for SmartPoolSettingsDatasource using mock response from PowerScale 9.5.
func TestAccSmartPoolSettingsDatasourceV16All(t *testing.T) {
	var data = "data.powerscale_smartpool_settings.settings"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetSmartPoolSettings).Return(mockV16StoragepoolSettingsAll, nil).Build()
				},
				Config: ProviderConfig + SmartPoolSettingsDatasourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(data, "settings.manage_io_optimization", "true"),
					resource.TestCheckResourceAttr(data, "settings.manage_io_optimization_apply_to_files", "true"),
					resource.TestCheckResourceAttr(data, "settings.manage_protection", "true"),
					resource.TestCheckResourceAttr(data, "settings.manage_protection_apply_to_files", "true"),
				),
			},
		},
		CheckDestroy: func(_ *terraform.State) error {
			if FunctionMocker != nil {
				FunctionMocker.UnPatch()
			}
			return nil
		},
	})
}

var SmartPoolSettingsDatasourceConfig = `data "powerscale_smartpool_settings" "settings" {
}

output "smartpool_settings" {
value = data.powerscale_smartpool_settings.settings
}
`

var mockV5StoragepoolSettingsNone = &powerscale.V5StoragepoolSettings{
	Settings: &powerscale.V5StoragepoolSettingsSettings{
		AutomaticallyManageIoOptimization: "none",
		AutomaticallyManageProtection:     "none",
	},
}

var mockV5StoragepoolSettingsAll = &powerscale.V5StoragepoolSettings{
	Settings: &powerscale.V5StoragepoolSettingsSettings{
		AutomaticallyManageIoOptimization: "all",
		AutomaticallyManageProtection:     "all",
	},
}

var mockV16StoragepoolSettingsNone = &powerscale.V16StoragepoolSettings{
	Settings: &powerscale.V16StoragepoolSettingsSettings{
		AutomaticallyManageIoOptimization: "none",
		AutomaticallyManageProtection:     "none",
	},
}

var mockV16StoragepoolSettingsAll = &powerscale.V16StoragepoolSettings{
	Settings: &powerscale.V16StoragepoolSettingsSettings{
		AutomaticallyManageIoOptimization: "all",
		AutomaticallyManageProtection:     "all",
	},
}

var mockV16StoragepoolSettingsFilesAtDefault = &powerscale.V16StoragepoolSettings{
	Settings: &powerscale.V16StoragepoolSettingsSettings{
		AutomaticallyManageIoOptimization: "files_at_default",
		AutomaticallyManageProtection:     "files_at_default",
	},
}

var limitPct float32 = 90
var limitState = ""
var mockV16StoragepoolSettings = &powerscale.V16StoragepoolSettings{
	Settings: &powerscale.V16StoragepoolSettingsSettings{
		AutomaticallyManageIoOptimization:  "files_at_default",
		AutomaticallyManageProtection:      "files_at_default",
		DefaultTransferLimitPct:            &limitPct,
		DefaultTransferLimitState:          &limitState,
		GlobalNamespaceAccelerationEnabled: false,
		GlobalNamespaceAccelerationState:   "inactive",
		ProtectDirectoriesOneLevelHigher:   false,
		SpilloverEnabled:                   false,
		SpilloverTarget:                    powerscale.V1StoragepoolSettingsSettingsSpilloverTarget{Id: 0, Name: "mockname", Type: "mocktype"},
		SsdL3CacheDefaultEnabled:           false,
		SsdQabMirrors:                      "one",
		SsdSystemBtreeMirrors:              "one",
		SsdSystemDeltaMirrors:              "one",
		VirtualHotSpareDenyWrites:          true,
		VirtualHotSpareHideSpare:           true,
		VirtualHotSpareLimitDrives:         0,
		VirtualHotSpareLimitPercent:        1,
	},
}
