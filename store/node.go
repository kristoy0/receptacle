package store

type Nodes struct {
	Hosts []struct {
		ID      string `json:"id"`
		Node    string `json:"node"`
		Address string `json:"address"`
	}
}
