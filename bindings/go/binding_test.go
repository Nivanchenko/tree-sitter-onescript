package tree_sitter_onescript_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_onescript "github.com/eightm/tree-sitter-onescript/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_onescript.Language())
	if language == nil {
		t.Errorf("Error loading OneScript grammar")
	}
}
