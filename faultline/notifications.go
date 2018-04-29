package faultline

type Slack struct {
	Type           string `json:"type"`
	NotifyInterval int    `json:"notifyInterval"`
	Threshold      int    `json:"threshold"`
	Endpoint       string `json:"endpoint"`
	Channel        string `json:"channel"`
	Username       string `json:"username"`
	IconEmoji      string `json:"iconEmoji,omitempty"`
	IconURL        string `json:"iconUrl,omitempty"`
	Timezone       int    `json:"timezone,omitempty"`
	LinkTemplate   string `json:"linkTemplate,omitempty"`
}

type GitHub struct {
	Type           string   `json:"type"`
	NotifyInterval int      `json:"notifyInterval"`
	Threshold      int      `json:"threshold"`
	Endpoint       string   `json:"endpoint,omitempty"`
	UserToken      string   `json:"userToken"`
	Owner          string   `json:"owner"`
	Repo           int      `json:"repo"`
	Labels         []string `json:"labels,omitempty"`
	IfExist        string   `json:"if_exist,omitempty"`
	Timezone       int      `json:"timezone,omitempty"`
	LinkTemplate   string   `json:"linkTemplate,omitempty"`
}

type GitLab struct {
	Type                string   `json:"type"`
	NotifyInterval      int      `json:"notifyInterval"`
	Threshold           int      `json:"threshold"`
	Endpoint            string   `json:"endpoint,omitempty"`
	PersonalAccessToken string   `json:"personalAccessToken"`
	Owner               string   `json:"owner"`
	Repo                int      `json:"repo"`
	Labels              []string `json:"labels,omitempty"`
	IfExist             string   `json:"if_exist,omitempty"`
	Timezone            int      `json:"timezone,omitempty"`
	LinkTemplate        string   `json:"linkTemplate,omitempty"`
}
