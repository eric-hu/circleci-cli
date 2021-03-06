package cmd_test

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("collapse", func() {
	var (
		command *exec.Cmd
		results []byte
	)

	Describe("a .circleci folder with config.yml and local orbs folder containing the hugo orb", func() {
		BeforeEach(func() {
			var err error
			command = exec.Command(pathCLI, "collapse", "-r", "testdata/hugo-collapse/.circleci")
			results, err = ioutil.ReadFile("testdata/hugo-collapse/result.yml")
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("collapse all YAML contents as expected", func() {
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			session.Wait()
			Expect(err).ShouldNot(HaveOccurred())
			Eventually(session.Err.Contents()).Should(BeEmpty())
			Eventually(session.Out.Contents()).Should(MatchYAML(results))
			Eventually(session).Should(gexec.Exit(0))
		})
	})

	Describe("local orbs folder with mixed inline and local commands, jobs, etc", func() {
		BeforeEach(func() {
			var err error
			var path string = "nested-orbs-and-local-commands-etc"
			command = exec.Command(pathCLI, "collapse", "-r", filepath.Join("testdata", path, "test"))
			results, err = ioutil.ReadFile(filepath.Join("testdata", path, "result.yml"))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("collapse all YAML contents as expected", func() {
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			session.Wait()
			Expect(err).ShouldNot(HaveOccurred())
			Eventually(session.Err.Contents()).Should(BeEmpty())
			Eventually(session.Out.Contents()).Should(MatchYAML(results))
			Eventually(session).Should(gexec.Exit(0))
		})
	})

	Describe("an orb containing local executors and commands in folder", func() {
		BeforeEach(func() {
			var err error
			command = exec.Command(pathCLI, "collapse", "-r", "testdata/myorb/test")
			results, err = ioutil.ReadFile("testdata/myorb/result.yml")
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("collapse all YAML contents as expected", func() {
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			session.Wait()
			Expect(err).ShouldNot(HaveOccurred())
			Eventually(session.Err.Contents()).Should(BeEmpty())
			Eventually(session.Out.Contents()).Should(MatchYAML(results))
			Eventually(session).Should(gexec.Exit(0))
		})
	})

	Describe("with a large nested config including rails orb", func() {
		BeforeEach(func() {
			var err error
			var path string = "test-with-large-nested-rails-orb"
			command = exec.Command(pathCLI, "collapse", "-r", filepath.Join("testdata", path, "test"))
			results, err = ioutil.ReadFile(filepath.Join("testdata", path, "result.yml"))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("collapse all YAML contents as expected", func() {
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			session.Wait()
			Expect(err).ShouldNot(HaveOccurred())
			Eventually(session.Err.Contents()).Should(BeEmpty())
			Eventually(session.Out.Contents()).Should(MatchYAML(results))
			Eventually(session).Should(gexec.Exit(0))
		})
	})
})
