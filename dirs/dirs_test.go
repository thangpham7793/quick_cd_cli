package dirs_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/thangpham7793/qcd/dirs"
)

var _ = Describe("Dirs", func() {

	BeforeEach(func() {
		os.Remove("dir.json")
	})

	Context("With no json file", func() {
		It("Should create one on init", func() {
			j := dirs.Dirs{}
			j.Init(dirs.FILEPATH)
			Expect(filepath.Abs("dir.json")).To(BeAnExistingFile())
		})
	})

	Context("With a valid alias and path", func() {
		It("Should create an entry", func() {
			j := dirs.Dirs{}
			j.Init(dirs.FILEPATH)
			j.AddOne("test", "/home/thang/projects/personal/qcd/dirs")
			Expect(j).To(HaveKeyWithValue("test", "/home/thang/projects/personal/qcd/dirs"))
		})
	})
})
