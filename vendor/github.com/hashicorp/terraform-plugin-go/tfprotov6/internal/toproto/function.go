// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func CallFunction_Response(in *tfprotov6.CallFunctionResponse) (*tfplugin6.CallFunction_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.CallFunction_Response{
		Diagnostics: diags,
	}

	if in.Result != nil {
		resp.Result = DynamicValue(in.Result)
	}

	return resp, nil
}

func Function(in *tfprotov6.Function) (*tfplugin6.Function, error) {
	if in == nil {
		return nil, nil
	}

	resp := &tfplugin6.Function{
		Description:        in.Description,
		DescriptionKind:    StringKind(in.DescriptionKind),
		DeprecationMessage: in.DeprecationMessage,
		Parameters:         make([]*tfplugin6.Function_Parameter, 0, len(in.Parameters)),
		Summary:            in.Summary,
	}

	for position, parameter := range in.Parameters {
		if parameter == nil {
			return nil, fmt.Errorf("missing function parameter definition at position: %d", position)
		}

		functionParameter, err := Function_Parameter(parameter)

		if err != nil {
			return nil, fmt.Errorf("unable to marshal function parameter at position %d: %w", position, err)
		}

		resp.Parameters = append(resp.Parameters, functionParameter)
	}

	if in.Return == nil {
		return nil, fmt.Errorf("missing function return definition")
	}

	functionReturn, err := Function_Return(in.Return)

	if err != nil {
		return nil, fmt.Errorf("unable to marshal function return: %w", err)
	}

	resp.Return = functionReturn

	if in.VariadicParameter != nil {
		variadicParameter, err := Function_Parameter(in.VariadicParameter)

		if err != nil {
			return nil, fmt.Errorf("unable to marshal variadic function parameter: %w", err)
		}

		resp.VariadicParameter = variadicParameter
	}

	return resp, nil
}

func Function_Parameter(in *tfprotov6.FunctionParameter) (*tfplugin6.Function_Parameter, error) {
	if in == nil {
		return nil, nil
	}

	resp := &tfplugin6.Function_Parameter{
		AllowNullValue:     in.AllowNullValue,
		AllowUnknownValues: in.AllowUnknownValues,
		Description:        in.Description,
		DescriptionKind:    StringKind(in.DescriptionKind),
		Name:               in.Name,
	}

	if in.Type == nil {
		return nil, fmt.Errorf("missing function parameter type definition")
	}

	ctyType, err := CtyType(in.Type)

	if err != nil {
		return resp, fmt.Errorf("error marshaling function parameter type: %w", err)
	}

	resp.Type = ctyType

	return resp, nil
}

func Function_Return(in *tfprotov6.FunctionReturn) (*tfplugin6.Function_Return, error) {
	if in == nil {
		return nil, nil
	}

	resp := &tfplugin6.Function_Return{}

	if in.Type == nil {
		return nil, fmt.Errorf("missing function return type definition")
	}

	ctyType, err := CtyType(in.Type)

	if err != nil {
		return resp, fmt.Errorf("error marshaling function return type: %w", err)
	}

	resp.Type = ctyType

	return resp, nil
}

func GetFunctions_Response(in *tfprotov6.GetFunctionsResponse) (*tfplugin6.GetFunctions_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.GetFunctions_Response{
		Diagnostics: diags,
		Functions:   make(map[string]*tfplugin6.Function, len(in.Functions)),
	}

	for name, functionPtr := range in.Functions {
		function, err := Function(functionPtr)

		if err != nil {
			return nil, fmt.Errorf("error marshaling function definition for %q: %w", name, err)
		}

		resp.Functions[name] = function
	}

	return resp, nil
}

func GetMetadata_FunctionMetadata(in *tfprotov6.FunctionMetadata) *tfplugin6.GetMetadata_FunctionMetadata {
	if in == nil {
		return nil
	}

	return &tfplugin6.GetMetadata_FunctionMetadata{
		Name: in.Name,
	}
}
