package main

import (
    "os"
    "fmt"
    "github.com/spf13/cobra"
)

func autocompletion() {
    rootcmd := &cobra.Command{
        Use: "absh",
        Short: "Kingrog's shell",
    }

    // bash-completion
    rootcmd.AddCommand(&cobra.Command{
        Use: "bash-completion",
        Short: "Generate bash completion script",
        Run: func (cmd *cobra.Command, args []string)  {
            rootcmd.GenBashCompletion(os.Stdout)
        },
    })

    // zsh-completion
    rootcmd.AddCommand(&cobra.Command{
        Use: "zsh-completion",
        Short: "Generate zsh completion script",
        Run: func (cmd *cobra.Command, args []string)  {
            rootcmd.GenZshCompletion(os.Stdout)
        },
    })

    // fish-completion
    rootcmd.AddCommand(&cobra.Command{
        Use: "fish-completion",
        Short: "Generate fish completion script",
        Run: func (cmd *cobra.Command, args []string)  {
            rootcmd.GenFishCompletion(os.Stdout, true)
        },
    })

    // powershell-completion
    rootcmd.AddCommand(&cobra.Command{
        Use: "powershell-completion",
        Short: "Generate powershell completion script",
        Run: func (cmd *cobra.Command, args []string)  {
            rootcmd.GenPowerShellCompletionWithDesc(os.Stdout)
        },
    })

    if err := rootcmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

}
