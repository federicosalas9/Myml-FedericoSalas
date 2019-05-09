package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
)

const (
	//urlUsers = "https://api.mercadolibre.com/users/"
	//urlSites="https://api.mercadolibre.com/sites/"
	urlUsers = "http://localhost:8082/user/"
	urlSites = "http://localhost:8082/sites/"
)

func (users *User) GetU() *apierrors.ApiError {
	if users.ID == 0 {
		return &apierrors.ApiError{
			Message: "userID is empty",
			Status:  http.StatusBadRequest,
		}
	}

	final := fmt.Sprintf("%s%d", urlUsers, users.ID)
	response, err := http.Get(final)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal([]byte(data), &users); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}

func (site *Site) GetS() *apierrors.ApiError {
	if site.ID == "" {
		return &apierrors.ApiError{
			Message: "siteID is empty",
			Status:  http.StatusBadRequest,
		}
	}
	final := fmt.Sprintf("%s%s", urlSites, site.ID)
	response, err := http.Get(final)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal([]byte(data), &site); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}

func (categories *Categories) GetC(siteID string) *apierrors.ApiError {
	if siteID == "" {
		return &apierrors.ApiError{
			Message: "siteID is empty",
			Status:  http.StatusBadRequest,
		}
	}
	final := fmt.Sprintf("%s%s%s", urlSites, siteID, "/categories")
	response, err := http.Get(final)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal([]byte(data), &categories); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}
