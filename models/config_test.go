package models_test

import (
	"encoding/json"
	"testing"
	"turboenigma/models"

	"github.com/stretchr/testify/assert"
)


func TestNotificationConfig(t *testing.T) {
	t.Run("Single Notification Config can be parsed", func(t *testing.T) {
		var jsonString string = "[{\"channel\":\"#abc\",\"labels\":[\"test-a\",\"test-b\"]}]"
		var config []models.NotificationConfig

		var err = json.Unmarshal([]byte(jsonString), &config)

		assert.Empty(t, err)
		assert.Equal(t, "#abc", config[0].Channel)
		assert.Equal(t, "test-a", config[0].Labels[0])
		assert.Equal(t, "test-b", config[0].Labels[1])
	})
}

func TestMultipleNotificationConfigs(t *testing.T) {
	t.Run("Multiple Notification Configs can be parsed", func(t *testing.T) {
		var jsonString string = "[{\"channel\":\"#123\",\"labels\":[\"test-c\"]},{\"channel\":\"#456\",\"labels\":[\"test-d\"]}]"
		var config []models.NotificationConfig

		var err = json.Unmarshal([]byte(jsonString), &config)

		assert.Empty(t, err)
		assert.Equal(t, "#123", config[0].Channel)
		assert.Equal(t, "test-c", config[0].Labels[0])
		assert.Equal(t, "#456", config[1].Channel)
		assert.Equal(t, "test-d", config[1].Labels[0])
	})
}
