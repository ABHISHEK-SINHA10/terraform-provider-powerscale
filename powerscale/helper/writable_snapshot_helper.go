/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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

package helper

import (
	"context"
	powerscale "dell/powerscale-go-client"
	"strconv"
	"terraform-provider-powerscale/client"
	"terraform-provider-powerscale/powerscale/models"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// GetWritableSnapshot retrieve writable snapshot object.
func GetWritableSnapshot(ctx context.Context, client *client.Client, path string) (*powerscale.V14SnapshotWritableExtended, error) {
	writableSnapshot, _, err := client.PscaleOpenAPIClient.SnapshotApi.GetSnapshotv14SnapshotWritableWspath(ctx, path).Execute()
	return writableSnapshot, err
}

// UpdateWritableSnapshot creates/updates writable snapshots.
func UpdateWritableSnapshot(ctx context.Context, client *client.Client, v14WritableSnapshot powerscale.V14SnapshotWritableItem) (*powerscale.Createv14SnapshotWritableItemResponse, error) {
	writableSnapshot, _, err := client.PscaleOpenAPIClient.SnapshotApi.CreateSnapshotv14SnapshotWritableItem(ctx).V14SnapshotWritableItem(v14WritableSnapshot).Execute()
	return writableSnapshot, err
}

// DeleteWritableSnapshot retrieve writable snapshot object.
func DeleteWritableSnapshot(ctx context.Context, client *client.Client, path string) error {
	_, err := client.PscaleOpenAPIClient.SnapshotApi.DeleteSnapshotv14SnapshotWritableWspath(ctx, path).Execute()
	return err
}

// UpdateWritableSnapshotState updates the state parameters based on the fetched computed values from the API.
func UpdateWritableSnapshotState(state *models.WritableSnapshot, fetchedState *powerscale.Createv14SnapshotWritableItemResponse) {
	state.DstPath = types.StringValue(fetchedState.DstPath)
	// since the create operation is performed using snap_id so converting the fetched int32 value to string for state and plan match
	state.SrcSnap = types.StringValue(strconv.Itoa(int(fetchedState.SrcId)))
	state.ID = types.Int32Value(fetchedState.Id)
	state.SrcPath = types.StringValue(fetchedState.SrcPath)
	state.State = types.StringValue(fetchedState.State)
	state.SnapName = types.StringValue(fetchedState.SrcSnap)
}
