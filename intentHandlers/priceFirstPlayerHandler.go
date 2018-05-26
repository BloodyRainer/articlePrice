package intentHandlers

import "github.com/BloodyRainer/articlePrice/df"

func SavePriceFristPlayerAskSecondPlayer(dfReq df.Request) (*df.Response, error) {

	gs, err := df.MakeGameSessionFromDfRequest(dfReq)
	if err != nil {
		return nil, err
	}

	

}
