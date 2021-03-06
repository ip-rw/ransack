// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	empty "github.com/golang/protobuf/ptypes/empty"
	ent "github.com/ip-rw/ransack/ent"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MachineService implements MachineServiceServer
type MachineService struct {
	client *ent.Client
	UnimplementedMachineServiceServer
}

// NewMachineService returns a new MachineService
func NewMachineService(client *ent.Client) *MachineService {
	return &MachineService{
		client: client,
	}
}

// toProtoMachine transforms the ent type to the pb type
func toProtoMachine(e *ent.Machine) *Machine {
	return &Machine{
		Fingerprint: string(e.Fingerprint),
		Hostname:    string(e.Hostname),
		Hwid:        string(e.Hwid),
		Id:          int32(e.ID),
	}
}

// Create implements MachineServiceServer.Create
func (svc *MachineService) Create(ctx context.Context, req *CreateMachineRequest) (*Machine, error) {
	machine := req.GetMachine()
	res, err := svc.client.Machine.Create().
		SetFingerprint(string(machine.GetFingerprint())).
		SetHostname(string(machine.GetHostname())).
		SetHwid(string(machine.GetHwid())).
		Save(ctx)

	switch {
	case err == nil:
		return toProtoMachine(res), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal: %s", err)
	}
}

// Get implements MachineServiceServer.Get
func (svc *MachineService) Get(ctx context.Context, req *GetMachineRequest) (*Machine, error) {
	get, err := svc.client.Machine.Get(ctx, int(req.GetId()))
	switch {
	case err == nil:
		return toProtoMachine(get), nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
}

// Update implements MachineServiceServer.Update
func (svc *MachineService) Update(ctx context.Context, req *UpdateMachineRequest) (*Machine, error) {
	machine := req.GetMachine()
	res, err := svc.client.Machine.UpdateOneID(int(machine.GetId())).
		SetFingerprint(string(machine.GetFingerprint())).
		SetHostname(string(machine.GetHostname())).
		SetHwid(string(machine.GetHwid())).
		Save(ctx)

	switch {
	case err == nil:
		return toProtoMachine(res), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal: %s", err)
	}
}

// Delete implements MachineServiceServer.Delete
func (svc *MachineService) Delete(ctx context.Context, req *DeleteMachineRequest) (*empty.Empty, error) {
	err := svc.client.Machine.DeleteOneID(int(req.GetId())).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
}
