# fedora-devops

- a fedora system utility for monitoring the fedora installations in cloud antive applications. you can install this as sudo or provide the password for the validation in case of some of the command needs
- a fedora system utility written using golang. 
- eases your installation, and package management with the fedora.
- writing a package manager later on for cross compilation. 
- go package is available here : [system devops golang](https://pkg.go.dev/os/exec)

- Task to do 
  - debug
  - a graphical http response server
  - add package where instead of providing the password everytime, it takes the password using the os.Stdin

```
[gauravsablok@fedora]~/Desktop/codecreatede/fedora-devops% go run main.go -help
Usage of /tmp/go-build1408020812/b001/exe/main:
  -command string
        command for the system check (default "comm")
  -help string
        boolean value (default "false")
  -idnumber string
        idnumber for the rollback (default "id")
  -package string
        name for the package (default "pkgname")
  -password string
        password for the system (default "passw")
  -rpmpackage string
        rpm package path (default "rpm")
```


Gaurav Sablok
