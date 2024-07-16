package flags

import "flag"

var (
	// HelpFlag prints the usage.
	HelpFlag = flag.Bool("help", false, "Print usage.")

	// ConfigurationDirectoryFlag is the directory where the configuration files are located.
	ConfigurationDirectoryFlag = flag.String("configuration-directory", "", "The directory where the configuration files are located.")

	// OutputDirectoryFlag is the directory where the output files will be located.
	OutputDirectoryFlag = flag.String("output-directory", "./out", "The directory where the output files will be located.")

	// RecurseFlag indicates whether to recurse into subdirectories.
	RecurseFlag = flag.Bool("recurse", true, "Recurse into subdirectories.")
)

const FlagsUsageString string = `
	[-h|--help]
	[--configuration-directory <directory>] [--output-directory <directory>]
	[--recurse]`
