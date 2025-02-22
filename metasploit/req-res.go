package metasploit

type sessionListReq struct {
	_msgpack      struct{} `msgpack:",asArray"`
	Method, Token string
}

type SessionListRes struct {
	ID          uint32 `msgpack:",omitempty"`
	Type        string `msgpack:"type"`
	TunnelLocal string `msgpack:"tunnel_local"`
	TunnelPeer  string `msgpack:"tunnel_peer"`
	ViaExploit  string `msgpack:"via_exploit"`
	ViaPayload  string `msgpack:"via_payload"`
	Description string `msgpack:"desc"`
	Workspace   string `msgpack:"workspace"`
	SessionHost string `msgpack:"session_host"`
	SessionPort string `msgpack:"session_port"`
	Username    string `msgpack:"username"`
	UUID        string `msgpack:"uuid"`
	ExploitUUID string `msgpack:"exploit_uuid"`
}

type loginReq struct {
	_msgpack                   struct{} `msgpack:",asArray"`
	Method, Username, Password string
}

type LoginRes struct {
	Result       string `msgpack:"result"`
	Token        string `msgpack:"token"`
	Error        string `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type logoutReq struct {
	_msgpack                   struct{} `msgpack:",asArray"`
	Method, Token, LogoutToken string
}
