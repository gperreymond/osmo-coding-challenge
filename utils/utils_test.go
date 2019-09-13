package utils_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gperreymond/osmo-coding-challenge/utils"
)

var _ = Describe("Utils", func() {
	It("should return a valid UUID.v4()", func() {
		uuid := UUID()
		ref := reflect.ValueOf(uuid)
		Expect(ref.Kind()).To(Equal(reflect.String))
	})
})
