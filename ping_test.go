package ping

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestF(t *testing.T) {
	var p Ping

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	h := http.HandlerFunc(F)
	h.ServeHTTP(w, r)

	resp := w.Result()
	// if resp.StatusCode != http.StatusOK {
	// 	t.Errorf("wrong statsu code: gpt %v instead of %v", resp.StatusCode, http.StatusOK)
	// }
	err = json.NewDecoder(resp.Body).Decode(&p)
	// assert on correctness of response
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, "Hello World!", p.Message)
}
