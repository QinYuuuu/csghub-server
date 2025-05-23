package imagebuilder

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

var _ Builder = (*RemoteBuilder)(nil)

type RemoteBuilder struct {
	remote *url.URL
	client *http.Client
}

func NewRemoteBuilder(remoteURL string) (*RemoteBuilder, error) {
	parsedURL, err := url.Parse(remoteURL)
	if err != nil {
		return nil, err
	}
	return &RemoteBuilder{
		remote: parsedURL,
		client: http.DefaultClient,
	}, nil
}

func (h *RemoteBuilder) Build(ctx context.Context, req *BuildRequest) (*BuildResponse, error) {
	rel := &url.URL{Path: "/api/v1/imagebuilder/builder"}
	u := h.remote.ResolveReference(rel)
	response, err := h.doRequest(http.MethodPost, u.String(), req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var buildResponse BuildResponse
	if err := json.NewDecoder(response.Body).Decode(&buildResponse); err != nil {
		return nil, err
	}

	return &buildResponse, nil
}

func (h *RemoteBuilder) Status(ctx context.Context, req *StatusRequest) (*StatusResponse, error) {
	u := fmt.Sprintf("%s/api/v1/imagebuilder/%s/%s/status?build_id=%s", h.remote, req.OrgName, req.SpaceName, req.BuildID)
	response, err := h.doRequest(http.MethodGet, u, req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result map[string]int = make(map[string]int)
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}
	var imageID string
	var code int
	for k, v := range result {
		imageID = k
		code = v
		break
	}

	var statusResponse StatusResponse
	statusResponse.ImageID = imageID
	statusResponse.Code = code
	return &statusResponse, nil
}

func (h *RemoteBuilder) Logs(ctx context.Context, req *LogsRequest) (<-chan string, error) {
	u := fmt.Sprintf("%s/api/v1/imagebuilder/%s/%s/logs?build_id=%s", h.remote, req.OrgName, req.SpaceName, req.BuildID)

	rc, err := h.doStreamRequest(ctx, http.MethodGet, u, req)
	if err != nil {
		return nil, err
	}

	return h.readToChannel(rc), nil
}

func (h *RemoteBuilder) readToChannel(rc io.ReadCloser) <-chan string {
	output := make(chan string, 2)
	scanner := bufio.NewScanner(rc) // Use Scanner to split by line

	go func() {
		defer func() {
			rc.Close()
			close(output)
		}()
		for scanner.Scan() {
			line := scanner.Text()
			output <- line
		}
		if err := scanner.Err(); err != nil {
			slog.Error("remote builder log read failed", slog.Any("error", err))
			return
		}
	}()

	return output
}

// Helper method to execute the actual HTTP request and read the response.
func (h *RemoteBuilder) doRequest(method, url string, data interface{}) (*http.Response, error) {
	var buf io.Reader
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errData interface{}
		err := json.NewDecoder(resp.Body).Decode(&errData)
		if err != nil {
			return nil, fmt.Errorf("unexpected http status: %d, error: %w", resp.StatusCode, err)
		} else {
			return nil, fmt.Errorf("unexpected http status: %d, error: %v", resp.StatusCode, errData)
		}
	}

	return resp, nil
}

func (h *RemoteBuilder) doStreamRequest(ctx context.Context, method, url string, data interface{}) (io.ReadCloser, error) {
	var buf io.Reader
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Connection", "keep-alive")

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected http status code:%d", resp.StatusCode)
	}

	return resp.Body, nil
}
