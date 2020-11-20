package dirs_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/thangpham7793/qcd/dirs"
)

func cleanUp() {
	os.Remove(dirs.PathToFile)
	os.Remove("proj1")
	os.Remove("proj2")
	os.Remove("proj3")
}

var d dirs.Dirs
var proj1, proj2, proj3, cwd string
var err error
var testFileName string = "test_dir.json"

func setUp() {

	cwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	proj1 = filepath.Join(cwd, "proj1")
	proj2 = filepath.Join(cwd, "proj2")
	proj3 = filepath.Join(cwd, "proj3")

	os.Mkdir(proj1, 0611)
	os.Mkdir(proj2, 0611)
	os.Mkdir(proj3, 0611)

	d = dirs.Dirs{}
	d.Init(testFileName)
}

func makeTestData() string {
	return fmt.Sprintf(`{"proj1": "%s", "proj2": "%s"}`, proj1, proj2)
}

var _ = Describe("Dirs", func() {

	BeforeSuite(func() {
		setUp()
	})

	BeforeEach(func() {
		ioutil.WriteFile(dirs.PathToFile, []byte(makeTestData()), 0611)
		d = dirs.Dirs{}
		d.Init(testFileName)
	})

	Describe("Init", func() {

		Context("With an existing json file", func() {
			It("Should parse the json correctly", func() {
				Expect(d).To(HaveKeyWithValue("proj1", proj1))
				Expect(d).To(HaveKeyWithValue("proj2", proj2))
			})
		})

		Context("With no json file", func() {
			It("Should create one", func() {
				os.Remove(dirs.PathToFile)
				d = dirs.Dirs{}
				d.Init(testFileName)
				Expect(dirs.PathToFile).To(BeAnExistingFile())
			})
		})

	})

	Describe("AddOne", func() {
		Context("With a valid alias and path", func() {
			It("Should add a valid alias-path pair", func() {
				d.AddOne("test", cwd)
				Expect(d).To(HaveKeyWithValue("test", cwd))
			})
		})
	})

	Describe("AddCurrent", func() {
		Context("With a valid alias", func() {
			It("Should add a alias-cwd pair", func() {
				d.AddCurrent("cwd")
				Expect(d).To(HaveKeyWithValue("cwd", cwd))
			})
		})
	})

	Describe("UpdateOne", func() {
		Context("With an existing alias and valid new relative path", func() {
			It("Should update the alias with the abs path", func() {
				d.UpdateOne("proj2", "./proj3")
				Expect(d).To(HaveKeyWithValue("proj2", proj3))
			})
		})

		Context("With an existing alias and valid new abs path", func() {
			It("Should update the alias with the abs path", func() {
				d.UpdateOne("proj1", proj2)
				Expect(d).To(HaveKeyWithValue("proj1", proj2))
			})
		})
	})

	Describe("DeleteOne", func() {
		Context("With a valid alias", func() {
			It("Should delete the matching entry", func() {
				d.DeleteOne("proj1")
				Expect(d).ToNot(HaveKeyWithValue("proj1", proj1))
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
				Expect(d).ToNot(HaveKeyWithValue("proj1", proj1))
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
