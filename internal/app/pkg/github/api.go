// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/clivern/hamster/internal/app/response"
	"github.com/clivern/hamster/internal/app/sender"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// GithubURL api url
const GithubURL = "https://api.github.com"

// API is a representation of a github api
type API struct {
	Token      string `json:"token"`
	Author     string `json:"author"`
	Repository string `json:"repository"`
}

/************************************/
/*********** Comment API ************/
/************************************/

// NewComment creates a new issue comment
func (e *API) NewComment(body string, issueID int) (response.CreatedComment, error) {

	var createdComment response.CreatedComment
	comment := &sender.Comment{Body: body}

	jsonBody, err := comment.ConvertToJSON()

	if err != nil {
		return createdComment, err
	}

	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/comments", GithubURL, e.Author, e.Repository, issueID),
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
		return createdComment, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	ok, _ := createdComment.LoadFromJSON(bodyByte)

	if ok && resp.StatusCode == 201 {
		return createdComment, nil
	}
	return createdComment, fmt.Errorf("Error: %s", string(bodyByte))
}

/************************************/
/*********** LABEL API **************/
/************************************/

// CreateLabel creates a label
func (e *API) CreateLabel(name string, color string) (response.Label, error) {

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
		return createdLabel, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	ok, _ := createdLabel.LoadFromJSON(bodyByte)

	if ok && resp.StatusCode == 201 {
		return createdLabel, nil
	}
	return createdLabel, fmt.Errorf("Error: %s", string(bodyByte))
}

// UpdateLabel updates a label
func (e *API) UpdateLabel(currentName string, name string, color string) (response.Label, error) {

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
		return updatedLabel, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	ok, _ := updatedLabel.LoadFromJSON(bodyByte)

	if ok && resp.StatusCode == 200 {
		return updatedLabel, nil
	}

	return updatedLabel, fmt.Errorf("Error: %s", string(bodyByte))
}

// DeleteLabel deletes a label
func (e *API) DeleteLabel(name string) (bool, error) {

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
		return false, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	if resp.StatusCode == 204 {
		return true, nil
	}
	return false, fmt.Errorf("Error: %s", string(bodyByte))
}

// GetRepositoryLabels lists a repository labels
func (e *API) GetRepositoryLabels() ([]response.Label, error) {

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
		return labels, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	err = json.Unmarshal(bodyByte, &labels)

	if err == nil && resp.StatusCode == 200 {
		return labels, nil
	}
	return labels, fmt.Errorf("Error: %s", string(bodyByte))
}

// GetRepositoryIssueLabels lists a repository issue labels
func (e *API) GetRepositoryIssueLabels(issueID int) ([]response.Label, error) {

	var labels []response.Label

	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, issueID),
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
		return labels, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	err = json.Unmarshal(bodyByte, &labels)

	if err == nil && resp.StatusCode == 200 {
		return labels, nil
	}
	return labels, fmt.Errorf("Error: %s", string(bodyByte))
}

// GetLabel returns a label data
func (e *API) GetLabel(name string) (response.Label, error) {

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
		return label, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	ok, _ := label.LoadFromJSON(bodyByte)

	if ok && resp.StatusCode == 200 {
		return label, nil
	}
	return label, fmt.Errorf("Error: %s", string(bodyByte))
}

// RemoveLabelFromIssue removes a label from issue
func (e *API) RemoveLabelFromIssue(issueID int, labelName string) (bool, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels/%s", GithubURL, e.Author, e.Repository, issueID, labelName),
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
		return false, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	if resp.StatusCode == 204 {
		return true, nil
	}
	return false, fmt.Errorf("Error: %s", string(bodyByte))
}

// RemoveAllLabelForIssue Removes all labels from an issue
func (e *API) RemoveAllLabelForIssue(issueID int) (bool, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, issueID),
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
		return false, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	if resp.StatusCode == 204 {
		return true, nil
	}
	return false, fmt.Errorf("Error: %s", string(bodyByte))
}

// GetRepositoryMilestoneLabels Gets labels for every issue in a milestone
func (e *API) GetRepositoryMilestoneLabels(milestoneID int) ([]response.Label, error) {

	var labels []response.Label

	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/repos/%s/%s/milestones/%d/labels", GithubURL, e.Author, e.Repository, milestoneID),
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
		return labels, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	err = json.Unmarshal(bodyByte, &labels)

	if err == nil && resp.StatusCode == 200 {
		return labels, nil
	}
	return labels, fmt.Errorf("Error: %s", string(bodyByte))
}

// AddLabelsToIssue Adds labels to an issue
func (e *API) AddLabelsToIssue(issueID int, labels []string) ([]response.Label, error) {

	var assignedLabels []response.Label

	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, issueID),
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
		return assignedLabels, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	err = json.Unmarshal(bodyByte, &assignedLabels)

	if err == nil && resp.StatusCode == 200 {
		return assignedLabels, nil
	}
	return assignedLabels, fmt.Errorf("Error: %s", string(bodyByte))
}

// ReplaceAllLabelsForIssue Replaces all labels for an issue
func (e *API) ReplaceAllLabelsForIssue(issueID int, labels []string) ([]response.Label, error) {

	var assignedLabels []response.Label

	client := &http.Client{}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels", GithubURL, e.Author, e.Repository, issueID),
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
		return assignedLabels, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	err = json.Unmarshal(bodyByte, &assignedLabels)

	if err == nil && resp.StatusCode == 200 {
		return assignedLabels, nil
	}
	return assignedLabels, fmt.Errorf("Error: %s", string(bodyByte))
}

/************************************/
/************* PR API ***************/
/************************************/

// GetRepositoryPRLabels Gets a List of labels on a PR
func (e *API) GetRepositoryPRLabels(PRId int) ([]response.Label, error) {

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
		return labels, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	err = json.Unmarshal(bodyByte, &labels)

	if err == nil && resp.StatusCode == 200 {
		return labels, nil
	}
	return labels, fmt.Errorf("Error: %s", string(bodyByte))
}

// RemoveLabelFromPR Removes a label from a PR
func (e *API) RemoveLabelFromPR(PRId int, labelName string) (bool, error) {
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
		return false, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	if resp.StatusCode == 204 {
		return true, nil
	}
	return false, fmt.Errorf("Error: %s", string(bodyByte))
}

// RemoveAllLabelForPR Removes all labels from PR
func (e *API) RemoveAllLabelForPR(PRId int) (bool, error) {
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
		return false, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	if resp.StatusCode == 204 {
		return true, nil
	}
	return false, fmt.Errorf("Error: %s", string(bodyByte))
}

// AddLabelsToPR Adds labels to PR
func (e *API) AddLabelsToPR(PRId int, labels []string) ([]response.Label, error) {

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
		return assignedLabels, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	err = json.Unmarshal(bodyByte, &assignedLabels)

	if err == nil && resp.StatusCode == 200 {
		return assignedLabels, nil
	}
	return assignedLabels, fmt.Errorf("Error: %s", string(bodyByte))
}

// ReplaceAllLabelsForPR Replaces all labels for PR
func (e *API) ReplaceAllLabelsForPR(PRId int, labels []string) ([]response.Label, error) {

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
		return assignedLabels, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	err = json.Unmarshal(bodyByte, &assignedLabels)

	if err == nil && resp.StatusCode == 200 {
		return assignedLabels, nil
	}
	return assignedLabels, fmt.Errorf("Error: %s", string(bodyByte))
}

/************************************/
/********* Check Runs API ***********/
/************************************/

// CreateCheckRun Creates a check run (https://developer.github.com/v3/checks/runs/#create-a-check-run)
func (e *API) CreateCheckRun(CheckRun sender.CheckRun) (response.CheckRun, error) {

	var checkRun response.CheckRun

	jsonBody, err := CheckRun.ConvertToJSON()
	fmt.Println(jsonBody)
	if err != nil {
		return checkRun, err
	}

	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/repos/%s/%s/check-runs", GithubURL, e.Author, e.Repository),
		bytes.NewBufferString(jsonBody),
	)

	if err != nil {
		return checkRun, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/vnd.github.antiope-preview+json")

	resp, err := client.Do(req)

	if err != nil {
		return checkRun, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return checkRun, err
	}

	if resp.StatusCode == 400 {
		return checkRun, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	ok, _ := checkRun.LoadFromJSON(bodyByte)

	if ok && resp.StatusCode == 201 {
		return checkRun, nil
	}

	return checkRun, fmt.Errorf("Error: %s", string(bodyByte))
}

// UpdateCheckRun Updates a check run (https://developer.github.com/v3/checks/runs/#update-a-check-run)
func (e *API) UpdateCheckRun(ID int, CheckRun sender.CheckRun) (response.CheckRun, error) {

	var checkRun response.CheckRun

	jsonBody, err := CheckRun.ConvertToJSON()

	if err != nil {
		return checkRun, err
	}

	client := &http.Client{}

	req, err := http.NewRequest(
		"PATCH",
		fmt.Sprintf("%s/repos/%s/%s/check-runs/%d", GithubURL, e.Author, e.Repository, ID),
		bytes.NewBufferString(jsonBody),
	)

	if err != nil {
		return checkRun, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/vnd.github.antiope-preview+json")

	resp, err := client.Do(req)

	if err != nil {
		return checkRun, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return checkRun, err
	}

	if resp.StatusCode == 400 {
		return checkRun, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	ok, _ := checkRun.LoadFromJSON(bodyByte)

	if ok && resp.StatusCode == 200 {
		return checkRun, nil
	}

	return checkRun, fmt.Errorf("Error: %s", string(bodyByte))
}

// ListRefCheckRuns lists check runs for a specific ref (https://developer.github.com/v3/checks/runs/#list-check-runs-for-a-specific-ref)
func (e *API) ListRefCheckRuns(Ref string, CheckName string, Status string, Filter string) (response.CheckRuns, error) {

	var checkRuns response.CheckRuns

	client := &http.Client{}

	u, err := url.Parse(fmt.Sprintf("%s/repos/%s/%s/commits/%s/check-runs", GithubURL, e.Author, e.Repository, Ref))

	if err != nil {
		return checkRuns, err
	}

	q := u.Query()
	if CheckName != "" {
		q.Set("check_name", CheckName)
	}
	if Status != "" {
		q.Set("status", Status)
	}
	if Filter != "" {
		q.Set("filter", Filter)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(
		"GET",
		u.String(),
		nil,
	)

	if err != nil {
		return checkRuns, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/vnd.github.antiope-preview+json")

	resp, err := client.Do(req)

	if err != nil {
		return checkRuns, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return checkRuns, err
	}

	if resp.StatusCode == 401 {
		return checkRuns, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	ok, _ := checkRuns.LoadFromJSON(bodyByte)

	if ok && resp.StatusCode == 200 {
		return checkRuns, nil
	}
	return checkRuns, fmt.Errorf("Error: %s", string(bodyByte))
}

// ListSuiteCheckRuns lists check runs in a check suite (https://developer.github.com/v3/checks/runs/#list-check-runs-in-a-check-suite)
func (e *API) ListSuiteCheckRuns(CheckSuiteID int, CheckName string, Status string, Filter string) (response.CheckRuns, error) {

	var checkRuns response.CheckRuns

	client := &http.Client{}

	u, err := url.Parse(fmt.Sprintf("%s/repos/%s/%s/check-suites/%d/check-runs", GithubURL, e.Author, e.Repository, CheckSuiteID))

	if err != nil {
		return checkRuns, err
	}

	q := u.Query()
	if CheckName != "" {
		q.Set("check_name", CheckName)
	}
	if Status != "" {
		q.Set("status", Status)
	}
	if Filter != "" {
		q.Set("filter", Filter)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(
		"GET",
		u.String(),
		nil,
	)

	if err != nil {
		return checkRuns, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/vnd.github.antiope-preview+json")

	resp, err := client.Do(req)

	if err != nil {
		return checkRuns, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return checkRuns, err
	}

	if resp.StatusCode == 401 {
		return checkRuns, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	ok, _ := checkRuns.LoadFromJSON(bodyByte)

	if ok && resp.StatusCode == 200 {
		return checkRuns, nil
	}
	return checkRuns, fmt.Errorf("Error: %s", string(bodyByte))
}

// GetCheckRun Gets a single check run (https://developer.github.com/v3/checks/runs/#get-a-single-check-run)
func (e *API) GetCheckRun(ID int) (response.CheckRun, error) {

	var checkRun response.CheckRun

	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/repos/%s/%s/check-runs/%d", GithubURL, e.Author, e.Repository, ID),
		nil,
	)

	if err != nil {
		return checkRun, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/vnd.github.antiope-preview+json")

	resp, err := client.Do(req)

	if err != nil {
		return checkRun, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return checkRun, err
	}

	if resp.StatusCode == 401 {
		return checkRun, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	ok, _ := checkRun.LoadFromJSON(bodyByte)

	if ok && resp.StatusCode == 200 {
		return checkRun, nil
	}
	return checkRun, fmt.Errorf("Error: %s", string(bodyByte))
}

// ListCheckRunAnnotations Lists annotations for a check run (https://developer.github.com/v3/checks/runs/#list-annotations-for-a-check-run)
func (e *API) ListCheckRunAnnotations(ID int) (response.Annotations, error) {

	var annotations response.Annotations

	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/repos/%s/%s/check-runs/%d/annotations", GithubURL, e.Author, e.Repository, ID),
		nil,
	)

	if err != nil {
		return annotations, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", e.Token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/vnd.github.antiope-preview+json")

	resp, err := client.Do(req)

	if err != nil {
		return annotations, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return annotations, err
	}

	if resp.StatusCode == 401 {
		return annotations, fmt.Errorf("Oops: %s", string(bodyByte))
	}

	ok, _ := annotations.LoadFromJSON(bodyByte)

	if ok && resp.StatusCode == 200 {
		return annotations, nil
	}
	return annotations, fmt.Errorf("Error: %s", string(bodyByte))
}
