// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package api4

import (
	"context"
	"os"
	"testing"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/require"
)

func TestCreateCPAField(t *testing.T) {
	os.Setenv("MM_FEATUREFLAGS_CUSTOMPROFILEATTRIBUTES", "true")
	defer os.Unsetenv("MM_FEATUREFLAGS_CUSTOMPROFILEATTRIBUTES")
	th := Setup(t)
	defer th.TearDown()

	t.Run("a user without admin permissions should not be able to create a field", func(t *testing.T) {
		field := &model.PropertyField{
			Name: model.NewId(),
			Type: model.PropertyFieldTypeText,
		}

		_, resp, err := th.Client.CreateCPAField(context.Background(), field)
		CheckForbiddenStatus(t, resp)
		require.Error(t, err)
	})

	th.TestForSystemAdminAndLocal(t, func(t *testing.T, client *model.Client4) {
		field := &model.PropertyField{
			Name:  model.NewId(),
			Type:  model.PropertyFieldTypeText,
			Attrs: map[string]any{"visibility": "default"},
		}

		createdField, resp, err := client.CreateCPAField(context.Background(), field)
		CheckCreatedStatus(t, resp)
		require.NoError(t, err)
		require.NotZero(t, createdField.ID)
		require.Equal(t, "default", createdField.Attrs["visibility"])
	}, "a user with admin permissions should be able to create the field")
}

func TestListCPAFields(t *testing.T) {
	os.Setenv("MM_FEATUREFLAGS_CUSTOMPROFILEATTRIBUTES", "true")
	defer os.Unsetenv("MM_FEATUREFLAGS_CUSTOMPROFILEATTRIBUTES")
	th := Setup(t)
	defer th.TearDown()

	field := &model.PropertyField{
		Name:  model.NewId(),
		Type:  model.PropertyFieldTypeText,
		Attrs: map[string]any{"visibility": "default"},
	}
	createdField, _, err := th.SystemAdminClient.CreateCPAField(context.Background(), field)
	require.NoError(t, err)
	require.NotNil(t, createdField)

	t.Run("any user should be able to list fields", func(t *testing.T) {
		fields, resp, err := th.Client.ListCPAFields(context.Background())
		CheckOKStatus(t, resp)
		require.NoError(t, err)
		require.NotEmpty(t, fields)
		require.Len(t, fields, 1)
		require.Equal(t, createdField.ID, fields[0].ID)
	})

	t.Run("the endpoint should only list non deleted fields", func(t *testing.T) {
		require.Nil(t, th.App.DeleteCPAField(createdField.ID))
		fields, resp, err := th.Client.ListCPAFields(context.Background())
		CheckOKStatus(t, resp)
		require.NoError(t, err)
		require.Empty(t, fields)
	})
}

func TestPatchCPAField(t *testing.T) {
	os.Setenv("MM_FEATUREFLAGS_CUSTOMPROFILEATTRIBUTES", "true")
	defer os.Unsetenv("MM_FEATUREFLAGS_CUSTOMPROFILEATTRIBUTES")
	th := Setup(t)
	defer th.TearDown()

	t.Run("a user without admin permissions should not be able to patch a field", func(t *testing.T) {
		field := &model.PropertyField{
			Name: model.NewId(),
			Type: model.PropertyFieldTypeText,
		}
		createdField, appErr := th.App.CreateCPAField(field)
		require.Nil(t, appErr)
		require.NotNil(t, createdField)

		patch := &model.PropertyFieldPatch{Name: model.NewPointer(model.NewId())}
		_, resp, err := th.Client.PatchCPAField(context.Background(), createdField.ID, patch)
		CheckForbiddenStatus(t, resp)
		require.Error(t, err)
	})

	th.TestForSystemAdminAndLocal(t, func(t *testing.T, client *model.Client4) {
		field := &model.PropertyField{
			Name: model.NewId(),
			Type: model.PropertyFieldTypeText,
		}
		createdField, appErr := th.App.CreateCPAField(field)
		require.Nil(t, appErr)
		require.NotNil(t, createdField)

		patch := &model.PropertyFieldPatch{Name: model.NewPointer(model.NewId())}
		patchedField, resp, err := client.PatchCPAField(context.Background(), createdField.ID, patch)
		CheckOKStatus(t, resp)
		require.NoError(t, err)
		require.Equal(t, *patch.Name, patchedField.Name)
	}, "a user with admin permissions should be able to patch the field")
}

func TestDeleteCPAField(t *testing.T) {
	os.Setenv("MM_FEATUREFLAGS_CUSTOMPROFILEATTRIBUTES", "true")
	defer os.Unsetenv("MM_FEATUREFLAGS_CUSTOMPROFILEATTRIBUTES")
	th := Setup(t)
	defer th.TearDown()

	t.Run("a user without admin permissions should not be able to delete a field", func(t *testing.T) {
		field := &model.PropertyField{
			Name: model.NewId(),
			Type: model.PropertyFieldTypeText,
		}
		createdField, _, err := th.SystemAdminClient.CreateCPAField(context.Background(), field)
		require.NoError(t, err)
		require.NotNil(t, createdField)

		resp, err := th.Client.DeleteCPAField(context.Background(), createdField.ID)
		CheckForbiddenStatus(t, resp)
		require.Error(t, err)
	})

	th.TestForSystemAdminAndLocal(t, func(t *testing.T, client *model.Client4) {
		field := &model.PropertyField{
			Name: model.NewId(),
			Type: model.PropertyFieldTypeText,
		}
		createdField, _, err := th.SystemAdminClient.CreateCPAField(context.Background(), field)
		require.NoError(t, err)
		require.NotNil(t, createdField)
		require.Zero(t, createdField.DeleteAt)

		resp, err := client.DeleteCPAField(context.Background(), createdField.ID)
		CheckOKStatus(t, resp)
		require.NoError(t, err)

		deletedField, appErr := th.App.GetCPAField(createdField.ID)
		require.Nil(t, appErr)
		require.NotZero(t, deletedField.DeleteAt)
	}, "a user with admin permissions should be able to delete the field")
}
