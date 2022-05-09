package policy

import (
	"context"

	"github.com/urfave/cli"

	"github.com/smallstep/cli/command/ca/policy/acme"
	"github.com/smallstep/cli/command/ca/policy/authority"
	"github.com/smallstep/cli/command/ca/policy/provisioner"
	"github.com/smallstep/cli/internal/command"
)

// Command returns the policy subcommand.
func Command() cli.Command {
	ctx := context.Background()
	return cli.Command{
		Name:        "policy",
		Usage:       "manage certificate issuance policies",
		UsageText:   "**step ca policy** <subcommand> [arguments] [global-flags] [subcommand-flags]",
		Description: `**step ca policy** command group provides facilities for managing certificate issuance policies.`,
		Hidden:      command.ShouldBeHidden(ctx), // the `step ca policy` command is not shown in help (for now), unless STEPBETA=1 is provided
		Subcommands: cli.Commands{
			authority.Command(ctx),
			provisioner.Command(ctx),
			acme.Command(ctx),
		},
	}
}
