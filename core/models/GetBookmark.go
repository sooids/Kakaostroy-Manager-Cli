// Code generated by go generate; DO NOT EDIT

package models

import "time"

type GetBookmark struct {
	Bookmarks []struct {
		CreatedAt time.Time `json:"created_at"`
		ID        string    `json:"id"`
		Message   string    `json:"message,omitempty"`
		Status    int       `json:"status"`
		Activity  struct {
			CommentCount int       `json:"comment_count"`
			Downloadable bool      `json:"downloadable"`
			Pinned       bool      `json:"pinned"`
			Deletable    bool      `json:"deletable"`
			CreatedAt    time.Time `json:"created_at"`
			WithTagCount int       `json:"with_tag_count"`
			Sharable     bool      `json:"sharable"`
			Media        []struct {
				URL           string `json:"url"`
				ThumbnailUrl2 string `json:"thumbnail_url2"`
				ThumbnailUrl3 string `json:"thumbnail_url3"`
				ThumbnailUrl  string `json:"thumbnail_url"`
				OriginUrl     string `json:"origin_url"`
				URL2          string `json:"url2"`
				JpgUrl        string `json:"jpg_url"`
				Key           string `json:"key"`
				ContentType   string `json:"content_type"`
				Width         int    `json:"width"`
				Height        int    `json:"height"`
				MediaPath     string `json:"media_path"`
				SqUrl         string `json:"sq_url"`
			} `json:"media"`
			Content           string    `json:"content"`
			Required          bool      `json:"required"`
			Liked             bool      `json:"liked"`
			FeedId            string    `json:"feed_id"`
			Sid               string    `json:"sid"`
			Sympathized       bool      `json:"sympathized"`
			UpdatedAt         time.Time `json:"updated_at"`
			MediaType         string    `json:"media_type"`
			ActivityType      string    `json:"activity_type"`
			ID                string    `json:"id"`
			Modifiable        bool      `json:"modifiable"`
			ContentDecorators []struct {
				Text string `json:"text"`
				Type string `json:"type"`
			} `json:"content_decorators"`
			Blinded bool `json:"blinded"`
			Likes   []struct {
				Actor struct {
					BgImageUrl2      string `json:"bg_image_url2"`
					IsFavorite       bool   `json:"is_favorite"`
					Gender           string `json:"gender"`
					BirthLeapType    bool   `json:"birth_leap_type"`
					BirthType        string `json:"birth_type"`
					ProfileImageUrl2 string `json:"profile_image_url2"`
					Type             string `json:"type"`
					FollowerCount    int    `json:"follower_count"`
					Relation         struct {
						FeedBlocked bool   `json:"feed_blocked"`
						Friend      string `json:"friend"`
						Self        bool   `json:"self"`
						Follow      string `json:"follow"`
						Favorite    bool   `json:"favorite"`
					} `json:"relation"`
					IsCelebratable        bool          `json:"is_celebratable"`
					DefaultBgId           int           `json:"default_bg_id"`
					ActivityCount         int           `json:"activity_count"`
					IsValidUser           bool          `json:"is_valid_user"`
					ID                    string        `json:"id"`
					Relationship          string        `json:"relationship"`
					Vip                   bool          `json:"vip"`
					FriendCount           int           `json:"friend_count"`
					IsBirthday            bool          `json:"is_birthday"`
					BgImageUrl            string        `json:"bg_image_url"`
					AllowFollowing        bool          `json:"allow_following"`
					ProfileThumbnailUrl   string        `json:"profile_thumbnail_url"`
					StatusObjects         []interface{} `json:"status_objects"`
					ProfileImageUrl       string        `json:"profile_image_url"`
					DisplayName           string        `json:"display_name"`
					Permalink             string        `json:"permalink"`
					IsDefaultProfileImage bool          `json:"is_default_profile_image"`
					IsFeedBlocked         bool          `json:"is_feed_blocked"`
				} `json:"actor"`
				Emotion   string    `json:"emotion"`
				CreatedAt time.Time `json:"created_at"`
				ID        string    `json:"id"`
			} `json:"likes"`
			Summary   string `json:"summary"`
			LikeCount int    `json:"like_count"`
			Comments  []struct {
				LikeCount  int `json:"like_count"`
				Decorators []struct {
					Text string `json:"text"`
					Type string `json:"type"`
				} `json:"decorators"`
				ActivityId string    `json:"activity_id"`
				Index      string    `json:"index"`
				CreatedAt  time.Time `json:"created_at"`
				ID         string    `json:"id"`
				Text       string    `json:"text"`
				Writer     struct {
					BgImageUrl2      string `json:"bg_image_url2"`
					IsFavorite       bool   `json:"is_favorite"`
					Gender           string `json:"gender"`
					BirthLeapType    bool   `json:"birth_leap_type"`
					BirthType        string `json:"birth_type"`
					ProfileImageUrl2 string `json:"profile_image_url2"`
					Type             string `json:"type"`
					FollowerCount    int    `json:"follower_count"`
					Relation         struct {
						FeedBlocked bool   `json:"feed_blocked"`
						Friend      string `json:"friend"`
						Self        bool   `json:"self"`
						Follow      string `json:"follow"`
						Favorite    bool   `json:"favorite"`
					} `json:"relation"`
					IsCelebratable      bool   `json:"is_celebratable"`
					DefaultBgId         int    `json:"default_bg_id"`
					ActivityCount       int    `json:"activity_count"`
					IsValidUser         bool   `json:"is_valid_user"`
					ID                  string `json:"id"`
					Relationship        string `json:"relationship"`
					Vip                 bool   `json:"vip"`
					FriendCount         int    `json:"friend_count"`
					IsBirthday          bool   `json:"is_birthday"`
					BgImageUrl          string `json:"bg_image_url"`
					AllowFollowing      bool   `json:"allow_following"`
					ProfileThumbnailUrl string `json:"profile_thumbnail_url"`
					StatusObjects       []struct {
						ObjectType string `json:"object_type"`
						Message    string `json:"message"`
					} `json:"status_objects"`
					ProfileImageUrl       string `json:"profile_image_url"`
					DisplayName           string `json:"display_name"`
					Permalink             string `json:"permalink"`
					IsDefaultProfileImage bool   `json:"is_default_profile_image"`
					IsFeedBlocked         bool   `json:"is_feed_blocked"`
				} `json:"writer"`
				Liked bool `json:"liked"`
			} `json:"comments"`
			Bookmarked         bool   `json:"bookmarked"`
			Verb               string `json:"verb"`
			Permission         string `json:"permission"`
			CommentAllWritable bool   `json:"comment_all_writable"`
			HasUnreadReaction  bool   `json:"has_unread_reaction"`
			ShareCount         int    `json:"share_count"`
			Actor              struct {
				Birthday          string `json:"birthday"`
				BgImageUrl2       string `json:"bg_image_url2"`
				ProfileVideoUrlHq string `json:"profile_video_url_hq"`
				IsFavorite        bool   `json:"is_favorite"`
				BirthdayLeft      int    `json:"birthday_left"`
				Gender            string `json:"gender"`
				BirthLeapType     bool   `json:"birth_leap_type"`
				BirthType         string `json:"birth_type"`
				ProfileImageUrl2  string `json:"profile_image_url2"`
				Type              string `json:"type"`
				FollowerCount     int    `json:"follower_count"`
				Relation          struct {
					FeedBlocked bool   `json:"feed_blocked"`
					Friend      string `json:"friend"`
					Self        bool   `json:"self"`
					Follow      string `json:"follow"`
					Favorite    bool   `json:"favorite"`
				} `json:"relation"`
				ProfileVideoUrlSquareSmall string `json:"profile_video_url_square_small"`
				IsCelebratable             bool   `json:"is_celebratable"`
				DefaultBgId                int    `json:"default_bg_id"`
				ActivityCount              int    `json:"activity_count"`
				IsValidUser                bool   `json:"is_valid_user"`
				ID                         string `json:"id"`
				Relationship               string `json:"relationship"`
				Vip                        bool   `json:"vip"`
				FriendCount                int    `json:"friend_count"`
				ProfileVideoUrlLq          string `json:"profile_video_url_lq"`
				IsBirthday                 bool   `json:"is_birthday"`
				BgImageUrl                 string `json:"bg_image_url"`
				ProfileVideoUrlSquare      string `json:"profile_video_url_square"`
				AllowFollowing             bool   `json:"allow_following"`
				ProfileThumbnailUrl        string `json:"profile_thumbnail_url"`
				StatusObjects              []struct {
					ObjectType string `json:"object_type"`
					Message    string `json:"message"`
				} `json:"status_objects"`
				ProfileVideoUrlSquareMicroSmall string `json:"profile_video_url_square_micro_small"`
				ProfileImageUrl                 string `json:"profile_image_url"`
				DisplayName                     string `json:"display_name"`
				Permalink                       string `json:"permalink"`
				IsDefaultProfileImage           bool   `json:"is_default_profile_image"`
				IsFeedBlocked                   bool   `json:"is_feed_blocked"`
			} `json:"actor"`
			LatestFriendEmotion []struct {
				Actor struct {
					BgImageUrl2      string `json:"bg_image_url2"`
					IsFavorite       bool   `json:"is_favorite"`
					Gender           string `json:"gender"`
					BirthLeapType    bool   `json:"birth_leap_type"`
					BirthType        string `json:"birth_type"`
					ProfileImageUrl2 string `json:"profile_image_url2"`
					Type             string `json:"type"`
					FollowerCount    int    `json:"follower_count"`
					Relation         struct {
						FeedBlocked bool   `json:"feed_blocked"`
						Friend      string `json:"friend"`
						Self        bool   `json:"self"`
						Follow      string `json:"follow"`
						Favorite    bool   `json:"favorite"`
					} `json:"relation"`
					IsCelebratable        bool          `json:"is_celebratable"`
					DefaultBgId           int           `json:"default_bg_id"`
					ActivityCount         int           `json:"activity_count"`
					IsValidUser           bool          `json:"is_valid_user"`
					ID                    string        `json:"id"`
					Relationship          string        `json:"relationship"`
					Vip                   bool          `json:"vip"`
					FriendCount           int           `json:"friend_count"`
					IsBirthday            bool          `json:"is_birthday"`
					BgImageUrl            string        `json:"bg_image_url"`
					AllowFollowing        bool          `json:"allow_following"`
					ProfileThumbnailUrl   string        `json:"profile_thumbnail_url"`
					StatusObjects         []interface{} `json:"status_objects"`
					ProfileImageUrl       string        `json:"profile_image_url"`
					DisplayName           string        `json:"display_name"`
					Permalink             string        `json:"permalink"`
					IsDefaultProfileImage bool          `json:"is_default_profile_image"`
					IsFeedBlocked         bool          `json:"is_feed_blocked"`
				} `json:"actor"`
				Emotion   string    `json:"emotion"`
				CreatedAt time.Time `json:"created_at"`
				ID        string    `json:"id"`
			} `json:"latest_friend_emotion"`
			PushMute       bool `json:"push_mute"`
			WithMe         bool `json:"with_me"`
			SympathyCount  int  `json:"sympathy_count"`
			LatestComments []struct {
				LikeCount  int `json:"like_count"`
				Decorators []struct {
					Text      string `json:"text"`
					ID        string `json:"id,omitempty"`
					Type      string `json:"type"`
					Permalink string `json:"permalink,omitempty"`
				} `json:"decorators"`
				ActivityId string    `json:"activity_id"`
				Index      string    `json:"index"`
				CreatedAt  time.Time `json:"created_at"`
				ID         string    `json:"id"`
				Text       string    `json:"text"`
				Writer     struct {
					Birthday          string `json:"birthday"`
					BgImageUrl2       string `json:"bg_image_url2"`
					ProfileVideoUrlHq string `json:"profile_video_url_hq"`
					IsFavorite        bool   `json:"is_favorite"`
					BirthdayLeft      int    `json:"birthday_left"`
					Gender            string `json:"gender"`
					BirthLeapType     bool   `json:"birth_leap_type"`
					BirthType         string `json:"birth_type"`
					ProfileImageUrl2  string `json:"profile_image_url2"`
					Type              string `json:"type"`
					FollowerCount     int    `json:"follower_count"`
					Relation          struct {
						FeedBlocked bool   `json:"feed_blocked"`
						Friend      string `json:"friend"`
						Self        bool   `json:"self"`
						Follow      string `json:"follow"`
						Favorite    bool   `json:"favorite"`
					} `json:"relation"`
					ProfileVideoUrlSquareSmall string `json:"profile_video_url_square_small"`
					IsCelebratable             bool   `json:"is_celebratable"`
					DefaultBgId                int    `json:"default_bg_id"`
					ActivityCount              int    `json:"activity_count"`
					IsValidUser                bool   `json:"is_valid_user"`
					ID                         string `json:"id"`
					Relationship               string `json:"relationship"`
					Vip                        bool   `json:"vip"`
					FriendCount                int    `json:"friend_count"`
					ProfileVideoUrlLq          string `json:"profile_video_url_lq"`
					IsBirthday                 bool   `json:"is_birthday"`
					BgImageUrl                 string `json:"bg_image_url"`
					ProfileVideoUrlSquare      string `json:"profile_video_url_square"`
					AllowFollowing             bool   `json:"allow_following"`
					ProfileThumbnailUrl        string `json:"profile_thumbnail_url"`
					StatusObjects              []struct {
						ObjectType string `json:"object_type"`
						Message    string `json:"message"`
					} `json:"status_objects"`
					ProfileVideoUrlSquareMicroSmall string `json:"profile_video_url_square_micro_small"`
					ProfileImageUrl                 string `json:"profile_image_url"`
					DisplayName                     string `json:"display_name"`
					Permalink                       string `json:"permalink"`
					IsDefaultProfileImage           bool   `json:"is_default_profile_image"`
					IsFeedBlocked                   bool   `json:"is_feed_blocked"`
				} `json:"writer"`
				Liked bool `json:"liked"`
			} `json:"latest_comments"`
			Permalink string `json:"permalink"`
			ViewCount int    `json:"view_count"`
		} `json:"activity,omitempty"`
	} `json:"bookmarks"`
	SectionInfo struct {
		Count        int    `json:"count"`
		Type         string `json:"type"`
		Relationship string `json:"relationship"`
		Relation     struct {
			FeedBlocked bool   `json:"feed_blocked"`
			Friend      string `json:"friend"`
			Self        bool   `json:"self"`
			Follow      string `json:"follow"`
			Favorite    bool   `json:"favorite"`
			Ban         string `json:"ban"`
		} `json:"relation"`
	} `json:"section_info"`
}
