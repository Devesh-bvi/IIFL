package class

//LoginVendorResponse : Response of LoginRequestMobileForVendor
type LoginVendorResponse struct {
	Body struct {
		EmailID     string `json:"EmailId"`
		Message     string `json:"Message"`
		Status      int    `json:"Status"`
		TTManagerID string `json:"TTManagerId"`
		UserName    string `json:"UserName"`
		ValidUpto   string `json:"ValidUpto"`
		VendorName  string `json:"VendorName"`
	} `json:"body"`
	Head struct {
		ResponseCode      string `json:"responseCode"`
		Status            string `json:"status"`
		StatusDescription string `json:"statusDescription"`
	} `json:"head"`
}

//LoginClientResponse : Response of LoginRequest
type LoginClientResponse struct {
	Body struct {
		BulkOrderAllowed      int    `json:"BulkOrderAllowed"`
		CleareDt              string `json:"CleareDt"`
		ClientName            string `json:"ClientName"`
		ClientType            int    `json:"ClientType"`
		DPID                  string `json:"DPID"`
		EmailID               string `json:"EmailId"`
		InteractiveLocalIP    string `json:"InteractiveLocalIP"`
		InteractivePort       int    `json:"InteractivePort"`
		InteractivePublicIP   string `json:"InteractivePublicIP"`
		IsExternal            string `json:"IsExternal"`
		IsIDBound             int    `json:"IsIDBound"`
		IsIDBound2            int    `json:"IsIDBound2"`
		IsPLM                 int    `json:"IsPLM"`
		IsPLMDefined          int    `json:"IsPLMDefined"`
		LastAccessedTime      string `json:"LastAccessedTime"`
		LastLogin             string `json:"LastLogin"`
		LastPasswordModify    string `json:"LastPasswordModify"`
		Msg                   string `json:"Msg"`
		OTPCredentialID       string `json:"OTPCredentialID"`
		PLMsAllowed           int    `json:"PLMsAllowed"`
		POAStatus             string `json:"POAStatus"`
		PasswordChangeFlag    int    `json:"PasswordChangeFlag"`
		PasswordChangeMessage string `json:"PasswordChangeMessage"`
		RunningAuthorization  int    `json:"RunningAuthorization"`
		ServerDt              string `json:"ServerDt"`
		Success               int    `json:"Success"`
		TCPBCastPort          int    `json:"TCPBCastPort"`
		TCPBcastLocalIP       string `json:"TCPBcastLocalIP"`
		TCPBcastPublicIP      string `json:"TCPBcastPublicIP"`
		UDPBCastPort          int    `json:"UDPBCastPort"`
		UDPBcastIP            string `json:"UDPBcastIP"`
		VersionChanged        int    `json:"VersionChanged"`
	} `json:"body"`
	Head struct {
		ResponseCode      string `json:"responseCode"`
		Status            string `json:"status"`
		StatusDescription string `json:"statusDescription"`
	} `json:"head"`
}

//HoldResponse : Response of Holding
type HoldResponse struct {
	Body struct {
		CacheTime int `json:"CacheTime"`
		Data      []struct {
			BseCode        int     `json:"BseCode"`
			CurrentPL      int     `json:"CurrentPL"`
			CurrentPrice   float64 `json:"CurrentPrice"`
			CurrentValue   int     `json:"CurrentValue"`
			Exch           string  `json:"Exch"`
			ExchType       string  `json:"ExchType"`
			FullName       string  `json:"FullName"`
			FundedQty      int     `json:"FundedQty"`
			NseCode        int     `json:"NseCode"`
			PerChange      float64 `json:"PerChange"`
			PreviousClose  float64 `json:"PreviousClose"`
			Quantity       int     `json:"Quantity"`
			Symbol         string  `json:"Symbol"`
			YesterDayValue int     `json:"YesterDayValue"`
			CollateralQty  int     `json:"collateralQty"`
		} `json:"Data"`
		Message string `json:"Message"`
		Status  int    `json:"Status"`
	} `json:"body"`
	Head struct {
		ResponseCode      string `json:"responseCode"`
		Status            string `json:"status"`
		StatusDescription string `json:"statusDescription"`
	} `json:"head"`
}

//MarketFeedResponse : Response of MarketFeed
type MarketFeedResponse struct {
	Body struct {
		CacheTime int `json:"CacheTime"`
		Data      []struct {
			Exch     string  `json:"Exch"`
			ExchType string  `json:"ExchType"`
			High     float64 `json:"High"`
			LastRate float64 `json:"LastRate"`
			Low      float64 `json:"Low"`
			Message  string  `json:"Message"`
			PClose   float64 `json:"PClose"`
			Status   int     `json:"Status"`
			TickDt   string  `json:"TickDt"`
			Time     int     `json:"Time"`
			Token    int     `json:"Token"`
			TotalQty int     `json:"TotalQty"`
		} `json:"Data"`
		Message   string `json:"Message"`
		Status    int    `json:"Status"`
		TimeStamp string `json:"TimeStamp"`
	} `json:"body"`
	Head struct {
		ResponseCode      string `json:"responseCode"`
		Status            string `json:"status"`
		StatusDescription string `json:"statusDescription"`
	} `json:"head"`
}

//OrderResponse : Reponse of OrderRequest
type OrderResponse struct {
	Body struct {
		BrokerOrderID   int    `json:"BrokerOrderID"`
		ClientCode      string `json:"ClientCode"`
		Exch            string `json:"Exch"`
		ExchOrderID     string `json:"ExchOrderID"`
		ExchType        string `json:"ExchType"`
		LocalOrderID    int    `json:"LocalOrderID"`
		Message         string `json:"Message"`
		RMSResponseCode int    `json:"RMSResponseCode"`
		ScripCode       int    `json:"ScripCode"`
		Status          int    `json:"Status"`
		Time            string `json:"Time"`
	} `json:"body"`
	Head struct {
		ResponseCode      string `json:"responseCode"`
		Status            string `json:"status"`
		StatusDescription string `json:"statusDescription"`
	} `json:"head"`
}

//OrderStatusResponse : Resposne OrderStatus
type OrderStatusResponse struct {
	Head struct {
		ResponseCode      string `json:"responseCode"`
		Status            string `json:"status"`
		StatusDescription string `json:"statusDescription"`
	} `json:"head"`
	Body struct {
		Message         string `json:"Message"`
		OrdStatusResLst []struct {
			Exch          string `json:"Exch"`
			ExchOrderID   int64  `json:"ExchOrderID"`
			ExchOrderTime string `json:"ExchOrderTime"`
			ExchType      string `json:"ExchType"`
			OrderQty      int    `json:"OrderQty"`
			OrderRate     int    `json:"OrderRate"`
			PendingQty    int    `json:"PendingQty"`
			ScripCode     int    `json:"ScripCode"`
			Status        string `json:"Status"`
			Symbol        string `json:"Symbol"`
			TradedQty     int    `json:"TradedQty"`
		} `json:"OrdStatusResLst"`
		Status int `json:"Status"`
	} `json:"body"`
}

//TradeInformationResponse : Response for TradeInformation
type TradeInformationResponse struct {
	Head struct {
		ResponseCode      string `json:"responseCode"`
		Status            string `json:"status"`
		StatusDescription string `json:"statusDescription"`
	} `json:"head"`
	Body struct {
		Message     string `json:"Message"`
		Status      int    `json:"Status"`
		TradeDetail []struct {
			BuySell           string  `json:"BuySell"`
			Exch              string  `json:"Exch"`
			ExchOrderID       int64   `json:"ExchOrderID"`
			ExchType          string  `json:"ExchType"`
			ExchangeTradeID   int     `json:"ExchangeTradeID"`
			ExchangeTradeTime string  `json:"ExchangeTradeTime"`
			OrgQty            int     `json:"OrgQty"`
			PendingQty        int     `json:"PendingQty"`
			Qty               int     `json:"Qty"`
			Rate              float64 `json:"Rate"`
			ScripCode         int     `json:"ScripCode"`
			ScripName         string  `json:"ScripName"`
		} `json:"TradeDetail"`
	} `json:"body"`
}
