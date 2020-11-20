package dirs_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/thangpham7793/qcd/dirs"
)

func cleanUp() {
	os.Remove(dirs.FILEPATH)
	os.Remove("proj1")
	os.Remove("proj2")
	os.Remove("proj3")
}

func setUp() {
	os.Mkdir("proj1", 0611)
	os.Mkdir("proj2", 0611)
	os.Mkdir("proj3", 0611)
}

var d dirs.Dirs

var testData string = `{
	"proj1": "/home/thang/projects/personal/qcd/dirs/proj1",
	"proj2": "/home/thang/projects/personal/qcd/dirs/proj2"
}`

var _ = Describe("Dirs", func() {

	BeforeSuite(func() {
		setUp()
	})

	BeforeEach(func() {
		ioutil.WriteFile(dirs.FILEPATH, []byte(testData), 0611)
		d = dirs.Dirs{}
		d.Init(dirs.FILEPATH)
	})

	Describe("Init", func() {
		Context("With no json file", func() {
			It("Should create one", func() {
				os.Remove("dir.json")
				d = dirs.Dirs{}
				d.Init(dirs.FILEPATH)
				Expect(filepath.Abs("dir.json")).To(BeAnExistingFile())
			})
		})
	})

	Describe("AddOne", func() {
		Context("With a valid alias and path", func() {
			It("Should add a valid alias-path pair", func() {
				d.AddOne("test", "/home/thang/projects/personal/qcd/dirs")
				Expect(d).To(HaveKeyWithValue("test", "/home/thang/projects/personal/qcd/dirs"))
			})
		})
	})

	Describe("AddCurrent", func() {
		Context("With a valid alias", func() {
			It("Should add a alias-cwd pair", func() {
				d.AddCurrent("cwd")
				Expect(d).To(HaveKeyWithValue("cwd", "/home/thang/projects/personal/qcd/dirs"))
			})
		})
	})

	Describe("UpdateOne", func() {
		Context("With an existing alias and valid new relative path", func() {
			It("Should update the alias with the abs path", func() {
				d.UpdateOne("proj2", "./proj3")
				Expect(d).To(HaveKeyWithValue("proj2", "/home/thang/projects/personal/qcd/dirs/proj3"))
			})
		})

		Context("With an existing alias and valid new abs path", func() {
			It("Should update the alias with the abs path", func() {
				d.UpdateOne("proj1", "/home/thang/projects/personal/qcd/dirs/proj2")
				Expect(d).To(HaveKeyWithValue("proj1", "/home/thang/projects/personal/qcd/dirs/proj2"))
			})
		})
	})

	Describe("DeleteOne", func() {
		Context("With a valid alias", func() {
			It("Should delete the matching entry", func() {
				Expect(len(d)).To(Equal(2))
				d.DeleteOne("proj1")
				Expect(d).ToNot(HaveKeyWithValue("proj1", "/home/thang/projects/personal/qcd/dirs/proj1"))
				Expect(len(d)).To(Equal(1))
			})
		})
	})

	Describe("Clean", func() {
		Context("With some stored paths are invalid", func() {
			It("Should delete such entries", func() {
				os.Remove("proj1")
				d.Clean()
				Expect(d).ToNot(HaveKey("proj1"))
				Expect(d).ToNot(HaveKeyWithValue("proj1", "/home/thang/projects/personal/qcd/dirs/proj1"))
				Expect(len(d)).To(Equal(1))
			})
		})

		Context("With all stored paths are invalid", func() {
			It("Should delete all entries", func() {
				os.Remove("proj1")
				os.Remove("proj2")
				d.Clean()
				Expect(len(d)).To(Equal(0))
			})
		})
	})

	AfterSuite(func() {
		cleanUp()
	})
})
