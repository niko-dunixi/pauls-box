package update

import "github.com/paul-nelson-baker/pauls-box/core/helper"

func RunUpdate() error {
	return helper.SimpleRun("choco", "upgrade", "all")
}
