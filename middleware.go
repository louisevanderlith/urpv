package urpv

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func VerifyAuthentication(host string) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.GetHeader("Authorization")

		if len(reqToken) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			log.Println("bad request")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		data := url.Values{
			"token": {strings.Trim(splitToken[1], " ")},
		}
		log.Println("Encode token:", data.Encode())
		resp, err := http.PostForm(fmt.Sprintf("https://urpv%sinfo", host), data)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Println("not ok", resp.StatusCode)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		result := make(map[string]interface{})
		dec := json.NewDecoder(resp.Body)
		err = dec.Decode(&result)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Set("info", result)
	}
}
