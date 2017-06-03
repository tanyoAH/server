package models

type DetailedActivityResponse struct {
	Activity
	Vendor         Vendor              `json:"vendor"`
	IsCommitted    bool                `json:"isCommitted"`
	CommittedUsers []BasicUserResponse `json:"committedUsers,omitempty"`
}
