package recorder

import (
	"os"
	"testing"

	livekit "github.com/livekit/protocol/proto"
	"github.com/stretchr/testify/require"

	"github.com/livekit/livekit-recorder/pkg/config"
)

func TestInputUrl(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	req := &livekit.StartRecordingRequest{
		Input: &livekit.StartRecordingRequest_Template{
			Template: &livekit.RecordingTemplate{
				Layout: "speaker-light",
				Room: &livekit.RecordingTemplate_Token{
					Token: token,
				},
			},
		},
	}
	recorderSiteAddress := os.Getenv("RECORDER_SITE_ADDRESS")
	if len(recorderSiteAddress) == 0 {
		recorderSiteAddress = "recorder.livekit.io"
	}

	expected := "https://" + recorderSiteAddress + "/#/speaker-light?url=wss%3A%2F%2Ftest.livekit.cloud&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	rec := &Recorder{
		conf: &config.Config{WsUrl: "wss://test.livekit.cloud"},
	}

	actual, err := rec.GetInputUrl(req)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}
