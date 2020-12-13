package issue

import (
	"bytes"
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"github.com/google/go-github/v33/github"
	"text/template"
)

type CommentStruct struct {
	Name string
}

func AuthGitHub(paToken string) (*github.Client, error) {
	if paToken == "" {
		return nil, fmt.Errorf("empty personal access token")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: paToken},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client, nil
}

func GetTrackedPRs(ctx context.Context,
	ghClient *github.Client,
	owner string,
	repo string,
	labels []string,
	milestone string) ([]*github.Issue, error){
	fmt.Println(milestone)
	issues, _, err := ghClient.Issues.ListByRepo(ctx, owner, repo, &github.IssueListByRepoOptions{
		Labels: labels,
		Milestone: milestone,
	})
	if err != nil {
		return nil, err
	}
	return issues, nil
}

func CreateTemplate(file string) (*template.Template, error){
	t, err := template.ParseFiles(file)
	if err != nil {
		return nil, err
	}

	return t, err
}

func ParseTemplate(tpl *template.Template, name string) (string, error) {
	commentVar := CommentStruct{
		Name: name,
	}

	var retBytes bytes.Buffer
	if err := tpl.Execute(&retBytes, commentVar); err != nil {
		return "", err
	}

	return retBytes.String(), nil
}

func CommentOnIssue(ctx context.Context,
	ghClient *github.Client,
	owner string,
	repo string,
	issueNumber int,
	message string) (error) {
	comment := &github.IssueComment{
		Body: &message,
	}
	_, _, err := ghClient.Issues.CreateComment(ctx, owner, repo, issueNumber, comment)
	return err
}