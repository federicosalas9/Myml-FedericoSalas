package myml

import (
	ginGonic "github.com/gin-gonic/gin"
	myml2 "github.com/mercadolibre/myml/src/api/domain/myml"
	"github.com/mercadolibre/myml/src/api/services/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"net/http"
	"strconv"
)

const (
	paramUserID = "userID"
)

func GetInfoC(c *ginGonic.Context) {

	//------------------------------------------------------------------------------------------------------------
	userID := c.Param(paramUserID)
	//convierto el id a entero
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}
	//obtengo el usuario segun el id enviado, como me retorna un apierr compruebo si no es nulo para mostrar
	user, apiErr := myml.GetUserFromAPI(id)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}
	//------------------------------------------GO RUTINES------------------------------------------------------
	cSites := make(chan myml2.Site)            //creo el canal de comunicacion de sitio
	cCategories := make(chan myml2.Categories) //creo el canal de comunicacion de categorias
	var site myml2.Site
	var categories myml2.Categories
	go func() { myml.GetUserSite(user.SiteID, cSites) }()
	go func() { myml.GetSiteCategories(user.SiteID, cCategories) }()
	go func() {
		site = <-cSites            //extraigo la info de sitio cargada en el canal
		categories = <-cCategories //extraigo la info de categorias cargada en el canal

	}()
	//------------------------------------------RESPUESTA JSON----------------------------------------------------
	response := &myml2.Response{
		User:       *user,
		Site:       site,
		Categories: categories,
	}
	c.JSON(http.StatusOK, response)
}
