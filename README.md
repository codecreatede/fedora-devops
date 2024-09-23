# fedora-devops

- a fedora system utility for monitoring the fedora installations in cloud antive applications. 
- eases your installation, and package management with the fedora.
- go package is available here : [system devops golang](https://pkg.go.dev/os/exec)


```
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang/fedora-devops% \
go run main.go  -h
This is a fedora devops application that allows you to do all the fedora system devops

Usage:
  options [flags]

Flags:
  -e, --autoremove string     remove all the dnf install (default "remove dnf")
  -c, --clean string          clean all the package repositories (default "clean all the packages")
  -x, --dnfupdate string      update all dnf repositories (default "update all dnf")
  -h, --help                  help for options
  -j, --history string        reports history of all the repositories (default "history of the packages")
  -i, --info string           package information (default "information about a specific package")
  -s, --install string        package install (default "install a specific package")
  -l, --listcommand string    list all the repositories (default "list dnf")
  -q, --listrecent string     list all the recent dnf (default "list recent dnf")
  -m, --minimal string        update all minimal dnf (default "update all dnf minimal")
  -n, --packageid string      install the specific package with the id number (default "install the package id with the id number")
  -b, --pkgdowngrade string   downgrade a specific package (default "downgrade a specific package")
  -p, --pkgupdate string      updates all the dnf packages (default "update all the packages")
  -w, --refresh string        refreshing packages (default "refresh all the packages")
  -z, --reinstall string      re-install a specific package (default "reinstall a package")
  -v, --removep string        remove a specific package (default "remove a specific package")
  -r, --rollback string       rollback a specific release for the package (default "rolback a specific release")
  -f, --rpm string            install a specific rpm package (default "install a specific rpm")
  -d, --rpmpackage string     install the specific rpm package (default "install the rpm package")
  -t, --search string         search for a specific package repository (default "search a specific package")
  -a, --searchall string      search all the package repository (default "search all the package")
  -u, --upgrade string        list all the dnf upgrade (default "upgrade dnf")
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang/fedora-devops% go run main.go  -e yes




```


Gaurav Sablok
