package container

import "github.com/docker/docker/api/types"

type Container struct {
	ID            string            `json:"id,omitempty"`
	Name          string            `json:"name,omitempty"`
	ImageName     string            `json:"image,omitempty"`
	Ports         []types.Port      `json:"ports,omitempty"`
	Created       int64             `json:"created,omitempty"`
	Status        string            `json:"status,omitempty"`
	State         string            `json:"state,omitempty"`
	Memory        string            `json:"memory,omitempty"`
	RestartPolicy string            `json:"restartPolicy,omitempty"`
	Env           []string          `json:"env,omitempty"`
	AutoRemove    bool              `json:"autoRemove,omitempty"`
	Volumes       []string          `json:"volumes,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	Links         []string          `json:"links,omitempty"`
}
