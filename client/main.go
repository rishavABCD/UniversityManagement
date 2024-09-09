package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"university-management/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := proto.NewUniversityServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a new student
	createStudentReq := &proto.CreateStudentRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	res, err := client.CreateStudent(ctx, createStudentReq)
	if err != nil {
		log.Fatalf("Failed to create student: %v", err)
	}
	fmt.Printf("Created Student: %v\n", res.Student)

	// Fetch the student details
	getStudentReq := &proto.GetStudentRequest{Id: res.Student.Id}
	studentRes, err := client.GetStudent(ctx, getStudentReq)
	if err != nil {
		log.Fatalf("Failed to get student: %v", err)
	}
	fmt.Printf("Fetched Student: %v\n", studentRes.Student)
}
