package tidb

import (
	"context"

	"github.com/urfave/cli/v3"
)

var InsertUserCmd = cli.Command{
	Name: "insert-user",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "user",
			Aliases: []string{"u"},
			Value:   1,
			Usage:   "number of users to insert",
		},
	},
	Action: func(ctx context.Context, cmd *cli.Command) error {
		userCount := cmd.Int("user")
		InsertUser(ctx, userCount)
		return nil
	},
}
