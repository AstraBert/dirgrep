package cmd

import (
	"context"
	"log"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestMcpServer(t *testing.T) {
	server := GetMcpServer()
	ctx := context.Background()
	serverTransport, clientTransport := mcp.NewInMemoryTransports()
	serverSession, err := server.Connect(ctx, serverTransport, nil)
	if err != nil {
		t.Errorf("Expecting server to be able to connect, got error %s", err.Error())
		log.Fatal(err)
	}
	client := mcp.NewClient(&mcp.Implementation{Name: "client"}, nil)
	clientSession, err := client.Connect(ctx, clientTransport, nil)
	if err != nil {
		t.Errorf("Expecting client to be able to connect, got error %s", err.Error())
		log.Fatal(err)
	}
	res, err := clientSession.CallTool(ctx, &mcp.CallToolParams{
		Name:      "dirgrep",
		Arguments: map[string]any{"directory": "../testfiles", "pattern": "in voluptate velit", "recursive": true, "skip_dirs": []string{"skip"}, "context": 1},
	})
	if err != nil {
		t.Errorf("Expecting tool call to succeed, got error %s", err.Error())
		log.Fatal(err)
	}
	if len(res.Content) != 2 {
		t.Errorf("Expecting the returned contet to have length 2, got %d", len(res.Content))
		log.Fatal("Wrong content length")
	}
	content := res.Content[0]
	typedContent, ok := content.(*mcp.TextContent)
	if !ok {
		t.Error("Expecting content to be of type TextContent, but it is not")
		log.Fatal("Wrong content type")
	}
	text := typedContent.Text
	if text != "Matches found in ../testfiles/loremipsum/part1.txt:\n\n```text\n [bold red]in voluptate velit[/] \n```" && text != "Matches found in ../testfiles/loremipsum/part2.txt:\n\n```text\n [bold red]in voluptate velit[/] \n```" {
		t.Errorf("Got an unexpected result: %s", text)
		log.Fatal("Wrong result")
	}

	_ = clientSession.Close()
	_ = serverSession.Wait()
}
