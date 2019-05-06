package myml

import (
	ginGonic "github.com/gin-gonic/gin"
	myml2 "github.com/mercadolibre/myml/src/api/domain/myml"
	"github.com/mercadolibre/myml/src/api/services/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"net/http"
	"strconv"
	"sync"
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
	cErrors := make(chan *apierrors.ApiError)
	var site myml2.Site
	var categories myml2.Categories
	var errors *apierrors.ApiError
	var wg sync.WaitGroup
	go func() { myml.GetUserSite(user.SiteID, cSites, cErrors) }()
	go func() { myml.GetSiteCategories(user.SiteID, cCategories, cErrors) }()
	wg.Add(3)
	go func() {
		site = <-cSites //extraigo la info de sitio cargada en el canal
		wg.Done()
		categories = <-cCategories //extraigo la info de categorias cargada en el canal
		wg.Done()
		errors = <-cErrors
		wg.Done()
	}()
	wg.Wait()
	//------------------------------------------RESPUESTA JSON----------------------------------------------------

	if errors != nil {
		c.JSON(errors.Status, errors)
		return
	}

	response := &myml2.Response{
		User:       *user,
		Site:       site,
		Categories: categories,
	}
	c.JSON(http.StatusOK, response)
}
