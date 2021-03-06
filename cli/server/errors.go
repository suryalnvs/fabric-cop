/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cloudflare/cfssl/api"
	cerr "github.com/cloudflare/cfssl/errors"
	"github.com/cloudflare/cfssl/log"
	cop "github.com/hyperledger/fabric-cop/api"
)

var (
	errNoAuthHdr           = errors.New("No Authorization header was found")
	errNoBasicAuthHdr      = errors.New("No Basic Authorization header was found")
	errNoTokenAuthHdr      = errors.New("No Token Authorization header was found")
	errBasicAuthNotAllowed = errors.New("Basic authorization is not permitted")
	errTokenAuthNotAllowed = errors.New("Token authorization is not permitted")
	errInvalidUserPass     = errors.New("Invalid user name or password")
	errInputNotSeeker      = errors.New("Input stream was not a seeker")
)

func badRequest(w http.ResponseWriter, err error) error {
	return httpError(w, 400, cop.IOError, err.Error())
}

func authErr(w http.ResponseWriter, err error) error {
	return httpError(w, 401, cop.AuthorizationFailure, err.Error())
}

func notFound(w http.ResponseWriter, err error) error {
	return httpError(w, 404, cerr.RecordNotFound, err.Error())
}

func dbErr(w http.ResponseWriter, err error) error {
	return httpError(w, 500, cop.IOError, err.Error())
}

func httpError(w http.ResponseWriter, scode, code int, msg string) error {
	response := api.NewErrorResponse(msg, code)
	jsonMessage, err := json.Marshal(response)
	if err != nil {
		log.Errorf("Failed to marshal JSON: %v", err)
	} else {
		msg = string(jsonMessage)
	}
	http.Error(w, msg, scode)
	return nil
}
