package presenter

import "time"

type Record struct {
	Key        		string `json:"key"`
	TotalCount 		int    `json:"totalCount"`
	CreatedAt  		time.Time `json:"createdAt"`
}

type Response struct {
	Code        	 int `json:"code"`
	Msg string      `json:"msg"`
	Records[]		 *Record `json:"records"`
}