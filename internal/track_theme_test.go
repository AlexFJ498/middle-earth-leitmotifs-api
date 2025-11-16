package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTrackTheme_ValidInput(t *testing.T) {
	trackID := "456e7890-e89b-12d3-a456-426614174117"
	themeID := "456e7890-e89b-12d3-a456-426614174118"
	startSecond := 30
	endSecond := 90
	isVariant := false

	trackTheme, err := NewTrackTheme(trackID, themeID, startSecond, endSecond, isVariant)

	assert.NoError(t, err)
	assert.Equal(t, startSecond, trackTheme.StartSecond().Int())
	assert.Equal(t, endSecond, trackTheme.EndSecond().Int())
	assert.Equal(t, isVariant, trackTheme.IsVariant().Bool())
}

func TestNewTrackTheme_EndSecondLessThanStartSecond(t *testing.T) {
	trackID := "456e7890-e89b-12d3-a456-426614174117"
	themeID := "456e7890-e89b-12d3-a456-426614174118"
	startSecond := 90
	endSecond := 30
	isVariant := false

	_, err := NewTrackTheme(trackID, themeID, startSecond, endSecond, isVariant)

	assert.Error(t, err)
	assert.Equal(t, ErrEndSecondMustBeGreaterThanStartSecond, err)
}

func TestNewTrackTheme_EndSecondEqualToStartSecond(t *testing.T) {
	trackID := "456e7890-e89b-12d3-a456-426614174117"
	themeID := "456e7890-e89b-12d3-a456-426614174118"
	startSecond := 60
	endSecond := 60
	isVariant := false

	_, err := NewTrackTheme(trackID, themeID, startSecond, endSecond, isVariant)

	assert.Error(t, err)
	assert.Equal(t, ErrEndSecondMustBeGreaterThanStartSecond, err)
}

func TestNewTrackTheme_InvalidTrackID(t *testing.T) {
	trackID := "invalid-uuid"
	themeID := "456e7890-e89b-12d3-a456-426614174118"
	startSecond := 30
	endSecond := 90
	isVariant := false

	_, err := NewTrackTheme(trackID, themeID, startSecond, endSecond, isVariant)

	assert.Error(t, err)
}

func TestNewTrackTheme_InvalidThemeID(t *testing.T) {
	trackID := "456e7890-e89b-12d3-a456-426614174117"
	themeID := "invalid-uuid"
	startSecond := 30
	endSecond := 90
	isVariant := false

	_, err := NewTrackTheme(trackID, themeID, startSecond, endSecond, isVariant)

	assert.Error(t, err)
}

func TestNewTrackTheme_NegativeStartSecond(t *testing.T) {
	trackID := "456e7890-e89b-12d3-a456-426614174117"
	themeID := "456e7890-e89b-12d3-a456-426614174118"
	startSecond := -10
	endSecond := 90
	isVariant := false

	_, err := NewTrackTheme(trackID, themeID, startSecond, endSecond, isVariant)

	assert.Error(t, err)
	assert.Equal(t, ErrInvalidStartSecond, err)
}

func TestNewTrackTheme_NegativeEndSecond(t *testing.T) {
	trackID := "456e7890-e89b-12d3-a456-426614174117"
	themeID := "456e7890-e89b-12d3-a456-426614174118"
	startSecond := 30
	endSecond := -10
	isVariant := false

	_, err := NewTrackTheme(trackID, themeID, startSecond, endSecond, isVariant)

	assert.Error(t, err)
	assert.Equal(t, ErrInvalidEndSecond, err)
}
