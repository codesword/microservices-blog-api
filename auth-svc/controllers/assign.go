package controllers

import (
	"github.com/codesword/microservices-blog-api/pb/authorization"
	"golang.org/x/net/context"
)

//AssignPermissionsToRole assigns the given permissions to the given role
func (s *Server) AssignPermissionsToRole(ctx context.Context, params *authorization.PermissionsToRoleRequest) (*authorization.Empty, error) {
	return &authorization.Empty{}, nil
}

//AssignRolesToPermission assigns the given roles the given permission
func (s *Server) AssignRolesToPermission(ctx context.Context, params *authorization.RolesToPermissionRequest) (*authorization.Empty, error) {
	return &authorization.Empty{}, nil
}

//AssignRolesToUser assigns the given roles to the given user
func (s *Server) AssignRolesToUser(ctx context.Context, params *authorization.RolesToUserRequest) (*authorization.Empty, error) {
	return &authorization.Empty{}, nil
}

//AssignUsersToRole assigns the given users the given role
func (s *Server) AssignUsersToRole(ctx context.Context, params *authorization.UsersToRoleRequest) (*authorization.Empty, error) {
	return &authorization.Empty{}, nil
}
