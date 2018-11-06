// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package plugin

import (
	"github.com/clivern/hamster/internal/app/event"
	"github.com/clivern/hamster/internal/app/pkg/logger"
)

// RawListener Action
func RawListener(raw event.Raw) (bool, error) {
	logger.Infof("Raw event listener fired [%v]!", raw)
	return true, nil
}

// StatusListener Action
func StatusListener(status event.Status) (bool, error) {
	logger.Infof("Status event listener fired [%v]!", status)
	return true, nil
}

// WatchListener Action
func WatchListener(watch event.Watch) (bool, error) {
	logger.Infof("Watch event listener fired [%v]!", watch)
	return true, nil
}

// IssuesListener Action
func IssuesListener(issues event.Issues) (bool, error) {
	logger.Infof("Issues event listener fired [%v]!", issues)
	return true, nil
}

// IssueCommentListener Action
func IssueCommentListener(issueComment event.IssueComment) (bool, error) {
	logger.Infof("IssueComment event listener fired [%v]!", issueComment)
	return true, nil
}

// CreateListener Action
func CreateListener(create event.Create) (bool, error) {
	logger.Infof("Create event listener fired [%v]!", create)
	return true, nil
}

// LabelListener Action
func LabelListener(label event.Label) (bool, error) {
	logger.Infof("Label event listener fired [%v]!", label)
	return true, nil
}

// DeleteListener Action
func DeleteListener(delete event.Delete) (bool, error) {
	logger.Infof("Delete event listener fired [%v]!", delete)
	return true, nil
}

// MilestoneListener Action
func MilestoneListener(milestone event.Milestone) (bool, error) {
	logger.Infof("Milestone event listener fired [%v]!", milestone)
	return true, nil
}

// PullRequestListener Action
func PullRequestListener(pullRequest event.PullRequest) (bool, error) {
	logger.Infof("PullRequest event listener fired [%v]!", pullRequest)
	return true, nil
}

// PullRequestReviewListener Action
func PullRequestReviewListener(pullRequestReview event.PullRequestReview) (bool, error) {
	logger.Infof("PullRequestReview event listener fired [%v]!", pullRequestReview)
	return true, nil
}

// PullRequestReviewCommentListener Action
func PullRequestReviewCommentListener(pullRequestReviewComment event.PullRequestReviewComment) (bool, error) {
	logger.Infof("PullRequestReviewComment event listener fired [%v]!", pullRequestReviewComment)
	return true, nil
}

// CheckRunListener Action (Not Supported Yet)
func CheckRunListener(checkRun event.CheckRun) (bool, error) {
	logger.Infof("CheckRun event listener fired [%v]!", checkRun)
	return true, nil
}

// CheckSuiteListener Action (Not Supported Yet)
func CheckSuiteListener(checkSuite event.CheckSuite) (bool, error) {
	logger.Infof("CheckSuite event listener fired [%v]!", checkSuite)
	return true, nil
}

// CommitCommentListener Action (Not Supported Yet)
func CommitCommentListener(commitComment event.CommitComment) (bool, error) {
	logger.Infof("CommitComment event listener fired [%v]!", commitComment)
	return true, nil
}

// DeploymentListener Action (Not Supported Yet)
func DeploymentListener(deployment event.Deployment) (bool, error) {
	logger.Infof("Deployment event listener fired [%v]!", deployment)
	return true, nil
}

// DeploymentStatusListener Action (Not Supported Yet)
func DeploymentStatusListener(deploymentStatus event.DeploymentStatus) (bool, error) {
	logger.Infof("DeploymentStatus event listener fired [%v]!", deploymentStatus)
	return true, nil
}

// ForkListener Action (Not Supported Yet)
func ForkListener(fork event.Fork) (bool, error) {
	logger.Infof("Fork event listener fired [%v]!", fork)
	return true, nil
}

// GithubAppAuthorizationListener Action (Not Supported Yet)
func GithubAppAuthorizationListener(githubAppAuthorization event.GithubAppAuthorization) (bool, error) {
	logger.Infof("GithubAppAuthorization event listener fired [%v]!", githubAppAuthorization)
	return true, nil
}

// GollumListener Action (Not Supported Yet)
func GollumListener(gollum event.Gollum) (bool, error) {
	logger.Infof("Gollum event listener fired [%v]!", gollum)
	return true, nil
}

// InstallationListener Action (Not Supported Yet)
func InstallationListener(installation event.Installation) (bool, error) {
	logger.Infof("Installation event listener fired [%v]!", installation)
	return true, nil
}

// InstallationRepositoriesListener Action (Not Supported Yet)
func InstallationRepositoriesListener(installationRepositories event.InstallationRepositories) (bool, error) {
	logger.Infof("InstallationRepositories event listener fired [%v]!", installationRepositories)
	return true, nil
}

// MarketplacePurchaseListener Action (Not Supported Yet)
func MarketplacePurchaseListener(marketplacePurchase event.MarketplacePurchase) (bool, error) {
	logger.Infof("MarketplacePurchase event listener fired [%v]!", marketplacePurchase)
	return true, nil
}

// MemberListener Action (Not Supported Yet)
func MemberListener(member event.Member) (bool, error) {
	logger.Infof("Member event listener fired [%v]!", member)
	return true, nil
}

// MembershipListener Action (Not Supported Yet)
func MembershipListener(membership event.Membership) (bool, error) {
	logger.Infof("Membership event listener fired [%v]!", membership)
	return true, nil
}

// OrgBlockListener Action (Not Supported Yet)
func OrgBlockListener(orgBlock event.OrgBlock) (bool, error) {
	logger.Infof("OrgBlock event listener fired [%v]!", orgBlock)
	return true, nil
}

// OrganizationListener Action (Not Supported Yet)
func OrganizationListener(organization event.Organization) (bool, error) {
	logger.Infof("Organization event listener fired [%v]!", organization)
	return true, nil
}

// PageBuildListener Action (Not Supported Yet)
func PageBuildListener(pageBuild event.PageBuild) (bool, error) {
	logger.Infof("PageBuild event listener fired [%v]!", pageBuild)
	return true, nil
}

// ProjectListener Action (Not Supported Yet)
func ProjectListener(project event.Project) (bool, error) {
	logger.Infof("Project event listener fired [%v]!", project)
	return true, nil
}

// ProjectCardListener Action (Not Supported Yet)
func ProjectCardListener(projectCard event.ProjectCard) (bool, error) {
	logger.Infof("ProjectCard event listener fired [%v]!", projectCard)
	return true, nil
}

// ProjectColumnListener Action (Not Supported Yet)
func ProjectColumnListener(projectColumn event.ProjectColumn) (bool, error) {
	logger.Infof("ProjectColumn event listener fired [%v]!", projectColumn)
	return true, nil
}

// PublicListener Action (Not Supported Yet)
func PublicListener(public event.Public) (bool, error) {
	logger.Infof("Public event listener fired [%v]!", public)
	return true, nil
}

// PushListener Action (Not Supported Yet)
func PushListener(push event.Push) (bool, error) {
	logger.Infof("Push event listener fired [%v]!", push)
	return true, nil
}

// ReleaseListener Action (Not Supported Yet)
func ReleaseListener(release event.Release) (bool, error) {
	logger.Infof("Release event listener fired [%v]!", release)
	return true, nil
}

// RepositoryListener Action (Not Supported Yet)
func RepositoryListener(repository event.Repository) (bool, error) {
	logger.Infof("Repository event listener fired [%v]!", repository)
	return true, nil
}

// RepositoryImportListener Action (Not Supported Yet)
func RepositoryImportListener(repositoryImport event.RepositoryImport) (bool, error) {
	logger.Infof("RepositoryImport event listener fired [%v]!", repositoryImport)
	return true, nil
}

// RepositoryVulnerabilityAlertListener Action (Not Supported Yet)
func RepositoryVulnerabilityAlertListener(repositoryVulnerabilityAlert event.RepositoryVulnerabilityAlert) (bool, error) {
	logger.Infof("RepositoryVulnerabilityAlert event listener fired [%v]!", repositoryVulnerabilityAlert)
	return true, nil
}

// TeamListener Action (Not Supported Yet)
func TeamListener(team event.Team) (bool, error) {
	logger.Infof("Team event listener fired [%v]!", team)
	return true, nil
}

// TeamAddListener Action (Not Supported Yet)
func TeamAddListener(teamAdd event.TeamAdd) (bool, error) {
	logger.Infof("TeamAdd event listener fired [%v]!", teamAdd)
	return true, nil
}

// IssuesTestCommandListener Command
func IssuesTestCommandListener(command event.Command, issues event.Issues) (bool, error) {
	logger.Infof("IssuesTestCommandListener event listener fired [%v] [%v]!", command, issues)
	return true, nil
}

// IssueCommentTestCommandListener Command
func IssueCommentTestCommandListener(command event.Command, issueComment event.IssueComment) (bool, error) {
	logger.Infof("IssueCommentTestCommandListener event listener fired! [%v] [%v]!", command, issueComment)
	return true, nil
}
