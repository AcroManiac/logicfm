package swagger

import (
	"github.com/ahamtat/logicfm/internal/adapters/swagger/builder/restapi/operations"
	"github.com/go-openapi/loads"

	"github.com/ahamtat/logicfm/internal/adapters/swagger/builder/restapi"
	"go.uber.org/zap"
)

// NewBuilder constructs server for swagger Builder API specification.
func NewBuilder(logger *zap.SugaredLogger) (*restapi.Server, error) {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		logger.Fatal("failed embedding swagger specification", "error", err, "caller", "NewBuilder")
	}

	api := operations.NewBuilderapiAPI(swaggerSpec)
	server := restapi.NewServer(api)

	server.ConfigureAPI()

	return server, nil
}
