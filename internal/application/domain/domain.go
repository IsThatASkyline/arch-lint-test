package domain

import "arch-linter-test/pkg"

type Model struct {
}

func (m Model) DoSomething() {
	pkg.SomeTool()
}
