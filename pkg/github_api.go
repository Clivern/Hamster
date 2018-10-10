package pkg

import (
    "github.com/clivern/hamster/internal/app/response"
    "github.com/clivern/hamster/internal/app/sender"
    "fmt"
    "net/http"
    "bytes"
    "io/ioutil"
    "errors"
    "encoding/json"
    "strings"
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

    if resp.StatusCode == 400 {
        return created_comment, errors.New(fmt.Sprintf("Oops: %s", string(body_byte)))
    }

    ok, err := created_comment.LoadFromJSON(body_byte)

    if ok && resp.StatusCode == 201 {
        return created_comment, nil
    }else{
        return created_comment, errors.New(fmt.Sprintf("Error: %s", string(body_byte)))
    }
}

/************************************/
/*********** LABEL API **************/
/************************************/

// Create a Label
func (e *GithubAPI) CreateLabel (name string, color string) (response.Label, error) {

    var created_label response.Label
    label := &sender.Label{Name:name, Color:color}

    json_body, err := label.ConvertToJSON()

    if err != nil{
        return created_label, err
    }

    client := &http.Client{}

    req, err := http.NewRequest(
        "POST",
        fmt.Sprintf("%s/repos/%s/%s/labels", GithubURL, e.Author, e.Repository),
        bytes.NewBufferString(json_body),
    )

    if err != nil{
        return created_label, err
    }

    req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

    resp, err := client.Do(req)

    if err != nil{
        return created_label, err
    }

    defer resp.Body.Close()

    body_byte, err := ioutil.ReadAll(resp.Body)

    if err != nil{
        return created_label, err
    }

    if resp.StatusCode == 400 {
        return created_label, errors.New(fmt.Sprintf("Oops: %s", string(body_byte)))
    }

    ok, err := created_label.LoadFromJSON(body_byte)

    if ok && resp.StatusCode == 201 {
        return created_label, nil
    }else{
        return created_label, errors.New(fmt.Sprintf("Error: %s", string(body_byte)))
    }
}

// Update a Label
func (e *GithubAPI) UpdateLabel (current_name string, name string, color string) (response.Label, error) {

    var updated_label response.Label
    label := &sender.Label{Name:name, Color:color}

    json_body, err := label.ConvertToJSON()

    if err != nil{
        return updated_label, err
    }

    client := &http.Client{}

    req, err := http.NewRequest(
        "PATCH",
        fmt.Sprintf("%s/repos/%s/%s/labels/%s", GithubURL, e.Author, e.Repository, current_name),
        bytes.NewBufferString(json_body),
    )

    if err != nil{
        return updated_label, err
    }

    req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

    resp, err := client.Do(req)

    if err != nil{
        return updated_label, err
    }

    defer resp.Body.Close()

    body_byte, err := ioutil.ReadAll(resp.Body)

    if err != nil{
        return updated_label, err
    }

    if resp.StatusCode == 400 {
        return updated_label, errors.New(fmt.Sprintf("Oops: %s", string(body_byte)))
    }

    ok, err := updated_label.LoadFromJSON(body_byte)

    if ok && resp.StatusCode == 200 {
        return updated_label, nil
    }else{
        return updated_label, errors.New(fmt.Sprintf("Error: %s", string(body_byte)))
    }
}

// Delete a Label
func (e *GithubAPI) DeleteLabel (name string) (bool, error) {

    client := &http.Client{}

    req, err := http.NewRequest(
        "DELETE",
        fmt.Sprintf("%s/repos/%s/%s/labels/%s", GithubURL, e.Author, e.Repository, name),
        nil,
    )

    if err != nil{
        return false, err
    }

    req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

    resp, err := client.Do(req)

    if err != nil{
        return false, err
    }

    defer resp.Body.Close()

    body_byte, err := ioutil.ReadAll(resp.Body)

    if err != nil{
        return false, err
    }

    if resp.StatusCode == 400 {
        return false, errors.New(fmt.Sprintf("Oops: %s", string(body_byte)))
    }

    if resp.StatusCode == 204 {
        return true, nil
    }else{
        return false, errors.New(fmt.Sprintf("Error: %s", string(body_byte)))
    }
}

// Get a List of Repository Labels
func (e *GithubAPI) GetRepositoryLabels () ([]response.Label, error) {

    var labels []response.Label

    client := &http.Client{}

    req, err := http.NewRequest(
        "GET",
        fmt.Sprintf("%s/repos/%s/%s/labels", GithubURL, e.Author, e.Repository),
        nil,
    )

    if err != nil{
        return labels, err
    }

    req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

    resp, err := client.Do(req)

    if err != nil{
        return labels, err
    }

    defer resp.Body.Close()

    body_byte, err := ioutil.ReadAll(resp.Body)

    if err != nil{
        return labels, err
    }

    if resp.StatusCode == 401 {
        return labels, errors.New(fmt.Sprintf("Oops: %s", string(body_byte)))
    }

    err = json.Unmarshal(body_byte, &labels)

    if err == nil && resp.StatusCode == 200 {
        return labels, nil
    }else{
        return labels, errors.New(fmt.Sprintf("Error: %s", string(body_byte)))
    }
}

// Get a List of labels on an issue
func (e *GithubAPI) GetRepositoryIssueLabels (issue_id int) ([]response.Label, error) {

    var labels []response.Label

    client := &http.Client{}

    req, err := http.NewRequest(
        "GET",
        fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, issue_id),
        nil,
    )

    if err != nil{
        return labels, err
    }

    req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

    resp, err := client.Do(req)

    if err != nil{
        return labels, err
    }

    defer resp.Body.Close()

    body_byte, err := ioutil.ReadAll(resp.Body)

    if err != nil{
        return labels, err
    }

    if resp.StatusCode == 401 {
        return labels, errors.New(fmt.Sprintf("Oops: %s", string(body_byte)))
    }

    err = json.Unmarshal(body_byte, &labels)

    if err == nil && resp.StatusCode == 200 {
        return labels, nil
    }else{
        return labels, errors.New(fmt.Sprintf("Error: %s", string(body_byte)))
    }
}

// Get a Label with name
func (e *GithubAPI) GetLabel (name string) (response.Label, error) {

    var label response.Label

    client := &http.Client{}

    req, err := http.NewRequest(
        "GET",
        fmt.Sprintf("%s/repos/%s/%s/labels/%s", GithubURL, e.Author, e.Repository, name),
        nil,
    )

    if err != nil{
        return label, err
    }

    req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

    resp, err := client.Do(req)

    if err != nil{
        return label, err
    }

    defer resp.Body.Close()

    body_byte, err := ioutil.ReadAll(resp.Body)

    if err != nil{
        return label, err
    }

    if resp.StatusCode == 401 {
        return label, errors.New(fmt.Sprintf("Oops: %s", string(body_byte)))
    }

    ok, err := label.LoadFromJSON(body_byte)

    if ok && resp.StatusCode == 200 {
        return label, nil
    }else{
        return label, errors.New(fmt.Sprintf("Error: %s", string(body_byte)))
    }
}

// Remove a label from an issue
func (e *GithubAPI) RemoveLabelFromIssue (issue_id int, label_name string) (bool, error) {
    client := &http.Client{}

    req, err := http.NewRequest(
        "DELETE",
        fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels/%s", GithubURL, e.Author, e.Repository, issue_id, label_name),
        nil,
    )

    if err != nil{
        return false, err
    }

    req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

    resp, err := client.Do(req)

    if err != nil{
        return false, err
    }

    defer resp.Body.Close()

    body_byte, err := ioutil.ReadAll(resp.Body)

    if err != nil{
        return false, err
    }

    if resp.StatusCode == 400 {
        return false, errors.New(fmt.Sprintf("Oops: %s", string(body_byte)))
    }

    if resp.StatusCode == 204 {
        return true, nil
    }else{
        return false, errors.New(fmt.Sprintf("Error: %s", string(body_byte)))
    }
}

// Remove all labels from an issue
func (e *GithubAPI) RemoveAllLabelForIssue (issue_id int) (bool, error) {
    client := &http.Client{}

    req, err := http.NewRequest(
        "DELETE",
        fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, issue_id),
        nil,
    )

    if err != nil{
        return false, err
    }

    req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

    resp, err := client.Do(req)

    if err != nil{
        return false, err
    }

    defer resp.Body.Close()

    body_byte, err := ioutil.ReadAll(resp.Body)

    if err != nil{
        return false, err
    }

    if resp.StatusCode == 400 {
        return false, errors.New(fmt.Sprintf("Oops: %s", string(body_byte)))
    }

    if resp.StatusCode == 204 {
        return true, nil
    }else{
        return false, errors.New(fmt.Sprintf("Error: %s", string(body_byte)))
    }
}

// Get labels for every issue in a milestone
func (e *GithubAPI) GetRepositoryMilestoneLabels (milestone_id int) ([]response.Label, error) {

    var labels []response.Label

    client := &http.Client{}

    req, err := http.NewRequest(
        "GET",
        fmt.Sprintf("%s/repos/%s/%s/milestones/%d/labels", GithubURL, e.Author, e.Repository, milestone_id),
        nil,
    )

    if err != nil{
        return labels, err
    }

    req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

    resp, err := client.Do(req)

    if err != nil{
        return labels, err
    }

    defer resp.Body.Close()

    body_byte, err := ioutil.ReadAll(resp.Body)

    if err != nil{
        return labels, err
    }

    if resp.StatusCode == 401 {
        return labels, errors.New(fmt.Sprintf("Oops: %s", string(body_byte)))
    }

    err = json.Unmarshal(body_byte, &labels)

    if err == nil && resp.StatusCode == 200 {
        return labels, nil
    }else{
        return labels, errors.New(fmt.Sprintf("Error: %s", string(body_byte)))
    }
}

// Add labels to an issue
func (e *GithubAPI) AddLabelsToIssue (issue_id int, labels []string) ([]response.Label, error) {

    var assigned_labels []response.Label

    client := &http.Client{}

    req, err := http.NewRequest(
        "POST",
        fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, issue_id),
        bytes.NewBufferString(fmt.Sprintf(`["%s"]`, strings.Join(labels,`","`))),
    )

    if err != nil{
        return assigned_labels, err
    }

    req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

    resp, err := client.Do(req)

    if err != nil{
        return assigned_labels, err
    }

    defer resp.Body.Close()

    body_byte, err := ioutil.ReadAll(resp.Body)

    if err != nil{
        return assigned_labels, err
    }

    if resp.StatusCode == 400 {
        return assigned_labels, errors.New(fmt.Sprintf("Oops: %s", string(body_byte)))
    }

    err = json.Unmarshal(body_byte, &assigned_labels)

    if err == nil && resp.StatusCode == 200 {
        return assigned_labels, nil
    }else{
        return assigned_labels, errors.New(fmt.Sprintf("Error: %s", string(body_byte)))
    }
}

// Replace all labels for an issue
func (e *GithubAPI) ReplaceAllLabelsForIssue (issue_id int, labels []string) ([]response.Label, error) {

    var assigned_labels []response.Label

    client := &http.Client{}

    req, err := http.NewRequest(
        "PUT",
        fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, issue_id),
        bytes.NewBufferString(fmt.Sprintf(`["%s"]`, strings.Join(labels,`","`))),
    )

    if err != nil{
        return assigned_labels, err
    }

    req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

    resp, err := client.Do(req)

    if err != nil{
        return assigned_labels, err
    }

    defer resp.Body.Close()

    body_byte, err := ioutil.ReadAll(resp.Body)

    if err != nil{
        return assigned_labels, err
    }

    if resp.StatusCode == 400 {
        return assigned_labels, errors.New(fmt.Sprintf("Oops: %s", string(body_byte)))
    }

    err = json.Unmarshal(body_byte, &assigned_labels)

    if err == nil && resp.StatusCode == 200 {
        return assigned_labels, nil
    }else{
        return assigned_labels, errors.New(fmt.Sprintf("Error: %s", string(body_byte)))
    }
}