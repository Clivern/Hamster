<p align="center">
  <img alt="Hamster Logo" src="https://raw.githubusercontent.com/Clivern/Hamster/feature/listen/logo/logo.png" height="80" />
  <h3 align="center">Hamster</h3>
  <p align="center">An Opinionated Github Bot!</p>
</p>

---

## Documentation

To create a comment

```bash
$ export GITHUB_TOKEN=b1...
```

```go
import (
    "github.com/clivern/hamster/pkg"
    "os"
)


github_api := &pkg.GithubAPI{
    Token: os.Getenv("GITHUB_TOKEN"),
    Author:"Clivern",
    Repository:"Hamster",
}

// Replace Message with the message and 1 with the issue id
created_comment, err := github_api.NewComment("Message", 1)

if err == nil {
    // created_comment.ID
}else{
    // err.Error()
}
```


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