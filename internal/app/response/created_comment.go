package response

import (
    "encoding/json"
)

type CreatedComment struct {
    URL      string `json:"url"`
    HTMLURL  string `json:"html_url"`
    IssueURL string `json:"issue_url"`
    ID       int    `json:"id"`
    NodeID   string `json:"node_id"`
    User     struct {
        Login             string `json:"login"`
        ID                int    `json:"id"`
        NodeID            string `json:"node_id"`
        AvatarURL         string `json:"avatar_url"`
        GravatarID        string `json:"gravatar_id"`
        URL               string `json:"url"`
        HTMLURL           string `json:"html_url"`
        FollowersURL      string `json:"followers_url"`
        FollowingURL      string `json:"following_url"`
        GistsURL          string `json:"gists_url"`
        StarredURL        string `json:"starred_url"`
        SubscriptionsURL  string `json:"subscriptions_url"`
        OrganizationsURL  string `json:"organizations_url"`
        ReposURL          string `json:"repos_url"`
        EventsURL         string `json:"events_url"`
        ReceivedEventsURL string `json:"received_events_url"`
        Type              string `json:"type"`
        SiteAdmin         bool   `json:"site_admin"`
    } `json:"user"`
    CreatedAt         time.Time `json:"created_at"`
    UpdatedAt         time.Time `json:"updated_at"`
    AuthorAssociation string    `json:"author_association"`
    Body              string    `json:"body"`
}


func (e *CreatedComment) LoadFromJSON (data []byte) bool {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false
    }
    return true
}