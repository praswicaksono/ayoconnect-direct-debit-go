package directdebit_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAyoconnect(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ayoconnect Suite")
}
