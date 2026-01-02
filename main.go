package main

import (
	"fmt"
	"log"

	"github.com/ralucas/mtpfs/device"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mtpfs",
	Short: "MTP device filesystem operations tool",
	Long:  "A command line tool for detecting MTP devices and performing filesystem operations",
}

var listDevicesCmd = &cobra.Command{
	Use:   "list-devices",
	Short: "List available MTP devices",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		device.List()
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List files on MTP device",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "/"
		if len(args) > 0 {
			path = args[0]
		}
		fmt.Printf("Listing: %s\n", path)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete <path>",
	Short: "Delete file or directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Deleting: %s\n", args[0])
	},
}

var moveCmd = &cobra.Command{
	Use:   "move <source> <destination>",
	Short: "Move file or directory",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Moving %s to %s\n", args[0], args[1])
	},
}

var addCmd = &cobra.Command{
	Use:   "add <local-path> <remote-path>",
	Short: "Add file to device",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding %s to %s\n", args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(listCmd, deleteCmd, moveCmd, addCmd, listDevicesCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
