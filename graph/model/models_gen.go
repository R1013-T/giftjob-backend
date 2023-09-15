// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Calendar struct {
	ID           string  `json:"id"`
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	StartTime    *string `json:"start_time,omitempty"`
	EndTime      *string `json:"end_time,omitempty"`
	Location     *string `json:"location,omitempty"`
	IsAllDay     *bool   `json:"is_all_day,omitempty"`
	IsFromGoogle *bool   `json:"is_from_google,omitempty"`
	CompanyID    string  `json:"company_id"`
	UserID       string  `json:"user_id"`
	CreatedAt    *string `json:"created_at,omitempty"`
	UpdatedAt    *string `json:"updated_at,omitempty"`
}

type Company struct {
	ID                  string                `json:"id"`
	Name                *string               `json:"name,omitempty"`
	Tell                *float64              `json:"tell,omitempty"`
	Email               *string               `json:"email,omitempty"`
	Address             *string               `json:"address,omitempty"`
	SiteURL             *string               `json:"site_url,omitempty"`
	Industry            *string               `json:"industry,omitempty"`
	EmployeesNumber     *float64              `json:"employees_number,omitempty"`
	IsPinned            *bool                 `json:"is_pinned,omitempty"`
	PinnedAt            *string               `json:"pinned_at,omitempty"`
	IsTrash             *bool                 `json:"is_trash,omitempty"`
	CompanyCustomFields []*CompanyCustomField `json:"CompanyCustomFields,omitempty"`
	CreatedAt           *string               `json:"created_at,omitempty"`
	UpdatedAt           *string               `json:"updated_at,omitempty"`
}

type CompanyCustomField struct {
	ID        string  `json:"id"`
	GroupName string  `json:"group_name"`
	Label     string  `json:"label"`
	Value     *string `json:"value,omitempty"`
	Type      string  `json:"type"`
	CompanyID string  `json:"company_id"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

type CompanyCustomTemplate struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Template  string  `json:"template"`
	IsTrash   *bool   `json:"is_trash,omitempty"`
	UserID    string  `json:"user_id"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

type CreateCalendarInput struct {
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	StartTime    *string `json:"start_time,omitempty"`
	EndTime      *string `json:"end_time,omitempty"`
	Location     *string `json:"location,omitempty"`
	IsAllDay     *bool   `json:"is_all_day,omitempty"`
	IsFromGoogle *bool   `json:"is_from_google,omitempty"`
	CompanyID    string  `json:"company_id"`
	UserID       string  `json:"user_id"`
}

type CreateCompanyInput struct {
	Name            *string  `json:"name,omitempty"`
	Tell            *float64 `json:"tell,omitempty"`
	Email           *string  `json:"email,omitempty"`
	Address         *string  `json:"address,omitempty"`
	SiteURL         *string  `json:"site_url,omitempty"`
	Industry        *string  `json:"industry,omitempty"`
	EmployeesNumber *float64 `json:"employees_number,omitempty"`
	IsPinned        *bool    `json:"is_pinned,omitempty"`
	PinnedAt        *string  `json:"pinned_at,omitempty"`
	IsTrash         *bool    `json:"is_trash,omitempty"`
}

type CreateCustomFieldInput struct {
	GroupName string  `json:"group_name"`
	Label     string  `json:"label"`
	Value     *string `json:"value,omitempty"`
	Type      string  `json:"type"`
	CompanyID string  `json:"company_id"`
}

type CreateNoteInput struct {
	Title    *string `json:"title,omitempty"`
	Content  *string `json:"content,omitempty"`
	IsPinned *bool   `json:"is_pinned,omitempty"`
	PinnedAt *string `json:"pinned_at,omitempty"`
	IsTrash  *bool   `json:"is_trash,omitempty"`
	UserID   string  `json:"user_id"`
}

type CreatePersonInput struct {
	Name       *string  `json:"name,omitempty"`
	Department *string  `json:"department,omitempty"`
	Position   *string  `json:"position,omitempty"`
	Tell       *float64 `json:"tell,omitempty"`
	Email      *string  `json:"email,omitempty"`
	Memo       *string  `json:"memo,omitempty"`
	IsTrash    *bool    `json:"is_trash,omitempty"`
	CompanyID  string   `json:"company_id"`
	UserID     string   `json:"user_id"`
}

type CreateTemplateInput struct {
	Name     string `json:"name"`
	Template string `json:"template"`
	UserID   string `json:"user_id"`
	IsTrash  *bool  `json:"is_trash,omitempty"`
}

type CreateUserInput struct {
	UUID     string `json:"uuid"`
	Provider string `json:"provider"`
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Note struct {
	ID        string  `json:"id"`
	Title     *string `json:"title,omitempty"`
	Content   *string `json:"content,omitempty"`
	IsPinned  *bool   `json:"is_pinned,omitempty"`
	PinnedAt  *string `json:"pinned_at,omitempty"`
	IsTrash   *bool   `json:"is_trash,omitempty"`
	UserID    string  `json:"user_id"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

type Person struct {
	ID         string   `json:"id"`
	Name       *string  `json:"name,omitempty"`
	Department *string  `json:"department,omitempty"`
	Position   *string  `json:"position,omitempty"`
	Tell       *float64 `json:"tell,omitempty"`
	Email      *string  `json:"email,omitempty"`
	Memo       *string  `json:"memo,omitempty"`
	IsTrash    *bool    `json:"is_trash,omitempty"`
	CompanyID  string   `json:"company_id"`
	UserID     string   `json:"user_id"`
	CreatedAt  *string  `json:"created_at,omitempty"`
	UpdatedAt  *string  `json:"updated_at,omitempty"`
}

type UpdateCalendarInput struct {
	ID           string  `json:"id"`
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	StartTime    *string `json:"start_time,omitempty"`
	EndTime      *string `json:"end_time,omitempty"`
	Location     *string `json:"location,omitempty"`
	IsAllDay     *bool   `json:"is_all_day,omitempty"`
	IsFromGoogle *bool   `json:"is_from_google,omitempty"`
	CompanyID    *string `json:"company_id,omitempty"`
}

type UpdateCompanyInput struct {
	ID              string   `json:"id"`
	Name            *string  `json:"name,omitempty"`
	Tell            *float64 `json:"tell,omitempty"`
	Email           *string  `json:"email,omitempty"`
	Address         *string  `json:"address,omitempty"`
	SiteURL         *string  `json:"site_url,omitempty"`
	Industry        *string  `json:"industry,omitempty"`
	EmployeesNumber *float64 `json:"employees_number,omitempty"`
	IsPinned        *bool    `json:"is_pinned,omitempty"`
	PinnedAt        *string  `json:"pinned_at,omitempty"`
	IsTrash         *bool    `json:"is_trash,omitempty"`
}

type UpdateCustomFieldInput struct {
	ID        string  `json:"id"`
	GroupName *string `json:"group_name,omitempty"`
	Label     *string `json:"label,omitempty"`
	Value     *string `json:"value,omitempty"`
	Type      *string `json:"type,omitempty"`
}

type UpdateNoteInput struct {
	ID       string  `json:"id"`
	Title    *string `json:"title,omitempty"`
	Content  *string `json:"content,omitempty"`
	IsPinned *bool   `json:"is_pinned,omitempty"`
	PinnedAt *string `json:"pinned_at,omitempty"`
	IsTrash  *bool   `json:"is_trash,omitempty"`
}

type UpdatePersonInput struct {
	ID         string   `json:"id"`
	Name       *string  `json:"name,omitempty"`
	Department *string  `json:"department,omitempty"`
	Position   *string  `json:"position,omitempty"`
	Tell       *float64 `json:"tell,omitempty"`
	Email      *string  `json:"email,omitempty"`
	Memo       *string  `json:"memo,omitempty"`
	CompanyID  *string  `json:"company_id,omitempty"`
	IsTrash    *bool    `json:"is_trash,omitempty"`
}

type UpdateTemplateInput struct {
	ID       string  `json:"id"`
	Name     *string `json:"name,omitempty"`
	Template *string `json:"template,omitempty"`
	UserID   *string `json:"user_id,omitempty"`
	IsTrash  *bool   `json:"is_trash,omitempty"`
}

type User struct {
	ID        string                   `json:"id"`
	UUID      string                   `json:"uuid"`
	Provider  string                   `json:"provider"`
	UID       string                   `json:"uid"`
	Name      string                   `json:"name"`
	Email     string                   `json:"email"`
	CreatedAt *string                  `json:"created_at,omitempty"`
	UpdatedAt *string                  `json:"updated_at,omitempty"`
	Companies []*Company               `json:"companies,omitempty"`
	Templates []*CompanyCustomTemplate `json:"templates,omitempty"`
	Notes     []*Note                  `json:"Notes,omitempty"`
	Calendars []*Calendar              `json:"Calendars,omitempty"`
	People    []*Person                `json:"People,omitempty"`
}
