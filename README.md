<p align="center">
  <img alt="Hamster Logo" src="https://raw.githubusercontent.com/Clivern/Hamster/feature/listen/logo/logo.png" height="80" />
  <h3 align="center">Hamster</h3>
  <p align="center">A Bot Toolkit for Github!</p>
</p>

## Documentation

### Config The Application:

```bash
$ cp config.json config.dist.json
```

Then add your `github_token`, `github_webhook_secret`, `repository_author` and `repository_name`

```json
{
    "github_token": "...",
    "github_webhook_secret": "...",
    "repository_author": "Clivern",
    "repository_name": "Hamster"
}
```

### Create a Comment:

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