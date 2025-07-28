package handlers

type CpuResponse struct {
	EntityType string `json:"type"`
	Payload    int    `json:"payload"`
}



type IfInfo struct {
	Name       string `json:"name"`
	IpAddress  string `json:"ip"`
}

type AddrIpV4Response struct {
	EntityType string `json:"type"`
	Count    string    `json:"count"`
	Payload  []*IfInfo    `json:"payload"`
}


type DateTimeNow struct {
	EntityType string `json:"type"`
	Payload    string    `json:"payload"`
}
