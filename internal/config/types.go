package config

// ProvisionRequest represents the provision request structure
type ProvisionRequest struct {
	NodeName  string `json:"nodeName"`
	NodeUID   string `json:"nodeUid"`
	NodeType  string `json:"nodeType"`
	Token     string `json:"token,omitempty"`
	MasterIP  string `json:"masterIP,omitempty"`
	CPUs      int    `json:"cpus"`
	DiskSize  string `json:"diskSize"`
	Memory    string `json:"memory"`
	ImageOCI  string `json:"imageOci"`
	EnableSSH bool   `json:"enableSsh"`
}

// ProvisionResponse represents the API response
type ProvisionResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	NodeID   string `json:"nodeId,omitempty"`
	Error    string `json:"error,omitempty"`
	MasterIP string `json:"masterIP,omitempty"`
}
