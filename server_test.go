package main

import (
	"context"
	"testing"
	"time"

	pb "github.com/preimmortal/actions-test/proto"
)

func TestEcho(t *testing.T) {
	var srv Server

	input := "preimmortal"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := srv.ActionEcho(ctx, &pb.EchoRequest{Input: input})
	if err != nil {
		t.Fatal("Could not run ActionsEcho: ", err)
	}

	if reply.GetMessage() != input {
		t.Error("Invalid ActionsEcho Reply")
	}
	t.Log("ActionsEcho Reply: ", reply)

}
