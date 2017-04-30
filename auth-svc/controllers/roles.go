package controllers

import (
	"github.com/codesword/microservices-blog-api/pb/authorization"
	"golang.org/x/net/context"
)

//GetRole returns the role with the given id
func (s *Server) GetRole(ctx context.Context, params *authorization.ID) (*authorization.Role, error) {
	return &authorization.Role{}, nil
}

// ListRoles returns all availabe roles
// Takes an empty argument and returns an array of roles
func (s *Server) ListRoles(ctx context.Context, params *authorization.Empty) (*authorization.Roles, error) {
	return &authorization.Roles{}, nil
}

//CreateRole creates a role with the given params
func (s *Server) CreateRole(ctx context.Context, params *authorization.Role) (*authorization.Role, error) {
	return &authorization.Role{}, nil
}

//UpdateRole updates the attributes of the roles with the given id with the params
func (s *Server) UpdateRole(ctx context.Context, params *authorization.Role) (*authorization.Role, error) {
	return &authorization.Role{}, nil
}
