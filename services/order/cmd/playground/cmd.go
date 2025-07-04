package playground

import (
	"context"

	"github.com/urfave/cli/v3"
)

var ConvertNumberTypeCmd = cli.Command{
	Name: "convert-number-type",
	Action: func(ctx context.Context, c *cli.Command) error {
		ConvertNumberType()
		return nil
	},
}
