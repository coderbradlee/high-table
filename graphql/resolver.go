// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package graphql

import (
	"context"

	"github.com/iotexproject/high-table/api"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver is hte resolver that handles GraphQL request
type Resolver struct {
	cli api.Protocol
}

// Query returns a query resolver
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

// Delegate handles delegate requests
func (r *queryResolver) Delegate(ctx context.Context, epochNum int, groupID int) ([]*Delegate, error) {
	return r.getDelegates()
}

func (r *queryResolver) getDelegates() ([]*Delegate, error) {
	//requestedFields := graphql.CollectAllFields(ctx)
	//argsMap := parseFieldArguments(ctx, "byContractAddress", "xrc20")
	//address, err := getStringArg(argsMap, "address")
	//if err != nil {
	//	return errors.Wrap(err, "failed to get address")
	//}
	//numPerPage, err := getIntArg(argsMap, "numPerPage")
	//if err != nil {
	//	return errors.Wrap(err, "failed to get numPerPage")
	//}
	//page, err := getIntArg(argsMap, "page")
	//if err != nil {
	//	return errors.Wrap(err, "failed to get page")
	//}
	//output := &Xrc20List{Exist: false}
	//actionResponse.ByContractAddress = output
	//xrc20InfoList, err := r.AP.GetXrc20(address, uint64(numPerPage), uint64(page))
	//switch {
	//case errors.Cause(err) == indexprotocol.ErrNotExist:
	//	return nil
	//case err != nil:
	//	return errors.Wrap(err, "failed to get contract information")
	//}
	//output.Exist = true
	//output.Count = len(xrc20InfoList)
	//output.Xrc20 = make([]*Xrc20Info, 0, len(xrc20InfoList))
	//for _, c := range xrc20InfoList {
	//	output.Xrc20 = append(output.Xrc20, &Xrc20Info{
	//		Hash:      c.Hash,
	//		Timestamp: c.Timestamp,
	//		From:      c.From,
	//		To:        c.To,
	//		Quantity:  c.Quantity,
	//		Contract:  c.Contract,
	//	})
	//}
	return nil, nil
}

//func containField(requestedFields []string, field string) bool {
//	for _, f := range requestedFields {
//		if f == field {
//			return true
//		}
//	}
//	return false
//}
//
//func parseFieldArguments(ctx context.Context, fieldName string, selectedFieldName string) map[string]*ast.Value {
//	fields := graphql.CollectFieldsCtx(ctx, nil)
//	var field graphql.CollectedField
//	for _, f := range fields {
//		if f.Name == fieldName {
//			field = f
//		}
//	}
//	arguments := field.Arguments
//	if selectedFieldName != "" {
//		fields = graphql.CollectFields(ctx, field.Selections, nil)
//		for _, f := range fields {
//			if f.Name == selectedFieldName {
//				field = f
//			}
//		}
//		arguments = append(arguments, field.Arguments...)
//	}
//	argsMap := make(map[string]*ast.Value)
//	for _, arg := range arguments {
//		argsMap[arg.Name] = arg.Value
//	}
//	parseVariables(ctx, argsMap, arguments)
//	return argsMap
//}
//func parseVariables(ctx context.Context, argsMap map[string]*ast.Value, arguments ast.ArgumentList) {
//	val := graphql.GetRequestContext(ctx)
//	if val != nil {
//		for _, arg := range arguments {
//			if arg == nil {
//				continue
//			}
//			switch arg.Value.ExpectedType.Name() {
//			case "String":
//				value, ok := val.Variables[arg.Name].(string)
//				if ok {
//					argsMap[arg.Name].Raw = value
//				}
//			case "Int":
//				valueJSON, ok := val.Variables[arg.Name].(json.Number)
//				if ok {
//					value, err := valueJSON.Int64()
//					if err != nil {
//						return
//					}
//					argsMap[arg.Name].Raw = fmt.Sprintf("%d", value)
//				}
//			case "Pagination":
//				value, ok := val.Variables[arg.Name].(map[string]interface{})
//				if ok {
//					for k, v := range value {
//						valueJSON, ok := v.(json.Number)
//						if ok {
//							valueInt64, err := valueJSON.Int64()
//							if err != nil {
//								continue
//							}
//							child := &ast.ChildValue{Name: k, Value: &ast.Value{Raw: fmt.Sprintf("%d", valueInt64)}}
//							argsMap[arg.Name].Children = append(argsMap[arg.Name].Children, child)
//						}
//					}
//				}
//			default:
//				return
//			}
//		}
//	}
//}
//func getIntArg(argsMap map[string]*ast.Value, argName string) (int, error) {
//	getStr, err := getStringArg(argsMap, argName)
//	if err != nil {
//		return 0, err
//	}
//	intVal, err := strconv.Atoi(getStr)
//	if err != nil {
//		return 0, fmt.Errorf("%s must be an integer", argName)
//	}
//	return intVal, nil
//}
//
//func getStringArg(argsMap map[string]*ast.Value, argName string) (string, error) {
//	val, ok := argsMap[argName]
//	if !ok {
//		return "", fmt.Errorf("%s is required", argName)
//	}
//	return string(val.Raw), nil
//}
//
//func getBoolArg(argsMap map[string]*ast.Value, argName string) (bool, error) {
//	getStr, err := getStringArg(argsMap, argName)
//	if err != nil {
//		return false, err
//	}
//	boolVal, err := strconv.ParseBool(getStr)
//	if err != nil {
//		return false, fmt.Errorf("%s must be a boolean value", argName)
//	}
//	return boolVal, nil
//}
//
//func getPaginationArgs(argsMap map[string]*ast.Value) (map[string]int, error) {
//	pagination, ok := argsMap["pagination"]
//	if !ok {
//		return nil, ErrPaginationNotFound
//	}
//	childValueList := pagination.Children
//	paginationMap := make(map[string]int)
//	for _, childValue := range childValueList {
//		intVal, err := strconv.Atoi(childValue.Value.Raw)
//		if err != nil {
//			return nil, errors.Wrap(err, "pagination value must be an integer")
//		}
//		paginationMap[childValue.Name] = intVal
//	}
//	return paginationMap, nil
//}
//
//func ethAddrToIoAddr(ethAddr string) (string, error) {
//	ethAddress := common.HexToAddress(ethAddr)
//	ioAddress, err := address.FromBytes(ethAddress.Bytes())
//	if err != nil {
//		return "", err
//	}
//	return ioAddress.String(), nil
//}
