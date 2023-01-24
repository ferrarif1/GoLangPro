package subgraph

type UserData struct {
	Users User `json:"user,omitempty"`
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID                  string         `json:"id"`
	Address             string         `json:"address,omitempty"`
	FirstTradeTimestamp int64          `json:"firstTradeTimestamp,omitempty"`
	IsSolver            bool           `json:"isSolver,omitempty"`
	NumberOfTrades      int64          `json:"numberOfTrades,omitempty"`
	SolvedAmountEth     string         `json:"solvedAmountEth,omitempty"`
	SolvedAmountUsd     string         `json:"solvedAmountUsd,omitempty"`
	TradedAmountUsd     string         `json:"tradedAmountUsd,omitempty"`
	TradedAmountEth     string         `json:"tradedAmountEth,omitempty"`
	OrdersPlaced        []OrdersPlaced `json:"ordersPlaced,omitempty"`
}

type OrdersPlaced struct {
	ID       *string `json:"id"`
	IsSigned *bool   `json:"isSigned,omitempty"`
}

type Tokens struct {
	Tokens []Token `json:"tokens"`
}

type Token struct {
	ID                  string `json:"id"`
	Address             string `json:"address,omitempty"`
	FirstTradeTimestamp int    `json:"firstTradeTimestamp,omitempty"`
	Name                string `json:"name,omitempty"`
	Symbol              string `json:"symbol,omitempty"`
	Decimals            int    `json:"decimals,omitempty"`
	TotalVolume         string `json:"totalVolume,omitempty"`
	PriceEth            string `json:"priceEth,omitempty"`
	PriceUsd            string `json:"priceUsd,omitempty"`
	NumberOfTrades      int    `json:"numberOfTrades,omitempty"`
	TotalVolumeUsd      string `json:"totalVolumeUsd,omitempty"`
	TotalVolumeEth      string `json:"totalVolumeEth,omitempty"`
}

type Orders struct {
	Orders []Order `json:"orders"`
}

type Order struct {
	ID                  string  `json:"id"`
	TradesTimestamp     int     `json:"tradesTimestamp,omitempty"`
	InvalidateTimestamp int     `json:"invalidateTimestamp,omitempty"`
	PresignTimestamp    int     `json:"presignTimestamp,omitempty"`
	IsSigned            bool    `json:"isSigned,omitempty"`
	IsValid             bool    `json:"isValid,omitempty"`
	Owner               User    `json:"owner,omitempty"`
	Trades              []Trade `json:"trades,omitempty"`
}

type Trades struct {
	Trades []Trade `json:"trades,omitempty"`
}

type Trade struct {
	ID            string     `json:"id"`
	Timestamp     int        `json:"timestamp,omitempty"`
	GasPrice      string     `json:"gasPrice,omitempty"`
	FeeAmount     string     `json:"feeAmount,omitempty"`
	TxHash        string     `json:"txHash,omitempty"`
	Settlement    Settlement `json:"settlement,omitempty"`
	BuyAmount     string     `json:"buyAmount,omitempty"`
	SellAmount    string     `json:"sellAmount,omitempty"`
	BuyAmountUsd  string     `json:"buyAmountUsd,omitempty"`
	SellAmountUsd string     `json:"sellAmountUsd,omitempty"`
	BuyAmountEth  string     `json:"buyAmountEth,omitempty"`
	SellAmountEth string     `json:"sellAmountEth,omitempty"`
	SellToken     Token      `json:"sellToken,omitempty"`
	BuyToken      Token      `json:"buyToken,omitempty"`
	Order         Order      `json:"order,omitempty"`
}

type Settlement struct {
	ID                  string  `json:"id"`
	TxHash              string  `json:"txHash,omitempty"`
	FirstTradeTimestamp int     `json:"firstTradeTimestamp,omitempty"`
	Trades              []Trade `json:"trades,omitempty"`
	Solver              User    `json:"solver,omitempty"`
}

type Settlements struct {
	Settlements []Settlement `json:"settlements"`
}

type Bundle struct {
	ID          string `json:"id"`
	EthPriceUSD string `json:"ethPriceUSD,omitempty"`
}

type Bundles struct {
	Bundles []Bundle `json:"bundles"`
}

type UniswapPools struct {
	UniswapPools []UniswapPool `json:"uniswapPools"`
}

type UniswapPool struct {
	ID                     string       `json:"id"`
	Liquidity              string       `json:"liquidity,omitempty"`
	Token0Price            string       `json:"token0Price,omitempty"`
	Token1Price            string       `json:"token1Price,omitempty"`
	Tick                   string       `json:"tick,omitempty"`
	TotalValueLockedToken0 string       `json:"totalValueLockedToken0,omitempty"`
	TotalValueLockedToken1 string       `json:"totalValueLockedToken1,omitempty"`
	Token0                 UniswapToken `json:"token0,omitempty"`
	Token1                 UniswapToken `json:"token1,omitempty"`
}

type UniswapToken struct {
	ID       string `json:"id"`
	Address  string `json:"address,omitempty"`
	Name     string `json:"name,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	Decimals int    `json:"decimals,omitempty"`
	PriceEth string `json:"priceEth,omitempty"`
	PriceUsd string `json:"priceUsd,omitempty"`
}

type UniswapTokens struct {
	UniswapTokens []UniswapToken `json:"uniswapTokens"`
}

type Totals struct {
	Totals []Total `json:"totals"`
}

type Total struct {
	ID             string `json:"id"`
	Tokens         string `json:"tokens,omitempty"`
	Traders        string `json:"traders,omitempty"`
	NumberOfTrades string `json:"numberOfTrades,omitempty"`
	Settlements    string `json:"settlements,omitempty"`
	VolumeUsd      string `json:"volumeUsd,omitempty"`
	VolumeEth      string `json:"volumeEth,omitempty"`
	FeesUsd        string `json:"feesUsd,omitempty"`
	FeesEth        string `json:"feesEth,omitempty"`
}

type DailyTotals struct {
	DailyTotals []DailyTotal `json:"dailyTotals"`
}

type DailyTotal struct {
	ID             string  `json:"id"`
	Timestamp      int     `json:"timestamp"`
	TotalTokens    string  `json:"totalTokens"`
	NumberOfTrades string  `json:"numberOfTrades"`
	Orders         string  `json:"orders"`
	Settlements    string  `json:"settlements"`
	VolumeUsd      string  `json:"volumeUsd"`
	VolumeEth      string  `json:"volumeEth"`
	FeesUsd        string  `json:"feesUsd"`
	FeesEth        string  `json:"feesEth"`
	Tokens         []Token `json:"tokens"`
}

type HourlyTotals struct {
	HourlyTotals []HourlyTotal `json:"hourlyTotals"`
}

type HourlyTotal struct {
	Timestamp      int     `json:"timestamp"`
	TotalTokens    string  `json:"totalTokens"`
	NumberOfTrades string  `json:"numberOfTrades"`
	Orders         string  `json:"orders"`
	Settlements    string  `json:"settlements"`
	VolumeUsd      string  `json:"volumeUsd"`
	VolumeEth      string  `json:"volumeEth"`
	FeesUsd        string  `json:"feesUsd"`
	FeesEth        string  `json:"feesEth"`
	Tokens         []Token `json:"tokens"`
}

type TokenDailyTotals struct {
	TokenDailyTotals []TokenDailyTotal `json:"tokenDailyTotals"`
}

type TokenDailyTotal struct {
	ID             string `json:"id"`
	Token          Token  `json:"token"`
	Timestamp      int    `json:"timestamp"`
	TotalVolume    string `json:"totalVolume"`
	TotalVolumeUsd string `json:"totalVolumeUsd"`
	TotalVolumeEth string `json:"totalVolumeEth"`
	TotalTrades    string `json:"totalTrades"`
	OpenPrice      string `json:"openPrice"`
	ClosePrice     string `json:"closePrice"`
	HigherPrice    string `json:"higherPrice"`
	LowerPrice     string `json:"lowerPrice"`
	AveragePrice   string `json:"averagePrice"`
}

type TokenHourlyTotals struct {
	TokenHourlyTotals []TokenHourlyTotal `json:"tokenHourlyTotals"`
}

type TokenHourlyTotal struct {
	ID             string `json:"id"`
	Token          Token  `json:"token"`
	Timestamp      int    `json:"timestamp"`
	TotalVolume    string `json:"totalVolume"`
	TotalVolumeUsd string `json:"totalVolumeUsd"`
	TotalVolumeEth string `json:"totalVolumeEth"`
	TotalTrades    string `json:"totalTrades"`
	OpenPrice      string `json:"openPrice"`
	ClosePrice     string `json:"closePrice"`
	HigherPrice    string `json:"higherPrice"`
	LowerPrice     string `json:"lowerPrice"`
	AveragePrice   string `json:"averagePrice"`
}

type TokenTradingEvents struct {
	TokenTradingEvents []TokenTradingEvent `json:"tokenTradingEvents"`
}

type TokenTradingEvent struct {
	ID        string `json:"id"`
	Token     Token  `json:"token"`
	Trade     Trade  `json:"trade"`
	Timestamp int    `json:"timestamp"`
	AmountEth string `json:"amountEth"`
	AmountUsd string `json:"amountUsd"`
}

type Pairs struct {
	Pairs []Pair `json:"pairs"`
}

type Pair struct {
	ID                  string `json:"id"`
	Token0              Token  `json:"token0"`
	Token1              Token  `json:"token1"`
	Token0Price         string `json:"token0Price"`
	Token1Price         string `json:"token1Price"`
	Token0RelativePrice string `json:"token0relativePrice"`
	Token1RelativePrice string `json:"token1relativePrice"`
	VolumeToken0        string `json:"volumeToken0"`
	VolumeToken1        string `json:"volumeToken1"`
	VolumeTradedEth     string `json:"volumeTradedEth"`
	VolumeTradedUsd     string `json:"volumeTradedUsd"`
}

type PairDailies struct {
	PairDailies []PairDaily `json:"pairDailies"`
}
type PairDaily struct {
	ID                  string `json:"id"`
	Token0              Token  `json:"token0"`
	Token1              Token  `json:"token1"`
	Token0Price         string `json:"token0Price"`
	Token1Price         string `json:"token1Price"`
	Token0RelativePrice string `json:"token0relativePrice"`
	Token1RelativePrice string `json:"token1relativePrice"`
	VolumeToken0        string `json:"volumeToken0"`
	VolumeToken1        string `json:"volumeToken1"`
	VolumeTradedEth     string `json:"volumeTradedEth"`
	VolumeTradedUsd     string `json:"volumeTradedUsd"`
}

type PairHourlies struct {
	PairHourlies []PairHourly `json:"pairHourlies"`
}

type PairHourly struct {
	ID                  string `json:"id"`
	Token0              Token  `json:"token0"`
	Token1              Token  `json:"token1"`
	Token0Price         string `json:"token0Price"`
	Token1Price         string `json:"token1Price"`
	Token0RelativePrice string `json:"token0relativePrice"`
	Token1RelativePrice string `json:"token1relativePrice"`
	VolumeToken0        string `json:"volumeToken0"`
	VolumeToken1        string `json:"volumeToken1"`
	VolumeTradedEth     string `json:"volumeTradedEth"`
	VolumeTradedUsd     string `json:"volumeTradedUsd"`
}
