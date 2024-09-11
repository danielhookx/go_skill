package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Person struct {
	RemoteHost string `form:"remote"`
	Name       string `form:"name"`
}

func main() {
	g := gin.Default()
	g.GET("/test", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindQuery(&person); err != nil {
			c.String(http.StatusBadRequest, "err: %w", err)
		}

		params := url.Values{}
		params.Add("name", person.Name)

		u, _ := url.ParseRequestURI(fmt.Sprintf("http://%s", person.RemoteHost))
		u.Path = "/test"
		u.RawQuery = params.Encode()
		urlStr := fmt.Sprintf("%v", u)

		proxyStr := "http://106.15.108.53:1337"
		proxyURL, err := url.Parse(proxyStr)
		if err != nil {
			fmt.Println("Error parsing proxy URL:", err)
			return
		}

		tr := &http.Transport{
			Proxy:           http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		// Create an HTTP client
		client := &http.Client{
			Transport: tr,
		}

		// // 创建HTTP CONNECT请求
		// connectReq, err := http.NewRequest(http.MethodConnect, fmt.Sprintf("http://%s", person.RemoteHost), nil)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating CONNECT request: %v", err)})
		// 	return
		// }

		// // 发送CONNECT请求并获取响应
		// connectResp, err := client.Do(connectReq)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error sending CONNECT request: %v", err)})
		// 	return
		// }
		// defer connectResp.Body.Close()

		// // 打印CONNECT请求的响应
		// fmt.Println("CONNECT Response Status:", connectResp.Status)

		// if connectResp.StatusCode != http.StatusOK {
		// 	return
		// }

		// Create a new GET request
		req, err := http.NewRequest(http.MethodGet, urlStr, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Send the GET request
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("Get Response Status:", resp.Status)

		// Set the response content type and write the response body
		// c.Header("Content-Type", "application/json")
		c.String(resp.StatusCode, "proxy get %s --- body: %s", urlStr, string(body))
	})
	http.ListenAndServe(":1830", g)
}
