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
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	list          string
	listrecent    string
	upgrade       string
	autoremove    string
	dnfupdate     string
	minimal       string
	pkgupdate     string
	pkgdowngrade  string
	clean         string
	history       string
	rollback      string
	info          string
	install       string
	reinstall     string
	removep       string
	search        string
	searchall     string
	rpm           string
	refresh       string
	alignment     string
	rpmpackage    string
	packageid     string
	lscpu         string
	lscpuextended string
)

var rootCmd = &cobra.Command{
	Use:  "options",
	Long: "This is a fedora devops application that allows you to do all the fedora system devops",
	Run:  flagsFunc,
}

func init() {
	rootCmd.Flags().
		StringVarP(&list, "listcommand", "l", "list dnf", "list all the repositories")
	rootCmd.Flags().
		StringVarP(&listrecent, "listrecent", "q", "list recent dnf", "list all the recent dnf")
	rootCmd.Flags().
		StringVarP(&refresh, "refresh", "w", "refresh all the packages", "refreshing packages")
	rootCmd.Flags().
		StringVarP(&upgrade, "upgrade", "u", "upgrade dnf", "list all the dnf upgrade")
	rootCmd.Flags().
		StringVarP(&autoremove, "autoremove", "e", "remove dnf", "remove all the dnf install")
	rootCmd.Flags().
		StringVarP(&dnfupdate, "dnfupdate", "x", "update all dnf", "update all dnf repositories")
	rootCmd.Flags().
		StringVarP(&minimal, "minimal", "m", "update all dnf minimal", "update all minimal dnf")
	rootCmd.Flags().
		StringVarP(&pkgupdate, "pkgupdate", "p", "update all the packages", "updates all the dnf packages")
	rootCmd.Flags().
		StringVarP(&pkgdowngrade, "pkgdowngrade", "b", "downgrade a specific package", "downgrade a specific package")
	rootCmd.Flags().
		StringVarP(&clean, "clean", "c", "clean all the packages", "clean all the package repositories")
	rootCmd.Flags().
		StringVarP(&history, "history", "j", "history of the packages", "reports history of all the repositories")
	rootCmd.Flags().
		StringVarP(&rollback, "rollback", "r", "rolback a specific release", "rollback a specific release for the package")
	rootCmd.Flags().
		StringVarP(&info, "info", "i", "information about a specific package", "package information")
	rootCmd.Flags().
		StringVarP(&install, "install", "s", "install a specific package", "package install")
	rootCmd.Flags().
		StringVarP(&reinstall, "reinstall", "z", "reinstall a package", "re-install a specific package")
	rootCmd.Flags().
		StringVarP(&removep, "removep", "v", "remove a specific package", "remove a specific package")
	rootCmd.Flags().
		StringVarP(&search, "search", "t", "search a specific package", "search for a specific package repository")
	rootCmd.Flags().
		StringVarP(&searchall, "searchall", "a", "search all the package", "search all the package repository")
	rootCmd.Flags().
		StringVarP(&rpm, "rpm", "f", "install a specific rpm", "install a specific rpm package")
	rootCmd.Flags().
		StringVarP(&rpmpackage, "rpmpackage", "d", "install the rpm package", "install the specific rpm package")
	rootCmd.Flags().
		StringVarP(&packageid, "packageid", "n", "install the package id with the id number", "install the specific package with the id number")
	rootCmd.Flags().StringVarP(&lscpu, "lscpu", "L", "check the lscpu", "lscpu etail")
	rootCmd.Flags().
		StringVarP(&lscpuextended, "lscpuextended", "E", "lscpu extended with json write", "lscpu extended")
}

func flagsFunc(cmd *cobra.Command, args []string) {
	if list == "yes" {
		out, err := exec.Command("dnf", "list", "--installed").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(out)
		}
	}
	if refresh == "yes" {
		out, err := exec.Command("sudo", "dnf", "makecache", "--refresh").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if listrecent == "yes" {
		out, err := exec.Command("sudo", "dnf", "list", "--recent").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if upgrade == "yes" {
		out, err := exec.Command("sudo", "dnf", "list", "--upgrade").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if removep == "yes" {
		out, err := exec.Command("sudo", "dnf", "remove", "&rpmpackage").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if autoremove == "yes" {
		out, err := exec.Command("sudo", "dnf", "list", "--autoremove").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if dnfupdate == "yes" {
		out, err := exec.Command("sudo", "dnf", "check-update").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if minimal == "yes" {
		out, err := exec.Command("sudo", "dnf", "upgrade-minimal").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if pkgupdate == "yes" {
		out, err := exec.Command("sudo", "dnf", "upgrade", "&rpmpackage").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if pkgdowngrade == "yes" {
		out, err := exec.Command("dnf", "downgrade", "&rpmpackage").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if clean == "yes" {
		out, err := exec.Command("sudo", "dnf", "cleanall").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
		out1, err1 := exec.Command("sudo", "dnf", "clean", "metadata").Output()
		if err1 != nil {
			log.Fatal(err1)
			fmt.Println(string(out1))
		}
	}
	if history == "yes" {
		out, err := exec.Command("sudo", "dnf", "history").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if rollback == "yes" {
		out, err := exec.Command("sudo", "dnf", "rollback", "&packageid").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if info == "yes" {
		out, err := exec.Command("sudo", "dnf", "info", "&rpmpackage").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if install == "yes" {
		out, err := exec.Command("sudo", "dnf", "install", "&rpmpackage").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if reinstall == "yes" {
		out, err := exec.Command("sudo", "dnf", "install", "&rpmpackage").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if search == "yes" {
		out, err := exec.Command("sudo", "dnf", "search", "&rpmpackage").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if searchall == "yes" {
		out, err := exec.Command("sudo", "dnf", "search", "--all", "&rpmpackage").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if rpm == "yes" {
		out, err := exec.Command("sudo", "dnf", "install", "&rpmpackage").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if lscpu == "yes" {
		out, err := exec.Command("lscpu").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
	if lscpuextended == "yes" {
		out, err := exec.Command("lscpu", "--extended", "--json").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(out))
		}
	}
}
