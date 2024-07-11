package iot

type IotResponseStatus struct {
	RequestID string `json:"request_id" db:"request_id"`
	Status    string `json:"status" db:"status"`
}

type IotInfoResponse struct {
	IotResponseStatus
	Rooms      []Room      `json:"rooms"`
	Groups     []InfoGroup `json:"groups"`
	Devices    []Device    `json:"devices"`
	Scenarios  []Scenario  `json:"scenarios"`
	Households []Household `json:"households"`
}

type IotDeviceResponse struct {
	IotResponseStatus
	Device
}

type IotGroupResponse struct {
	IotResponseStatus
	Group
}

type Room struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	HouseholdId string   `json:"household_id"`
	Devices     []string `json:"devices"`
}
type Group struct {
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	Aliases      []string     `json:"aliases"`
	Type         string       `json:"type"`
	Devices      []Device     `json:"devices"`
	Capabilities []Capability `json:"capabilities"`
	Actions      []Capability `json:"actions"`
}
type InfoGroup struct {
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	Aliases      []string     `json:"aliases"`
	HouseholdId  string       `json:"household_id"`
	Type         string       `json:"type"`
	Devices      []string     `json:"devices"`
	Capabilities []Capability `json:"capabilities"`
}
type Capability struct {
	Reportable  bool        `json:"reportable"`
	Retrievable bool        `json:"retrievable"`
	Type        string      `json:"type"`
	Parameters  interface{} `json:"parameters"`
	State       interface{} `json:"state"`
}

type Device struct {
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	Aliases      []string     `json:"aliases"`
	Type         string       `json:"type"`
	State        string       `json:"state"`
	ExternalId   string       `json:"external_id"`
	SkillId      string       `json:"skill_id"`
	HouseholdId  string       `json:"household_id"`
	Room         string       `json:"room"`
	Groups       []string     `json:"groups"`
	Capabilities []Capability `json:"capabilities"`
	Properties   []Property   `json:"properties"`
	Actions      []Capability `json:"actions"`
	CustomData   interface{}  `json:"custom_data"`
}

type Property struct {
	Retrievable bool        `json:"retrievable"`
	Type        string      `json:"type"`
	Parameters  interface{} `json:"parameters"`
	State       interface{} `json:"state"`
	LastUpdated float64     `json:"last_updated"`
}
type Scenario struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
type Household struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
