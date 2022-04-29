package lib

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//BringMeLonger is requests for actual URL
func BringMeLonger(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
	defer cancel()
	shortURLPath := c.Param("shortURL")
	ok, resp := getByUrlPath(ctx, shortURLPath)
	if !ok {
		c.JSON(http.StatusNoContent, gin.H{
			"error": "there is no such shortlink or expired",
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, resp.Url)
}

//MakeThisShorter is handler function for shortening URLs
func MakeThisShorter(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
	defer cancel()

	//get POSTed request
	req := request{
		url:    c.PostForm("url"),
		custom: c.PostForm("custom"),
	}
	dur := c.PostForm("validfor")
	if dur != "" {
		durP, err := time.ParseDuration(dur)
		if err != nil {
			log.Println("duration couldn't be parsed:", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  err.Error(),
				"detail": "duration should be in hours, minutes or sec like following format: '23h30m05s' ",
			})
			return
		}
		req.validFor = durP
	} else {
		req.validFor, _ = time.ParseDuration("8760h")
	}
	//check if custom URL was requested
	if req.custom != "" {
		if checkIfExist(ctx, req.custom) {
			c.JSON(http.StatusConflict, gin.H{
				"error":  "already exist",
				"detail": "custom url is already exist: " + req.custom,
			})
			return
		}

	} else {
		//if not, new custom URL is generated
		req.custom = generateUniqueShortLinkString(ctx, 7)
	}

	// data to be sent
	resp := Response{
		Url:      req.url,
		UrlPath:  req.custom,
		ShortURL: DOMAIN + "/" + req.custom,
		ExpireAt: time.Now().Add(req.validFor),
	}
	//save to db
	err := insertNewURL(ctx, &resp)
	if err != nil {
		log.Println("insertion failed:", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"inserted URL":      resp.Url,
		"created short URL": resp.ShortURL,
		"valid until":       resp.ExpireAt.Format("2 Jan 2006 on Monday at 15:04:05"),
	})

}
