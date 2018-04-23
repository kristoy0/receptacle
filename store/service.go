package store

type Service struct {
	List []struct {
		Host  string `json:"host,omitempty"`
		Name  string `json:"name,omitempty"`
		Image string `json:"image,omitempty"`
		Port  string `json:"port,omitempty"`
		IP    string `json:"ip,omitempty"`
	}
}
