package image

type Image struct {
	Name string `json:"name,omitempty"`
	Id   string `json:"id,omitempty"`
	Size string `json:"size,omitempty"`
}
