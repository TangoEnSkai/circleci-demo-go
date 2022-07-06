package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandlerOne(t *testing.T) {
	a := App{}
	a.Init()

	testCases := map[string]struct {
		// req *http.Request
		req        func() *http.Request
		wantStatus int
		wantErr    bool
		err        error
	}{
		"successfuly handle get with ok": {
			req: func() *http.Request {
				req, err := http.NewRequest("GET", "/", nil)
				if err != nil {
					t.Fatal(err)
				}
				return req
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
			err:        nil,
		},
	}

	for n, tt := range testCases {
		name, tc := n, tt

		t.Run(name, func(t *testing.T) {
			r := httptest.NewRecorder()

			a.Router.ServeHTTP(r, tc.req())

			if status := r.Code; status != tc.wantStatus {
				t.Errorf(
					"unexpected status: got (%v), want (%v)",
					status,
					tc.wantStatus,
				)
			}
		})
	}
}

func TestIndexHandlerTwo(t *testing.T) {
	a := App{}
	a.Init()

	testCases := map[string]struct {
		// req *http.Request
		req        func() *http.Request
		wantStatus int
		wantErr    bool
		err        error
	}{
		"successfuly handle get with ok": {
			req: func() *http.Request {
				req, err := http.NewRequest("GET", "/", nil)
				if err != nil {
					t.Fatal(err)
				}
				return req
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
			err:        nil,
		},
	}

	for n, tt := range testCases {
		name, tc := n, tt

		t.Run(name, func(t *testing.T) {
			r := httptest.NewRecorder()

			a.Router.ServeHTTP(r, tc.req())

			if status := r.Code; status != tc.wantStatus {
				t.Errorf(
					"unexpected status: got (%v), want (%v)",
					status,
					tc.wantStatus,
				)
			}
		})
	}
}

func TestIndexHandlerThree(t *testing.T) {
	a := App{}
	a.Init()

	testCases := map[string]struct {
		// req *http.Request
		req        func() *http.Request
		wantStatus int
		wantErr    bool
		err        error
	}{
		"successfuly handle get with ok": {
			req: func() *http.Request {
				req, err := http.NewRequest("GET", "/", nil)
				if err != nil {
					t.Fatal(err)
				}
				return req
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
			err:        nil,
		},
	}

	for n, tt := range testCases {
		name, tc := n, tt

		t.Run(name, func(t *testing.T) {
			r := httptest.NewRecorder()

			a.Router.ServeHTTP(r, tc.req())

			if status := r.Code; status != tc.wantStatus {
				t.Errorf(
					"unexpected status: got (%v), want (%v)",
					status,
					tc.wantStatus,
				)
			}
		})
	}
}
