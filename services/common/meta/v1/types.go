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

type ListOptions struct {
	TypeMeta       `json:",inline"`
	LabelSelector  string `json:"labelSelector,omitempty" from:"labelSelector"`
	FieldSelector  string `json:"fieldSelector,omitempty" from:"fieldSelector"`
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty"`
	Offset         *int64 `json:"offset,omitempty" from:"offset"`
	Limit          *int64 `json:"limit,omitempty" from:"limit"`
}

type ExportOptions struct {
	TypeMeta `json:",inline"`
	Export   bool `json:"export"`
	Exact    bool `json:"exact"`
}

type GetOptions struct {
	TypeMeta `json:",inline"`
}

type DeleteOptions struct {
	TypeMeta `json:",inline"`
	Unscoped bool `json:"unscoped"`
}

type CreateOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dryRun,omitempty"`
}

type PatchOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dryRun,omitempty"`
	Force    bool     `json:"force,omitempty"`
}

type UpdateOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dryRun,omitempty"`
}

type AuthorizeOptions struct {
	TypeMeta `json:",inline"`
}

type TableOptions struct {
	TypeMeta  `json:",inline"`
	NoHeaders bool `json:"-"`
}
