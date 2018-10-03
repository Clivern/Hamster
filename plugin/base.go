package plugin


// Here we define all our custom actions and it will
// get executed once we get a request from github
import (
    "github.com/clivern/hamster/internal/app/event"
    "fmt"
)

// Status Action
func StatusListener(status event.Status)(bool, error){
    fmt.Printf("StatusListener Fired: %s \n", status.Sha)
    return true, nil
}

// Watch Action
func WatchListener(watch event.Watch)(bool, error){
    fmt.Printf("WatchListener Fired: %s \n", watch.Action)
    return true, nil
}

// Issue Action
func IssuesListener(issues event.Issues)(bool, error){
    fmt.Printf("IssuesListener Fired: %s \n", issues.Action)
    return true, nil
}

// Issue Comment Action
func IssueCommentListener(issue_comment event.IssueComment)(bool, error){
    fmt.Printf("IssueCommentListener Fired: %s \n", issue_comment.Action)
    return true, nil
}


// Not Supported Yet
func CheckRunListener(check_run event.CheckRun)(bool, error){
    fmt.Println("CheckRunListener Fired")
    return true, nil
}

// Not Supported Yet
func CheckSuiteListener(check_suite event.CheckSuite)(bool, error){
    fmt.Printf("CheckSuiteListener Fired")
    return true, nil
}

// Not Supported Yet
func CommitCommentListener(commit_comment event.CommitComment)(bool, error){
    fmt.Printf("CommitCommentListener Fired")
    return true, nil
}

// Not Supported Yet
func CreateListener(create event.Create)(bool, error){
    fmt.Printf("CreateListener Fired")
    return true, nil
}

// Not Supported Yet
func DeleteListener(delete event.Delete)(bool, error){
    fmt.Printf("DeleteListener Fired")
    return true, nil
}

// Not Supported Yet
func DeploymentListener(deployment event.Deployment)(bool, error){
    fmt.Printf("DeploymentListener Fired")
    return true, nil
}

// Not Supported Yet
func DeploymentStatusListener(deployment_status event.DeploymentStatus)(bool, error){
    fmt.Printf("DeploymentStatusListener Fired")
    return true, nil
}

// Not Supported Yet
func ForkListener(fork event.Fork)(bool, error){
    fmt.Printf("ForkListener Fired")
    return true, nil
}

// Not Supported Yet
func GithubAppAuthorizationListener(github_app_authorization event.GithubAppAuthorization)(bool, error){
    fmt.Printf("GithubAppAuthorizationListener Fired")
    return true, nil
}

// Not Supported Yet
func GollumListener(gollum event.Gollum)(bool, error){
    fmt.Printf("GollumListener Fired")
    return true, nil
}

// Not Supported Yet
func InstallationListener(installation event.Installation)(bool, error){
    fmt.Printf("InstallationListener Fired")
    return true, nil
}

// Not Supported Yet
func InstallationRepositoriesListener(installation_repositories event.InstallationRepositories)(bool, error){
    fmt.Printf("InstallationRepositoriesListener Fired")
    return true, nil
}

// Not Supported Yet
func LabelListener(label event.Label)(bool, error){
    fmt.Printf("LabelListener Fired")
    return true, nil
}

// Not Supported Yet
func MarketplacePurchaseListener(marketplace_purchase event.MarketplacePurchase)(bool, error){
    fmt.Printf("MarketplacePurchaseListener Fired")
    return true, nil
}

// Not Supported Yet
func MemberListener(member event.Member)(bool, error){
    fmt.Printf("MemberListener Fired")
    return true, nil
}

// Not Supported Yet
func MembershipListener(membership event.Membership)(bool, error){
    fmt.Printf("MembershipListener Fired")
    return true, nil
}

// Not Supported Yet
func MilestoneListener(milestone event.Milestone)(bool, error){
    fmt.Printf("MilestoneListener Fired")
    return true, nil
}

// Not Supported Yet
func OrgBlockListener(org_block event.OrgBlock)(bool, error){
    fmt.Printf("OrgBlockListener Fired")
    return true, nil
}

// Not Supported Yet
func OrganizationListener(organization event.Organization)(bool, error){
    fmt.Printf("OrganizationListener Fired")
    return true, nil
}

// Not Supported Yet
func PageBuildListener(page_build event.PageBuild)(bool, error){
    fmt.Printf("PageBuildListener Fired")
    return true, nil
}

// Not Supported Yet
func ProjectListener(project event.Project)(bool, error){
    fmt.Printf("ProjectListener Fired")
    return true, nil
}

// Not Supported Yet
func ProjectCardListener(project_card event.ProjectCard)(bool, error){
    fmt.Printf("ProjectCardListener Fired")
    return true, nil
}

// Not Supported Yet
func ProjectColumnListener(project_column event.ProjectColumn)(bool, error){
    fmt.Printf("ProjectColumnListener Fired")
    return true, nil
}

// Not Supported Yet
func PublicListener(public event.Public)(bool, error){
    fmt.Printf("PublicListener Fired")
    return true, nil
}

// Not Supported Yet
func PullRequestListener(pull_request event.PullRequest)(bool, error){
    fmt.Printf("PullRequestListener Fired")
    return true, nil
}

// Not Supported Yet
func PullRequestReviewListener(pull_request_review event.PullRequestReview)(bool, error){
    fmt.Printf("PullRequestReviewListener Fired")
    return true, nil
}

// Not Supported Yet
func PullRequestReviewCommentListener(pull_request_review_comment event.PullRequestReviewComment)(bool, error){
    fmt.Printf("PullRequestReviewCommentListener Fired")
    return true, nil
}

// Not Supported Yet
func PushListener(push event.Push)(bool, error){
    fmt.Printf("PushListener Fired")
    return true, nil
}

// Not Supported Yet
func ReleaseListener(release event.Release)(bool, error){
    fmt.Printf("ReleaseListener Fired")
    return true, nil
}

// Not Supported Yet
func RepositoryListener(repository event.Repository)(bool, error){
    fmt.Printf("RepositoryListener Fired")
    return true, nil
}

// Not Supported Yet
func RepositoryImportListener(repository_import event.RepositoryImport)(bool, error){
    fmt.Printf("RepositoryImportListener Fired")
    return true, nil
}

// Not Supported Yet
func RepositoryVulnerabilityAlertListener(repository_vulnerability_alert event.RepositoryVulnerabilityAlert)(bool, error){
    fmt.Printf("RepositoryVulnerabilityAlertListener Fired")
    return true, nil
}

// Not Supported Yet
func TeamListener(team event.Team)(bool, error){
    fmt.Printf("TeamListener Fired")
    return true, nil
}

// Not Supported Yet
func TeamAddListener(team_add event.TeamAdd)(bool, error){
    fmt.Printf("TeamAddListener Fired")
    return true, nil
}