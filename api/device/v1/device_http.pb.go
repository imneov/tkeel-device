// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http 0.1.0

package v1

import (
	context "context"
	go_restful "github.com/emicklei/go-restful"
	errors "github.com/tkeel-io/kit/errors"
	result "github.com/tkeel-io/kit/result"
	protojson "google.golang.org/protobuf/encoding/protojson"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
)

import transportHTTP "github.com/tkeel-io/kit/transport/http"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the tkeel package it is being compiled against.
// import package.context.http.anypb.result.protojson.go_restful.errors.emptypb.

var (
	_ = protojson.MarshalOptions{}
	_ = anypb.Any{}
	_ = emptypb.Empty{}
)

type DeviceHTTPServer interface {
	AddDeviceExt(context.Context, *AddDeviceExtRequest) (*emptypb.Empty, error)
	CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceResponse, error)
	CreateDeviceDataRelation(context.Context, *CreateDeviceDataRelationRequest) (*emptypb.Empty, error)
	DeleteDevice(context.Context, *DeleteDeviceRequest) (*DeleteDeviceResponse, error)
	DeleteDeviceDataRelation(context.Context, *DeleteDeviceDataRelationRequest) (*emptypb.Empty, error)
	DeleteDeviceExt(context.Context, *DeleteDeviceExtRequest) (*emptypb.Empty, error)
	GetDevice(context.Context, *GetDeviceRequest) (*GetDeviceResponse, error)
	ListDeviceDataRelation(context.Context, *ListDeviceDataRelationRequest) (*emptypb.Empty, error)
	SaveDeviceConfAsOtherTemplte(context.Context, *SaveDeviceConfAsOtherTemplateRequest) (*CreateTemplateResponse, error)
	SaveDeviceConfAsSelfTemplte(context.Context, *SaveDeviceConfAsSelfTemplteRequest) (*emptypb.Empty, error)
	SearchEntity(context.Context, *ListDeviceRequest) (*ListDeviceResponse, error)
	SetDeviceAttribte(context.Context, *SetDeviceAttributeRequest) (*emptypb.Empty, error)
	SetDeviceCommand(context.Context, *SetDeviceCommandRequest) (*emptypb.Empty, error)
	SetDeviceRaw(context.Context, *SetDeviceRawRequest) (*emptypb.Empty, error)
	UpdateDevice(context.Context, *UpdateDeviceRequest) (*UpdateDeviceResponse, error)
	UpdateDeviceDataRelation(context.Context, *UpdateDeviceDataRelationRequest) (*emptypb.Empty, error)
	UpdateDeviceExt(context.Context, *UpdateDeviceExtRequest) (*emptypb.Empty, error)
}

type DeviceHTTPHandler struct {
	srv DeviceHTTPServer
}

func newDeviceHTTPHandler(s DeviceHTTPServer) *DeviceHTTPHandler {
	return &DeviceHTTPHandler{srv: s}
}

func (h *DeviceHTTPHandler) AddDeviceExt(req *go_restful.Request, resp *go_restful.Response) {
	in := AddDeviceExtRequest{}
	if err := transportHTTP.GetBody(req, &in.Ext); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.AddDeviceExt(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) CreateDevice(req *go_restful.Request, resp *go_restful.Response) {
	in := CreateDeviceRequest{}
	if err := transportHTTP.GetBody(req, &in.DevBasicInfo); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.CreateDevice(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) CreateDeviceDataRelation(req *go_restful.Request, resp *go_restful.Response) {
	in := CreateDeviceDataRelationRequest{}
	if err := transportHTTP.GetBody(req, &in.Relation); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.CreateDeviceDataRelation(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) DeleteDevice(req *go_restful.Request, resp *go_restful.Response) {
	in := DeleteDeviceRequest{}
	if err := transportHTTP.GetBody(req, &in.Ids); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.DeleteDevice(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) DeleteDeviceDataRelation(req *go_restful.Request, resp *go_restful.Response) {
	in := DeleteDeviceDataRelationRequest{}
	if err := transportHTTP.GetBody(req, &in.Relation); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.DeleteDeviceDataRelation(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) DeleteDeviceExt(req *go_restful.Request, resp *go_restful.Response) {
	in := DeleteDeviceExtRequest{}
	if err := transportHTTP.GetBody(req, &in.Keys); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.DeleteDeviceExt(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) GetDevice(req *go_restful.Request, resp *go_restful.Response) {
	in := GetDeviceRequest{}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.GetDevice(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) ListDeviceDataRelation(req *go_restful.Request, resp *go_restful.Response) {
	in := ListDeviceDataRelationRequest{}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.ListDeviceDataRelation(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) SaveDeviceConfAsOtherTemplte(req *go_restful.Request, resp *go_restful.Response) {
	in := SaveDeviceConfAsOtherTemplateRequest{}
	if err := transportHTTP.GetBody(req, &in.OtherTemplateInfo); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.SaveDeviceConfAsOtherTemplte(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) SaveDeviceConfAsSelfTemplte(req *go_restful.Request, resp *go_restful.Response) {
	in := SaveDeviceConfAsSelfTemplteRequest{}
	if err := transportHTTP.GetBody(req, &in.Id); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.SaveDeviceConfAsSelfTemplte(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) SearchEntity(req *go_restful.Request, resp *go_restful.Response) {
	in := ListDeviceRequest{}
	if err := transportHTTP.GetBody(req, &in.ListEntityQuery); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.SearchEntity(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) SetDeviceAttribte(req *go_restful.Request, resp *go_restful.Response) {
	in := SetDeviceAttributeRequest{}
	if err := transportHTTP.GetBody(req, &in.Content); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.SetDeviceAttribte(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) SetDeviceCommand(req *go_restful.Request, resp *go_restful.Response) {
	in := SetDeviceCommandRequest{}
	if err := transportHTTP.GetBody(req, &in.Content); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.SetDeviceCommand(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) SetDeviceRaw(req *go_restful.Request, resp *go_restful.Response) {
	in := SetDeviceRawRequest{}
	if err := transportHTTP.GetBody(req, &in.Value); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.SetDeviceRaw(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) UpdateDevice(req *go_restful.Request, resp *go_restful.Response) {
	in := UpdateDeviceRequest{}
	if err := transportHTTP.GetBody(req, &in.DevBasicInfo); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.UpdateDevice(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) UpdateDeviceDataRelation(req *go_restful.Request, resp *go_restful.Response) {
	in := UpdateDeviceDataRelationRequest{}
	if err := transportHTTP.GetBody(req, &in.Relation); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.UpdateDeviceDataRelation(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *DeviceHTTPHandler) UpdateDeviceExt(req *go_restful.Request, resp *go_restful.Response) {
	in := UpdateDeviceExtRequest{}
	if err := transportHTTP.GetBody(req, &in.Ext); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.UpdateDeviceExt(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func RegisterDeviceHTTPServer(container *go_restful.Container, srv DeviceHTTPServer) {
	var ws *go_restful.WebService
	for _, v := range container.RegisteredWebServices() {
		if v.RootPath() == "/v1" {
			ws = v
			break
		}
	}
	if ws == nil {
		ws = new(go_restful.WebService)
		ws.ApiVersion("/v1")
		ws.Path("/v1").Produces(go_restful.MIME_JSON)
		container.Add(ws)
	}

	handler := newDeviceHTTPHandler(srv)
	ws.Route(ws.POST("/devices").
		To(handler.CreateDevice))
	ws.Route(ws.PUT("/devices/{id}").
		To(handler.UpdateDevice))
	ws.Route(ws.POST("/devices/delete").
		To(handler.DeleteDevice))
	ws.Route(ws.GET("/devices/{id}").
		To(handler.GetDevice))
	ws.Route(ws.POST("/search").
		To(handler.SearchEntity))
	ws.Route(ws.POST("/devices/{id}/ext").
		To(handler.AddDeviceExt))
	ws.Route(ws.POST("/devices/{id}/ext/delete").
		To(handler.DeleteDeviceExt))
	ws.Route(ws.PUT("/devices/{id}/ext").
		To(handler.UpdateDeviceExt))
	ws.Route(ws.POST("/devices/{id}/relations").
		To(handler.CreateDeviceDataRelation))
	ws.Route(ws.PUT("/devcies/{id}/relations").
		To(handler.UpdateDeviceDataRelation))
	ws.Route(ws.POST("/devices/{id}/relations/delete").
		To(handler.DeleteDeviceDataRelation))
	ws.Route(ws.GET("/devices/{id}/relations").
		To(handler.ListDeviceDataRelation))
	ws.Route(ws.POST("/devices/{id}/raw/set").
		To(handler.SetDeviceRaw))
	ws.Route(ws.POST("/devices/{id}/attribute/set").
		To(handler.SetDeviceAttribte))
	ws.Route(ws.POST("/devices/{id}/command/set").
		To(handler.SetDeviceCommand))
	ws.Route(ws.POST("/devices/{id}/configs/saveAsSelfTemplate").
		To(handler.SaveDeviceConfAsSelfTemplte))
	ws.Route(ws.POST("/devices/{id}/configs/saveAsOtherTemplate").
		To(handler.SaveDeviceConfAsOtherTemplte))
}
