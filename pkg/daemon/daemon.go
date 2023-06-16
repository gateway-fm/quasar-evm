package daemon

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/misnaged/annales/logger"
)

type Client struct {
	addr string
	ws   *websocket.Conn
	stop chan struct{}
}

func NewClient(addr string, handshakeTimeout time.Duration) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), handshakeTimeout)
	defer cancel()

	var h http.Header

	conn, _, err := websocket.DefaultDialer.DialContext(ctx, addr, h)
	if err != nil {
		return nil, fmt.Errorf("weboscket dial to daemon at %s: %w", addr, err)
	}

	client := &Client{
		addr: addr,
		ws:   conn,
		stop: make(chan struct{}),
	}

	go client.listenWsResponses()

	return client, nil
}

const closeConErr = "use of closed network connection"

// listenWsRequests start ws response listener loop
func (c *Client) listenWsResponses() {
	//defer close(c.wsResponse)
	logger.Log().Info("start listening incoming daemon ws messages")

	for {
		select {
		case <-c.stop:
			logger.Log().Info("stop listening incoming daemon ws messages")
			return
		default:
			_, message, err := c.ws.ReadMessage()
			if err != nil {
				if !strings.Contains(err.Error(), closeConErr) {
					logger.Log().Error(fmt.Errorf("read ws message %T: %w, %T", err, err, errors.Unwrap(err)).Error())
				}

				c.Close()
				return
			}

			fmt.Println(string(message))

			/*if err := c.handleWsMsg(message); err != nil {
				logger.Log().Error(fmt.Errorf("handle ws message %T: %w", err, err).Error())
				c.Close()
				return
			}*/
		}
	}
}

// Close all connections to node and stop listeners
func (c *Client) Close() {
	close(c.stop)

	if err := c.ws.Close(); err != nil {
		logger.Log().Error(fmt.Errorf("error to close ws connection to daemon at %s: %w", c.addr, err).Error())
	}
}
