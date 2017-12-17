package container

type Container struct {
	Name          string            `json:"name,omitempty"`
	ImageName     string            `json:"image,omitempty"`
	Ports         []string          `json:"ports,omitempty"`
	Memory        string            `json:"memory,omitempty"`
	RestartPolicy string            `json:"restartPolicy,omitempty"`
	Env           []string          `json:"env,omitempty"`
	AutoRemove    bool              `json:"autoRemove,omitempty"`
	Volumes       []string          `json:"volumes,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	Links         []string          `json:"links,omitempty"`
}