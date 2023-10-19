package services

import (
	"context"

	"github.com/cosmos/gogoproto/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
)

// ReflectionService implements the cosmos.reflection.v1 service.
type ReflectionService struct {
	reflectionv1.UnimplementedReflectionServiceServer
	files *descriptorpb.FileDescriptorSet
}

func NewReflectionService() (*ReflectionService, error) {
	fds, err := proto.MergedGlobalFileDescriptors()
	if err != nil {
		return nil, err
	}

	// load any protoregistry file descriptors not in gogo
	protoregistry.GlobalFiles.RangeFiles(func(fileDescriptor protoreflect.FileDescriptor) bool {
		if !haveFileDescriptor[fileDescriptor.Path()] {
			fds.File = append(fds.File, protodesc.ToFileDescriptorProto(fileDescriptor))
		}
		return true
	})

	slices.SortFunc(fds.File, func(x, y *descriptorpb.FileDescriptorProto) int {
		if *x.Name < *y.Name {
			return -1
		}
		if *x.Name > *y.Name {
			return 1
		}
		return 0
	})
	return &ReflectionService{files: fds}, nil
}

func (r ReflectionService) FileDescriptors(_ context.Context, _ *reflectionv1.FileDescriptorsRequest) (*reflectionv1.FileDescriptorsResponse, error) {
	return &reflectionv1.FileDescriptorsResponse{
		Files: r.files.File,
	}, nil
}

var _ reflectionv1.ReflectionServiceServer = &ReflectionService{}
