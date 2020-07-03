package event_service

import (
	"code-test/server/model"
)

func (scrEvent *ScreenResizeEvent) Map() *model.Data {

	dataToReturn := &model.Data{
		WebsiteUrl: scrEvent.WebsiteUrl,
		SessionId:  scrEvent.SessionId,

		ResizeFrom: scrEvent.ResizeFrom,
		ResizeTo:   scrEvent.ResizeTo,
	}

	return dataToReturn
}
