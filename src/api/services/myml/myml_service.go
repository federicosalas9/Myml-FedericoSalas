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

func GetUserSite(siteID string, c chan myml.Response, cErrors chan *apierrors.ApiError) {

	site := &myml.Site{
		ID: siteID,
	}
	if apiErr := site.GetS(); apiErr != nil {
		cErrors <- apiErr
	}
	response := <-c
	response.Site = *site
	c <- response
	cErrors <- nil
}

func GetSiteCategories(siteID string, c chan myml.Response, cErrors chan *apierrors.ApiError) {
	categories := &myml.Categories{}
	categories.GetC(siteID)
	if apiErr := categories.GetC(siteID); apiErr != nil {
		cErrors <- apiErr
	}
	response := <-c
	response.Categories = *categories
	c <- response
	cErrors <- nil
}
