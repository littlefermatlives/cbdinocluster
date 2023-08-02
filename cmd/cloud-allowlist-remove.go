package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var cloudAllowListRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"delete"},
	Short:   "Removes an allowed CIDRs",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		helper := CmdHelper{}
		logger := helper.GetLogger()
		ctx := helper.GetContext()
		prov := helper.GetCloudProvisioner(ctx)

		cluster, err := helper.IdentifyCloudCluster(ctx, prov, args[0])
		if err != nil {
			logger.Fatal("failed to identify cluster", zap.Error(err))
		}

		err = prov.RemoveAllowListEntry(ctx, cluster.ClusterID, args[1])
		if err != nil {
			logger.Fatal("failed to remove allow list entry", zap.Error(err))
		}
	},
}

func init() {
	cloudAllowListCmd.AddCommand(cloudAllowListRemoveCmd)
}