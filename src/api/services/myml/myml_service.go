package myml

import (
	"github.com/mercadolibre/Myml - Federico Salas/src/api/domain/myml"
	"github.com/mercadolibre/Myml - Federico Salas/src/api/utils/apierrors"
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

	response := <-c //extraigo del canal el response con el user modificado
	response.Site = *site
	c <- response //modifico el site y devuelvo al canal el response que ahora tiene el user y el site modificado
	cErrors <- nil
}

func GetSiteCategories(siteID string, c chan myml.Response, cErrors chan *apierrors.ApiError) {
	categories := &myml.Categories{}
	categories.GetC(siteID)
	if apiErr := categories.GetC(siteID); apiErr != nil {
		cErrors <- apiErr
	}
	response := <-c //extraigo del canal el response con el user y el site modificado
	response.Categories = *categories
	c <- response //modifico las categorias y devuelvo al canal el response que ahora tiene el user, el site y las cat. modificadas
	cErrors <- nil
}
