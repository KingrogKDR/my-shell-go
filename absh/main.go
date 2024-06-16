package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func exists(slice []string, element string) bool{
    for _, x := range slice{
        if x == element{
            return true
        }
    }
    return false
}

func executablePath(commands []string) {
    env, ok := os.LookupEnv("PATH")
    if ok {
        paths := strings.Split(env, ":")
        for _, path := range paths {
            execPath := path + "/" + commands[1]
            if _, err := os.Stat(execPath); err == nil {
                fmt.Fprintf(os.Stdout, "%s is %s\n", commands[1], execPath)
                return
            }
        }
    } 
    fmt.Fprintf(os.Stdout, "%s: not found\n", commands[1])
    
}

func main() {
    // Uncomment this block to pass the first stage

    builtins := []string{"echo", "exit", "type"}

    for{

        fmt.Fprint(os.Stdout, "$ ")

        // Wait for user input
        command, err := bufio.NewReader(os.Stdin).ReadString('\n')
        if err != nil {
            log.Fatal(err);
        }

        command = strings.TrimSpace(command)

        cmd := strings.Split(command, " ")

        
        switch cmd[0] {
        case "absh":
            autocompletion()
        case "exit":
            if len(cmd) == 1{
                os.Exit(0)
            }
            code, err := strconv.Atoi(cmd[1])
            if err != nil {
                log.Fatal(err)
            }
            os.Exit(code)
        case "echo":
            fmt.Fprintln(os.Stdout, strings.Join(cmd[1:], " "))
        case "type" :
            if exists(builtins, cmd[1]) {
                fmt.Printf("%s is a shell builtin\n", cmd[1])    
            } else {
                executablePath(cmd)
            }
        case "cd":
            if cmd[1] == "~" {
                env, ok := os.LookupEnv("HOME")
                if ok {
                    os.Chdir(env)
                }
            } else if err:=os.Chdir(cmd[1]); err!=nil {
                fmt.Fprintf(os.Stdout, "%s: No such file or directory\n", cmd[1])
            }
        default:
            cmand := exec.Command(cmd[0], cmd[1:]...)
            cmand.Stdout = os.Stdout
            cmand.Stderr = os.Stderr
            cmand.Stdin = os.Stdin
            if err:=cmand.Run(); err!=nil {
                fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd[0])
            }
        }
    }
}
