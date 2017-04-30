package controllers

import (
	"github.com/codesword/microservices-blog-api/pb/authorization"
	"golang.org/x/net/context"
)

// GetPermission returns a single permission with a given ID
func (s *Server) GetPermission(ctx context.Context, params *authorization.ID) (*authorization.Permission, error) {
	return &authorization.Permission{}, nil
}

//ListPermissions returns a list of all permissions
func (s *Server) ListPermissions(ctx context.Context, params *authorization.Empty) (*authorization.Permissions, error) {
	return &authorization.Permissions{}, nil
}

//AlterPermissions refreshes available permission with the list provided
func (s *Server) AlterPermissions(ctx context.Context, params *authorization.AlterPermissionsRequest) (*authorization.Empty, error) {
	return &authorization.Empty{}, nil
}
