package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Dadard29/geopolitics/api"
	"github.com/Dadard29/geopolitics/managers"
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/go-api-utils/auth"
	"io/ioutil"
	"net/http"
)

// POST
// Authorization: 	token
// Params: 			None
// Body: 			None
func RelationshipPost(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	from := r.URL.Query().Get("from")
	if from == "" {
		api.Api.BuildMissingParameter(w)
		return
	}

	to := r.URL.Query().Get("to")
	if to == "" {
		api.Api.BuildMissingParameter(w)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusBadRequest, "invalid body", w)
		return
	}

	var rel models.RelationshipInput
	err = json.Unmarshal(body, &rel)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusBadRequest, "invalid json body", w)
		return
	}

	if err := rel.CheckSanity(); err != nil {
		logger.Error(err.Error())
		api.Api.BuildMissingParameter(w)
		return
	}

	out, err := managers.RelationshipManagerCreate(rel, from, to)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, "failed to create the relationship", w)
		return
	}

	api.Api.BuildJsonResponse(true, "relationship created", out, w)
}

// GET
// Authorization: 	token
// Params: 			country
// Body: 			None
func RelationshipGet(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	country := r.URL.Query().Get("country")
	if country == "" {
		api.Api.BuildMissingParameter(w)
		return
	}

	rel, err := managers.RelationshipManagerGet(country)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, "couldn't process request", w)
		return
	}

	api.Api.BuildJsonResponse(true, fmt.Sprintf("relationships for country %s retrieved", country), rel, w)
}

// GET
// Authorization: 	token
// Params: 			None
// Body: 			None
func RelationshipAllGet(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	// TODO

	api.Api.BuildJsonResponse(true, "ok", nil, w)
}
