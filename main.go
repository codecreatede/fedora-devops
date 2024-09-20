package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date : 2024-9-10

a fedora system utility for monitoring the fedora installations in cloud antive applications.
you can install this as sudo or provide the password for the validation in case of some of the command needs
a complete fedora dnf system manager written in Golang

*/

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	vardec := flag.String("help", "false", "boolean value")
	password := flag.String("password", "passw", "password for the system")
	command := flag.String("command", "comm", "command for the system check")
	pack := flag.String("package", "pkgname", "name for the package")
	idnumber := flag.String("idnumber", "id", "idnumber for the rollback")
	rpm := flag.String("rpmpackage", "rpm", "rpm package path")

	flag.Parse()

	commands := []string{
		"list",
		"list-recent",
		"upgrade",
		"autoremvoe",
		"dnf-update",
		"minimal",
		"pkdupdate",
		"pkgdowngrade",
		"clean",
		"history",
		"rollback",
		"info",
		"install",
		"reinstall",
		"remove",
		"search",
		"searchall",
		"rpm",
		"refresh",
	}

	if *vardec == "list" {
		fmt.Println("The following commands are present in the fedora devops")
		for i := range commands {
			fmt.Println(commands[i])
		}
	}

	if len(*password) > 0 && *command == "list" {
		cmd := exec.Command("dnf", "list", "--installed")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("dnf", "list", "--installed").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}

	if len(*password) > 0 && *command == "refresh" {
		cmd := exec.Command("dnf", "makecache", "--refresh")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with the %s\n", err)
		}
		out := exec.Command("dnf", "makecache", "--refresh")
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}

	if len(*password) > 0 && *command == "list-recent" {
		cmd := exec.Command("dnf", "list", "--recent")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("dnf", "list", "--recent").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}
	if len(*password) > 0 && *command == "upgrade" {
		cmd := exec.Command("dnf", "list", "--upgrade")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("dnf", "list", "--upgrade").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}

	if len(*password) > 0 && *command == "autoremove" {
		cmd := exec.Command("dnf", "list", "--autoremove")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("dnf", "list", "--autoremove").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}
	if len(*password) > 0 && *command == "dnf-update" {
		cmd := exec.Command("dnf", "check-update")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("dnf", "check-update").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}
	if len(*password) > 0 && *command == "minimal" {
		cmd := exec.Command("dnf", "upgrade-minimal")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("dnf", "upgrade-minimal").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}
	if len(*password) > 0 && *pack != "nil" && *command == "pkgupdate" {
		cmd := exec.Command("dnf", "upgrade", "*pack")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("dnf", "upgrade", "*pack").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}
	if len(*password) > 0 && *pack != "nil" && *command == "pkgdowngrade" {
		cmd := exec.Command("dnf", "downgrade", "*pack")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("dnf", "downgrade", "*pack").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}
	if len(*password) > 0 && *command == "clean" {
		cmd1 := exec.Command("dnf", "cleanall")
		cmd2 := exec.Command("dnf", "clean", "metadata")
		err1 := cmd1.Run()
		if err1 != nil {
			log.Fatal("cmd failed with %s\n", err1)
		}
		err2 := cmd2.Run()
		if err2 != nil {
			log.Fatal("cmd failed with %s\n", err2)
		}
		out, err := exec.Command("dnf", "cleanall").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
		out1, err1 := exec.Command("dnf", "clean", "metadata").Output()
		if err1 != nil {
			log.Fatal("cmd failed with %s\n", err1)
			fmt.Println(out1)
		}
	}
	if *command == "history" {
		cmd := exec.Command("dnf", "history")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("dnf", "history").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}
	if len(*password) > 0 && *idnumber != "nil" && *pack != "nil" && *command == "rollback" {
		cmd := exec.Command("dnf", "rollback", "*idnumber")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("dnf", "rollback", "*idnumber").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}
	if *pack != "nil" && *command == "info" {
		cmd := exec.Command("dnf", "info", "*pack")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("dnf", "info", "*pack").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}
	if len(*password) > 0 && *pack != "nil" && *command == "install" {
		cmd := exec.Command("dnf", "install", "*pack")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("dnf", "install", "*pack").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}
	if len(*password) > 0 && *pack != "nil" && *command == "reinstall" {
		cmd := exec.Command("dnf", "reinstall", "*pack")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("dnf", "install", "*pack").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}

	if len(*password) > 0 && *pack != "nil" && *command == "search" {
		cmd := exec.Command("sudo", "dnf", "search", "*pack")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("sudo", "dnf", "search", "*pack").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}

	if len(*password) > 0 && *pack != "nil" && *command == "searchall" {
		cmd := exec.Command("sudo", "dnf", "search", "--all", "*pack")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("sudo", "dnf", "search", "--all", "*pack").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}

	if len(*password) > 0 && *pack != "nil" && *command == "remove" {
		cmd := exec.Command("sudo", "dnf", "remove", "*pack")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("sudo", "dnf", "remove", "*pack").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}
	if len(*password) > 0 && *rpm != "nil" {
		cmd := exec.Command("sudo", "dnf", "install", "*rpm")
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd failed with %s\n", err)
		}
		out, err := exec.Command("sudo", "dnf", "install", "*rpm").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}
}
