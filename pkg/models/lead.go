package models

type AuditDetails struct {
	HTTPStatus       int      `json:"http_status"`
	HasSSL           bool     `json:"has_ssl"`
	IsMobileFriendly bool     `json:"is_mobile_friendly"`
	TechStack        []string `json:"tech_stack"`
	ErrorMessage     string   `json:"error_message,omitempty"`
}

type Lead struct {
	OsmID    int64        `json:"osm_id"`
	Name     string       `json:"name"`
	Phone    string       `json:"phone"`
	Website  string       `json:"website"`
	Street   string       `json:"street"`
	City     string       `json:"city"`
	Category string       `json:"category"`
	Tier     string       `json:"tier"`
	Audit    AuditDetails `json:"audit"`
}

func (l Lead) DisplayPhone() string {
	if l.Phone != "" {
		return l.Phone
	}
	return "(555) 019-2831"
}

func (l Lead) FormattedAddress() string {
	if l.Street != "" && l.City != "" {
		return l.Street + ", " + l.City
	}
	if l.Street != "" {
		return l.Street
	}
	if l.City != "" {
		return l.City
	}
	return "Metropolitan Area"
}
