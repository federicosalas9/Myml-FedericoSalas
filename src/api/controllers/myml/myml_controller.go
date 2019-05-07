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
	//-----------------------------------------------------------------------------------------------------------
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
	wg.Add(5)
	go func() {
		wg.Done()
		cResponse <- response //cargo en el canal el response con el user  modificado
		wg.Done()
		response = <-cResponse //extraigo del canal el response con el user y el site modificado
		wg.Done()
		cResponse <- response //cargo en el canal el response con el user y el site modificado
		wg.Done()
		response = <-cResponse //extraigo del canal el response con el user, el site y las cat modificadas
		wg.Done()
		errors = <-cErrors //si hay errores los cargo en el canal de errores
		if errors != nil {
			c.JSON(errors.Status, errors)
			wg.Done()
			return
		}
	}()
	go func() { myml.GetUserSite(response.User.SiteID, cResponse, cErrors) }()
	go func() { myml.GetSiteCategories(response.User.SiteID, cResponse, cErrors) }()
	wg.Wait()
	//------------------------------------------RESPUESTA JSON----------------------------------------------------
	c.JSON(http.StatusOK, response)
}
