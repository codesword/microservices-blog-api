package controllers

import (
	"context"
	"github.com/codesword/microservices-blog-api/pb/authorization"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Assign Controller", func() {
	Describe("AssignPermissionsToRole", func() {
		It("returns Empty struct", func() {
			response, err := client.AssignPermissionsToRole(context.Background(), &authorization.PermissionsToRoleRequest{})
			Expect(err).NotTo(HaveOccurred())
			Expect(response).To(Equal(&authorization.Empty{}))
		})
	})

	Describe("AssignRolesToPermission", func() {
		It("returns Empty struct", func() {
			response, err := client.AssignRolesToPermission(context.Background(), &authorization.RolesToPermissionRequest{})
			Expect(err).NotTo(HaveOccurred())
			Expect(response).To(Equal(&authorization.Empty{}))
		})
	})

	Describe("AssignUsersToRole", func() {
		It("returns Empty struct", func() {
			response, err := client.AssignUsersToRole(context.Background(), &authorization.UsersToRoleRequest{})
			Expect(err).NotTo(HaveOccurred())
			Expect(response).To(Equal(&authorization.Empty{}))
		})
	})

	Describe("AssignRolesToUser", func() {
		It("returns Empty struct", func() {
			response, err := client.AssignRolesToUser(context.Background(), &authorization.RolesToUserRequest{})
			Expect(err).NotTo(HaveOccurred())
			Expect(response).To(Equal(&authorization.Empty{}))
		})
	})
})
