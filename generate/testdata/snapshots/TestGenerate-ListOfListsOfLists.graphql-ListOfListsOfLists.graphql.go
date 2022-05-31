// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
)

// ListOfListsOfListsResponse is returned by ListOfListsOfLists on success.
type ListOfListsOfListsResponse struct {
	ListOfListsOfLists [][][]string `json:"listOfListsOfLists"`
}

// GetListOfListsOfLists returns ListOfListsOfListsResponse.ListOfListsOfLists, and is useful for accessing the field via an interface.
func (v *ListOfListsOfListsResponse) GetListOfListsOfLists() [][][]string {
	return v.ListOfListsOfLists
}

func ListOfListsOfLists(
	client graphql.Client,
	opts ...graphql.RequestOption,
) (*ListOfListsOfListsResponse, error) {
	req := &graphql.Request{
		OpName: "ListOfListsOfLists",
		Query: `
query ListOfListsOfLists {
	listOfListsOfLists
}
`,
	}
	var err error

	var data ListOfListsOfListsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
		opts...,
	)

	return &data, err
}

