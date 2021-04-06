package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"os"
)

type WINRepository struct {
	ES *elasticsearch.Client
}

// Index name
var INDEX = os.Getenv("ELS_INDEX")

func ProvideWINRepository(ES *elasticsearch.Client) WINRepository {
	return WINRepository{ES: ES}
}

// search ElasticSearch Query를 파라미터로 받아 요청 후 결과를 반환한다
func (w *WINRepository) search(query map[string]interface{}) (map[string]interface{}, error){
	var buf bytes.Buffer
	var r  map[string]interface{}

	// 쿼리 인코딩에 실패한 경우 리턴
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("Error encoding query: %s\n", err)
	}
	// 검색 실행
	resp, err := w.ES.Search(
		w.ES.Search.WithContext(context.Background()),
		w.ES.Search.WithIndex(INDEX),
		w.ES.Search.WithBody(&buf),
		w.ES.Search.WithTrackTotalHits(true),
		w.ES.Search.WithPretty(),
	)
	if err != nil {
		return nil, fmt.Errorf("Error getting responses: %s\n", err)
	}
	defer resp.Body.Close()

	// 검색 응답에 에러가 있는 경우, 에러 리턴
	if resp.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			return nil, fmt.Errorf("Error parsing the response body: %s\n", err)
		} else {
			// Print the response status and error information.
			return nil, fmt.Errorf("[%s] %s: %s",
				resp.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	// 응답 디코딩에 실패한 경우 에러 리턴
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("Error parsing the response body: %s\n", err)
	}

	return r, nil
}

// List 현재시간 기준 6시간 전까지의 상위 이슈 30개를 반환한다
func (w *WINRepository) List() (map[string]interface{}, error){
	// Search Query
	query := map[string]interface{}{
		"size": 0,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{"match_all": map[string]interface{}{}},
				},
				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"createdAt": map[string]interface{}{
							"gte": "now-3h/H",
							"lte": "now",
						},
					},
				},
			},
		},
		"aggs": map[string]interface{}{
			"countByWord": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "word",
					"size": 30,
				},
			},
		},
	}

	// 검색을 요청한다
	r, err := w.search(query)
	if err != nil {
		return nil, err
	}


	return r, nil
}


// FindWordToTagPercent 이슈 Word의 태그별 점유율을 반환한다
func (w *WINRepository) FindWordToTagPercent(word string) (map[string]interface{}, error){
	query := map[string]interface{}{
		"size": 0,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match": map[string]string{
							"word": word,
						},
					},
				},
				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"createdAt": map[string]interface{}{
							"gte": "now-3h/H",
							"lte": "now",
						},
					},
				},
			},
		},
		"aggs": map[string]interface{}{
			"countPerTag": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "tag",
					"size": 10,
				},
			},
		},
	}

	r, err := w.search(query)
	if err != nil {
		return nil, err
	}

	return r, nil
}
