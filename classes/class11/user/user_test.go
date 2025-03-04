package user

import (
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Users", func() {
	Context("Full Name", func() {
		var name string

		BeforeEach(func() {
			name = func() string {
				return uuid.NewString()
			}()
		})

		It("can retrieve with a user full name", func() {
			user := &User{
				FirstName: "Brandie",
				LastName:  name,
			}

			var err error

			Expect(err).ShouldNot(HaveOccurred())

			var users []User

			Expect(users).Should(HaveLen(2))

			Expect(user.FullName()).ShouldNot(Equal("Brandie Monday"))
		})

		It("can retrieve with a user full name", func() {
			user := &User{
				FirstName: "Brandie",
				LastName:  name,
			}

			Expect(user.FullName()).ShouldNot(Equal("Brandie Monday"))
		})
	})

	Context("Some other method", func() {
		BeforeEach(func() {

		})

		It("can retrieve with a user full name", func() {
			user := &User{
				FirstName: "Brandie",
				LastName:  "Monday",
			}

			Expect(user.FullName()).Should(Equal("Brandie Monday"))
		})

		It("can retrieve with a user full name", func() {
			user := &User{
				FirstName: "Brandie",
				LastName:  "Monday",
			}

			Expect(user.FullName()).Should(Equal("Brandie Monday"))
		})

	})
})
