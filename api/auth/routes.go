package auth

import "pipenator-backend/common"

var Routes = []common.Route{
	common.Route{
		Path:     "/auth/autenticar",
		Method:   "GET",
		Handler:  autenticar,
		Security: false,
	},
}
