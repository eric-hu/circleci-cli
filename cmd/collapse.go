package cmd

import (
	"github.com/CircleCI-Public/circleci-cli/filetree"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var collapseCommand = &cobra.Command{
	Use:   "collapse",
	Short: "Collapse your CircleCI configuration to a single file",
	Run:   collapse,
}

var root string

func init() {
	collapseCommand.Flags().StringVarP(&root, "root", "r", ".circleci", "path to your configuration (default is .circleci)")
}

func collapse(cmd *cobra.Command, args []string) {
	tree, err := filetree.NewTree(root)
	if err != nil {
		Logger.FatalOnError("An error occurred trying to build the tree", err)
	}

	y, err := yaml.Marshal(&tree)
	if err != nil {
		Logger.FatalOnError("Failed trying to marshal the tree to YAML ", err)
	}
	Logger.Infof("%s\n", string(y))
}
