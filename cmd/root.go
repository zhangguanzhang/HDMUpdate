package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"hdm/HDM"
	"os"
	"sync"
)

var user, password, tfpIP, filename  string
var processLimit uint8



var rootCmd = &cobra.Command{
	Use:   "HDMUpdate",
	Short: "H3C servers' HDM batch upgrade the firmware",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}
		processWg := new(sync.WaitGroup)
		processWg.Add(len(args))

		ProcessLimit := make(chan struct{}, processLimit)
		for i := uint8(0); i < processLimit; i++ {
			ProcessLimit <- struct{}{}
		}
		for _, k := range args {
			<- ProcessLimit
			go func(ip string) {
				defer func() {
					ProcessLimit <- struct{}{}
					processWg.Done()
				}()
				h,err := HDM.NewHDM(ip, user, password)
				if err != nil {
					log.Errorf("login failure %s", ip)
					return
				}
				if err := h.Up(filename, tfpIP);err != nil {
					log.Errorf("update failre: %s",err.Error())
					return
				}
				log.Infof("%s update successed",ip)
			}(k)
		}
		processWg.Wait()

	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init(){
	rootCmd.Flags().StringVarP(&user,"user","u","admin","HDM login user")
	rootCmd.Flags().StringVarP(&password,"password","p","Password@_","HDM login password")
	rootCmd.Flags().StringVarP(&tfpIP,"tftp","t","","HDM get bin file from the tftp ip")
	rootCmd.Flags().StringVarP(&filename,"filename","f","","the HDM bin filename from the tftp root path")
	rootCmd.Flags().Uint8Var(&processLimit, "processlimit", 1, "update process limit")
	_ = rootCmd.MarkFlagRequired("tftp")
}
