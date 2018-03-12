package store

// Resources - Resource spec
type Resources struct {
	Memory    string   `json:"memory,omitempty"`
	CPU       float32  `json:"cpu,omitempty"`
	Instances int      `json:"instances,omitempty"`
	Volumes   []string `json:"volumes,omitempty"`
}

// Task - task spec
type Task struct {
	Name      string    `json:"name,omitempty"`
	Image     string    `json:"image,omitempty"`
	Command   []string  `json:"command,omitempty"`
	Resources Resources `json:"resources,omitempty"`
	Env       []string  `json:"env,omitempty"`
}
