package info

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo_Validate(t *testing.T) {
	info := Info{
		AppName:     "pkg",
		Description: "pkg description",
		Version:     "1.0.0",
	}

	assert.Nil(t, info.Validate())
}

func TestInfo_Invalidate(t *testing.T) {
	infoTable := []Info{{AppName: "pkg", Description: "pkg description"}, {Description: "pkg description"}, {Version: "1.0.0"}}

	for i := range infoTable {
		assert.Error(t, infoTable[i].Validate())
	}
}
