package event_service

import (
	"code-test/server/model"
)

// this mapping functions isolate the datamode object construction to a single file, if the data model changes, the mapping related changes will happen in the same place

func (scrEvent *ScreenResizeEvent) Map() *model.Data {

	dataToReturn := &model.Data{
		WebsiteUrl: scrEvent.WebsiteUrl,
		SessionId:  scrEvent.SessionId,

		ResizeFrom: scrEvent.ResizeFrom,
		ResizeTo:   scrEvent.ResizeTo,
	}

	return dataToReturn
}

func (timeEvent *TimeTakenEvent) Map() *model.Data {

	dataToReturn := &model.Data{
		WebsiteUrl: timeEvent.WebsiteUrl,
		SessionId:  timeEvent.SessionId,

		FormCompletionTime: timeEvent.FormCompletionTime,
	}

	return dataToReturn
}

func (cpEvent *CopyPasteEvent) Map() *model.Data {

	dataToReturn := &model.Data{
		WebsiteUrl: cpEvent.WebsiteUrl,
		SessionId:  cpEvent.SessionId,

		CopyAndPaste: map[string]bool{
			cpEvent.FormId: cpEvent.Pasted,
		},
	}

	return dataToReturn
}
