package main

import "time"

type User struct {
	ID                   string    `json:"id"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	Email                string    `json:"email"`
	GoogleRefreshToken   string    `json:"google_refresh_token"`
	FacebookRefreshToken string    `json:"facebook_refresh_token"`
	TwitterRefreshToken  string    `json:"twitter_refresh_token"`
	GithubRefreshToken   string    `json:"github_refresh_token"`
	LastLoginAt          time.Time `json:"-"`
	CreatedAt            time.Time `json:"-"`
	UpdatedAt            time.Time `json:"-"`
}

func NewGoogleUser(id, email, refreshToken string) (*User, error) {
	return nil, nil
}
