package pkg

import "time"

type OpActionType string

const (
	BindWallet   OpActionType = "create account"                    // 绑定钱包指令
	QueryBalance OpActionType = "query"                             // 查询钱包余额
	TransferTo   OpActionType = `transfer\s+([\d\.]+)\s+to\s+(\d+)` // 转账
)

type Op struct {
	Action    OpActionType `json:"action"`
	Timestamp time.Time    `json:"timestamp"`
	From      string       `json:"from"`
	To        string       `json:"to"`
	Value     string       `json:"value"`
	Message   string       `json:"message"`
	OpId      string       `json:"opId"`     // op hashed value, unique global
	PrevOpId  string       `json:"prevOpId"` // represents the op which trigger this op produced if not empty.
}
