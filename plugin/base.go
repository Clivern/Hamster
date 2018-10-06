package plugin


// Here we define all our custom actions and it will
// get executed once we get a request from github
import (
    "github.com/clivern/hamster/internal/app/event"
    "github.com/clivern/hamster/pkg"
)

// Any Action
func RawListener(raw event.Raw)(bool, error){
    pkg.Info("Raw event listener fired!")
    return true, nil
}

// Status Action
func StatusListener(status event.Status)(bool, error){
    pkg.Info("Status event listener fired!")
    return true, nil
}

// Watch Action
func WatchListener(watch event.Watch)(bool, error){
    pkg.Info("Watch event listener fired!")
    return true, nil
}

// Issue Action
func IssuesListener(issues event.Issues)(bool, error){
    pkg.Info("Issues event listener fired!")
    return true, nil
}

// Issue Comment Action
func IssueCommentListener(issue_comment event.IssueComment)(bool, error){
    pkg.Info("IssueComment event listener fired!")
    return true, nil
}

// Create Action
func CreateListener(create event.Create)(bool, error){
    pkg.Info("Create event listener fired!")
    return true, nil
}

// Label Action
func LabelListener(label event.Label)(bool, error){
    pkg.Info("Label event listener fired!")
    return true, nil
}

// Delete Action
func DeleteListener(delete event.Delete)(bool, error){
    pkg.Info("Delete event listener fired!")
    return true, nil
}

// Milestone Action
func MilestoneListener(milestone event.Milestone)(bool, error){
    pkg.Info("Milestone event listener fired!")
    return true, nil
}

// Not Supported Yet
func CheckRunListener(check_run event.CheckRun)(bool, error){
    pkg.Info("CheckRun event listener fired!")
    return true, nil
}

// Not Supported Yet
func CheckSuiteListener(check_suite event.CheckSuite)(bool, error){
    pkg.Info("CheckSuite event listener fired!")
    return true, nil
}

// Not Supported Yet
func CommitCommentListener(commit_comment event.CommitComment)(bool, error){
    pkg.Info("CommitComment event listener fired!")
    return true, nil
}

// Not Supported Yet
func DeploymentListener(deployment event.Deployment)(bool, error){
    pkg.Info("Deployment event listener fired!")
    return true, nil
}

// Not Supported Yet
func DeploymentStatusListener(deployment_status event.DeploymentStatus)(bool, error){
    pkg.Info("DeploymentStatus event listener fired!")
    return true, nil
}

// Not Supported Yet
func ForkListener(fork event.Fork)(bool, error){
    pkg.Info("Fork event listener fired!")
    return true, nil
}

// Not Supported Yet
func GithubAppAuthorizationListener(github_app_authorization event.GithubAppAuthorization)(bool, error){
    pkg.Info("GithubAppAuthorization event listener fired!")
    return true, nil
}

// Not Supported Yet
func GollumListener(gollum event.Gollum)(bool, error){
    pkg.Info("Gollum event listener fired!")
    return true, nil
}

// Not Supported Yet
func InstallationListener(installation event.Installation)(bool, error){
    pkg.Info("Installation event listener fired!")
    return true, nil
}

// Not Supported Yet
func InstallationRepositoriesListener(installation_repositories event.InstallationRepositories)(bool, error){
    pkg.Info("InstallationRepositories event listener fired!")
    return true, nil
}

// Not Supported Yet
func MarketplacePurchaseListener(marketplace_purchase event.MarketplacePurchase)(bool, error){
    pkg.Info("MarketplacePurchase event listener fired!")
    return true, nil
}

// Not Supported Yet
func MemberListener(member event.Member)(bool, error){
    pkg.Info("Member event listener fired!")
    return true, nil
}

// Not Supported Yet
func MembershipListener(membership event.Membership)(bool, error){
    pkg.Info("Membership event listener fired!")
    return true, nil
}

// Not Supported Yet
func OrgBlockListener(org_block event.OrgBlock)(bool, error){
    pkg.Info("OrgBlock event listener fired!")
    return true, nil
}

// Not Supported Yet
func OrganizationListener(organization event.Organization)(bool, error){
    pkg.Info("Organization event listener fired!")
    return true, nil
}

// Not Supported Yet
func PageBuildListener(page_build event.PageBuild)(bool, error){
    pkg.Info("PageBuild event listener fired!")
    return true, nil
}

// Not Supported Yet
func ProjectListener(project event.Project)(bool, error){
    pkg.Info("Project event listener fired!")
    return true, nil
}

// Not Supported Yet
func ProjectCardListener(project_card event.ProjectCard)(bool, error){
    pkg.Info("ProjectCard event listener fired!")
    return true, nil
}

// Not Supported Yet
func ProjectColumnListener(project_column event.ProjectColumn)(bool, error){
    pkg.Info("ProjectColumn event listener fired!")
    return true, nil
}

// Not Supported Yet
func PublicListener(public event.Public)(bool, error){
    pkg.Info("Public event listener fired!")
    return true, nil
}

// Not Supported Yet
func PullRequestListener(pull_request event.PullRequest)(bool, error){
    pkg.Info("PullRequest event listener fired!")
    return true, nil
}

// Not Supported Yet
func PullRequestReviewListener(pull_request_review event.PullRequestReview)(bool, error){
    pkg.Info("PullRequestReview event listener fired!")
    return true, nil
}

// Not Supported Yet
func PullRequestReviewCommentListener(pull_request_review_comment event.PullRequestReviewComment)(bool, error){
    pkg.Info("PullRequestReviewComment event listener fired!")
    return true, nil
}

// Not Supported Yet
func PushListener(push event.Push)(bool, error){
    pkg.Info("Push event listener fired!")
    return true, nil
}

// Not Supported Yet
func ReleaseListener(release event.Release)(bool, error){
    pkg.Info("Release event listener fired!")
    return true, nil
}

// Not Supported Yet
func RepositoryListener(repository event.Repository)(bool, error){
    pkg.Info("Repository event listener fired!")
    return true, nil
}

// Not Supported Yet
func RepositoryImportListener(repository_import event.RepositoryImport)(bool, error){
    pkg.Info("RepositoryImport event listener fired!")
    return true, nil
}

// Not Supported Yet
func RepositoryVulnerabilityAlertListener(repository_vulnerability_alert event.RepositoryVulnerabilityAlert)(bool, error){
    pkg.Info("RepositoryVulnerabilityAlert event listener fired!")
    return true, nil
}

// Not Supported Yet
func TeamListener(team event.Team)(bool, error){
    pkg.Info("Team event listener fired!")
    return true, nil
}

// Not Supported Yet
func TeamAddListener(team_add event.TeamAdd)(bool, error){
    pkg.Info("TeamAdd event listener fired!")
    return true, nil
}