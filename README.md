<p align="center">
    <img alt="Hamster Logo" src="https://raw.githubusercontent.com/Clivern/Hamster/master/logo/logo.png" height="80" />
    <h3 align="center">Hamster</h3>
    <p align="center">A Bot Toolkit for Github!</p>
</p>

## Documentation

### Config & Run The Application

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
    "repository_name": "Hamster"
}
```

Add a new webhook from `Settings > Webhooks`, Set the `Payload URL` to be `https://hamster.com/listen`, `Content type` as `JSON` and Add Your Webhook Secret.

And then run the application

```bash
$ go build squeal.go
$ ./squeal

// OR

$ go run squeal.go
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
func RawListener(raw event.Raw)(bool, error){
    pkg.Info("Raw event listener fired!")
    return true, nil
}
```

**[status event:](https://developer.github.com/v3/activity/events/types/#statusevent)** any time a Repository has a status update from the API, The following callback get called.
```go
// plugin/base.go

// Status Action
func StatusListener(status event.Status)(bool, error){
    pkg.Info("Status event listener fired!")
    return true, nil
}
```

**[watch event:](https://developer.github.com/v3/activity/events/types/#watchevent)** any time a User stars a Repository.
```go
// plugin/base.go

// Watch Action
func WatchListener(watch event.Watch)(bool, error){
    pkg.Info("Watch event listener fired!")
    return true, nil
}
```

**[issues event:](https://developer.github.com/v3/activity/events/types/#issuesevent)** any time an Issue is assigned, unassigned, labeled, unlabeled, opened, edited, milestoned, demilestoned, closed, or reopened.
```go
// plugin/base.go

// Issue Action
func IssuesListener(issues event.Issues)(bool, error){
    pkg.Info("Issues event listener fired!")
    return true, nil
}
```

**[issue_comment event:](https://developer.github.com/v3/activity/events/types/#issuecommentevent)** any time a comment on an issue is created, edited, or deleted.
```go
// plugin/base.go

// Issue Comment Action
func IssueCommentListener(issue_comment event.IssueComment)(bool, error){
    pkg.Info("IssueComment event listener fired!")
    return true, nil
}
```

**[push event:](https://developer.github.com/v3/activity/events/types/#pushevent)** Any Git push to a Repository, including editing tags or branches. Commits via API actions that update references are also counted. This is the default event.
```go
// plugin/base.go

// Push Action
func PushListener(push event.Push)(bool, error){
    pkg.Info("Push event listener fired!")
    return true, nil
}
```

**[create event:](https://developer.github.com/v3/activity/events/types/#createevent)** Any time a Branch or Tag is created.

```go
// plugin/base.go

// Create Action
func CreateListener(create event.Create)(bool, error){
    pkg.Info("Create event listener fired!")
    return true, nil
}
```

**[label event:](https://developer.github.com/v3/activity/events/types/#labelevent)** Any time a Label is created, edited, or deleted.

```go
// plugin/base.go

// Label Action
func LabelListener(label event.Label)(bool, error){
    pkg.Info("Label event listener fired!")
    return true, nil
}
```

**[delete event:](https://developer.github.com/v3/activity/events/types/#deleteevent)** Any time a branch or tag is deleted.

```go
// plugin/base.go

// Delete Action
func DeleteListener(delete event.Delete)(bool, error){
    pkg.Info("Delete event listener fired!")
    return true, nil
}
```

**[milestone event:](https://developer.github.com/v3/activity/events/types/#milestoneevent)** Any time a Milestone is created, closed, opened, edited, or deleted.

```go
// plugin/base.go

// Milestone Action
func MilestoneListener(milestone event.Milestone)(bool, error){
    pkg.Info("Milestone event listener fired!")
    return true, nil
}
```

All current supported events and the future events will be available on `plugin/base.go`. Also it is handy to add aditional callbacks so each event can have any number of callbacks.

Also please check [the latest github webhooks guide](https://developer.github.com/webhooks/).


### Build a Custom Commands

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
func IssuesTestCommandListener(command event.Command, issues event.Issues)(bool, error){
    pkg.Info("IssuesTestCommandListener event listener fired!")
    return true, nil
}

// Test Command Listener for Issues Comments
func IssueCommentTestCommandListener(command event.Command, issue_comment event.IssueComment)(bool, error){
    pkg.Info("IssueCommentTestCommandListener event listener fired!")
    return true, nil
}

// Run Command Callbacks
// Run Command Listener for Issues
func IssuesRunCommandListener(command event.Command, issues event.Issues)(bool, error){
    pkg.Info("IssuesTestCommandListener event listener fired!")
    return true, nil
}

// Run Command Listener for Issues Comments
func IssueCommentRunCommandListener(command event.Command, issue_comment event.IssueComment)(bool, error){
    pkg.Info("IssueCommentTestCommandListener event listener fired!")
    return true, nil
}
```

Now if you create a new issue or issue comment, the related callbacks will get notified with command object:

```go
/test
/test{option1}
/test{option1,option2}
/test{option1,option2,option3

/run
/run{option1}
/run{option1,option2}
/run{option1,option2,option3

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
    "github.com/clivern/hamster/pkg"
    "os"
)


github_api := &pkg.GithubAPI{
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

### Logging

We use [google/logger](https://github.com/google/logger) under the hood, make use of it or use these simple functions:

```go
import (
    "github.com/clivern/hamster/pkg"
)

pkg.Info("Info Goes Here!")
pkg.Warning("Warning Goes Here!")
pkg.Error("Error Goes Here!")
pkg.Fatal("Fatal Error Goes Here!")
```


## Badges

[![Build Status](https://travis-ci.org/Clivern/Hamster.svg?branch=master)](https://travis-ci.org/Clivern/Hamster)
[![GitHub license](https://img.shields.io/github/license/Clivern/Hamster.svg)](https://github.com/Clivern/Hamster/blob/master/LICENSE)


## Changelog

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