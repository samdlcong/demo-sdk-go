package v1

import "time"

type TypeMeta struct {
	Kind string `json:"kind,omitempty"`

	APIVersion string `json:"apiversion,omitempty"`
}

type ListMeta struct {
	TotalCount int64 `json:"totalCount,omitempty"`
}

type ObjectMeta struct {
	ID        uint64    `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Name      string    `json:"name,omitempty" gorm:"column:name" validate:"name"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gomrm:"column:updatedAt"`
}
