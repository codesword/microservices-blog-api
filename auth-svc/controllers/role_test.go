package controllers

import (
	"context"

	"github.com/codesword/microservices-blog-api/pb/authorization"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Roles Controller", func() {

	Describe("ListRoles", func() {
		It("should return all empty role list", func() {
			res, err := client.ListRoles(context.Background(), &authorization.Empty{})
			Expect(err).ToNot(HaveOccurred())
			Expect(len(res.Values)).To(Equal(0))
		})
	})

	Describe("GetRole", func() {
		It("should return the requested role", func() {
			response, err := client.GetRole(context.Background(), &authorization.ID{Id: "any-id"})
			Expect(err).ToNot(HaveOccurred())
			Expect(response).To(Equal(&authorization.Role{}))
		})
	})

	Describe("CreateRole", func() {
		Context("When role params is valid", func() {
			It("should return an empty role", func() {
				role, err := client.CreateRole(context.Background(), &authorization.Role{Name: "New Role"})
				Expect(err).NotTo(HaveOccurred())
				Expect(role).To(Equal(&authorization.Role{}))
			})
		})
	})

	Describe("UpdateRole", func() {
		It("should update the role information", func() {
			role, err := client.UpdateRole(context.Background(), &authorization.Role{Name: "Updating Role"})
			Expect(err).ToNot(HaveOccurred())
			Expect(role).To(Equal(&authorization.Role{}))
		})
	})
})
