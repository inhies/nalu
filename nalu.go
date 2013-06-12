// Package nalu allows the easy creation of a STOMP API server using WebSockets
package nalu

import (
	"fmt"
	websocket "github.com/inhies/gowebsocket"
	"github.com/jjeffery/stomp"
	stompServer "github.com/jjeffery/stomp/server"
	"net"
	"net/http"
	"time"
)

// API represents a STOMP API server
type API struct {
	Stomp    *stomp.Conn // Server side connection to the API
	listener listener    // Send new connections here
}

// Starts a STOMP server and returns a new API. Use the Upgrade() method to
// upgrade an http.Request to a STOMP WebSocket. Set heartbeat to the minimum
// interval you want to send and receive heartbeats and  StompOpts to your
// desired settings for the STOMP server.
func NewAPI(heartbeat time.Duration, StompOpts stomp.Options) (conn *API,
	err error) {

	// Create the API and initialize the connection listener
	conn = &API{
		listener: make(chan net.Conn),
	}

	stompCfg := stompServer.Server{
		HeartBeat: heartbeat,
	}
	// Start a new stomp server and wait for new connections
	go stompCfg.Serve(conn.listener)

	// Connect to the server locally
	err = conn.connectLocal(StompOpts)
	if err != nil {
		fmt.Println("Errrr")
		return nil, err
	}
	return
}

// Upgrades an HTTP connection to a WebSocket and passes the WebSocket to the
// STOMP server for protocol negotiation.
func (a *API) Upgrade(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	if req.Header.Get("Origin") != "http://"+req.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	// Create a net.Conn from the websocket
	conn, err := websocket.NewConn(w, req, req.Header, 4096, 4096)
	if err != nil {
		// TODO(inhies): Figure out a way to return an error from this
		fmt.Println("error getting websocket conn:", err)
		return
	}

	// Send this connection to our STOMP server
	a.listener <- conn
}

// Creates the local STOMP cilent connection for our program to communicate with
// the API
func (a *API) connectLocal(opts stomp.Options) (err error) {
	// Create a pipe from our code to the API without sockets
	conn1, conn2 := net.Pipe()

	// Send the first connection to the STOMP server
	a.listener <- conn1

	// Turn the second conenction in to a STOMP client that we'll use locally
	a.Stomp, err = stomp.Connect(conn2, opts)
	if err != nil {
		fmt.Println("err:", err)
	}
	return
}
