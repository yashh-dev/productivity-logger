// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.19.1
// source: tracker/v1/tracker.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTrackerCreateBlockFunc = "/helloworld.v1.tracker.Tracker/CreateBlockFunc"
const OperationTrackerDeletBlockFunc = "/helloworld.v1.tracker.Tracker/DeletBlockFunc"
const OperationTrackerListBlockFunc = "/helloworld.v1.tracker.Tracker/ListBlockFunc"
const OperationTrackerUpdateBlockFunc = "/helloworld.v1.tracker.Tracker/UpdateBlockFunc"

type TrackerHTTPServer interface {
	CreateBlockFunc(context.Context, *CreateBlock) (*BlockResp, error)
	DeletBlockFunc(context.Context, *DeletBlock) (*BlockResp, error)
	ListBlockFunc(context.Context, *GetListReq) (*BlockResp, error)
	UpdateBlockFunc(context.Context, *UpdateBlock) (*BlockResp, error)
}

func RegisterTrackerHTTPServer(s *http.Server, srv TrackerHTTPServer) {
	r := s.Route("/")
	r.POST("/createblock", _Tracker_CreateBlockFunc0_HTTP_Handler(srv))
	r.DELETE("/deleteblock/{UserId}/{BlockId}", _Tracker_DeletBlockFunc0_HTTP_Handler(srv))
	r.PUT("/updateblock", _Tracker_UpdateBlockFunc0_HTTP_Handler(srv))
	r.GET("/getblocks/{UserId}", _Tracker_ListBlockFunc0_HTTP_Handler(srv))
}

func _Tracker_CreateBlockFunc0_HTTP_Handler(srv TrackerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateBlock
		if err := ctx.Bind(&in.CreateBlockRequestBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTrackerCreateBlockFunc)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateBlockFunc(ctx, req.(*CreateBlock))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BlockResp)
		return ctx.Result(200, reply)
	}
}

func _Tracker_DeletBlockFunc0_HTTP_Handler(srv TrackerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeletBlock
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTrackerDeletBlockFunc)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletBlockFunc(ctx, req.(*DeletBlock))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BlockResp)
		return ctx.Result(200, reply)
	}
}

func _Tracker_UpdateBlockFunc0_HTTP_Handler(srv TrackerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateBlock
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTrackerUpdateBlockFunc)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateBlockFunc(ctx, req.(*UpdateBlock))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BlockResp)
		return ctx.Result(200, reply)
	}
}

func _Tracker_ListBlockFunc0_HTTP_Handler(srv TrackerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTrackerListBlockFunc)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListBlockFunc(ctx, req.(*GetListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BlockResp)
		return ctx.Result(200, reply)
	}
}

type TrackerHTTPClient interface {
	CreateBlockFunc(ctx context.Context, req *CreateBlock, opts ...http.CallOption) (rsp *BlockResp, err error)
	DeletBlockFunc(ctx context.Context, req *DeletBlock, opts ...http.CallOption) (rsp *BlockResp, err error)
	ListBlockFunc(ctx context.Context, req *GetListReq, opts ...http.CallOption) (rsp *BlockResp, err error)
	UpdateBlockFunc(ctx context.Context, req *UpdateBlock, opts ...http.CallOption) (rsp *BlockResp, err error)
}

type TrackerHTTPClientImpl struct {
	cc *http.Client
}

func NewTrackerHTTPClient(client *http.Client) TrackerHTTPClient {
	return &TrackerHTTPClientImpl{client}
}

func (c *TrackerHTTPClientImpl) CreateBlockFunc(ctx context.Context, in *CreateBlock, opts ...http.CallOption) (*BlockResp, error) {
	var out BlockResp
	pattern := "/createblock"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTrackerCreateBlockFunc))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.CreateBlockRequestBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TrackerHTTPClientImpl) DeletBlockFunc(ctx context.Context, in *DeletBlock, opts ...http.CallOption) (*BlockResp, error) {
	var out BlockResp
	pattern := "/deleteblock/{UserId}/{BlockId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTrackerDeletBlockFunc))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TrackerHTTPClientImpl) ListBlockFunc(ctx context.Context, in *GetListReq, opts ...http.CallOption) (*BlockResp, error) {
	var out BlockResp
	pattern := "/getblocks/{UserId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTrackerListBlockFunc))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TrackerHTTPClientImpl) UpdateBlockFunc(ctx context.Context, in *UpdateBlock, opts ...http.CallOption) (*BlockResp, error) {
	var out BlockResp
	pattern := "/updateblock"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTrackerUpdateBlockFunc))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
