package models

type Resident struct {
	NIC     string `json:"nic"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
