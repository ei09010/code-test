package event_service

import (
	"code-test/server/model"
	"code-test/server/repository"
	"log"
)

// this mapping functions isolate the datamodel object construction to a single file - if the data model changes, the mapping related changes will happen in the same place
// also, status validation related to the event in database to which I am mapping to
func (scrEvent *ScreenResizeEvent) Map() (*model.Data, error) {

	dataToReturn, err := repository.SessionsData.Get(scrEvent.SessionId, scrEvent.WebsiteUrl)

	if err != nil {
		log.Println(errorRetrievingObjectToMap, "with error", err)
		return nil, err
	}

	// Since only one re-size happens, I'm assuming that if already stored resize data is empty, we can override with valid (non zero-value) received resize data
	dataToReturnResizeFromInvalid := dataToReturn.ResizeFrom.Height == "" && dataToReturn.ResizeFrom.Width == ""

	if dataToReturnResizeFromInvalid {
		dataToReturn.ResizeFrom = scrEvent.ResizeFrom
	}

	dataToReturnResizeToInvalid := dataToReturn.ResizeTo.Height == "" && dataToReturn.ResizeTo.Width == ""

	if dataToReturnResizeToInvalid {
		dataToReturn.ResizeTo = scrEvent.ResizeTo
	}

	return dataToReturn, nil
}

func (timeEvent *TimeTakenEvent) Map() (*model.Data, error) {

	dataToReturn, err := repository.SessionsData.Get(timeEvent.SessionId, timeEvent.WebsiteUrl)

	if err != nil {
		log.Println(errorRetrievingObjectToMap, "with error", err)
		return nil, err
	}

	dataToReturn.Time = timeEvent.Time

	return dataToReturn, nil
}

func (cpEvent *CopyPasteEvent) Map() (*model.Data, error) {

	dataToReturn, err := repository.SessionsData.Get(cpEvent.SessionId, cpEvent.WebsiteUrl)

	if err != nil {
		log.Println(errorRetrievingObjectToMap, "with error", err)
		return nil, err
	}

	// store copy paste events
	// Given that the paste operation will only change from false to true once, I'm only adding to the dictionary

	if _, ok := dataToReturn.CopyAndPaste[""]; ok {
		delete(dataToReturn.CopyAndPaste, "")
	}

	dataToReturn.CopyAndPaste[cpEvent.FormId] = cpEvent.Pasted

	return dataToReturn, nil
}
