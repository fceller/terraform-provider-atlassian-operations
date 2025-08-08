package dataModels

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ServiceModel struct {
	ID              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	Description     types.String `tfsdk:"description"`
	Tier            types.Int32  `tfsdk:"tier"`
	Type            types.String `tfsdk:"type"`
	Owner           types.String `tfsdk:"owner"`
	ChangeApprovers types.Object `tfsdk:"change_approvers"`
	Responders      types.Object `tfsdk:"responders"`
	Stakeholders    types.Object `tfsdk:"stakeholders"`
	Projects        types.Object `tfsdk:"projects"`
}

type ChangeApproversModel struct {
	Groups types.List `tfsdk:"groups"`
}

type RespondersModel struct {
	Users types.List `tfsdk:"users"`
	Teams types.List `tfsdk:"teams"`
}

type StakeholdersModel struct {
	Users types.List `tfsdk:"users"`
}

type ProjectsModel struct {
	IDs types.List `tfsdk:"ids"`
}

var ChangeApproversModelMap = map[string]attr.Type{
	"groups": types.ListType{ElemType: types.StringType},
}

var RespondersModelMap = map[string]attr.Type{
	"users": types.ListType{ElemType: types.StringType},
	"teams": types.ListType{ElemType: types.StringType},
}

var StakeholdersModelMap = map[string]attr.Type{
	"users": types.ListType{ElemType: types.StringType},
}

var ProjectsModelMap = map[string]attr.Type{
	"ids": types.ListType{ElemType: types.StringType},
}

var ServiceModelMap = map[string]attr.Type{
	"id":               types.StringType,
	"name":             types.StringType,
	"description":      types.StringType,
	"tier":             types.Int32Type,
	"type":             types.StringType,
	"owner":            types.StringType,
	"change_approvers": types.ObjectType{AttrTypes: ChangeApproversModelMap},
	"responders":       types.ObjectType{AttrTypes: RespondersModelMap},
	"stakeholders":     types.ObjectType{AttrTypes: StakeholdersModelMap},
	"projects":         types.ObjectType{AttrTypes: ProjectsModelMap},
}

func (m *ServiceModel) AsValue() types.Object {
	return types.ObjectValueMust(ServiceModelMap, map[string]attr.Value{
		"id":               m.ID,
		"name":             m.Name,
		"description":      m.Description,
		"tier":             m.Tier,
		"type":             m.Type,
		"owner":            m.Owner,
		"change_approvers": m.ChangeApprovers,
		"responders":       m.Responders,
		"stakeholders":     m.Stakeholders,
		"projects":         m.Projects,
	})
}
