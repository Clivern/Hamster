<p align="center">
  <img alt="Hamster Logo" src="https://raw.githubusercontent.com/Clivern/Hamster/feature/listen/logo/logo.png" height="80" />
  <h3 align="center">Hamster</h3>
  <p align="center">An Opinionated Github Bot!</p>
</p>

---

## Documentation

### Create a Comment:

```bash
$ export GITHUB_TOKEN=b1...
```

```go
// for more info https://developer.github.com/v3/issues/comments/#create-a-comment

import (
    "github.com/clivern/hamster/pkg"
    "os"
)


github_api := &pkg.GithubAPI{
    Token: os.Getenv("GITHUB_TOKEN"),
    Author: "Clivern",
    Repository: "Hamster",
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

### Register Actions or Listeners:

Once any action happen on github like a new commit, new issue, new comment .... etc. You will get a `POST` request to your defined Hamster URL.

In order to create custom github listeners, please check the following

```go
import (
    "github.com/clivern/hamster/internal/app/receiver"
    "github.com/clivern/hamster/internal/app/listener"
    "fmt"
)


// Let's assume we get the following commit object with ID=1 and SHA is "Hi"
var commit receiver.Commit

commit.ID = 1
commit.Sha = "Hi"


var actions listener.Action

actions.RegisterCommitAction(func(commit receiver.Commit)(bool, error){
    // Interact with commit object
    fmt.Printf("Action 1 ID: %d\n", commit.ID) // ~ returns 1
    fmt.Printf("Action 1 SHA: %s\n", commit.Sha) // ~ returns Hi
    return true, nil
})


actions.RegisterCommitAction(func(commit receiver.Commit)(bool, error){
    // Interact with commit object
    fmt.Printf("Action 2 ID: %d\n", commit.ID) // ~ returns 1
    fmt.Printf("Action 2 SHA: %s\n", commit.Sha) // ~ returns Hi
    return true, nil
})

// Execute all actions
actions.ExecuteCommitActions(commit)
```

### Verify the Incoming Github Actions

```go
import (
    "github.com/gin-gonic/gin"
    "github.com/clivern/hamster/internal/app/listener"
    "os"
    "fmt"
)

// using gin framework to create a new route
r.POST("/parser-test", func (c *gin.Context) {

    rawBody, _ := c.GetRawData()

    parser := &listener.Parser{
        UserAgent: c.GetHeader("User-Agent"),
        GithubDelivery: c.GetHeader("X-GitHub-Delivery"),
        GitHubEvent: c.GetHeader("X-GitHub-Event"),
        HubSignature: c.GetHeader("X-Hub-Signature"),
        Body: string(rawBody),
    }

    fmt.Println(parser.GetUserAgent())
    fmt.Println(parser.GetGithubDelivery())
    fmt.Println(parser.GetGitHubEvent())
    fmt.Println(parser.GetHubSignature())
    // if true this means that request coming from github with your secret
    fmt.Println(parser.VerifySignature(os.Getenv("GITHUB_WEBHOOK_SECRET")))

    c.JSON(200, gin.H{
        "status": "ok",
    })
})
```

```bash
// Github request is like that
$ curl -X POST \
    -H 'User-Agent: GitHub-Hookshot/997d58f' \
    -H 'X-GitHub-Delivery: 06743070-c63e-11e8-9f16-48c86eb6bbe5' \
    -H 'X-GitHub-Event: issue_comment' \
    -H 'X-Hub-Signature: sha1=fr4254ba00119b76734567de793fadcef43c9865' \
    -d '{"action":"created",...}' 'http://localhost:8080/parser-test'
```

### Working Examples


## Badges

[![Build Status](https://travis-ci.org/Clivern/Hamster.svg?branch=master)](https://travis-ci.org/Clivern/Hamster)
[![GitHub license](https://img.shields.io/github/license/Clivern/Hamster.svg)](https://github.com/Clivern/Hamster/blob/master/LICENSE)


## Changelog

* Version 1.0.0:
```
Initial Release.
```


## Acknowledgements

© 2018, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Hamster** is authored and maintained by [@clivern](http://github.com/clivern).