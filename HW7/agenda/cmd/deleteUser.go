package cmd

import(
	"../entity"
	"github.com/spf13/cobra"
	"log"
)

var deleteUserCmd = &cobra.Command{
	Use: "deleteUser",
	Short: "Delete an account using username and password",
	Long: "Usage: agenda deleteUser -u [username] -p [password]",
	Run: func(cmd *cobra.Command, args []string){
		u, _ := cmd.Flags().GetString("username")
		p, _ := cmd.Flags().GetString("password")
		if entity.DeleteUser(u, p) == true {
			log.Println("Delete user successfully")
		} else {
			log.Println("Delete user failed")
			log.Println("Error: no such user exists or user info incorrect")
			log.Println("Error: try './agenda deleteUser -h' for help")
		}
	},
}

func init() {
	entity.Init()
	rootCmd.AddCommand(deleteUserCmd)
	deleteUserCmd.Flags().StringP("password", "p", "", "")
	deleteUserCmd.Flags().StringP("username", "u", "", "")
}
