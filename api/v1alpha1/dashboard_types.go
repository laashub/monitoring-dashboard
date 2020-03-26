/*
Copyright 2020 The KubeSphere authors.

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

package v1alpha1

import (
	"github.com/monitoring-dashboard/api/v1alpha1/panels"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DashboardSpec defines the desired state of Dashboard
type DashboardSpec struct {
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	Time        Time         `json:"time,omitempty"`
	Datasource  string       `json:"datasource,omitempty"`
	Panels      []Panel      `json:"panels,omitempty"`
	Templatings []Templating `json:"templatings,omitempty"`
}

type Time struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type Panel struct {
	Row        panels.Row        `json:",inline"`
	Graph      panels.Graph      `json:",inline"`
	SingleStat panels.SingleStat `json:",inline"`
	Table      panels.Table      `json:",inline"`
}

type Templating struct {
	Name    string   `json:"name,omitempty"`
	Label   string   `json:"label,omitempty"`
	Options []Option `json:"options,omitempty"`
	Query   string   `json:"query,omitempty"`
	Type    string   `json:"type,omitempty"`
	Sort    int      `json:"sort,omitempty"`
}

type Option struct {
	Selected bool   `json:"selected,omitempty"`
	Text     string `json:"text,omitempty"`
	Value    string `json:"value,omitempty"`
}

// +kubebuilder:object:root=true

// Dashboard is the Schema for the dashboards API
type Dashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DashboardSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// DashboardList contains a list of Dashboard
type DashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dashboard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Dashboard{}, &DashboardList{})
}
