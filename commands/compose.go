package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var templateType string
var constellationFilePath string
var mapsFilePath string
var outputPath string

var composeCmd = &cobra.Command{
	Use:   "compose",
	Short: "compose a package into requested template type",
	Long: "compose is for composing a package based on mapsFilePath and constellationFilePath and template[default value is helm].\n" +
		"abstrakt compose -t [template type] -f [constellationFilePath] -m [mapsFilePath] -o [outputPath]",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args: " + strings.Join(args, " "))
		fmt.Println("template: " + templateType)
		fmt.Println("constellationFilePath: " + constellationFilePath)
		fmt.Println("mapsFilePath: " + mapsFilePath)
		fmt.Println("outputPath: " + outputPath)
	},
}

func init() {
	composeCmd.Flags().StringVarP(&constellationFilePath, "constellationFilePath", "f", "", "constellation file path")
	err := composeCmd.MarkFlagRequired("constellationFilePath")
	if err != nil {
		fmt.Println("constellationFilePath is required")
	}
	composeCmd.Flags().StringVarP(&mapsFilePath, "mapsFilePath", "m", "", "maps file path")

	err = composeCmd.MarkFlagRequired("mapsFilePath")
	if err != nil {
		fmt.Println("mapsFilePath is required")
	}
	composeCmd.Flags().StringVarP(&outputPath, "outputPath", "o", "", "destination directory")

	err = composeCmd.MarkFlagRequired("outputPath")
	if err != nil {
		fmt.Println("outputPath is required")
	}
	composeCmd.Flags().StringVarP(&templateType, "template type", "t", "helm", "output template type")
}
