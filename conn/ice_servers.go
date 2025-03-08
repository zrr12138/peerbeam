package conn

import "github.com/pion/webrtc/v4"

var iceServers = []webrtc.ICEServer{
	{
		URLs:           []string{"turn:101.34.238.168:3478"},
		Username:       "ranber",
		Credential:     "12138",
		CredentialType: webrtc.ICECredentialTypePassword,
	},
}
