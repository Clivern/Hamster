// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/clivern/hamster/internal/app/response"
	"github.com/clivern/hamster/internal/app/sender"
	"io/ioutil"
	"net/http"
	"strings"
)

const GithubURL = "https://api.github.com"

type GithubAPI struct {
	Token      string `json:"token"`
	Author     string `json:"author"`
	Repository string `json:"repository"`
}

/************************************/
/*********** Comment API ************/
/************************************/

func (e *GithubAPI) NewComment(body string, issueId int) (response.CreatedComment, error) {

	var createdComment response.CreatedComment
	comment := &sender.Comment{Body: body}

	jsonBody, err := comment.ConvertToJSON()

	if err != nil {
		return createdComment, err
	}

	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/comments", GithubURL, e.Author, e.Repository, issueId),
		bytes.NewBufferString(jsonBody),
	)

	if err != nil {
		return createdComment, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return createdComment, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return createdComment, err
	}

	if resp.StatusCode == 400 {
		return createdComment, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	ok, err := createdComment.LoadFromJSON(bodyByte)

	if ok && resp.StatusCode == 201 {
		return createdComment, nil
	} else {
		return createdComment, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

/************************************/
/*********** LABEL API **************/
/************************************/

// Create a Label
func (e *GithubAPI) CreateLabel(name string, color string) (response.Label, error) {

	var createdLabel response.Label
	label := &sender.Label{Name: name, Color: color}

	jsonBody, err := label.ConvertToJSON()

	if err != nil {
		return createdLabel, err
	}

	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/repos/%s/%s/labels", GithubURL, e.Author, e.Repository),
		bytes.NewBufferString(jsonBody),
	)

	if err != nil {
		return createdLabel, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return createdLabel, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return createdLabel, err
	}

	if resp.StatusCode == 400 {
		return createdLabel, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	ok, err := createdLabel.LoadFromJSON(bodyByte)

	if ok && resp.StatusCode == 201 {
		return createdLabel, nil
	} else {
		return createdLabel, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

// Update a Label
func (e *GithubAPI) UpdateLabel(currentName string, name string, color string) (response.Label, error) {

	var updatedLabel response.Label
	label := &sender.Label{Name: name, Color: color}

	jsonBody, err := label.ConvertToJSON()

	if err != nil {
		return updatedLabel, err
	}

	client := &http.Client{}

	req, err := http.NewRequest(
		"PATCH",
		fmt.Sprintf("%s/repos/%s/%s/labels/%s", GithubURL, e.Author, e.Repository, currentName),
		bytes.NewBufferString(jsonBody),
	)

	if err != nil {
		return updatedLabel, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return updatedLabel, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return updatedLabel, err
	}

	if resp.StatusCode == 400 {
		return updatedLabel, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	ok, err := updatedLabel.LoadFromJSON(bodyByte)

	if ok && resp.StatusCode == 200 {
		return updatedLabel, nil
	} else {
		return updatedLabel, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

// Delete a Label
func (e *GithubAPI) DeleteLabel(name string) (bool, error) {

	client := &http.Client{}

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/repos/%s/%s/labels/%s", GithubURL, e.Author, e.Repository, name),
		nil,
	)

	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	if resp.StatusCode == 400 {
		return false, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else {
		return false, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

// Get a List of Repository Labels
func (e *GithubAPI) GetRepositoryLabels() ([]response.Label, error) {

	var labels []response.Label

	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/repos/%s/%s/labels", GithubURL, e.Author, e.Repository),
		nil,
	)

	if err != nil {
		return labels, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return labels, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return labels, err
	}

	if resp.StatusCode == 401 {
		return labels, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	err = json.Unmarshal(bodyByte, &labels)

	if err == nil && resp.StatusCode == 200 {
		return labels, nil
	} else {
		return labels, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

// Get a List of labels on an issue
func (e *GithubAPI) GetRepositoryIssueLabels(issueId int) ([]response.Label, error) {

	var labels []response.Label

	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, issueId),
		nil,
	)

	if err != nil {
		return labels, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return labels, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return labels, err
	}

	if resp.StatusCode == 401 {
		return labels, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	err = json.Unmarshal(bodyByte, &labels)

	if err == nil && resp.StatusCode == 200 {
		return labels, nil
	} else {
		return labels, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

// Get a Label with name
func (e *GithubAPI) GetLabel(name string) (response.Label, error) {

	var label response.Label

	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/repos/%s/%s/labels/%s", GithubURL, e.Author, e.Repository, name),
		nil,
	)

	if err != nil {
		return label, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return label, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return label, err
	}

	if resp.StatusCode == 401 {
		return label, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	ok, err := label.LoadFromJSON(bodyByte)

	if ok && resp.StatusCode == 200 {
		return label, nil
	} else {
		return label, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

// Remove a label from an issue
func (e *GithubAPI) RemoveLabelFromIssue(issueId int, labelName string) (bool, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels/%s", GithubURL, e.Author, e.Repository, issueId, labelName),
		nil,
	)

	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	if resp.StatusCode == 400 {
		return false, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else {
		return false, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

// Remove all labels from an issue
func (e *GithubAPI) RemoveAllLabelForIssue(issueId int) (bool, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, issueId),
		nil,
	)

	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	if resp.StatusCode == 400 {
		return false, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else {
		return false, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

// Get labels for every issue in a milestone
func (e *GithubAPI) GetRepositoryMilestoneLabels(milestoneId int) ([]response.Label, error) {

	var labels []response.Label

	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/repos/%s/%s/milestones/%d/labels", GithubURL, e.Author, e.Repository, milestoneId),
		nil,
	)

	if err != nil {
		return labels, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return labels, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return labels, err
	}

	if resp.StatusCode == 401 {
		return labels, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	err = json.Unmarshal(bodyByte, &labels)

	if err == nil && resp.StatusCode == 200 {
		return labels, nil
	} else {
		return labels, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

// Add labels to an issue
func (e *GithubAPI) AddLabelsToIssue(issueId int, labels []string) ([]response.Label, error) {

	var assignedLabels []response.Label

	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, issueId),
		bytes.NewBufferString(fmt.Sprintf(`["%s"]`, strings.Join(labels, `","`))),
	)

	if err != nil {
		return assignedLabels, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return assignedLabels, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return assignedLabels, err
	}

	if resp.StatusCode == 400 {
		return assignedLabels, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	err = json.Unmarshal(bodyByte, &assignedLabels)

	if err == nil && resp.StatusCode == 200 {
		return assignedLabels, nil
	} else {
		return assignedLabels, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

// Replace all labels for an issue
func (e *GithubAPI) ReplaceAllLabelsForIssue(issueId int, labels []string) ([]response.Label, error) {

	var assignedLabels []response.Label

	client := &http.Client{}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, issueId),
		bytes.NewBufferString(fmt.Sprintf(`["%s"]`, strings.Join(labels, `","`))),
	)

	if err != nil {
		return assignedLabels, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return assignedLabels, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return assignedLabels, err
	}

	if resp.StatusCode == 400 {
		return assignedLabels, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	err = json.Unmarshal(bodyByte, &assignedLabels)

	if err == nil && resp.StatusCode == 200 {
		return assignedLabels, nil
	} else {
		return assignedLabels, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

/************************************/
/************* PR API ***************/
/************************************/

// Get a List of labels on a PR
func (e *GithubAPI) GetRepositoryPRLabels(PRId int) ([]response.Label, error) {

	var labels []response.Label

	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, PRId),
		nil,
	)

	if err != nil {
		return labels, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return labels, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return labels, err
	}

	if resp.StatusCode == 401 {
		return labels, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	err = json.Unmarshal(bodyByte, &labels)

	if err == nil && resp.StatusCode == 200 {
		return labels, nil
	} else {
		return labels, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

// Remove a label from a PR
func (e *GithubAPI) RemoveLabelFromPR(PRId int, labelName string) (bool, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels/%s", GithubURL, e.Author, e.Repository, PRId, labelName),
		nil,
	)

	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	if resp.StatusCode == 400 {
		return false, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else {
		return false, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

// Remove all labels from PR
func (e *GithubAPI) RemoveAllLabelForPR(PRId int) (bool, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, PRId),
		nil,
	)

	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	if resp.StatusCode == 400 {
		return false, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else {
		return false, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

// Add labels to PR
func (e *GithubAPI) AddLabelsToPR(PRId int, labels []string) ([]response.Label, error) {

	var assignedLabels []response.Label

	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, PRId),
		bytes.NewBufferString(fmt.Sprintf(`["%s"]`, strings.Join(labels, `","`))),
	)

	if err != nil {
		return assignedLabels, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return assignedLabels, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return assignedLabels, err
	}

	if resp.StatusCode == 400 {
		return assignedLabels, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	err = json.Unmarshal(bodyByte, &assignedLabels)

	if err == nil && resp.StatusCode == 200 {
		return assignedLabels, nil
	} else {
		return assignedLabels, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}

// Replace all labels for PR
func (e *GithubAPI) ReplaceAllLabelsForPR(PRId int, labels []string) ([]response.Label, error) {

	var assignedLabels []response.Label

	client := &http.Client{}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, PRId),
		bytes.NewBufferString(fmt.Sprintf(`["%s"]`, strings.Join(labels, `","`))),
	)

	if err != nil {
		return assignedLabels, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))

	resp, err := client.Do(req)

	if err != nil {
		return assignedLabels, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return assignedLabels, err
	}

	if resp.StatusCode == 400 {
		return assignedLabels, errors.New(fmt.Sprintf("Oops: %s", string(bodyByte)))
	}

	err = json.Unmarshal(bodyByte, &assignedLabels)

	if err == nil && resp.StatusCode == 200 {
		return assignedLabels, nil
	} else {
		return assignedLabels, errors.New(fmt.Sprintf("Error: %s", string(bodyByte)))
	}
}
