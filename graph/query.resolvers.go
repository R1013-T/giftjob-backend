package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"giftjob-backend/graph/model"
)

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetTemplate is the resolver for the getTemplate field.
func (r *queryResolver) GetTemplate(ctx context.Context, id string) (*model.CompanyCustomTemplate, error) {
	var template model.CompanyCustomTemplate
	if err := r.DB.First(&template, id).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

// GetCompany is the resolver for the getCompany field.
func (r *queryResolver) GetCompany(ctx context.Context, id string) (*model.Company, error) {
	var company model.Company
	if err := r.DB.First(&company, id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

// GetPerson is the resolver for the getPerson field.
func (r *queryResolver) GetPerson(ctx context.Context, id string) (*model.Person, error) {
	var person model.Person
	if err := r.DB.First(&person, id).Error; err != nil {
		return nil, err
	}
	return &person, nil
}

// GetNote is the resolver for the getNote field.
func (r *queryResolver) GetNote(ctx context.Context, id string) (*model.Note, error) {
	var note model.Note
	if err := r.DB.First(&note, id).Error; err != nil {
		return nil, err
	}
	return &note, nil
}

// GetCalendar is the resolver for the getCalendar field.
func (r *queryResolver) GetCalendar(ctx context.Context, id string) (*model.Calendar, error) {
	var calendar model.Calendar
	if err := r.DB.First(&calendar, id).Error; err != nil {
		return nil, err
	}
	return &calendar, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
