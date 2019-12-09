package internal

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion(t *testing.T) {
	qMatcher := NewBasicQuestionMatcher()
	question, err := qMatcher.Match("que seria tal coisa")

	assert.Nil(t, err)
	assert.Equal(t, "tal coisa", question.Term)
}

func TestControllerPostgres(t *testing.T) {

	aController := NewController(nil, nil)
	out := new(bytes.Buffer)
	err := aController.Answer("que coisa", out)

	assert.Nil(t, err)
	assert.Equal(t, "", out.String())
}
