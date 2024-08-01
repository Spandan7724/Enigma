package main

import (
	"fmt"
    "os"
	"github.com/Spandan7724/enigma/internal/cracker"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
    "runtime/pprof"
)


var cpuprofile = "cpuprofile.prof"

func main() {
    f, err := os.Create(cpuprofile)
    if err != nil {
        fmt.Println(err)
    }
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()

    printHeader()

    var rootCmd = &cobra.Command{
        Use:   "enigma",
        Short: "An offline password cracker",
    }

    var hashType string
    var targetHash string
    var wordlist string

    var crackCmd = &cobra.Command{
        Use:   "crack",
        Short: "Crack a password hash",
        Run: func(cmd *cobra.Command, args []string) {
            cracker.DictionaryAttack(hashType, targetHash, wordlist)
        },
    }

    crackCmd.Flags().StringVarP(&hashType, "hash", "t", "", "Type of hash (md5, sha1, sha256)")
    crackCmd.Flags().StringVarP(&targetHash, "target", "a", "", "Hash to crack")
    crackCmd.Flags().StringVarP(&wordlist, "wordlist", "w", "", "Path to wordlist")
    crackCmd.MarkFlagRequired("target")
    crackCmd.MarkFlagRequired("wordlist")

    rootCmd.AddCommand(crackCmd)
    rootCmd.Execute()
}

func printHeader() {
    p := termenv.ColorProfile()
    fmt.Println(termenv.String("Enigma").Foreground(p.Color("5")).Bold().String())
    fmt.Println(termenv.String("An offline password cracker").Foreground(p.Color("2")).String())
    fmt.Println()
}
