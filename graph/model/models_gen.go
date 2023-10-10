// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Calendar struct {
	ID           string    `json:"id"`
	Title        *string   `json:"title,omitempty"`
	Description  *string   `json:"description,omitempty"`
	Color        *string   `json:"color,omitempty"`
	StartTime    *string   `json:"start_time,omitempty"`
	EndTime      *string   `json:"end_time,omitempty"`
	Location     *string   `json:"location,omitempty"`
	IsAllDay     *bool     `json:"is_all_day,omitempty"`
	IsFromGoogle *bool     `json:"is_from_google,omitempty"`
	CompanyID    *string   `json:"company_id,omitempty"`
	UserID       string    `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Company struct {
	ID                  string                `json:"id"`
	Name                *string               `json:"name,omitempty"`
	Color               *string               `json:"color,omitempty"`
	Tell                *string               `json:"tell,omitempty"`
	Email               *string               `json:"email,omitempty"`
	Address             *string               `json:"address,omitempty"`
	SiteURL             *string               `json:"site_url,omitempty"`
	Industry            *string               `json:"industry,omitempty"`
	EmployeesNumber     *float64              `json:"employees_number,omitempty"`
	IsPinned            *bool                 `json:"is_pinned,omitempty"`
	PinnedAt            *string               `json:"pinned_at,omitempty"`
	IsTrash             *bool                 `json:"is_trash,omitempty"`
	CompanyCustomFields []*CompanyCustomField `json:"CompanyCustomFields,omitempty"`
	UserID              string                `json:"user_id"`
	CreatedAt           time.Time             `json:"created_at"`
	UpdatedAt           time.Time             `json:"updated_at"`
}

type CompanyCustomField struct {
	ID        string    `json:"id"`
	GroupName string    `json:"group_name"`
	Label     string    `json:"label"`
	Value     *string   `json:"value,omitempty"`
	Type      string    `json:"type"`
	CompanyID string    `json:"company_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CompanyCustomTemplate struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	IsTrash   *bool     `json:"is_trash,omitempty"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CompanyCustomTemplateField struct {
	ID         string    `json:"id"`
	GroupName  string    `json:"group_name"`
	Label      string    `json:"label"`
	Type       string    `json:"type"`
	TemplateID string    `json:"template_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateCalendarInput struct {
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	Color        *string `json:"color,omitempty"`
	StartTime    *string `json:"start_time,omitempty"`
	EndTime      *string `json:"end_time,omitempty"`
	Location     *string `json:"location,omitempty"`
	IsAllDay     *bool   `json:"is_all_day,omitempty"`
	IsFromGoogle *bool   `json:"is_from_google,omitempty"`
	CompanyID    *string `json:"company_id,omitempty"`
	UserID       string  `json:"user_id"`
}

type CreateCompanyInput struct {
	Name            *string  `json:"name,omitempty"`
	Color           *string  `json:"color,omitempty"`
	Tell            *string  `json:"tell,omitempty"`
	Email           *string  `json:"email,omitempty"`
	Address         *string  `json:"address,omitempty"`
	SiteURL         *string  `json:"site_url,omitempty"`
	Industry        *string  `json:"industry,omitempty"`
	EmployeesNumber *float64 `json:"employees_number,omitempty"`
	IsPinned        *bool    `json:"is_pinned,omitempty"`
	PinnedAt        *string  `json:"pinned_at,omitempty"`
	IsTrash         *bool    `json:"is_trash,omitempty"`
	UserID          string   `json:"user_id"`
}

type CreateCustomFieldInput struct {
	GroupName string  `json:"group_name"`
	Label     string  `json:"label"`
	Value     *string `json:"value,omitempty"`
	Type      string  `json:"type"`
	CompanyID string  `json:"company_id"`
}

type CreateNoteInput struct {
	Title     *string `json:"title,omitempty"`
	Content   *string `json:"content,omitempty"`
	Color     *string `json:"color,omitempty"`
	IsPinned  *bool   `json:"is_pinned,omitempty"`
	PinnedAt  *string `json:"pinned_at,omitempty"`
	IsTrash   *bool   `json:"is_trash,omitempty"`
	CompanyID *string `json:"company_id,omitempty"`
	UserID    string  `json:"user_id"`
}

type CreatePersonInput struct {
	Name       *string  `json:"name,omitempty"`
	Department *string  `json:"department,omitempty"`
	Position   *string  `json:"position,omitempty"`
	Tell       *float64 `json:"tell,omitempty"`
	Email      *string  `json:"email,omitempty"`
	Memo       *string  `json:"memo,omitempty"`
	IsTrash    *bool    `json:"is_trash,omitempty"`
	CompanyID  *string  `json:"company_id,omitempty"`
	UserID     string   `json:"user_id"`
}

type CreateTemplateFieldInput struct {
	GroupName  string `json:"group_name"`
	Label      string `json:"label"`
	Type       string `json:"type"`
	TemplateID string `json:"template_id"`
	UserID     string `json:"user_id"`
}

type CreateTemplateInput struct {
	Name    string `json:"name"`
	UserID  string `json:"user_id"`
	IsTrash *bool  `json:"is_trash,omitempty"`
}

type CreateUserInput struct {
	Provider string `json:"provider"`
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Image    string `json:"image"`
}

type Note struct {
	ID        string    `json:"id"`
	Title     *string   `json:"title,omitempty"`
	Content   *string   `json:"content,omitempty"`
	Color     *string   `json:"color,omitempty"`
	IsPinned  *bool     `json:"is_pinned,omitempty"`
	PinnedAt  *string   `json:"pinned_at,omitempty"`
	IsTrash   *bool     `json:"is_trash,omitempty"`
	CompanyID *string   `json:"company_id,omitempty"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Person struct {
	ID         string    `json:"id"`
	Name       *string   `json:"name,omitempty"`
	Department *string   `json:"department,omitempty"`
	Position   *string   `json:"position,omitempty"`
	Tell       *float64  `json:"tell,omitempty"`
	Email      *string   `json:"email,omitempty"`
	Memo       *string   `json:"memo,omitempty"`
	IsTrash    *bool     `json:"is_trash,omitempty"`
	CompanyID  *string   `json:"company_id,omitempty"`
	UserID     string    `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type SignInInput struct {
	Provider string `json:"provider"`
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Image    string `json:"image"`
}

type UpdateCalendarInput struct {
	ID           string  `json:"id"`
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	Color        *string `json:"color,omitempty"`
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
	Color           *string  `json:"color,omitempty"`
	Tell            *string  `json:"tell,omitempty"`
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
	ID        string  `json:"id"`
	Title     *string `json:"title,omitempty"`
	Content   *string `json:"content,omitempty"`
	Color     *string `json:"color,omitempty"`
	IsPinned  *bool   `json:"is_pinned,omitempty"`
	PinnedAt  *string `json:"pinned_at,omitempty"`
	CompanyID *string `json:"company_id,omitempty"`
	IsTrash   *bool   `json:"is_trash,omitempty"`
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

type UpdateTemplateFieldInput struct {
	ID         string  `json:"id"`
	Label      *string `json:"label,omitempty"`
	Type       *string `json:"type,omitempty"`
	TemplateID *string `json:"template_id,omitempty"`
}

type UpdateTemplateInput struct {
	ID      string  `json:"id"`
	Name    *string `json:"name,omitempty"`
	UserID  *string `json:"user_id,omitempty"`
	IsTrash *bool   `json:"is_trash,omitempty"`
}

type User struct {
	ID        string                   `json:"id"`
	Provider  string                   `json:"provider"`
	UID       string                   `json:"uid"`
	Name      string                   `json:"name"`
	Email     string                   `json:"email"`
	Image     string                   `json:"image"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
	Companies []*Company               `json:"companies,omitempty"`
	Templates []*CompanyCustomTemplate `json:"templates,omitempty"`
	Notes     []*Note                  `json:"Notes,omitempty"`
	Calendars []*Calendar              `json:"Calendars,omitempty"`
	People    []*Person                `json:"People,omitempty"`
}
