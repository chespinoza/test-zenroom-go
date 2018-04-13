package main

import (
	"github.com/spf13/cobra"

)

var startCmd = &cobra.Command{
	Use:"start",
	Short: "Start the service ...",
	RunE: func(cmd *cobra.Command, args[])error{
		// config ...
		
	}
}