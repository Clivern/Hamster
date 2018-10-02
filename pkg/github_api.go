package pkg

import (
    "github.com/clivern/hamster/internal/app/response"
    "github.com/clivern/hamster/internal/app/sender"
    "fmt"
    "net/http"
    "bytes"
    "io/ioutil"
    "errors"
)

const GithubURL = "https://api.github.com"

type GithubAPI struct {
    Token       string `json:"token"`
    Author      string `json:"author"`
    Repository  string `json:"repository"`
}

func (e *GithubAPI) NewComment (body string, issueId int) (response.CreatedComment, error) {

    var created_comment response.CreatedComment
    comment := &sender.Comment{Body:body}

    json_body, err := comment.ConvertToJSON()

    if err != nil{
        return created_comment, err
    }

    client := &http.Client{}

    req, err := http.NewRequest(
        "POST",
        fmt.Sprintf("%s/repos/%s/%s/issues/%d/comments", GithubURL, e.Author, e.Repository, issueId),
        bytes.NewBufferString(json_body),
    )

    if err != nil{
        return created_comment, err
    }

    req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

    resp, err := client.Do(req)

    if err != nil{
        return created_comment, err
    }

    defer resp.Body.Close()

    body_byte, err := ioutil.ReadAll(resp.Body)

    if err != nil{
        return created_comment, err
    }

    if resp.StatusCode == 401 {
        return created_comment, errors.New(fmt.Sprintf("Unauthorized Access: %s", string(body_byte)))
    }

    ok, err := created_comment.LoadFromJSON(body_byte)

    if ok && resp.StatusCode == 201 {
        return created_comment, nil
    }else{
        return created_comment, errors.New(fmt.Sprintf("Error: %s", string(body_byte)))
    }
}