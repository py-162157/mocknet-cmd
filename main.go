package main

import (
	"context"
	"front-end/rpctest/rpctest"
	"front-end/topology"

	"github.com/desertbit/grumble"
	"google.golang.org/grpc"
)

type Infos struct {
	remoteAaddress string
	topology       rpctest.Emunet
	connection     *grpc.ClientConn
	client         rpctest.MocknetClient
}

func main() {
	var infos Infos

	app := grumble.New(&grumble.Config{
		Name:        "MockNet Cmd",
		Description: "A demo front-end cmd for MockNet",
	})

	createCommand := &grumble.Command{
		Name: "create",
		Help: "create a topoloty",

		Flags: func(f *grumble.Flags) {
			f.String("n", "network", "minimal", "Create a mocknet")
		},

		Args: func(a *grumble.Args) {
			a.String("topo", "topology to be created", grumble.Default("fat-tree,4"))
		},

		Run: func(c *grumble.Context) error {
			// Args.
			topo, err := topology.Make_Topo(c.Args.String("topo"))
			if err != nil {
				c.App.PrintError(err)
			} else {
				infos.topology = topo
				// origin: if infos.remoteAaddress == "" {
				if infos.remoteAaddress != "" {
					c.App.Println("please first input the remote address!")
				} else {
					message := rpctest.Message{
						Type: 0,
						Command: &rpctest.Command{
							EmunetCreation: &rpctest.EmunetCreation{
								Emunet: &infos.topology,
							},
						},
					}
					//c.App.Println(message)
					//client := rpctest.NewMocknetClient(infos.connection)
					if _, err := infos.client.Process(context.Background(), &message); err != nil {
						c.App.PrintError(err)
					}
				}
			}
			return nil
		},
	}

	setCommand := &grumble.Command{
		Name: "set",
		Help: "specify the setting",

		Flags: func(f *grumble.Flags) {
			f.String("r", "remote", "192.168.66.35:10010", "ip address and port")
		},

		Run: func(c *grumble.Context) error {
			infos.remoteAaddress = c.Flags.String("remote")
			c.App.Println("remote address:", infos.remoteAaddress)
			if conn, err := grpc.Dial(infos.remoteAaddress, grpc.WithInsecure(), grpc.WithBlock()); err != nil {
				c.App.PrintError(err)
			} else {
				infos.connection = conn
				infos.client = rpctest.NewMocknetClient(infos.connection)
				c.App.Println("connected to remote controller!")
			}
			return nil
		},
	}

	testCommand := &grumble.Command{
		Name: "test",
		Help: "test the network",

		Run: func(c *grumble.Context) error {
			c.App.Println("begin to test the network")
			message := rpctest.Message{
				Type: 3,
				Command: &rpctest.Command{
					EmunetTest: &rpctest.EmunetTest{},
				},
			}
			if _, err := infos.client.Process(context.Background(), &message); err != nil {
				c.App.PrintError(err)
			}
			return nil
		},
	}

	app.AddCommand(createCommand)
	app.AddCommand(setCommand)
	app.AddCommand(testCommand)

	app.Run()
}
