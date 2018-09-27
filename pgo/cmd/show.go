package cmd

/*
 Copyright 2017-2018 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

import (
	"fmt"

	"github.com/crunchydata/postgres-operator/util"
	"github.com/spf13/cobra"
)

const TreeBranch = "\t"
const TreeTrunk = "\t"

var ShowPVC bool
var PVCRoot string

var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the description of a cluster",
	Long: `Show allows you to show the details of a policy, backup, pvc, or cluster. For example:

	pgo show backup mycluster
	pgo show backup mycluster --backup-type=pgbackrest
	pgo show cluster mycluster
	pgo show config
	pgo show ingest myingest
	pgo show policy policy1
	pgo show pvc mycluster
	pgo show user mycluster`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(`Error: You must specify the type of resource to show.
Valid resource types include:
	* backup
	* cluster
	* config
	* ingest
	* policy
	* pvc
	* upgrade
	* user
	`)
		} else {
			switch args[0] {
			case "cluster":
			case "pvc":
			case "policy":
			case "ingest":
			case "user":
			case "config":
				//		case "schedule":
			case "upgrade":
			case "backup":
				break
			default:
				fmt.Println(`Error: You must specify the type of resource to show.
Valid resource types include:
	* backup
	* cluster
	* config
	* ingest
	* policy
	* pvc
	* upgrade
	* user`)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(ShowCmd)
	ShowCmd.AddCommand(ShowBackupCmd)
	ShowCmd.AddCommand(ShowClusterCmd)
	ShowCmd.AddCommand(ShowConfigCmd)
	ShowCmd.AddCommand(ShowIngestCmd)
	ShowCmd.AddCommand(ShowPolicyCmd)
	ShowCmd.AddCommand(ShowPVCCmd)
	//	ShowCmd.AddCommand(ShowScheduleCmd)
	ShowCmd.AddCommand(ShowUpgradeCmd)
	ShowCmd.AddCommand(ShowUserCmd)

	ShowBackupCmd.Flags().StringVarP(&BackupType, "backup-type", "", util.LABEL_BACKUP_TYPE_BASEBACKUP, "The backup type output to list. Valid choices are pgbasebackup (default) or pgbackrest.")
	ShowClusterCmd.Flags().StringVarP(&CCPImageTag, "ccp-image-tag", "", "", "Filter the results based on the image tag of the cluster.")
	ShowClusterCmd.Flags().StringVarP(&OutputFormat, "output", "o", "", "The output format. Currently, JSON is supported.")
	ShowClusterCmd.Flags().StringVarP(&Selector, "selector", "s", "", "The selector to use for cluster filtering.")
	ShowPVCCmd.Flags().StringVarP(&PVCRoot, "pvc-root", "", "", "The PVC directory to list.")
	ShowScheduleCmd.Flags().StringVarP(&Selector, "selector", "s", "", "The selector to use for cluster filtering.")
	ShowScheduleCmd.Flags().StringVarP(&ScheduleName, "schedule-name", "", "", "The name of the schedule to show.")
	ShowScheduleCmd.Flags().BoolVarP(&NoPrompt, "no-prompt", "n", false, "No command line confirmation.")
	ShowUserCmd.Flags().StringVarP(&Selector, "selector", "s", "", "The selector to use for cluster filtering.")
}

var ShowConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Show configuration information",
	Long: `Show configuration information for the Operator. For example:

	pgo show config`,
	Run: func(cmd *cobra.Command, args []string) {
		showConfig(args)
	},
}

var ShowPolicyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Show policy information",
	Long: `Show policy information. For example:

	pgo show policy policy1`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: Policy name(s) required for this command.")
		} else {
			showPolicy(args)
		}
	},
}

var ShowPVCCmd = &cobra.Command{
	Use:   "pvc",
	Short: "Show PVC information",
	Long: `Show PVC information. For example:

	pgo show pvc all
	pgo show pvc mycluster-backup
	pgo show pvc mycluster-xlog
	pgo show pvc mycluster`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: PVC name(s) required for this command.")
		} else {
			showPVC(args)
		}
	},
}

var ShowUpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Show upgrade information",
	Long: `Show upgrade information. For example:

	pgo show upgrade mycluster`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: cluster name(s) required for this command.")
		} else {
			showUpgrade(args)
		}
	},
}

// showBackupCmd represents the show backup command
var ShowBackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Show backup information",
	Long: `Show backup information. For example:

	pgo show backup mycluser`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: cluster name(s) required for this command.")
		} else {
			if BackupType == util.LABEL_BACKUP_TYPE_BACKREST {

				showBackrest(args)
			} else if BackupType == util.LABEL_BACKUP_TYPE_BASEBACKUP {
				showBackup(args)
			} else {
				fmt.Println("Error: Valid backup-type values are pgbasebackup and pgbackrest. The default if not supplied is pgbasebackup.")
			}
		}
	},
}

// ShowClusterCmd represents the show cluster command
var ShowClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Show cluster information",
	Long: `Show a PostgreSQL cluster. For example:

	pgo show cluster mycluster`,
	Run: func(cmd *cobra.Command, args []string) {
		if Selector == "" && len(args) == 0 {
			fmt.Println("Error: Cluster name(s) required for this command.")
		} else {
			showCluster(args)
		}
	},
}

// ShowIngestCmd represents the show ingest command
var ShowIngestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "Show ingest information",
	Long: `Show an ingest. For example:

	pgo show ingest myingest`,
	Run: func(cmd *cobra.Command, args []string) {
		if Selector == "" && len(args) == 0 {
			fmt.Println("Error: Ingest name(s) required for this command.")
		} else {
			showIngest(args)
		}
	},
}

// ShowUserCmd represents the show user command
var ShowUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Show user information",
	Long: `Show users on a cluster. For example:

	pgo show user mycluster`,
	Run: func(cmd *cobra.Command, args []string) {
		if Selector == "" && len(args) == 0 {
			fmt.Println("Error: Cluster name(s) required for this command.")
		} else {
			showUser(args)
		}
	},
}

// ShowUserCmd represents the show user command
var ShowScheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Show schedule information",
	Long: `Show cron-like schedules.  For example:

	pgo show schedule mycluster
	pgo show schedule --selector=pg-cluster=mycluster
	pgo show schedule --schedule-name=mycluster-pgbackrest-full`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 && Selector == "" && ScheduleName == "" {
			fmt.Println("Error: cluster name, schedule name or selector is required to show a schedule.")
			return
		}
		showSchedule(args)
	},
}
