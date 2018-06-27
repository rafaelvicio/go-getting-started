package pipelines

import "pipenator-backend/common"

var Routes = []common.Route{
	common.Route{
		Path:     "/pipelines",
		Method:   "GET",
		Handler:  ListPipelines,
		Security: true,
	},
	common.Route{
		Path:     "/pipelines",
		Method:   "POST",
		Handler:  CreatePipeline,
		Security: false,
	},
	common.Route{
		Path:     "/pipelines/{id}",
		Method:   "GET",
		Handler:  ShowPipeline,
		Security: false,
	},
	common.Route{
		Path:     "/pipelines/{id}",
		Method:   "DELETE",
		Handler:  DeletePipeline,
		Security: false,
	},
	common.Route{
		Path:     "/pipelines/{id}",
		Method:   "PUT",
		Handler:  UpdatePipeline,
		Security: false,
	},
}
