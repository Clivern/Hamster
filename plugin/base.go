// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package plugin

//revive:disable:unused-parameter

import (
	"github.com/clivern/hamster/internal/app/event"
	"github.com/clivern/hamster/internal/app/pkg/logger"
)

// Any Action
func RawListener(raw event.Raw) (bool, error) {
	logger.Info("Raw event listener fired!")
	return true, nil
}

// Status Action
func StatusListener(status event.Status) (bool, error) {
	logger.Info("Status event listener fired!")
	return true, nil
}

// Watch Action
func WatchListener(watch event.Watch) (bool, error) {
	logger.Info("Watch event listener fired!")
	return true, nil
}

// Issue Action
func IssuesListener(issues event.Issues) (bool, error) {
	logger.Info("Issues event listener fired!")
	return true, nil
}

// Issue Comment Action
func IssueCommentListener(issueComment event.IssueComment) (bool, error) {
	logger.Info("IssueComment event listener fired!")
	return true, nil
}

// Create Action
func CreateListener(create event.Create) (bool, error) {
	logger.Info("Create event listener fired!")
	return true, nil
}

// Label Action
func LabelListener(label event.Label) (bool, error) {
	logger.Info("Label event listener fired!")
	return true, nil
}

// Delete Action
func DeleteListener(delete event.Delete) (bool, error) {
	logger.Info("Delete event listener fired!")
	return true, nil
}

// Milestone Action
func MilestoneListener(milestone event.Milestone) (bool, error) {
	logger.Info("Milestone event listener fired!")
	return true, nil
}

// Pull Request Action
func PullRequestListener(pullRequest event.PullRequest) (bool, error) {
	logger.Info("PullRequest event listener fired!")
	return true, nil
}

// Pull Request Review Action
func PullRequestReviewListener(pullRequestReview event.PullRequestReview) (bool, error) {
	logger.Info("PullRequestReview event listener fired!")
	return true, nil
}

// Pull Request Review Comment Action
func PullRequestReviewCommentListener(pullRequestReviewComment event.PullRequestReviewComment) (bool, error) {
	logger.Info("PullRequestReviewComment event listener fired!")
	return true, nil
}

// Not Supported Yet
func CheckRunListener(checkRun event.CheckRun) (bool, error) {
	logger.Info("CheckRun event listener fired!")
	return true, nil
}

// Not Supported Yet
func CheckSuiteListener(checkSuite event.CheckSuite) (bool, error) {
	logger.Info("CheckSuite event listener fired!")
	return true, nil
}

// Not Supported Yet
func CommitCommentListener(commitComment event.CommitComment) (bool, error) {
	logger.Info("CommitComment event listener fired!")
	return true, nil
}

// Not Supported Yet
func DeploymentListener(deployment event.Deployment) (bool, error) {
	logger.Info("Deployment event listener fired!")
	return true, nil
}

// Not Supported Yet
func DeploymentStatusListener(deploymentStatus event.DeploymentStatus) (bool, error) {
	logger.Info("DeploymentStatus event listener fired!")
	return true, nil
}

// Not Supported Yet
func ForkListener(fork event.Fork) (bool, error) {
	logger.Info("Fork event listener fired!")
	return true, nil
}

// Not Supported Yet
func GithubAppAuthorizationListener(githubAppAuthorization event.GithubAppAuthorization) (bool, error) {
	logger.Info("GithubAppAuthorization event listener fired!")
	return true, nil
}

// Not Supported Yet
func GollumListener(gollum event.Gollum) (bool, error) {
	logger.Info("Gollum event listener fired!")
	return true, nil
}

// Not Supported Yet
func InstallationListener(installation event.Installation) (bool, error) {
	logger.Info("Installation event listener fired!")
	return true, nil
}

// Not Supported Yet
func InstallationRepositoriesListener(installationRepositories event.InstallationRepositories) (bool, error) {
	logger.Info("InstallationRepositories event listener fired!")
	return true, nil
}

// Not Supported Yet
func MarketplacePurchaseListener(marketplacePurchase event.MarketplacePurchase) (bool, error) {
	logger.Info("MarketplacePurchase event listener fired!")
	return true, nil
}

// Not Supported Yet
func MemberListener(member event.Member) (bool, error) {
	logger.Info("Member event listener fired!")
	return true, nil
}

// Not Supported Yet
func MembershipListener(membership event.Membership) (bool, error) {
	logger.Info("Membership event listener fired!")
	return true, nil
}

// Not Supported Yet
func OrgBlockListener(orgBlock event.OrgBlock) (bool, error) {
	logger.Info("OrgBlock event listener fired!")
	return true, nil
}

// Not Supported Yet
func OrganizationListener(organization event.Organization) (bool, error) {
	logger.Info("Organization event listener fired!")
	return true, nil
}

// Not Supported Yet
func PageBuildListener(pageBuild event.PageBuild) (bool, error) {
	logger.Info("PageBuild event listener fired!")
	return true, nil
}

// Not Supported Yet
func ProjectListener(project event.Project) (bool, error) {
	logger.Info("Project event listener fired!")
	return true, nil
}

// Not Supported Yet
func ProjectCardListener(projectCard event.ProjectCard) (bool, error) {
	logger.Info("ProjectCard event listener fired!")
	return true, nil
}

// Not Supported Yet
func ProjectColumnListener(projectColumn event.ProjectColumn) (bool, error) {
	logger.Info("ProjectColumn event listener fired!")
	return true, nil
}

// Not Supported Yet
func PublicListener(public event.Public) (bool, error) {
	logger.Info("Public event listener fired!")
	return true, nil
}

// Not Supported Yet
func PushListener(push event.Push) (bool, error) {
	logger.Info("Push event listener fired!")
	return true, nil
}

// Not Supported Yet
func ReleaseListener(release event.Release) (bool, error) {
	logger.Info("Release event listener fired!")
	return true, nil
}

// Not Supported Yet
func RepositoryListener(repository event.Repository) (bool, error) {
	logger.Info("Repository event listener fired!")
	return true, nil
}

// Not Supported Yet
func RepositoryImportListener(repositoryImport event.RepositoryImport) (bool, error) {
	logger.Info("RepositoryImport event listener fired!")
	return true, nil
}

// Not Supported Yet
func RepositoryVulnerabilityAlertListener(repositoryVulnerabilityAlert event.RepositoryVulnerabilityAlert) (bool, error) {
	logger.Info("RepositoryVulnerabilityAlert event listener fired!")
	return true, nil
}

// Not Supported Yet
func TeamListener(team event.Team) (bool, error) {
	logger.Info("Team event listener fired!")
	return true, nil
}

// Not Supported Yet
func TeamAddListener(teamAdd event.TeamAdd) (bool, error) {
	logger.Info("TeamAdd event listener fired!")
	return true, nil
}

// Test Command Listener for Issues
func IssuesTestCommandListener(command event.Command, issues event.Issues) (bool, error) {
	logger.Info("IssuesTestCommandListener event listener fired!")
	return true, nil
}

// Test Command Listener for Issues Comments
func IssueCommentTestCommandListener(command event.Command, issueComment event.IssueComment) (bool, error) {
	logger.Info("IssueCommentTestCommandListener event listener fired!")
	return true, nil
}
