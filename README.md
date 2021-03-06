<p align="center">
    <img alt="Hamster Logo" src="https://raw.githubusercontent.com/Clivern/Hamster/master/assets/img/logo.png" height="80" />
    <h3 align="center">Hamster</h3>
    <p align="center">A Bot Toolkit for Github!</p>
</p>

## Documentation

### Config & Run The Application

Hamster uses [Go Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies. First Create a dist config file.

```bash
$ cp config.json config.dist.json
```

Then add your `app_mode`, `app_port`, `app_log_level`, `github_token`, `github_webhook_secret`, `repository_author` and `repository_name`

```json
{
    "app_mode": "prod",
    "app_port": "8080",
    "app_log_level": "info",
    "github_token": "...",
    "github_webhook_secret": "...",
    "repository_author": "Clivern",
    "repository_name": "Hamster",

    "app_domain": "example.com",
    "github_app_client_id": "..",
    "github_app_redirect_uri": "..",
    "github_app_allow_signup": "true",
    "github_app_scope": "..",
    "github_app_client_secret": ".."
}
```

You can config `app_domain` and the rest of github app configs `github_app_*` in case you need a github app not a personal bot.

Add a new webhook from `Settings > Webhooks`, Set the `Payload URL` to be `https://hamster.com/listen`, `Content type` as `JSON` and Add Your Webhook Secret.

And then run the application

```bash
$ go build hamster.go
$ ./hamster

// OR

$ go run hamster.go
```

Also running hamster with docker still an option. Just don't forget to update `GithubToken`, `GithubWebhookSecret`, `RepositoryAuthor` and `RepositoryName` inside `docker-compose.yml` file. Then run the following stuff

```bash
$ docker-compose build
$ docker-compose up -d
```

### Customize the Default Event Listeners

Anytime github call hamster listen endpoint, there will be a callback that get called with incoming data. For example when you get a status change call from github, the `StatusListener(status event.Status)` will get called. So do whatever you need inside this callback.

**any event:** any time listen endpoint get a call, the following callback get called.

```go
// plugin/base.go

// Any Action
func RawListener(raw event.Raw) (bool, error) {
    logger.Info("Raw event listener fired!")
    return true, nil
}
```

**[status event:](https://developer.github.com/v3/activity/events/types/#statusevent)** any time a Repository has a status update from the API, The following callback get called.
```go
// plugin/base.go

// Status Action
func StatusListener(status event.Status) (bool, error) {
    logger.Info("Status event listener fired!")
    return true, nil
}
```

**[watch event:](https://developer.github.com/v3/activity/events/types/#watchevent)** any time a User stars a Repository.
```go
// plugin/base.go

// Watch Action
func WatchListener(watch event.Watch) (bool, error) {
    logger.Info("Watch event listener fired!")
    return true, nil
}
```

**[issues event:](https://developer.github.com/v3/activity/events/types/#issuesevent)** any time an Issue is assigned, unassigned, labeled, unlabeled, opened, edited, milestoned, demilestoned, closed, or reopened.
```go
// plugin/base.go

// Issue Action
func IssuesListener(issues event.Issues) (bool, error) {
    logger.Info("Issues event listener fired!")
    return true, nil
}
```

**[issue_comment event:](https://developer.github.com/v3/activity/events/types/#issuecommentevent)** any time a comment on an issue is created, edited, or deleted.
```go
// plugin/base.go

// Issue Comment Action
func IssueCommentListener(issueComment event.IssueComment) (bool, error) {
    logger.Info("IssueComment event listener fired!")
    return true, nil
}
```

**[push event:](https://developer.github.com/v3/activity/events/types/#pushevent)** Any Git push to a Repository, including editing tags or branches. Commits via API actions that update references are also counted. This is the default event.
```go
// plugin/base.go

// Push Action
func PushListener(push event.Push) (bool, error) {
    logger.Info("Push event listener fired!")
    return true, nil
}
```

**[create event:](https://developer.github.com/v3/activity/events/types/#createevent)** Any time a Branch or Tag is created.

```go
// plugin/base.go

// Create Action
func CreateListener(create event.Create) (bool, error) {
    logger.Info("Create event listener fired!")
    return true, nil
}
```

**[label event:](https://developer.github.com/v3/activity/events/types/#labelevent)** Any time a Label is created, edited, or deleted.

```go
// plugin/base.go

// Label Action
func LabelListener(label event.Label) (bool, error) {
    logger.Info("Label event listener fired!")
    return true, nil
}
```

**[delete event:](https://developer.github.com/v3/activity/events/types/#deleteevent)** Any time a branch or tag is deleted.

```go
// plugin/base.go

// Delete Action
func DeleteListener(delete event.Delete) (bool, error) {
    logger.Info("Delete event listener fired!")
    return true, nil
}
```

**[milestone event:](https://developer.github.com/v3/activity/events/types/#milestoneevent)** Any time a Milestone is created, closed, opened, edited, or deleted.

```go
// plugin/base.go

// Milestone Action
func MilestoneListener(milestone event.Milestone) (bool, error) {
    logger.Info("Milestone event listener fired!")
    return true, nil
}
```

**[pull_request event:](https://developer.github.com/v3/activity/events/types/#pullrequestevent)** Any time a pull request is assigned, unassigned, labeled, unlabeled, opened, edited, closed, reopened, or synchronized (updated due to a new push in the branch that the pull request is tracking). Also any time a pull request review is requested, or a review request is removed.

```go
// plugin/base.go

// Pull Request Action
func PullRequestListener(pullRequest event.PullRequest) (bool, error) {
    logger.Info("PullRequest event listener fired!")
    return true, nil
}
```

**[pull_request_review event:](https://developer.github.com/v3/activity/events/types/#pullrequestreviewevent)** Any time a pull request review is submitted, edited, or dismissed.

```go
// plugin/base.go

// Pull Request Review Action
func PullRequestReviewListener(pullRequestReview event.PullRequestReview) (bool, error) {
    logger.Info("PullRequestReview event listener fired!")
    return true, nil
}
```

**[pull_request_review_comment event:](https://developer.github.com/v3/activity/events/types/#pullrequestreviewcommentevent)** Any time a comment on a pull request's unified diff is created, edited, or deleted (in the Files Changed tab).

```go
// plugin/base.go

// Pull Request Review Comment Action
func PullRequestReviewCommentListener(pullRequestReviewComment event.PullRequestReviewComment) (bool, error) {
    logger.Info("PullRequestReviewComment event listener fired!")
    return true, nil
}
```

All current supported events and the future events will be available on `plugin/base.go`. Also it is handy to add aditional callbacks so each event can have any number of callbacks.

Also please check [the latest github webhooks guide](https://developer.github.com/webhooks/).


### Build Custom Commands

In order to build an interactive bot, you will need to listen to a pre-defined commands that once your repo users type on an issue or a comment, your application get notified. Github don't support this by default but it is still possible to achieve this manually.

First you need to define you command and the callback on `internal/app/controller/listener.go`, exactly like the `test` command:

```go
// The default test command for issue comments
commands.RegisterIssueCommentAction("test", plugin.IssueCommentTestCommandListener)

//The new run command for issue comments
commands.RegisterIssueCommentAction("run", plugin.IssueCommentRunCommandListener)
```

```go
// The default test command for issues
commands.RegisterIssuesAction("test", plugin.IssuesTestCommandListener)

//The new run command for issues
commands.RegisterIssuesAction("run", plugin.IssuesRunCommandListener)
```

Then define the callbacks on `plugin/base.go` same as `test` commands callbacks:

```go
// Test Command Callbacks
// Test Command Listener for Issues
func IssuesTestCommandListener(command event.Command, issues event.Issues) (bool, error) {
    logger.Info("IssuesTestCommandListener event listener fired!")
    return true, nil
}

// Test Command Listener for Issues Comments
func IssueCommentTestCommandListener(command event.Command, issue_comment event.IssueComment) (bool, error) {
    logger.Info("IssueCommentTestCommandListener event listener fired!")
    return true, nil
}

// Run Command Callbacks
// Run Command Listener for Issues
func IssuesRunCommandListener(command event.Command, issues event.Issues) (bool, error) {
    logger.Info("IssuesTestCommandListener event listener fired!")
    return true, nil
}

// Run Command Listener for Issues Comments
func IssueCommentRunCommandListener(command event.Command, issue_comment event.IssueComment) (bool, error) {
    logger.Info("IssueCommentTestCommandListener event listener fired!")
    return true, nil
}
```

Now if you create a new issue or issue comment, the related callbacks will get notified with command object:

```go
/test
/test{option1}
/test{option1,option2}
/test{option1,option2,option3}

/run
/run{option1}
/run{option1,option2}
/run{option1,option2,option3}

The command object will be

event.Command{Name=test, Parameters=[]}
event.Command{Name=test, Parameters=[option1]}
event.Command{Name=test, Parameters=[option1 option2]}
event.Command{Name=test, Parameters=[option1 option2 option3]}

event.Command{Name=run, Parameters=[]}
event.Command{Name=run, Parameters=[option1]}
event.Command{Name=run, Parameters=[option1 option2]}
event.Command{Name=run, Parameters=[option1 option2 option3]}
```

### Create a Github Comment

```go
// for more info https://developer.github.com/v3/issues/comments/#create-a-comment

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Replace Message with the message and 1 with the issue id
created_comment, err := github_api.NewComment("Message", 1)

if err == nil {
    // created_comment.ID
    // check github.com/clivern/hamster/internal/app/response/created_comment.CreatedComment for available data
}else{
    // err.Error()
}
```

### Create a Label

```go
// for more info https://developer.github.com/v3/issues/labels/#create-a-label

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Get Repository label with name
// github_api.CreateLabel (name string, color string) (response.Label, error)
label, err := github_api.CreateLabel("Bug", "f29513")

if err == nil {
    // label of type response.Label
}else{
    // err.Error()
}
```

### Get a Label

```go
// for more info https://developer.github.com/v3/issues/labels/#get-a-single-label

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Get Repository label with name
// github_api.GetLabel (name string) (response.Label, error)
label, err := github_api.GetLabel("Bug")

if err == nil {
    // label of type response.Label
}else{
    // err.Error()
}
```

### Update a Label with Name

```go
// for more info https://developer.github.com/v3/issues/labels/#update-a-label

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Update label name and color
// github_api.UpdateLabel (currentName string, name string, color string) (response.Label, error)
label, err := github_api.UpdateLabel("CurrentName", "NewName", "b01f26")

if err == nil {
    // label of type response.Label
}else{
    // err.Error()
}
```

### Delete a Label with Name

```go
// for more info https://developer.github.com/v3/issues/labels/#delete-a-label

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Delete label with name
// github_api.DeleteLabel (name string) (bool, error)
ok, err := github_api.DeleteLabel("CurrentName")

if ok && err == nil {
    // label deleted
}else{
    // err.Error()
}
```

### Get Repository Labels List

```go
// for more info https://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Get Repository labels
// github_api.GetRepositoryLabels () ([]response.Label, error)
labels, err := github_api.GetRepositoryLabels()

if err == nil {
    // labels of type []response.Label
}else{
    // err.Error()
}
```

### Get Issue Labels List

```go
// for more info https://developer.github.com/v3/issues/labels/#list-labels-on-an-issue

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Get Repository issue labels with issue_id
// github_api.GetRepositoryIssueLabels (issueId int) ([]response.Label, error)
labels, err := github_api.GetRepositoryIssueLabels(9)

if err == nil {
    // labels of type []response.Label
}else{
    // err.Error()
}
```

### Remove Label from an Issue

```go
// for more info https://developer.github.com/v3/issues/labels/#remove-a-label-from-an-issue

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Remove a Label from an Issue
// github_api.RemoveLabelFromIssue (issueId int, labelName string) (bool, error)
ok, err := github_api.RemoveLabelFromIssue(9, "bug")

if ok && err == nil {
    // Label Removed
}else{
    // err.Error()
}
```

### Remove All Labels from an Issue

```go
// for more info https://developer.github.com/v3/issues/labels/#remove-all-labels-from-an-issue

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Remove a Label from an Issue
// github_api.RemoveAllLabelForIssue (issueId int) (bool, error)
ok, err := github_api.RemoveAllLabelForIssue(9)

if ok && err == nil {
    // All Labels Removed
}else{
    // err.Error()
}
```

### Get Milestone Labels List

```go
// for more info https://developer.github.com/v3/issues/labels/#get-labels-for-every-issue-in-a-milestone

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Get Milestone Labels List
// github_api.GetRepositoryMilestoneLabels (milestoneId int) ([]response.Label, error)
labels, err := github_api.GetRepositoryMilestoneLabels(9)

if err == nil {
    // labels of type []response.Label
}else{
    // err.Error()
}
```

### Add Labels to an Issue

```go
// for more info https://developer.github.com/v3/issues/labels/#add-labels-to-an-issue

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Add Labels to an Issue
// github_api.AddLabelsToIssue (issueId int, labels []string) ([]response.Label, error)
labels, err := github_api.AddLabelsToIssue(9, []string{"new-label", "another-label"})

if err == nil {
    // labels of type []response.Label
}else{
    // err.Error()
}
```

### Replace all Labels for an Issue

```go
// for more info https://developer.github.com/v3/issues/labels/#replace-all-labels-for-an-issue

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Replace all Labels for an Issue
// github_api.ReplaceAllLabelsForIssue (issueId int, labels []string) ([]response.Label, error)
labels, err := github_api.ReplaceAllLabelsForIssue(9, []string{"new-label", "another-label"})

if err == nil {
    // labels of type []response.Label
}else{
    // err.Error()
}
```


### Get PR Labels List

```go
// for more info https://developer.github.com/v3/issues/labels/#list-labels-on-an-issue

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Get Repository PR labels with PRId
// github_api.GetRepositoryPRLabels (PRId int) ([]response.Label, error)
labels, err := github_api.GetRepositoryPRLabels(9)

if err == nil {
    // labels of type []response.Label
}else{
    // err.Error()
}
```

### Remove Label from PR

```go
// for more info https://developer.github.com/v3/issues/labels/#remove-a-label-from-an-issue

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Remove a Label from PR
// github_api.RemoveLabelFromPR (PRId int, labelName string) (bool, error)
ok, err := github_api.RemoveLabelFromPR(9, "bug")

if ok && err == nil {
    // Label Removed
}else{
    // err.Error()
}
```

### Remove All Labels from PR

```go
// for more info https://developer.github.com/v3/issues/labels/#remove-all-labels-from-an-issue

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Remove a Label from PR
// github_api.RemoveAllLabelForPR (PRId int) (bool, error)
ok, err := github_api.RemoveAllLabelForPR(9)

if ok && err == nil {
    // All Labels Removed
}else{
    // err.Error()
}
```

### Add Labels to PR

```go
// for more info https://developer.github.com/v3/issues/labels/#add-labels-to-an-issue

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Add Labels to PR
// github_api.AddLabelsToPR (PRId int, labels []string) ([]response.Label, error)
labels, err := github_api.AddLabelsToPR(9, []string{"new-label", "another-label"})

if err == nil {
    // labels of type []response.Label
}else{
    // err.Error()
}
```

### Replace all Labels for PR

```go
// for more info https://developer.github.com/v3/issues/labels/#replace-all-labels-for-an-issue

import (
    "github.com/clivern/hamster/internal/app/pkg/github"
    "os"
)


github_api := &github.API{
    Token: os.Getenv("GithubToken"),
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// Replace all Labels for PR
// github_api.ReplaceAllLabelsForPR (PRId int, labels []string) ([]response.Label, error)
labels, err := github_api.ReplaceAllLabelsForPR(9, []string{"new-label", "another-label"})

if err == nil {
    // labels of type []response.Label
}else{
    // err.Error()
}
```

### Authorizing OAuth Apps

You can enable other users to authorize your OAuth App. First configure the app credentials on `config.dist.json` or `docker-compose.yml`

```
// config.dist.json
    "app_domain": "example.com",
    "github_app_client_id": "..",
    "github_app_redirect_uri": "..",
    "github_app_allow_signup": "false",
    "github_app_scope": "", // It can be empty
    "github_app_client_secret": ".."
```
```
// docker-compose.yml
     - GithubAppClientID=ValueHere
     - GithubAppRedirectURI=ValueHere
     - GithubAppAllowSignup=true
     - GithubAppScope= // It can be empty
     - GithubAppClientSecret=ValueHere
     - AppDomain=example.com
```

The github app should configured to use `http://example.com/auth` as redirect URL.

If you run the application, the authorize URL on `/login` page should be something like that:
```
https://github.com/login/oauth/authorize?allow_signup=true&client_id=Iv1.eda..&redirect_uri=https%3A%2F%2F3fa8b997.ngrok.io%2Fauth&scope=&state=5a0a8c973e93ed82820f7896dddb1df70a3dce62
```

If you click authorize and authorized the app, github will send you back to hamster `/auth` route with a code and the state. Hamster will use that code to fetch the `accessToken` for you or any user. You can use the `accessToken` to do all subsequent github API Calls.

### Github Check Runs

To create a status check:

```go

import (
    "github.com/clivern/hamster/internal/app/response"
    "github.com/clivern/hamster/internal/app/sender"
    "github.com/clivern/hamster/internal/app/pkg/github"

    "os"
    "time"
    "fmt"
)

output := sender.Output{
    Title: "CI Report",
    Summary: "Output Summary",
    Text: "Some Text Goes Here",
}

checkRun := sender.CheckRun{
    Name: "CI Status",
    HeadSha: "6c46684560f9fed86be6fd87f9dbaa91e4b242c9",
    Status: "in_progress",
    DetailsURL: "http://clivern.com/ci/5",
    ExternalID: "43",
    StartedAt: time.Now().UTC().Format(time.RFC3339),
    Output: output,
}

var checkRunResponse response.CheckRun

github_api := &github.API{
    Token: "5688665c9184800e...", # Token via a GitHub App.
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// method CreateCheckRun(CheckRun sender.CheckRun) (response.CheckRun, error)
checkRunResponse, err := github_api.CreateCheckRun(checkRun)

if err == nil{
    fmt.Println(checkRunResponse.ID)
}else{
    fmt.Println(err.Error())
}
```

To update a status check:

```go
import (
    "github.com/clivern/hamster/internal/app/response"
    "github.com/clivern/hamster/internal/app/sender"
    "github.com/clivern/hamster/internal/app/pkg/github"

    "os"
    "time"
    "fmt"
)

output := sender.Output{
    Title: "CI Report",
    Summary: "Final Output Summary",
    Text: "Some Final Text Goes Here",
}

checkRun := sender.CheckRun{
    Name: "CI Status",
    HeadSha: "6c46684560f9fed86be6fd87f9dbaa91e4b242c9",
    Status: "completed",
    DetailsURL: "http://clivern.com/ci/5",
    ExternalID: "43",
    CompletedAt: time.Now().UTC().Format(time.RFC3339),
    Conclusion: "success",
    Output: output,
}

var checkRunResponse response.CheckRun

github_api := &github.API{
    Token: "5688665c9184800e...", // Token via a GitHub App.
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

// method UpdateCheckRun(ID int, CheckRun sender.CheckRun) (response.CheckRun, error)
checkRunResponse, err := github_api.UpdateCheckRun(25165135, checkRun)

if err == nil{
    fmt.Println(checkRunResponse.ID)
}else{
    fmt.Println(err.Error())
}
```

To get a status check with ID:

```go
import (
    "github.com/clivern/hamster/internal/app/response"
    "github.com/clivern/hamster/internal/app/pkg/github"

    "os"
    "fmt"
)

var checkRunResponse response.CheckRun

github_api := &github.API{
    Token: "5688665c9184800e...", // Token via a GitHub App.
    Author: os.Getenv("RepositoryAuthor"),
    Repository: os.Getenv("RepositoryName"),
}

checkRunResponse, err := github_api.GetCheckRun(25165135)

if err == nil{
    fmt.Println(checkRunResponse.ID)
}else{
    fmt.Println(err.Error())
}
```

### Logging

We use [google/logger](https://github.com/google/logger) under the hood, make use of it or use these simple functions:

```go
import (
    "github.com/clivern/hamster/internal/app/pkg/logger"
)

logger.Info("Info Goes Here!")
logger.Infoln("Infoln Goes Here!")
logger.Infof("Infof %s Here!", "Goes")

logger.Warning("Warning Goes Here!")
logger.Warningln("Warningln Goes Here!")
logger.Warningf("Warningf %s Here!", "Goes")

logger.Error("Error Goes Here!")
logger.Errorln("Errorln Goes Here!")
logger.Errorf("Errorf %s Here!", "Goes")

logger.Fatal("Fatal Goes Here!")
logger.Fatalln("Fatalln Goes Here!")
logger.Fatalf("Fatalf %s Here!", "Goes")
```


## Badges

[![Build Status](https://travis-ci.org/Clivern/Hamster.svg?branch=master)](https://travis-ci.org/Clivern/Hamster)
[![GitHub license](https://img.shields.io/github/license/Clivern/Hamster.svg)](https://github.com/Clivern/Hamster/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-3.1.0-red.svg)](https://github.com/Clivern/Hamster/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/Clivern/Hamster)](https://goreportcard.com/report/github.com/Clivern/Hamster)


## Changelog

* Version 3.1.0:
```
Switch to go 1.11 modules.
```

* Version 3.0.1:
```
Fix ineffassign for some vars.
```

* Version 3.0.0:
```
More Enhancements.
```

* Version 2.0.0:
```
Add More Events.
Add Labels & Comments API to Github pkg.
Custom Commands.
OAuth Apps Support.
Check Runs Support.
```

* Version 1.1.1:
```
Add Logger Package.
```

* Version 1.1.0:
```
Add new events watch, issues and issue_comment.
Fix dockerfile & docker-compose.
```

* Version 1.0.0:
```
Initial Release.
```


## Acknowledgements

© 2018, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Hamster** is authored and maintained by [@clivern](http://github.com/clivern).
