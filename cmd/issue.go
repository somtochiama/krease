package cmd

import (
	"github.com/SomtochiAma/krease/pkg/issue"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

var (
	file string
	milestone string
	labels []string
)

func init() {
	rootCmd.AddCommand(IssueCmd)
	IssueCmd.Flags().StringVar(&file, "file", "", "The file that contains the go template of the comment")
	IssueCmd.Flags().StringVar(&milestone, "milestone", "none", "Milestone of the issue you want to comment on")
	IssueCmd.Flags().StringArrayVar(&labels, "labels", []string{""}, "label of the issue you want to comment on")
	_ = IssueCmd.MarkFlagRequired("file")
	_ = IssueCmd.MarkFlagRequired("owner")
}

var IssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Comment on issues with particular labels + milestone",
	Args:  cobra.MinimumNArgs(1),
	RunE:   commentIssues,
}

func commentIssues(cmd *cobra.Command, args []string) error {
	repo := args[0]
	ghClient, err := issue.AuthGitHub(token)
	if err != nil {
		klog.Errorf("error authenticating to github", err)
		return err
	}

	issues, err := issue.GetTrackedPRs(cmd.Context(), ghClient, name, repo, labels, milestone)
	if err != nil {
		klog.Info("error getting issues: ", err)
		return err
	}
	klog.Infof("Number of issues gotten with label %v and milestone %v is %v", labels, milestone, len(issues))

	template, err := issue.CreateTemplate(file)
	if err != nil {
		klog.Error("error creating template from comment file: ", err)
		return err
	}

	for _, eachIssue := range issues {
		klog.Info("Commenting on issue ", eachIssue.GetNumber())
		comment, err := issue.ParseTemplate(template, *eachIssue.Assignee.Login)
		if err != nil {
			return err
		}
		err = issue.CommentOnIssue(cmd.Context(), ghClient, name, repo, eachIssue.GetNumber(), comment)
		if err != nil {
			klog.Error("error while commenting on issue: ", err)
		}
	}

	return nil
}

