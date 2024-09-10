if *password && *command == "list" {
		cmd := exec.Command("dnf", "list", "--installed")
		cmd.Stdout = os.Getenv()
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal("cmd.Run() failed with %s\n", err)
		}
		out, err := exec.Command("dnf", "list", "--installed").Output()
		if err != nil {
			log.Fatal(err)
		}

	if *password && *command == "list-recent" {
			cmd := exec.Command("dnf", "list", "--recent")
			cmd.Stdout = os.GetEnv()
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				log.Fatal("cmd failed with %s\n", err)
			}
			out,err := exec.Command("dnf", "list", "--recent").Output()
			if err != nil {
				log.Fatal(err)
			}
	if *password && *command == "upgrade" {
				cmd := exec.Command("dnf", "list", "--upgrade")
				cmd.StdOut = os.GetEnv()
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatal("cmd failed with %s\n", err)
				}
			}
			out, err := exec.Command("dnf", "list", "--upgrade").Output()
			if err != nil {
				log.Fatal(err)
			}

	if *password && *command == "autoremove" {
       	                        cmd := exec.Command("dnf", "list", "--autoremove")
				cmd.Stdout = os.GetEnv()
				cmd.Sterr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatal("cmd failed with %s\n", err)
				}
				out, err := exec.Command("dnf", "list", "--autoremove").Output()
				if err != nil {
					log.Fatal(err)
				}
			}
	if *password && *command == "dnf-update" {
				cmd := exec.Command("dnf", "check-update")
				cmd.Stdout = os.GetEnv()
				cmd.Sterr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatal("cmd failed with %s\n", err)
				}
				out, err := exec.Command("dnf", "check-update").Output()
				if err != nil {
					log.Fatal(err)
				}
			}
	if *password && *command == "minimal" {
				cmd := exec.Command("dnf", "upgrade-minimal")
				cmd.Stdout = os.GetEnv()
				cmd.Sterr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatal("cmd failed with %s\n", err)
				}
				out, err := exec.Command("dnf", "upgrade-minimal").Output()
				if err != nil {
					log.Fatal(err)
				}
			}
	if *password &&  *pack && *command == "pkgupdate" {
				cmd := exec.Command("dnf", "upgrade", "*pack")
				cmd.Stdout = os.GetEnv()
				cmd.Sterr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatal("cmd failed with %s\n", err)
				}
				out, err := exec.Command("dnf", "upgrade", "*pack").Output()
				if err != nil {
					log.Fatal(err)
				}
			}
	if *password && *pack && *command == "pkgdowngrade" {
				cmd := exec.Command("dnf", "downgrade", "*pack")
				cmd.Stdout = os.GetEnv()
				cmd.Sterr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatal("cmd failed with %s\n", err)
				}
				out, err := exec.Command("dnf", "downgrade", "*pack").Output()
				if err != nil {
					log.Fatal(err)
				}
			}
	if *password && *command == "clean" {
				cmd1 := exec.Command("dnf", "cleanall")
				cmd2 := exec.Command("dnf", "clean", "metadata")
				cmd1.Stdout = os.GetEnv()
				cmd1.Stdout = os.GetEnv()
				cmd1.Sterr = os.Stderr
				cmd2.Sterr = os.Stderr
				err1 := cmd.Run()
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
				out1, err1 := exec.Command("dnf", "clean", "metadata")
				if err1 != nil {
					log.Fata("cmd failed with %s\n", err1)
				}
			}
	if *command == "history" {
				cmd := exec.Command("dnf", "history")
				cmd.Stdout = os.GetEnv()
				cmd.Sterr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatal("cmd failed with %s\n", err)
				}
				out, err := exec.Command("dnf", "history").Output()
				if err != nil {
					log.Fatal(err)
				}
			}
	if *password && *idnumber && *pack && *command == "rollback" {
				cmd := exec.Command("dnf", "rollback", "*idnumber")
				cmd.Stdout = os.GetEnv()
				cmd.Sterr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatal("cmd failed with %s\n", err)
				}
				out, err := exec.Command("dnf", "rollback", "*idnumber").Output()
				if err != nil {
					log.Fatal(err)
				}
			}
	if *pack && *command == "info" {
				cmd := exec.Command("dnf", "info", "*pack")
				cmd.Stdout = os.GetEnv()
				cmd.Sterr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatal("cmd failed with %s\n", err)
				}
				out, err := exec.Command("dnf", "info", "*pack").Output()
				if err != nil {
					log.Fatal(err)
				}
			}
	if *password && *pack && *command == "install" {
				cmd := exec.Command("dnf", "install", "*pack")
				cmd.Stdout = os.GetEnv()
				cmd.Sterr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatal("cmd failed with %s\n", err)
				}
				out, err := exec.Command("dnf", "install", "*pack").Output()
				if err != nil {
					log.Fatal(err)
				}
			}
        if *password && *pack && *command == "reinstall" {
		                cmd := exec.Command("dnf", "reinstall", "*pack")
		                cmd.Stdout = os.GetEnv()
		                cmd.Sterr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatal("cmd failed with %s\n", err)
				}
				out, err := exec.Command("dnf", "install", "*pack").Output()
				if err != nil {
					log.Fatal(err)
				}
			}

	}
