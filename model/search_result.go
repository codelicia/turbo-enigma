package model

type SearchResult struct {
	Ok       bool   `json:"ok"`
	Query    string `json:"query"`
	Messages struct {
		Total      int `json:"total"`
		Pagination struct {
			TotalCount int `json:"total_count"`
			Page       int `json:"page"`
			PerPage    int `json:"per_page"`
			PageCount  int `json:"page_count"`
			First      int `json:"first"`
			Last       int `json:"last"`
		} `json:"pagination"`
		Paging struct {
			Count int `json:"count"`
			Total int `json:"total"`
			Page  int `json:"page"`
			Pages int `json:"pages"`
		} `json:"paging"`
		Matches []struct {
			Iid     string  `json:"iid"`
			Team    string  `json:"team"`
			Score   float64 `json:"score"`
			Channel struct {
				ID                 string        `json:"id"`
				IsChannel          bool          `json:"is_channel"`
				IsGroup            bool          `json:"is_group"`
				IsIm               bool          `json:"is_im"`
				Name               string        `json:"name"`
				IsShared           bool          `json:"is_shared"`
				IsOrgShared        bool          `json:"is_org_shared"`
				IsExtShared        bool          `json:"is_ext_shared"`
				IsPrivate          bool          `json:"is_private"`
				IsMpim             bool          `json:"is_mpim"`
				PendingShared      []interface{} `json:"pending_shared"`
				IsPendingExtShared bool          `json:"is_pending_ext_shared"`
			} `json:"channel"`
			Type        string      `json:"type"`
			User        interface{} `json:"user"`
			Username    string      `json:"username"`
			Ts          string      `json:"ts"`
			Text        string      `json:"text"`
			Permalink   string      `json:"permalink"`
			NoReactions bool        `json:"no_reactions,omitempty"`
			Attachments []struct {
				FromURL       string   `json:"from_url"`
				Fallback      string   `json:"fallback"`
				Ts            string   `json:"ts"`
				MsgSubtype    string   `json:"msg_subtype"`
				AuthorSubname string   `json:"author_subname"`
				ChannelID     string   `json:"channel_id"`
				ChannelName   string   `json:"channel_name"`
				IsMsgUnfurl   bool     `json:"is_msg_unfurl"`
				Text          string   `json:"text"`
				AuthorIcon    string   `json:"author_icon"`
				AuthorLink    string   `json:"author_link"`
				MrkdwnIn      []string `json:"mrkdwn_in"`
				ID            int      `json:"id"`
				OriginalURL   string   `json:"original_url"`
				Footer        string   `json:"footer"`
			} `json:"attachments,omitempty"`
			Blocks []struct {
				Type     string `json:"type"`
				BlockID  string `json:"block_id"`
				Elements []struct {
					Type     string `json:"type"`
					Elements []struct {
						Type string `json:"type"`
						URL  string `json:"url"`
					} `json:"elements"`
				} `json:"elements"`
			} `json:"blocks,omitempty"`
		} `json:"matches"`
	} `json:"messages"`
}
