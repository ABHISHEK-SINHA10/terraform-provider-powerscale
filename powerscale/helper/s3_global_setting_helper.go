package helper

import (
	"context"
	powerscale "dell/powerscale-go-client"
	"terraform-provider-powerscale/client"
	"terraform-provider-powerscale/powerscale/models"
)

// GetS3GlobalSetting gets S3 GlobalSetting.
func GetS3GlobalSetting(ctx context.Context, client *client.Client) (*powerscale.V10S3SettingsGlobal, error) {
	param := client.PscaleOpenAPIClient.ProtocolsApi.GetProtocolsv10S3SettingsGlobal(ctx)
	response, _, err := param.Execute()
	return response, err
}

// UpdateS3GlobalSetting update s3 GlobalSetting.
func UpdateS3GlobalSetting(ctx context.Context, client *client.Client, GlobalSettingToUpdate powerscale.V10S3SettingsGlobalSettings) error {
	updateParam := client.PscaleOpenAPIClient.ProtocolsApi.UpdateProtocolsv10S3SettingsGlobal(ctx)
	updateParam = updateParam.V10S3SettingsGlobal(GlobalSettingToUpdate)
	_, err := updateParam.Execute()
	return err
}

// SetGlobalSetting updates the S3 Global Setting.
func SetGlobalSetting(ctx context.Context, client *client.Client, s3GSPlan models.S3GlobalSettingResource) (models.S3GlobalSettingResource, error) {
	var toUpdate powerscale.V10S3SettingsGlobalSettings
	err := ReadFromState(ctx, &s3GSPlan, &toUpdate)
	if err != nil {
		return models.S3GlobalSettingResource{}, err
	}
	err = UpdateS3GlobalSetting(ctx, client, toUpdate)
	if err != nil {
		return models.S3GlobalSettingResource{}, err
	}
	globalSettings, err := GetS3GlobalSetting(ctx, client)
	if err != nil {
		return models.S3GlobalSettingResource{}, err
	}
	var state models.S3GlobalSettingResource
	err = CopyFieldsToNonNestedModel(ctx, globalSettings.GetSettings(), &state)
	if err != nil {
		return models.S3GlobalSettingResource{}, err
	}
	return state, nil
}

// GetGlobalSetting reads the S3 Global Setting.
func GetGlobalSetting(ctx context.Context, client *client.Client, s3GlobalSettingState models.S3GlobalSettingResource) error {
	globalSettings, err := GetS3GlobalSetting(ctx, client)
	if err != nil {
		return err
	}
	err = CopyFieldsToNonNestedModel(ctx, globalSettings.GetSettings(), &s3GlobalSettingState)
	if err != nil {
		return err
	}
	return nil
}
