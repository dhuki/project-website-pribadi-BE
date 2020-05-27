package infrastructure

import (
	"context"

	"github.com/website-pribadi/pkg/bookmark/domain/entity"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Describe at the top level without having to wrap it in a func init() {}. it is just trick
var _ = Describe("Infrastructure", func() { // blocks define the thing you’re testing.
	// Context blocks define context of testing
	Context("when topic find by id", func() {
		Context("when id is 'asu'", func() {
			// It blocks run the code to test and state what you expect to happen.
			It("should return object with id 'asu'", func() {
				got, err := repo.FindById(context.TODO(), "asu")
				Expect(err).NotTo(HaveOccurred()) // expect not to error
				Expect(got).To(Equal(entity.Topic{
					ID:          "asu",
					Name:        "asu",
					Description: "asu",
				}))
			})
		})
	})

})
