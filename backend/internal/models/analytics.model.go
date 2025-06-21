package models

type AnalyticsEventRequest struct {
	Event      string  `json:"event"`                 // event type
	Username   string  `json:"username,omitempty"`    // for registered user
	SessionID  string  `json:"session_id,omitempty"`  // for anonymous sessions
	Path       string  `json:"path,omitempty"`        // page_visited
	Referrer   string  `json:"referrer,omitempty"`    // page_visited
	RecipeID   int32   `json:"recipe_id,omitempty"`   // recipe_opened
	Device     string  `json:"device"`                // device info
	IPAddress  string  `json:"-"`                     // server will add (for login)
}

type PageVisitedRecord struct {
	SessionID  string `json:"session_id"`
	Username   string `json:"username,omitempty"`
	Path       string `json:"path"`
	VisitTime  string `json:"visit_time"` // RFC3339 timestamp
	Referrer   string `json:"referrer,omitempty"`
	Device     string `json:"device"`
}

type UserRegisteredRecord struct {
	Username          string `json:"username"`
	RegistrationTime  string `json:"registration_time"` // RFC3339
	Device            string `json:"device"`
}

type UserLoggedInRecord struct {
	Username  string `json:"username"`
	LoginTime string `json:"login_time"` // RFC3339
	Device    string `json:"device"`
	IPAddress string `json:"ip_address"`
}

type RecipeOpenedRecord struct {
	Username  string `json:"username,omitempty"`
	RecipeID  int32  `json:"recipe_id"`
	OpenTime  string `json:"open_time"` // RFC3339
	Device    string `json:"device"`
}