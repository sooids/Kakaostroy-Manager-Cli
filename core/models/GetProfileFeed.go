// Code generated by go generate; DO NOT EDIT

package models

type GetProfileFeed struct {
	MutualFriend interface{} `json:"mutual_friend"`
	Profile      struct {
		BgImageUrl2          string `json:"bg_image_url2"`
		WaitingQuestionCount int    `json:"waiting_question_count"`
		IsFavorite           bool   `json:"is_favorite"`
		Gender               string `json:"gender"`
		BirthLeapType        bool   `json:"birth_leap_type"`
		BirthType            string `json:"birth_type"`
		ProfileImageUrl2     string `json:"profile_image_url2"`
		Type                 string `json:"type"`
		FollowerCount        int    `json:"follower_count"`
		Relation             struct {
			FeedBlocked bool   `json:"feed_blocked"`
			Friend      string `json:"friend"`
			Self        bool   `json:"self"`
			Follow      string `json:"follow"`
			Favorite    bool   `json:"favorite"`
			Ban         string `json:"ban"`
			Foaf        bool   `json:"foaf"`
		} `json:"relation"`
		IsCelebratable      bool   `json:"is_celebratable"`
		DefaultBgId         int    `json:"default_bg_id"`
		ActivityCount       int    `json:"activity_count"`
		IsValidUser         bool   `json:"is_valid_user"`
		FriendAcceptLevel   string `json:"friend_accept_level"`
		ID                  string `json:"id"`
		Relationship        string `json:"relationship"`
		Vip                 bool   `json:"vip"`
		FriendCount         int    `json:"friend_count"`
		IsBirthday          bool   `json:"is_birthday"`
		MessageRejectee     bool   `json:"message_rejectee"`
		BgImageUrl          string `json:"bg_image_url"`
		AllowFollowing      bool   `json:"allow_following"`
		ProfileThumbnailUrl string `json:"profile_thumbnail_url"`
		StatusObjects       []struct {
			ObjectType string `json:"object_type"`
			Message    string `json:"message"`
		} `json:"status_objects"`
		MessageSendable       bool     `json:"message_sendable"`
		ProfileImageUrl       string   `json:"profile_image_url"`
		DisplayName           string   `json:"display_name"`
		GenderPermission      string   `json:"gender_permission"`
		BiographySummary      []string `json:"biography_summary"`
		Permalink             string   `json:"permalink"`
		IsDefaultProfileImage bool     `json:"is_default_profile_image"`
		IsFeedBlocked         bool     `json:"is_feed_blocked"`
	} `json:"profile"`
}
