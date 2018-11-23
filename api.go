/*
   Copyright 2018 Sysdig

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

package rhc_manager

import (
	"net/http"
	"fmt"
	"github.com/pkg/errors"
	"encoding/json"
	"bytes"
)

const apiBaseEndpoint = "https://connect.redhat.com/api/v2/"

type ApiClient struct {
	HttpClient http.Client
}

// Retrieves a project information using it's ProjectResponse ID
func (api ApiClient) GetProject(id string) (project ProjectResponse, err error) {
	resp, err := api.HttpClient.Get(apiBaseEndpoint + fmt.Sprintf("projects/%s", id))
	if err != nil {
		err = errors.Wrapf(err, "error while trying to get the project info for %s", id)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		err = errors.Errorf("error code %d: A project with the specified PID (%s) was not found.", resp.StatusCode, id)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&project)
	if err != nil {
		err = errors.Wrapf(err, "error while trying to decode json for project %s", id)
		return
	}

	return
}

// Triggers a build for a project using the specified tag
func (api ApiClient) BuildProject(id, tag string) (response BuildResponse, err error) {

	if tag == "" {
		err = errors.New("tag cannot be empty")
		return
	}

	var requestModel = struct {
		Tag string `json:"tag"`
	}{Tag: tag}

	bdata, err := json.Marshal(requestModel)
	if err != nil {
		errors.Wrapf(err, "error marshaling query to trigger the build (project: %s, tag: %s)", id, tag)
		return
	}
	data := bytes.NewReader(bdata)

	resp, err := api.HttpClient.Post(apiBaseEndpoint+fmt.Sprintf("/projects/%s/build", id), "application/json", data)
	if err != nil {
		errors.Wrapf(err, "error triggering build (project: %s, tag: %s)", id, tag)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 403:
		err = errors.Errorf("the project build service is disabled (project: %s, tag: %s)", id, tag)
		return
	case 404:
		err = errors.Errorf("no build configuration found (project: %s, tag: %s)", id, tag)
		return
	case 428:
		err = errors.Errorf("unable to create build (project: %s, tag: %s)", id, tag)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		err = errors.Wrapf(err, "error while unmarshaling the response (project: %s, tag: %s)", id, tag)
		return
	}

	return

}
