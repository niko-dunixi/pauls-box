package update

import (
	"github.com/paul-nelson-baker/pauls-box/core/helper"
)

func RunUpdate() error {
	if err := helper.SimpleRun("brew", "update"); err != nil {
		return err
	}
	if err := helper.SimpleRun("brew", "upgrade"); err != nil {
		return err
	}
	if err := helper.SimpleRun("brew", "upgrade", "--cask"); err != nil {
		return err
	}
	if err := helper.SimpleRun("brew", "cleanup"); err != nil {
		return err
	}
	return nil
}
