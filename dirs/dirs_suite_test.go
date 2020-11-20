package dirs_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDirs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dirs Suite")
}
