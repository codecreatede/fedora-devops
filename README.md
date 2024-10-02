# fedora-devops

- a fedora system utility for monitoring the fedora installations in cloud antive applications. 
- main menu for the fedoradevops
```
gauravsablok@gaurav-sablok ~/Desktop/codecreatede/golang/fedora-devops ±main⚡ » \
go run main.go -h
This is a fedora devops application that allows you to do all the fedora system devops

Usage:
  options [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  install
  run

Flags:
  -h, --help   help for options

Use "options [command] --help" for more information about a command.
```
- system wide install 
```
gauravsablok@gaurav-sablok ~/Desktop/codecreatede/golang/fedora-devops ±main⚡ » \
go run main.go run -h
Only system maintainence task needs to be done with this flag

Usage:
  options run [flags]

Flags:
  -e, --autoremove string      remove all the dnf install (default "remove dnf")
  -c, --clean string           clean all the package repositories (default "clean all the packages")
  -x, --dnfupdate string       update all dnf repositories (default "update all dnf")
  -h, --help                   help for run
  -j, --history string         reports history of all the repositories (default "history of the packages")
  -l, --listcommand string     list all the repositories (default "list dnf")
  -q, --listrecent string      list all the recent dnf (default "list recent dnf")
  -L, --lscpu string           lscpu detail (default "check the lscpu")
  -E, --lscpuextended string   lscpu extended (default "lscpu extended with json write")
  -m, --minimal string         update all minimal dnf (default "update all dnf minimal")
  -w, --refresh string         refreshing packages (default "refresh all the packages")
  -u, --upgrade string         list all the dnf upgrade (default "upgrade dnf")
```
- system wide installation 
```
gauravsablok@gaurav-sablok ~/Desktop/codecreatede/golang/fedora-devops ±main⚡ » go run main.go install -h
system wide installation needs to be done with this flag only. Use the -f flag to define the rpm package or the packageid

Usage:
  options install [flags]

Flags:
  -h, --help                  help for install
  -i, --info string           package information (default "information about a specific package")
  -s, --install string        package install (default "install a specific package")
  -b, --pkgdowngrade string   downgrade a specific package (default "downgrade a specific package")
  -p, --pkgupdate string      update the dnf packages (default "update the packages")
  -z, --reinstall string      re-install a specific package (default "reinstall a package")
  -v, --removep string        remove a specific package (default "remove a specific package")
  -r, --rollback string       rollback a specific release for the package (default "rolback a specific release")
  -f, --rpm string            install a specific rpm (default "install a specific rpm")
  -t, --search string         search for a specific package repository (default "search a specific package")
  -a, --searchall string      search all the package repository (default "search all the package")
gauravsablok@gaurav-sablok ~/Desktop/codecreatede/golang/fedora-devops ±main⚡ » \
go run main.go install -s yes -f \
    ~/Downloads/warp-terminal-v0.2024.09.24.08.02.stable_01-1.x86_64.rpm

```

Gaurav Sablok
