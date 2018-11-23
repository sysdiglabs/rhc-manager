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
	"time"
	"fmt"
)

type ProjectResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
	Code    int         `json:"code"`
	Project Project     `json:"data"`
}

type Project struct {
	Rebuild               string      `json:"rebuild"`
	RepoHealthIndex       string      `json:"repo_health_index"`
	LatestPublishedDigest string      `json:"latest_published_digest"`
	LatestPublishedTag    string      `json:"latest_published_tag"`
	ProjectType           interface{} `json:"project_type"`
	BuildService          bool        `json:"build_service"`
	AutoRebuild           bool        `json:"auto_rebuild"`
	AutoPublish           bool        `json:"auto_publish"`
	Tags []struct {
		Digest              string      `json:"digest"`
		Name                string      `json:"name"`
		HealthIndex         string      `json:"health_index"`
		Published           bool        `json:"published"`
		ScanResults         string      `json:"scan_results"`
		ScanStatus          string      `json:"scan_status"`
		ApplicationPlaybook interface{} `json:"application_playbook"`
	} `json:"tags"`
}

func (p Project) String() string {
	header := fmt.Sprintf(`Project:
  Rebuild:            %s
  Health Index:       %s
  Last Published Tag: %s
  Auto Rebuild:       %t
  Auto Publish:       %t
`, p.Rebuild, p.RepoHealthIndex, p.LatestPublishedTag, p.AutoRebuild, p.AutoPublish)

	tagHeader := "  Tags: "
	tags := ""
	for _, tag := range p.Tags {
		tags = tags + fmt.Sprintf(`
    Name:         %s
    Health Index: %s
    Published:    %t
    Scan Status:  %s
`, tag.Name, tag.HealthIndex, tag.Published, tag.ScanStatus)
	}

	return header + tagHeader + tags
}

type BuildResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Build   Build  `json:"data"`
}

type Build struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata struct {
		Name              string    `json:"name"`
		Namespace         string    `json:"namespace"`
		SelfLink          string    `json:"selfLink"`
		UID               string    `json:"uid"`
		ResourceVersion   string    `json:"resourceVersion"`
		CreationTimestamp time.Time `json:"creationTimestamp"`
		Labels struct {
			Buildconfig                 string `json:"buildconfig"`
			OpenshiftIoBuildConfigName  string `json:"openshift.io/build-config.name"`
			OpenshiftIoBuildStartPolicy string `json:"openshift.io/build.start-policy"`
		} `json:"labels"`
		Annotations struct {
			OpenshiftIoBuildConfigName string `json:"openshift.io/build-config.name"`
			OpenshiftIoBuildNumber     string `json:"openshift.io/build.number"`
		} `json:"annotations"`
		OwnerReferences []struct {
			APIVersion string `json:"apiVersion"`
			Kind       string `json:"kind"`
			Name       string `json:"name"`
			UID        string `json:"uid"`
			Controller bool   `json:"controller"`
		} `json:"ownerReferences"`
	} `json:"metadata"`
	Spec struct {
		ServiceAccount string `json:"serviceAccount"`
		Source struct {
			Type string `json:"type"`
			Git struct {
				URI string `json:"uri"`
				Ref string `json:"ref"`
			} `json:"git"`
			SourceSecret struct {
				Name string `json:"name"`
			} `json:"sourceSecret"`
		} `json:"source"`
		Strategy struct {
			Type string `json:"type"`
			DockerStrategy struct {
				ForcePull      bool   `json:"forcePull"`
				DockerfilePath string `json:"dockerfilePath"`
			} `json:"dockerStrategy"`
		} `json:"strategy"`
		Output struct {
			To struct {
				Kind string `json:"kind"`
				Name string `json:"name"`
			} `json:"to"`
			PushSecret struct {
				Name string `json:"name"`
			} `json:"pushSecret"`
		} `json:"output"`
		Resources    []interface{} `json:"resources"`
		PostCommit   []interface{} `json:"postCommit"`
		NodeSelector interface{}   `json:"nodeSelector"`
		TriggeredBy []struct {
			Message string `json:"message"`
		} `json:"triggeredBy"`
	} `json:"spec"`
	Status struct {
		Phase string `json:"phase"`
		Config struct {
			Kind      string `json:"kind"`
			Namespace string `json:"namespace"`
			Name      string `json:"name"`
		} `json:"config"`
		Output []interface{} `json:"output"`
	} `json:"status"`
}
