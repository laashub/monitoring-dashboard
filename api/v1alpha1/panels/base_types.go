// +kubebuilder:object:generate=true

package panels

type BasePanel struct {
	Title   string   `json:"title,omitempty"`
	Type    string   `json:"type,omitempty"`
	Targets []Target `json:"targets,omitempty"`
}

type Target struct {
	Expression   string `json:"expr,omitempty"`
	LegendFormat string `json:"legendFormat,omitempty"`
	Instant      bool   `json:"instant,omitempty"`
}
