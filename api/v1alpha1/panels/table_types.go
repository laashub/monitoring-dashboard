// +kubebuilder:object:generate:=true

package panels

type Table struct {
	BasePanel
	Sort   Sort    `json:"type,omitempty"`
	Styles []Style `json:"style,omitempty"`
}

type Sort struct {
	Col  string `json:"col,omitempty"`
	Desc bool   `json:"desc,omitempty"`
}

type Style struct {
	Alias    string `json:"alias,omitempty"`
	Pattern  string `json:"pattern,omitempty"`
	Type     string `json:"type,omitempty"`
	Unit     string `json:"unit,omitempty"`
	Decimals int    `json:"decimals,omitempty"`
}
