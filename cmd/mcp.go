package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type DirGrepParams struct {
	Directory string   `json:"directory" jsonschema:"Path to the directory to perform the grep operations within"`
	Pattern   string   `json:"pattern" jsonschema:"Pattern to use for grep operations"`
	Context   int      `json:"context" jsonschema:"Context padding around the match (number of charachters)."`
	SkipDirs  []string `json:"skip_dirs" jsonschema:"Directories to skip for the grep operations. Common choices might be '.venv', 'node_modules', '.git'"`
	Recursive bool     `json:"recursive" jsonschema:"Whether or not to perform the grep operations recursively within the specified directory"`
}

func DirGrep(ctx context.Context, req *mcp.CallToolRequest, args DirGrepParams) (*mcp.CallToolResult, any, error) {
	res, err := GrepMany(args.Pattern, args.Directory, args.Recursive, args.SkipDirs, args.Context)
	if err != nil {
		return &mcp.CallToolResult{
			IsError: true,
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("An error occurred: %s\n", err.Error())},
			},
		}, nil, nil
	}
	contents := []mcp.Content{}
	for k := range res {
		if len(res[k]) > 0 {
			content := &mcp.TextContent{Text: fmt.Sprintf("Matches found in %s:\n\n```text\n%s\n```", k, strings.Join(res[k], "\n\n---\n\n"))}
			contents = append(contents, content)
		}
	}
	return &mcp.CallToolResult{
		Content: contents,
	}, nil, nil
}

func GetMcpServer() *mcp.Server {
	server := mcp.NewServer(&mcp.Implementation{Name: "dirgrep-mcp", Version: "1.0.0"}, &mcp.ServerOptions{Instructions: "Use this MCP server in order to perform directory-wide grep operations.", HasTools: true, HasPrompts: false, HasResources: false})
	mcp.AddTool(server, &mcp.Tool{Name: "dirgrep", Description: "`dirgrep` is a tool that allows directory-wide grep operations. It takes, as arguments, the directory where to perform the grep operations, the pattern, the context (number of charachters) for the returned matches, a `recursive` parameter (whether or not to go through the directory recursively) and a `skip_dirs` parameter with the list of sub-directories to skip while performing the grep search."}, DirGrep)
	return server
}
