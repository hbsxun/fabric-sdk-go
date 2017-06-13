package cert

import "github.com/spf13/cobra"

var certCmd = &cobra.Command{
	Use:   "cert",
	Short: "Certificate commands",
	Long:  "Certificate commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

//Cmd returns the certificate command
func Cmd() *cobra.Command {
	//flags := certCmd.Flags()
	certCmd.AddCommand(getCertRegisterCmd())
	certCmd.AddCommand(getCertEnrollCmd())

	return certCmd
}
