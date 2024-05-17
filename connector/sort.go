package connector

import (
	"github.com/hasura/ndc-elasticsearch/types"
	"github.com/hasura/ndc-sdk-go/schema"
)

func prepareSortQuery(orderBy *schema.OrderBy, state *types.State) ([]map[string]interface{}, error) {
	elements := orderBy.Elements
	sort := make([]map[string]interface{}, len(elements))
	for i, element := range elements {
		target := element.Target["name"].(string)
		// check if the target field is orderable or not
		if _, ok := state.UnsupportedSortFields[target]; ok {
			return nil, schema.BadRequestError("sorting not supported on this field", map[string]interface{}{
				"value": element.Target["name"].(string),
			})
		}
		sort[i] = make(map[string]interface{})
		name := element.Target["name"].(string)
		sort[i][name] = map[string]interface{}{"order": element.OrderDirection}
	}
	return sort, nil
}
