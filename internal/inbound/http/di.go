package http

import (
	"github.com/asnur/vocagame-be-interview/internal/inbound/http/transaction"
	"github.com/asnur/vocagame-be-interview/internal/inbound/http/user"
	"github.com/asnur/vocagame-be-interview/internal/inbound/http/wallet"
	"go.uber.org/dig"
)

type (
	Inbound struct {
		dig.In

		Transaction transaction.Controller

		User user.Controller

		Wallet wallet.Controller
	}
)
