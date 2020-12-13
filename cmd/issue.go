package cmd

import (
	"fmt"
	"github.com/SomtochiAma/krease/pkg/issue"
	"github.com/spf13/cobra"
	//log "k8s.io/klog/v2"
	//"github.com/SomtochiAma/krease/pkg/issue"
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
		return err
	}

	issues, err := issue.GetTrackedPRs(cmd.Context(), ghClient, name, repo, labels, milestone)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(issues))

	template, err := issue.CreateTemplate(file)
	for _, eachIssue := range issues {
		comment, err := issue.ParseTemplate(template, eachIssue.Assignee.GetName())
		if err != nil {
			return err
		}
		err = issue.CommentOnIssue(cmd.Context(), ghClient, name, repo, eachIssue.GetNumber(), comment)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

