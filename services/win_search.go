package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/99-66/NaiaBackendApi/repositories"
)

// search ElasticSearch Query를 파라미터로 받아 요청 후 결과를 반환한다
func search(query map[string]interface{}) (map[string]interface{}, error) {
	var buf bytes.Buffer
	var r map[string]interface{}

	// 쿼리 인코딩에 실패한 경우 리턴
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("Error encoding query: %s\n", err)
	}

	// 검색 실행
	es := repositories.Connections.ES
	resp, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(repositories.ESMeta.Index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
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
