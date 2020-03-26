// +kubebuilder:object:generate=true

package panels

type Row struct {
	Title  string  `json:"title,omitempty"`
	Type   string  `json:"type,omitempty"`
	Panels []Panel `json:"panels,omitempty"`
}

type Panel struct {
	Graph Graph `json:",inline"`
	Table Table `json:",inline"`
}
