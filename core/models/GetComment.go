// Code generated by go generate; DO NOT EDIT

package models

type GetComment []struct {
	LikeCount int `json:"like_count"`
	Decorators []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"decorators"`
	ActivityId string `json:"activity_id"`
	Index string `json:"index"`
	CreatedAt time.Time `json:"created_at"`
	ID string `json:"id"`
	Text string `json:"text"`
	Writer struct {
		Birthday string `json:"birthday"`
		BgImageUrl2 string `json:"bg_image_url2"`
		ProfileVideoUrlHq string `json:"profile_video_url_hq"`
		IsFavorite bool `json:"is_favorite"`
		BirthdayLeft int `json:"birthday_left"`
		Gender string `json:"gender"`
		BirthLeapType bool `json:"birth_leap_type"`
		BirthType string `json:"birth_type"`
		ProfileImageUrl2 string `json:"profile_image_url2"`
		Type string `json:"type"`
		FollowerCount int `json:"follower_count"`
		Relation struct {
			FeedBlocked bool `json:"feed_blocked"`
			Friend string `json:"friend"`
			Self bool `json:"self"`
			Follow string `json:"follow"`
			Favorite bool `json:"favorite"`
		} `json:"relation"`
		ProfileVideoUrlSquareSmall string `json:"profile_video_url_square_small"`
		IsCelebratable bool `json:"is_celebratable"`
		DefaultBgId int `json:"default_bg_id"`
		ActivityCount int `json:"activity_count"`
		IsValidUser bool `json:"is_valid_user"`
		ID string `json:"id"`
		Relationship string `json:"relationship"`
		Vip bool `json:"vip"`
		FriendCount int `json:"friend_count"`
		ProfileVideoUrlLq string `json:"profile_video_url_lq"`
		IsBirthday bool `json:"is_birthday"`
		BgImageUrl string `json:"bg_image_url"`
		ProfileVideoUrlSquare string `json:"profile_video_url_square"`
		AllowFollowing bool `json:"allow_following"`
		ProfileThumbnailUrl string `json:"profile_thumbnail_url"`
		StatusObjects []struct {
			ObjectType string `json:"object_type"`
			Message string `json:"message"`
		} `json:"status_objects"`
		ProfileVideoUrlSquareMicroSmall string `json:"profile_video_url_square_micro_small"`
		ProfileImageUrl string `json:"profile_image_url"`
		DisplayName string `json:"display_name"`
		Permalink string `json:"permalink"`
		IsDefaultProfileImage bool `json:"is_default_profile_image"`
		IsFeedBlocked bool `json:"is_feed_blocked"`
	} `json:"writer"`
	Liked bool `json:"liked"`
}