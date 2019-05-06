package myml

import (
	"github.com/mercadolibre/myml/src/api/domain/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
)

func GetUserFromAPI(userID int64) (*myml.User, *apierrors.ApiError) {

	user := &myml.User{
		ID: userID,
	}
	if apiErr := user.GetU(); apiErr != nil {
		return nil, apiErr
	}
	return user, nil
}

func GetUserSite(siteID string, c chan myml.Site, cErrors chan *apierrors.ApiError) {

	site := &myml.Site{
		ID: siteID,
	}
	if apiErr := site.GetS(); apiErr != nil {
		cErrors <- apiErr
	}
	c <- *site
	cErrors <- nil
}

func GetSiteCategories(siteID string, c chan myml.Categories) {
	categories := &myml.Categories{}
	categories.GetC(siteID)
	if apiErr := categories.GetC(siteID); apiErr != nil {
		//--
	}
	c <- *categories
}
