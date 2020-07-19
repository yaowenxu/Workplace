package main

import (
	pb "./pb_pkg"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	test := &pb.Student{
		Name: "yaowenxu",
		Male:  true,
		Scores: []int32{98, 97, 96},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &pb.Student{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}else {
		fmt.Println("运行成功！")
	}
}
