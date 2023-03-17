package data

import (
	"github.com/Yostardev/cmdb-sdk/general"
	"github.com/Yostardev/json"
)

func NewDistinctData(f []any) []byte {
	b, _ := json.Marshal(&f)
	return b
}

func ParseDistinctData(b []byte) []any {
	var f []any
	_ = json.Unmarshal(b, &f)
	return f
}

func NewFilter(f map[string]any) []byte {
	b, _ := json.Marshal(&f)
	return b
}

func ParseFilter(b []byte) map[string]any {
	var f = make(map[string]any)
	_ = json.Unmarshal(b, &f)
	return f
}

func NewDocumentList(d []map[string]any) []byte {
	b, _ := json.Marshal(&d)
	return b
}

func ParseDocumentList(b []byte) []map[string]any {
	var f []map[string]any
	_ = json.Unmarshal(b, &f)
	return f
}

func NewDocument(d map[string]any) []byte {
	b, _ := json.Marshal(&d)
	return b
}

func ParseDocument(b []byte) map[string]any {
	var f map[string]any
	_ = json.Unmarshal(b, &f)
	return f
}

type ListDocument struct {
	Paginate *general.PaginateInfo `json:"paginate"`
	Items    any                   `json:"items"`
}

func NewListDocument(d *ListDocumentResponse) *ListDocument {
	return &ListDocument{
		d.Paginate,
		ParseDocumentList(d.Data),
	}
}

func NewTreeNodes(d []*TreeNode) []byte {
	b, _ := json.Marshal(&d)
	return b
}

func ParseTreeNodes(b []byte) []*TreeNode {
	var f []*TreeNode
	_ = json.Unmarshal(b, &f)
	return f
}

type TreeNode struct {
	Title         string         `json:"title"`
	Value         string         `json:"value"`
	NextFilter    map[string]any `json:"next_filter"`
	NextTag       string         `json:"next_tag"`
	ResourceCount int64          `json:"resource_count"`
	Icon          string         `json:"icon"`
	IsLeaf        bool           `json:"is_leaf"`
}
