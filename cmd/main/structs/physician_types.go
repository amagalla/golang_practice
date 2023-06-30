package structs

type InsertPhysiciansRequestData struct {
	PhysicianFirstName string `json:"physician_first_name"`
	PhysicianLastName string `json:"physician_last_name"`
}

type PhysicianListResponse struct {
	PhysicianList []Physicians `json:"results"`
}

type Physicians struct {
	PhysicianID   uint64 `json:"physician_id"`
	FirstName string `json:"physician_first_name"`
	LastName string `json:"physician_last_name"`
}
