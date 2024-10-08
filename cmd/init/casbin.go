package main

import (
	"github.com/oj-lab/platform/cmd/web_server/handler"
	casbin_agent "github.com/oj-lab/platform/modules/agent/casbin"
)

func loadCasbinPolicies() {
	enforcer := casbin_agent.GetDefaultCasbinEnforcer()

	_, err := enforcer.AddGroupingPolicies([][]string{
		{`role:super`, `role:admin`, `system`},
	})
	if err != nil {
		panic(err)
	}

	err = handler.AddUserCasbinPolicies()
	if err != nil {
		panic(err)
	}
	err = handler.AddProblemCasbinPolicies()
	if err != nil {
		panic(err)
	}
	err = handler.AddFrontendPagePolicies()
	if err != nil {
		panic(err)
	}

	err = enforcer.SavePolicy()
	if err != nil {
		panic(err)
	}
}
