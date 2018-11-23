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
	"testing"
)

func TestApiClient_GetProject(t *testing.T) {
	cases := []struct {
		pid                 string
		expectedErrorString string
	}{
		{pid: "pNotCorrectId",  expectedErrorString: "error code 404: A project with the specified PID (pNotCorrectId) was not found."},
		{pid: "pNotCorrect/id", expectedErrorString: "error code 404: A project with the specified PID (pNotCorrect/id) was not found."},
	}

	client := ApiClient{}

	for _, c := range cases {
		_, err := client.GetProject(c.pid)
		if err == nil {
			t.Errorf("this call should return an error, and it doesn't (pid %s)", c.pid)
		}
		if err.Error() !=  c.expectedErrorString {
			t.Errorf(`incorrect error message "%s", expected "%s"`, err.Error(), c.expectedErrorString)
		}
	}
}

func TestApiClient_BuildProject(t *testing.T) {

	cases := []struct {
		pid                 string
		tag                 string
		expectedErrorString string
	}{
		{pid: "pNotCorrectId", tag: "", expectedErrorString: "tag cannot be empty"},
		{pid: "pNotCorrectId", tag: "0.0.0-testing", expectedErrorString: "no build configuration found (project: pNotCorrectId, tag: 0.0.0-testing)"},
	}

	client := ApiClient{}

	for _, c := range cases {
		_, err := client.BuildProject(c.pid, c.tag)
		if err == nil {
			t.Fatalf("this call should return an error, and it doesn't (pid: %s, tag: %s, err: %s)", c.pid, c.tag, c.expectedErrorString)
		}

		if err.Error() != c.expectedErrorString {
			t.Errorf(`incorrect error message "%s" expected "%s" (pid: %s, tag: %s)`, err.Error(), c.expectedErrorString, c.pid, c.tag)
		}
	}
}

func TestProject_String(t *testing.T) {
	tags := []struct {
		Digest              string      `json:"digest"`
		Name                string      `json:"name"`
		HealthIndex         string      `json:"health_index"`
		Published           bool        `json:"published"`
		ScanResults         string      `json:"scan_results"`
		ScanStatus          string      `json:"scan_status"`
		ApplicationPlaybook interface{} `json:"application_playbook"`
	}{
		{Name: "0.0.1", HealthIndex: "B", Published: false, ScanStatus: "Passing"},
	}

	project := Project{
		Rebuild:            "Recommended",
		LatestPublishedTag: "0.1.0-testing",
		RepoHealthIndex:    "B",
		AutoRebuild:        false,
		AutoPublish:        true,
		Tags:               tags,
	}

	expectedString := `Project:
  Rebuild:            Recommended
  Health Index:       B
  Last Published Tag: 0.1.0-testing
  Auto Rebuild:       false
  Auto Publish:       true
  Tags: 
    Name:         0.0.1
    Health Index: B
    Published:    false
    Scan Status:  Passing
`

	if project.String() != expectedString {
		t.Error("project string format is not correct")
	}
}
