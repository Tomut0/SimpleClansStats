package domain

type DashboardLayout struct {
	Version int               `json:"version"`
	Widgets []DashboardWidget `json:"widgets"`
}

type DashboardWidget struct {
	ID     string `json:"id"`
	Type   string `json:"type,omitempty"`
	Field  string `json:"field,omitempty"`
	Format string `json:"format,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
}
