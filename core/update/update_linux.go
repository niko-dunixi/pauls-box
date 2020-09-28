package update

import "github.com/paul-nelson-baker/pauls-box/core/helper"

func RunUpdate() error {
	if err := helper.SimpleRun("sudo", "apt-get", "update", "-y"); err != nil {
		return err
	}
	if err := helper.SimpleRun("sudo", "apt-get", "upgrade", "-y", "--with-new-pkgs"); err != nil {
		return err
	}
	if err := helper.SimpleRun("sudo", "apt-get", "autoremove", "-y"); err != nil {
		return err
	}
	return nil
}
