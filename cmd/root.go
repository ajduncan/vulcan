package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"

  "github.com/ajduncan/vulcan/pkg/beacon"
  "github.com/ajduncan/vulcan/pkg/katric"
  "github.com/ajduncan/vulcan/pkg/scuttlebutt"
  "github.com/ajduncan/vulcan/pkg/ellipsis"
)

var rootCmd = &cobra.Command{
  Use:   "vulcan",
  Short: "Vulcan web analytics.",
  Long: `An explorative way of learning how web analytics works at scale. Built
         with lack of feelings by ajduncan and friends in Go.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Live long and prosper.")
  },
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Vulcan",
  Long:  `A number greater than 0, with prefix 'v', and possible suffixes like
          'a', 'b' or 'RELEASE'`,
  Run: func(cmd *cobra.Command, args []string) {
    // todo ...
    fmt.Println("Vulcan v0.1a")
  },
}

var beaconCmd = &cobra.Command{
  Use:   "beacon",
  Short: "Run the beacon analytics endpoint.",
  Long:  `Run the beacon analytics endpoint.`,
  Run: func(cmd *cobra.Command, args []string) {
    beacon.RunBeaconService()
  },
}

var katricCmd = &cobra.Command{
  Use:   "katric",
  Short: "Run the katric db interface api.",
  Long:  `Run the katric db interface api.`,
  Run: func(cmd *cobra.Command, args []string) {
    katric.RunKatricService()
  },
}

var scuttlebuttCmd = &cobra.Command{
  Use:   "scuttlebutt",
  Short: "Run the analytics reporting service.",
  Long:  `Run the analytics reporting service.`,
  Run: func(cmd *cobra.Command, args []string) {
    scuttlebutt.RunScuttlebuttService()
  },
}

var ellipsisCmd = &cobra.Command{
  Use:   "ellipsis",
  Short: "Run the ellipsis documentation and testing api.",
  Long:  `Run the ellipsis documentation and testing api.`,
  Run: func(cmd *cobra.Command, args []string) {
    ellipsis.RunEllipsisService()
  },
}

func init() {
  rootCmd.AddCommand(versionCmd)
  rootCmd.AddCommand(beaconCmd)
  rootCmd.AddCommand(katricCmd)
  rootCmd.AddCommand(scuttlebuttCmd)
  rootCmd.AddCommand(ellipsisCmd)
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
