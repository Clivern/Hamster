package pkg


import (
    "io/ioutil"
    "os"
    "encoding/json"
    "errors"
    "fmt"
)

type Config struct {
    GithubToken string `json:"github_token"`
    GithubWebhookSecret string `json:"github_webhook_secret"`
    RepositoryAuthor string `json:"repository_author"`
    RepositoryName string `json:"repository_name"`
    AppMode string `json:"app_mode"`
    AppPort string `json:"app_port"`
}


func (e *Config) Load(file string) (bool, error) {

    _, err := os.Stat(file)

    if err != nil{
        return false, errors.New(fmt.Sprintf("Config file %s not found, Hamster will read Env variables!", file))
    }

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
    if os.Getenv("AppMode") == "" {
        os.Setenv("GithubToken", e.GithubToken)
        os.Setenv("GithubWebhookSecret", e.GithubWebhookSecret)
        os.Setenv("RepositoryAuthor", e.RepositoryAuthor)
        os.Setenv("RepositoryName", e.RepositoryName)
        os.Setenv("AppMode", e.AppMode)
        os.Setenv("AppPort", e.AppPort)
    }
}

func (e *Config) GinEnv () {
    // Used by gin framework
    // https://github.com/gin-gonic/gin/blob/d510595aa58c2417373d89a8d8ffa21cf58673cb/utils.go#L140
    os.Setenv("PORT", os.Getenv("AppPort"))
}
