package model_test

import (
	"encoding/json"
	"testing"
	"turboenigma/model"

	"github.com/stretchr/testify/assert"
)

func TestNotificationConfig(t *testing.T) {
	t.Run("Single Notification Config can be parsed", func(t *testing.T) {
		var jsonString string = "[{\"channel\":\"#abc\",\"labels\":[\"test-a\",\"test-b\"]}]"
		var notificationRules []model.NotificationRule

		var err = json.Unmarshal([]byte(jsonString), &notificationRules)

		assert.Empty(t, err)
		assert.Equal(t, "#abc", notificationRules[0].Channel)
		assert.Equal(t, "test-a", notificationRules[0].Labels[0])
		assert.Equal(t, "test-b", notificationRules[0].Labels[1])
	})
}

func TestMultipleNotificationConfigs(t *testing.T) {
	t.Run("Multiple Notification Configs can be parsed", func(t *testing.T) {
		var jsonString string = "[{\"channel\":\"#123\",\"labels\":[\"test-c\"]},{\"channel\":\"#456\",\"labels\":[\"test-d\"]}]"

		var notificationRules []model.NotificationRule

		var err = json.Unmarshal([]byte(jsonString), &notificationRules)

		assert.Empty(t, err)
		assert.Equal(t, "#123", notificationRules[0].Channel)
		assert.Equal(t, "test-c", notificationRules[0].Labels[0])
		assert.Equal(t, "#456", notificationRules[1].Channel)
		assert.Equal(t, "test-d", notificationRules[1].Labels[0])
	})
}
