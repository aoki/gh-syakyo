package main

import (
	"log"
)

var cmdPullRequest = &Command{
	Run:   pullRequest,
	Usage: "pull-request [-f] [TITLE|-i ISSUE] [-b BASE] [-h HEAD]",
	Short: "Open a pull request on GitHub",
	Long: `Opens apull request on GitHub for the project the "origin" remote
	points to. The default head of the pull request is the current branch.
	Both base and head of the pull request can be explicitly given in one of
	the following formats: "brach", "owner:branch", "owner/repo:branch".
	This command will abort operation if it detects that the current topic
	branch has local commits that are not yet pushed to its upstream branch
	on the remote. To skip this check, use -f.

	If instead of normal TITLE an issue number is given with -i, the pull
	request will be attached to an existing GitHub issue. Alternatively, instead
	of title you can paste a full URL to an issue on GitHub.
	`,
}

var flagPullRequestBase, flagPullRequestHead string

func init() {
	// TODO: make base default as USER:master, head ddefault as USER:HEAD
	cmdPullRequest.Flag.StringVar(&flagPullRequestBase, "b", "master", "BASE")
	cmdPullRequest.Flag.StringVar(&flagPullRequestHead, "h", "HEAD", "HEAD")
}

func pullRequest(cmd *Command, args []string) {
	log.Println(args)
	log.Println(flagPullRequestBase)
	log.Println(flagPullRequestHead)
}
