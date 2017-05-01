package controllers

import (
	"context"

	"github.com/codesword/microservices-blog-api/pb/authorization"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Permissions Controller", func() {

	Describe("ListPermissions", func() {
		It("returns available empty permission list", func() {
			permissions, err := client.ListPermissions(context.Background(), &authorization.Empty{})
			Expect(err).To(BeNil())
			Expect(len(permissions.Values)).To(Equal(0))
		})
	})

	Describe("GetPermission", func() {
		Context("When the permission exists", func() {
			It("should return the requested permission", func() {
				response, err := client.GetPermission(context.Background(), &authorization.ID{Id: "no-id"})
				Expect(err).ToNot(HaveOccurred())
				Expect(response).To(Equal(&authorization.Permission{}))
			})
		})
	})
})
