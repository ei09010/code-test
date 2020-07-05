package repository

import (
	"code-test/server/model"
	"code-test/server/repository"
	"testing"
)

func TestInitUserSession_initSuccess(t *testing.T) {
	// Arrange

	expectedWebSiteUrl := "https://ravelin.com"
	expectedSessionId := "1235"

	sessionDataStore := repository.GetInstance()

	expectedResult := &model.Data{
		WebsiteUrl: expectedWebSiteUrl,
		SessionId:  expectedSessionId,
		CopyAndPaste: map[string]bool{
			"": false,
		},
	}

	// Act

	res, err := sessionDataStore.InitUserSession(expectedSessionId, expectedWebSiteUrl)

	// Assert

	if expectedResult.SessionId != res.SessionId {
		t.Errorf("Expected %v, got %v", expectedResult.SessionId, res.SessionId)
	}

	if expectedResult.WebsiteUrl != res.WebsiteUrl {
		t.Errorf("Expected %v, got %v", expectedResult.WebsiteUrl, res.WebsiteUrl)
	}

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

}

func TestSave_saveSuccess(t *testing.T) {
	// Arrange

	expectedWebSiteUrl := "https://ravelin.com"
	expectedSessionId := "1235"

	sessionDataStore := repository.GetInstance()

	dataToSave := &model.Data{
		WebsiteUrl: expectedWebSiteUrl,
		SessionId:  expectedSessionId,
		CopyAndPaste: map[string]bool{
			"": false,
		},
	}

	// Act

	err := sessionDataStore.Save(dataToSave)

	// Assert

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

	// ideally I would access directly to the object storage to check the storage status, but the compiler seems to think that I'm in a different package
	dataStored, err := sessionDataStore.Get(dataToSave.SessionId, dataToSave.WebsiteUrl)

	if dataStored == nil {
		t.Errorf("Expected %v, got %v", dataToSave, nil)
	}

	if dataToSave.WebsiteUrl != dataStored.WebsiteUrl {
		t.Errorf("Expected %v, got %v", dataToSave.WebsiteUrl, dataStored.WebsiteUrl)
	}

	if dataToSave.SessionId != dataStored.SessionId {
		t.Errorf("Expected %v, got %v", dataToSave.SessionId, dataStored.SessionId)
	}
}

func TestGet_getSuccess(t *testing.T) {
	// Arrange

	expectedWebSiteUrl := "https://ravelin.com"
	expectedSessionId := "1235"

	sessionDataStore := repository.GetInstance()

	dataToSave := &model.Data{
		WebsiteUrl: expectedWebSiteUrl,
		SessionId:  expectedSessionId,
		CopyAndPaste: map[string]bool{
			"": false,
		},
	}

	err := sessionDataStore.Save(dataToSave)

	if err != nil {
		t.Errorf("Error arranging test scenario: %v, got %v", nil, err)
	}

	// Act

	dataStored, err := sessionDataStore.Get(dataToSave.SessionId, dataToSave.WebsiteUrl)

	// Assert

	if dataStored == nil {
		t.Errorf("Expected %v, got %v", dataToSave, nil)
	}

	if dataToSave.WebsiteUrl != dataStored.WebsiteUrl {
		t.Errorf("Expected %v, got %v", dataToSave.WebsiteUrl, dataStored.WebsiteUrl)
	}

	if dataToSave.SessionId != dataStored.SessionId {
		t.Errorf("Expected %v, got %v", dataToSave.SessionId, dataStored.SessionId)
	}
}

func TestUpdate_updateWithTimeSuccess(t *testing.T) {
	// Arrange

	expectedWebSiteUrl := "https://ravelin.com"
	expectedSessionId := "1235"

	sessionDataStore := repository.GetInstance()

	savedData := &model.Data{
		WebsiteUrl: expectedWebSiteUrl,
		SessionId:  expectedSessionId,
		Time:       0,
		CopyAndPaste: map[string]bool{
			"": false,
		},
	}

	err := sessionDataStore.Save(savedData)

	if err != nil {
		t.Errorf("Error arranging test scenario: %v, got %v", nil, err)
	}

	updateData := &model.Data{
		WebsiteUrl: expectedWebSiteUrl,
		SessionId:  expectedSessionId,
		Time:       12,
		CopyAndPaste: map[string]bool{
			"": false,
		},
	}

	// Act

	dataStored, err := sessionDataStore.Update(updateData)

	// Assert

	if dataStored == nil {
		t.Errorf("Expected %v, got %v", updateData, nil)
	}

	if updateData.Time != dataStored.Time {
		t.Errorf("Expected %v, got %v", updateData.Time, dataStored.Time)
	}

	if updateData.SessionId != dataStored.SessionId {
		t.Errorf("Expected %v, got %v", updateData.SessionId, dataStored.SessionId)
	}

	if updateData.WebsiteUrl != dataStored.WebsiteUrl {
		t.Errorf("Expected %v, got %v", updateData.SessionId, dataStored.SessionId)
	}
}

func TestUpdate_noObjectToUpdate_updateError(t *testing.T) {
	// Arrange

	expectedWebSiteUrl := "https://ravelin.com"
	expectedSessionId := "5678"

	sessionDataStore := repository.GetInstance()

	updateData := &model.Data{
		WebsiteUrl: expectedWebSiteUrl,
		SessionId:  expectedSessionId,
		Time:       12,
		CopyAndPaste: map[string]bool{
			"": false,
		},
	}

	// Act

	dataStored, err := sessionDataStore.Update(updateData)

	// Assert

	if err == nil {
		t.Errorf("Expected %v, got %v", err, nil)
	}
	if dataStored != nil {
		t.Errorf("Expected %v, got %v", nil, dataStored)
	}

}

func TestUpdate_updateWithAlreadyUpdateTime_UpdateCopyPasteSuccess(t *testing.T) {
	// Arrange

	expectedWebSiteUrl := "https://ravelin.com"
	expectedSessionId := "78945"
	oldTime := 12

	sessionDataStore := repository.GetInstance()

	savedData := &model.Data{
		WebsiteUrl: expectedWebSiteUrl,
		SessionId:  expectedSessionId,
		Time:       oldTime,
		CopyAndPaste: map[string]bool{
			"": false,
		},
	}

	err := sessionDataStore.Save(savedData)

	if err != nil {
		t.Errorf("Error arranging test scenario: %v, got %v", nil, err)
	}

	newTime := 15
	newFormId := "formId1"
	pasted := true

	updateData := &model.Data{
		WebsiteUrl: expectedWebSiteUrl,
		SessionId:  expectedSessionId,
		Time:       newTime,
		CopyAndPaste: map[string]bool{
			newFormId: pasted,
		},
	}

	// Act

	dataStored, err := sessionDataStore.Update(updateData)

	// Assert

	if dataStored == nil {
		t.Errorf("Expected %v, got %v", updateData, nil)
	}

	if updateData.Time == oldTime {
		t.Errorf("Expected %v, got %v", updateData.Time, savedData.Time)
	}

	if updateData.CopyAndPaste[newFormId] != dataStored.CopyAndPaste[newFormId] {
		t.Errorf("Expected %v, got %v", updateData.CopyAndPaste[newFormId], dataStored.CopyAndPaste[newFormId])
	}

	if _, ok := dataStored.CopyAndPaste[""]; ok {
		t.Errorf("Expected %v, got %v", "not available", dataStored.CopyAndPaste[""])
	}

	if updateData.SessionId != dataStored.SessionId {
		t.Errorf("Expected %v, got %v", updateData.SessionId, dataStored.SessionId)
	}

	if updateData.WebsiteUrl != dataStored.WebsiteUrl {
		t.Errorf("Expected %v, got %v", updateData.SessionId, dataStored.SessionId)
	}
}
