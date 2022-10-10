package routers

import (
	"testing"

	"github.com/gorilla/mux"
	corev2 "github.com/sensu/core/v2"
	"github.com/sensu/sensu-go/testing/mockstore"
)

func TestMutatorsRouter(t *testing.T) {
	// Setup the router
	s := &mockstore.MockStore{}
	router := NewMutatorsRouter(s)
	parentRouter := mux.NewRouter().PathPrefix(corev2.URLPrefix).Subrouter()
	router.Mount(parentRouter)

	empty := &corev2.Mutator{}
	fixture := corev2.FixtureMutator("foo")

	tests := []routerTestCase{}
	tests = append(tests, getTestCases(fixture)...)
	tests = append(tests, listTestCases(empty)...)
	tests = append(tests, createTestCases(empty)...)
	tests = append(tests, updateTestCases(fixture)...)
	tests = append(tests, deleteTestCases(fixture)...)
	for _, tt := range tests {
		run(t, tt, parentRouter, s)
	}
}
