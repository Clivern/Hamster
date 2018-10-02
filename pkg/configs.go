package pkg


import (
    "io/ioutil"
    "os"
    "encoding/json"
)

type Config struct {
    GithubToken string `json:"github_token"`
    GithubWebhookSecret string `json:"github_webhook_secret"`
    RepositoryAuthor string `json:"repository_author"`
    RepositoryName string `json:"repository_name"`
}


func (e *Config) Load(file string) (bool, error) {

    data, err := ioutil.ReadFile(file)

    if err != nil {
        return false, err
    }

    err = json.Unmarshal(data, &e)

    if err != nil {
        return false, err
    }

    return true, nil
}

func (e *Config) Cache () {
    os.Setenv("GithubToken", e.GithubToken)
    os.Setenv("GithubWebhookSecret", e.GithubWebhookSecret)
    os.Setenv("RepositoryAuthor", e.RepositoryAuthor)
    os.Setenv("RepositoryName", e.RepositoryName)
}