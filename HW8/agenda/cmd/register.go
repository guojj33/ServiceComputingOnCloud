package cmd

import (
	"../entity"
	"github.com/spf13/cobra"
	"log"
)

var registerCmd = &cobra.Command{
	Use: "register",
	Short: "Register an account",
	Long: "Usage : agenda register -u [username] -p [password] -e [email] -t [telephone]",
	Run: func(cmd *cobra.Command, args []string){
		u, _ := cmd.Flags().GetString("username")
		p, _ := cmd.Flags().GetString("password")
		e, _ := cmd.Flags().GetString("email")
		t, _ := cmd.Flags().GetString("telephone")
		if entity.CreateUser(u, p, e, t) == true {
			log.Println("Register successfully")
			log.Println("Register info: username: " + u + "\npassword: " + p + "\nemail: " + e + "\ntelephone: " + t)
		} else {
			log.Println("Create user failed")
			log.Println("Error: user info invalid or username already exist")
			log.Println("Error: try './agenda register -h' for help")
		}
	},
}

func init() {
	entity.Init()
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("telephone", "t", "", "")
	registerCmd.Flags().StringP("email", "e", "", "")
	registerCmd.Flags().StringP("password", "p", "", "")
	registerCmd.Flags().StringP("username", "u", "", "")
}
