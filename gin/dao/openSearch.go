package dao

import (
	"context"
	"fmt"
	"net/http"
	"time"

	//commonDB "git.89trillion.com/89t/server/base-service-common/system/db"
	"github.com/gin-gonic/gin"
	"github.com/opensearch-project/opensearch-go/v3"
	"github.com/opensearch-project/opensearch-go/v3/opensearchapi"
)

var osConn = &opensearch.Client{}

func InitOpenSearch() {

	cfg := opensearch.Config{
		Addresses: []string{
			"http://10.0.1.84:32000",
			//"http://localhost:9201",
		},
		Username: "admin",
		Password: "admin",
		Transport: &http.Transport{
			MaxIdleConns:        300,
			IdleConnTimeout:     30 * time.Second,
			DisableCompression:  true,
			MaxIdleConnsPerHost: 100,
		},
	}

	var err error
	osConn, err = opensearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
}

func GetFromCommonOpenSearch(ctx *gin.Context) {
	response, err := osConn.Do(context.TODO(), opensearchapi.DocumentGetReq{
		Index:      "offline_match_iw",
		DocumentID: "userId10",
	}, nil)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("CommonOpenSearch 查询到结果:%v", response),
	})
}

func GetFromOpenSearch(ctx *gin.Context) {
	response, err := osConn.Do(context.TODO(), opensearchapi.DocumentGetReq{
		Index:      "offline_match_iw",
		DocumentID: "userId10",
	}, nil)
	if err != nil {
		panic(err)
	}

	response, err = osConn.Do(context.TODO(), opensearchapi.DocumentGetReq{
		Index:      "offline_match_iw",
		DocumentID: "userId10",
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("OpenSearch 查询到结果:%v", response))

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("OpenSearch 查询到结果:%v", response),
	})
}

func WriteToOpenSearch(ctx *gin.Context) {
	response, err := osConn.Do(context.TODO(), opensearchapi.IndicesCreateReq{Index: "offline_match_iw"}, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("OpenSearch 查询到结果:%v", response))

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("OpenSearch 查询到结果:%v", response),
	})
}
