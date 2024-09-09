package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"university-management/proto"

	"google.golang.org/grpc"
)

type universityServer struct {
	proto.UnimplementedUniversityServiceServer
	students map[string]*proto.Student
	courses  map[string]*proto.Course
}

func newServer() *universityServer {
	return &universityServer{
		students: make(map[string]*proto.Student),
		courses:  make(map[string]*proto.Course),
	}
}

// Implement CreateStudent method
func (s *universityServer) CreateStudent(ctx context.Context, req *proto.CreateStudentRequest) (*proto.CreateStudentResponse, error) {
	student := &proto.Student{
		Id:    fmt.Sprintf("student-%d", len(s.students)+1),
		Name:  req.Name,
		Email: req.Email,
	}
	s.students[student.Id] = student
	return &proto.CreateStudentResponse{Student: student}, nil
}

// Implement GetStudent method
func (s *universityServer) GetStudent(ctx context.Context, req *proto.GetStudentRequest) (*proto.GetStudentResponse, error) {
	student, exists := s.students[req.Id]
	if !exists {
		return nil, fmt.Errorf("student not found")
	}
	return &proto.GetStudentResponse{Student: student}, nil
}

// Implement ListStudents method
func (s *universityServer) ListStudents(ctx context.Context, req *proto.ListStudentsRequest) (*proto.ListStudentsResponse, error) {
	students := make([]*proto.Student, 0, len(s.students))
	for _, student := range s.students {
		students = append(students, student)
	}
	return &proto.ListStudentsResponse{Students: students}, nil
}

// Implement other methods (CreateCourse, GetCourse, ListCourses, EnrollStudent) similarly

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterUniversityServiceServer(grpcServer, newServer())
	log.Println("Server is running on port :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
