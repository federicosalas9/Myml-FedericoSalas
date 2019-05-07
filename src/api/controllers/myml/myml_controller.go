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
	response := myml2.Response{}
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
	response.User = *user
	//------------------------------------------GO RUTINES------------------------------------------------------
	cResponse := make(chan myml2.Response)
	cErrors := make(chan *apierrors.ApiError)
	var errors *apierrors.ApiError
	var wg sync.WaitGroup
	go func() { myml.GetUserSite(user.SiteID, cResponse, cErrors) }()
	go func() { myml.GetSiteCategories(user.SiteID, cResponse, cErrors) }()
	wg.Add(5)
	go func() {
		cResponse <- response
		wg.Done()
		response = <-cResponse
		wg.Done()
		cResponse <- response
		wg.Done()
		response = <-cResponse
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
	c.JSON(http.StatusOK, response)
}
